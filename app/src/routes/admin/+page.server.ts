import { admin } from '$lib/modules/admin/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ request }) => {
	const cookie = request.headers.get('cookie') ?? undefined;
	const data   = await admin.stats(cookie);

	return {
		stats: {
			plants:  data.plants,
			orders:  data.orders,
			reviews: data.reviews,
			users:   data.users
		},
		recentOrders: data.recent_orders.map(o => ({
			id:         o.id,
			status:     o.status,
			totalPrice: o.total_price,
			createdAt:  o.created_at,
			userName:   o.user_name
		}))
	};
};
