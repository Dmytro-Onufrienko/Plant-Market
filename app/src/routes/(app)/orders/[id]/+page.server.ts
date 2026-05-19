import { error }          from '@sveltejs/kit';
import { ApiError } from '$lib/modules/common/api/fetcher';
import { orders as ordersApi } from '$lib/modules/orders/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params, url, request }) => {
	const cookie  = request.headers.get('cookie') ?? undefined;
	const orderId = Number(params.id);

	let data;
	try {
		data = await ordersApi.get(orderId, cookie);
	} catch (e) {
		if (e instanceof ApiError && e.status === 404) error(404, 'Order not found');
		throw e;
	}

	const o = data.order;

	// Translate Go shipping address fields → frontend field names
	const shippingAddress = {
		name:       o.shipping_address.full_name,
		address:    o.shipping_address.street,
		city:       o.shipping_address.city,
		country:    o.shipping_address.country,
		postalCode: o.shipping_address.zip
	};

	const items = o.items.map(item => ({
		id:              item.id,
		quantity:        item.quantity,
		priceAtPurchase: item.price_at_purchase,
		plantName:       item.plant_name,
		plantSlug:       item.plant_slug,
		plantImages:     item.plant_images
	}));

	const justPlaced = url.searchParams.get('success') === '1';

	return {
		order: {
			id:              o.id,
			status:          o.status,
			totalPrice:      o.total_price,
			shippingAddress: o.shipping_address,
			createdAt:       o.created_at,
			userId:          o.user_id
		},
		items,
		shippingAddress,
		justPlaced
	};
};
