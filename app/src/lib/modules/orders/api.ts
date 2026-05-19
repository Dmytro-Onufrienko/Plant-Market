import { fetcher } from '$lib/modules/common/api/fetcher';

export const orders = {
	list: (cookie?: string) =>
		fetcher<{ orders: App.Order[] }>('/api/orders', {}, cookie),

	get: (id: number, cookie?: string) =>
		fetcher<{ order: App.OrderWithItems }>(`/api/orders/${id}`, {}, cookie),

	checkout: (data: { items: { plant_id: number; quantity: number }[]; shipping: App.ShippingAddress }, cookie?: string) =>
		fetcher<{ order: App.Order }>('/api/orders', {
			method: 'POST',
			body: JSON.stringify(data)
		}, cookie)
};
