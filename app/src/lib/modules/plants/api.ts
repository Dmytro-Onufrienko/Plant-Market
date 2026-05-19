import { fetcher } from '$lib/modules/common/api/fetcher';

export const plants = {
	list: (params?: { category?: string; page?: number; featured?: boolean; limit?: number }, cookie?: string) => {
		const q = new URLSearchParams();
		if (params?.category) q.set('category', params.category);
		if (params?.page) q.set('page', String(params.page));
		if (params?.featured) q.set('featured', 'true');
		if (params?.limit) q.set('limit', String(params.limit));
		const qs = q.toString();
		return fetcher<{ plants: App.Plant[]; total: number; page: number }>(
			`/api/plants${qs ? '?' + qs : ''}`,
			{},
			cookie
		);
	},

	getBySlug: (slug: string, cookie?: string) =>
		fetcher<{ plant: App.Plant; reviews: App.ReviewWithUser[]; avg_rating: number }>(
			`/api/plants/${slug}`,
			{},
			cookie
		),

	getByID: (id: number, cookie?: string) =>
		fetcher<{ plant: App.Plant }>(`/api/plants/by-id/${id}`, {}, cookie)
};

export const categories = {
	list: (cookie?: string) =>
		fetcher<{ categories: App.Category[] }>('/api/categories', {}, cookie)
};
