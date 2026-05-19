import { fetcher } from '$lib/modules/common/api/fetcher';

export const blog = {
	list: (cookie?: string) =>
		fetcher<{ posts: App.BlogPost[] }>('/api/blog', {}, cookie),

	getBySlug: (slug: string, cookie?: string) =>
		fetcher<{ post: App.BlogPost }>(`/api/blog/${slug}`, {}, cookie)
};
