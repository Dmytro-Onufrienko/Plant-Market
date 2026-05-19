import { fail, redirect }  from '@sveltejs/kit';
import { PUBLIC_API_URL }  from '$env/static/public';
import { loginSchema }     from '$lib/server/validations/schemas';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals, url }) => {
	if (locals.user) redirect(302, '/');
	return { redirectTo: url.searchParams.get('redirectTo') ?? '/' };
};

export const actions: Actions = {
	default: async ({ request, cookies, url }) => {
		const form = await request.formData();
		const raw  = {
			email:    String(form.get('email')    ?? ''),
			password: String(form.get('password') ?? '')
		};

		const parsed = loginSchema.safeParse(raw);
		if (!parsed.success) {
			return fail(400, {
				errors: parsed.error.flatten().fieldErrors,
				values: { email: raw.email }
			});
		}

		const res = await fetch(`${PUBLIC_API_URL}/api/auth/login`, {
			method:  'POST',
			headers: { 'Content-Type': 'application/json' },
			body:    JSON.stringify(parsed.data)
		});

		if (!res.ok) {
			return fail(400, {
				errors: { email: ['Invalid email or password'], password: [] as string[] },
				values: { email: raw.email }
			});
		}

		// Forward the httpOnly JWT cookie from Go to the browser
		const setCookie = res.headers.get('set-cookie');
		if (setCookie) {
			const [nameValue, ...attrs] = setCookie.split(';');
			const eqIdx = nameValue.indexOf('=');
			const name  = nameValue.slice(0, eqIdx).trim();
			const value = nameValue.slice(eqIdx + 1).trim();
			const maxAgeAttr = attrs.find(a => a.trim().toLowerCase().startsWith('max-age'));
			const maxAge = maxAgeAttr ? parseInt(maxAgeAttr.split('=')[1]) : 86400;
			cookies.set(name, value, { path: '/', httpOnly: true, maxAge, sameSite: 'lax' });
		}

		const redirectTo = url.searchParams.get('redirectTo');
		redirect(302, redirectTo?.startsWith('/') ? redirectTo : '/');
	}
};
