import { fail }   from '@sveltejs/kit';
import { ApiError } from '$lib/modules/common/api/fetcher';
import { PUBLIC_API_URL } from '$env/static/public';
import { changeEmailSchema, changePasswordSchema } from '$lib/server/validations/schemas';
import type { PageServerLoad, Actions } from './$types';

export const load: PageServerLoad = async ({ locals }) => {
	return { user: locals.user! };
};

export const actions: Actions = {
	changeEmail: async ({ request, locals }) => {
		const form = await request.formData();
		const raw  = {
			newEmail:        String(form.get('newEmail')        ?? ''),
			currentPassword: String(form.get('currentPassword') ?? '')
		};

		const parsed = changeEmailSchema.safeParse(raw);
		if (!parsed.success) {
			return fail(400, {
				action: 'changeEmail',
				errors: parsed.error.flatten().fieldErrors,
				values: raw
			});
		}

		const cookie = request.headers.get('cookie') ?? '';
		const res = await fetch(`${PUBLIC_API_URL}/api/users/email`, {
			method:  'PUT',
			headers: { 'Content-Type': 'application/json', Cookie: cookie },
			body:    JSON.stringify({
				new_email:        parsed.data.newEmail,
				current_password: parsed.data.currentPassword
			})
		});

		if (res.status === 403) {
			return fail(400, {
				action: 'changeEmail',
				errors: { currentPassword: ['Incorrect password'], newEmail: [] as string[] },
				values: raw
			});
		}
		if (res.status === 409) {
			return fail(400, {
				action: 'changeEmail',
				errors: { newEmail: ['Email already in use'], currentPassword: [] as string[] },
				values: raw
			});
		}
		if (!res.ok) {
			return fail(400, { action: 'changeEmail', errors: { currentPassword: ['Update failed'] }, values: raw });
		}

		return { action: 'changeEmail', success: true };
	},

	changePassword: async ({ request }) => {
		const form = await request.formData();
		const raw  = {
			currentPassword: String(form.get('currentPassword') ?? ''),
			newPassword:     String(form.get('newPassword')     ?? ''),
			confirmPassword: String(form.get('confirmPassword') ?? '')
		};

		const parsed = changePasswordSchema.safeParse(raw);
		if (!parsed.success) {
			return fail(400, {
				action: 'changePassword',
				errors: parsed.error.flatten().fieldErrors
			});
		}

		const cookie = request.headers.get('cookie') ?? '';
		const res = await fetch(`${PUBLIC_API_URL}/api/users/password`, {
			method:  'PUT',
			headers: { 'Content-Type': 'application/json', Cookie: cookie },
			body:    JSON.stringify({
				current_password: parsed.data.currentPassword,
				new_password:     parsed.data.newPassword,
				confirm_password: parsed.data.confirmPassword
			})
		});

		if (res.status === 403) {
			return fail(400, {
				action: 'changePassword',
				errors: {
					currentPassword: ['Incorrect password'],
					newPassword:     [] as string[],
					confirmPassword: [] as string[]
				}
			});
		}
		if (!res.ok) {
			return fail(400, { action: 'changePassword', errors: { currentPassword: ['Update failed'] } });
		}

		return { action: 'changePassword', success: true };
	}
};
