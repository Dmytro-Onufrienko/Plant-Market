import { fetcher } from '$lib/modules/common/api/fetcher';

export const auth = {
	me: (cookie?: string) =>
		fetcher<App.UserPublic>('/api/auth/me', {}, cookie),

	register: (data: { email: string; password: string; name: string }) =>
		fetcher<App.UserPublic>('/api/auth/register', {
			method: 'POST',
			body: JSON.stringify(data)
		}),

	login: (data: { email: string; password: string }) =>
		fetcher<{ user: App.UserPublic }>('/api/auth/login', {
			method: 'POST',
			body: JSON.stringify(data)
		}),

	logout: (cookie?: string) =>
		fetcher<null>('/api/auth/logout', { method: 'POST' }, cookie)
};
