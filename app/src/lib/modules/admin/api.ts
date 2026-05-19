import { fetcher, ApiError } from '$lib/modules/common/api/fetcher';
import { PUBLIC_API_URL } from '$env/static/public';

export const admin = {
	stats: (cookie?: string) =>
		fetcher<App.AdminStats>('/api/admin/stats', {}, cookie),

	plants: {
		list: (cookie?: string) =>
			fetcher<{ plants: App.Plant[] }>('/api/admin/plants', {}, cookie),
		create: (data: unknown, cookie?: string) =>
			fetcher<{ plant: App.Plant }>('/api/admin/plants', {
				method: 'POST',
				body: JSON.stringify(data)
			}, cookie),
		update: (id: number, data: unknown, cookie?: string) =>
			fetcher<{ plant: App.Plant }>(`/api/admin/plants/${id}`, {
				method: 'PUT',
				body: JSON.stringify(data)
			}, cookie),
		delete: (id: number, cookie?: string) =>
			fetcher<null>(`/api/admin/plants/${id}`, { method: 'DELETE' }, cookie),
		upload: async (formData: FormData, cookie?: string): Promise<{ url: string }> => {
			const headers: Record<string, string> = {};
			if (cookie) headers['Cookie'] = cookie;
			const res = await fetch(`${PUBLIC_API_URL}/api/admin/upload`, {
				method: 'POST',
				body: formData,
				headers,
				credentials: 'include'
			});
			const body = await res.json();
			if (!res.ok) throw new ApiError(res.status, body?.error ?? 'Upload failed');
			return body.data as { url: string };
		}
	},

	blog: {
		list: (cookie?: string) =>
			fetcher<{ posts: App.BlogPost[] }>('/api/admin/blog', {}, cookie),
		create: (data: unknown, cookie?: string) =>
			fetcher<{ post: App.BlogPost }>('/api/admin/blog', {
				method: 'POST',
				body: JSON.stringify(data)
			}, cookie),
		update: (id: number, data: unknown, cookie?: string) =>
			fetcher<{ post: App.BlogPost }>(`/api/admin/blog/${id}`, {
				method: 'PUT',
				body: JSON.stringify(data)
			}, cookie),
		delete: (id: number, cookie?: string) =>
			fetcher<null>(`/api/admin/blog/${id}`, { method: 'DELETE' }, cookie)
	},

	orders: {
		list: (cookie?: string) =>
			fetcher<{ orders: App.OrderWithUser[] }>('/api/admin/orders', {}, cookie),
		updateStatus: (id: number, status: string, cookie?: string) =>
			fetcher<null>(`/api/admin/orders/${id}/status`, {
				method: 'PUT',
				body: JSON.stringify({ status })
			}, cookie)
	},

	reviews: {
		list: (cookie?: string) =>
			fetcher<{ reviews: App.ReviewWithDetails[] }>('/api/admin/reviews', {}, cookie),
		delete: (id: number, cookie?: string) =>
			fetcher<null>(`/api/admin/reviews/${id}`, { method: 'DELETE' }, cookie)
	}
};
