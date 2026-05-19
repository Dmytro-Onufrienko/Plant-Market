import { fail, redirect }  from '@sveltejs/kit';
import { PUBLIC_API_URL }  from '$env/static/public';
import { registerSchema }  from '$lib/server/validations/schemas';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals }) => {
	if (locals.user) redirect(302, '/');
};

export const actions: Actions = {
	default: async ({ request, cookies }) => {
		const form = await request.formData();
		const raw  = {
			name:     String(form.get('name')     ?? ''),
			email:    String(form.get('email')    ?? ''),
			password: String(form.get('password') ?? '')
		};

		const parsed = registerSchema.safeParse(raw);
		if (!parsed.success) {
			return fail(400, {
				errors: parsed.error.flatten().fieldErrors,
				values: raw
			});
		}

		const res = await fetch(`${PUBLIC_API_URL}/api/auth/register`, {
			method:  'POST',
			headers: { 'Content-Type': 'application/json' },
			body:    JSON.stringify(parsed.data)
		});

		if (res.status === 409) {
			return fail(400, {
				errors: { name: [] as string[], email: ['This email is already registered'], password: [] as string[] },
				values: raw
			});
		}

		if (!res.ok) {
			return fail(400, {
				errors: { name: [] as string[], email: ['Registration failed'], password: [] as string[] },
				values: raw
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

		redirect(302, '/');
	}
};
