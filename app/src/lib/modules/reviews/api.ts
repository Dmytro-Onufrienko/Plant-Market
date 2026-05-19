import { fetcher } from '$lib/modules/common/api/fetcher';

export const reviews = {
	latest: (limit = 3, cookie?: string) =>
		fetcher<{ reviews: App.ReviewWithDetails[] }>(`/api/reviews/latest?limit=${limit}`, {}, cookie),

	create: (data: { plant_id: number; rating: number; text: string }, cookie?: string) =>
		fetcher<null>('/api/reviews', { method: 'POST', body: JSON.stringify(data) }, cookie)
};
