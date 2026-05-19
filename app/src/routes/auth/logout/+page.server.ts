import { redirect }     from '@sveltejs/kit';
import { auth } from '$lib/modules/auth/api';
import type { Actions }  from './$types';

export const actions: Actions = {
	default: async ({ locals, cookies, request }) => {
		if (!locals.user) redirect(302, '/');

		const cookie = request.headers.get('cookie') ?? undefined;
		await auth.logout(cookie).catch(() => {});

		cookies.delete('monstera_token', { path: '/' });
		redirect(302, '/');
	}
};
