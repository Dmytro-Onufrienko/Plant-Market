import { redirect }           from '@sveltejs/kit';
import { orders as ordersApi } from '$lib/modules/orders/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ request, locals }) => {
	if (locals.user?.role === 'admin') redirect(302, '/admin/orders');

	const cookie = request.headers.get('cookie') ?? undefined;
	const data   = await ordersApi.list(cookie);

	const mappedOrders = data.orders.map(o => ({
		id:         o.id,
		status:     o.status,
		totalPrice: o.total_price,
		createdAt:  o.created_at
	}));

	return { orders: mappedOrders };
};
