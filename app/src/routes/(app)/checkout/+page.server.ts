import { fail, redirect }    from '@sveltejs/kit';
import { ApiError } from '$lib/modules/common/api/fetcher';
import { orders as ordersApi } from '$lib/modules/orders/api';
import { checkoutSchema }    from '$lib/server/validations/schemas';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals }) => {
	if (locals.user?.role === 'admin') redirect(302, '/admin');
	return {};
};

export const actions: Actions = {
	default: async ({ request, locals }) => {
		if (!locals.user) return fail(401, { error: 'Login required' });

		const cookie = request.headers.get('cookie') ?? undefined;
		const form   = await request.formData();

		let cartItems: Array<{ plantId: number; quantity: number; price: number }> = [];
		try {
			cartItems = JSON.parse(String(form.get('cartItems') ?? '[]'));
		} catch {
			return fail(400, { error: 'Invalid cart data' });
		}

		if (!cartItems.length) return fail(400, { error: 'Cart is empty' });

		const raw = {
			name:       String(form.get('name')       ?? ''),
			address:    String(form.get('address')    ?? ''),
			city:       String(form.get('city')       ?? ''),
			country:    String(form.get('country')    ?? ''),
			postalCode: String(form.get('postalCode') ?? '')
		};

		const parsed = checkoutSchema.safeParse(raw);
		if (!parsed.success) {
			return fail(400, { errors: parsed.error.flatten().fieldErrors, values: raw });
		}

		let order;
		try {
			const result = await ordersApi.checkout({
				items: cartItems.map(i => ({ plant_id: i.plantId, quantity: i.quantity })),
				shipping: {
					full_name: parsed.data.name,
					street:    parsed.data.address,
					city:      parsed.data.city,
					zip:       parsed.data.postalCode,
					country:   parsed.data.country
				}
			}, cookie);
			order = result.order;
		} catch (e) {
			if (e instanceof ApiError) {
				if (e.message.includes('stock'))   return fail(400, { error: e.message });
				if (e.message.includes('not found')) return fail(400, { error: e.message });
			}
			return fail(500, { error: 'Checkout failed' });
		}

		redirect(302, `/orders/${order.id}?success=1`);
	}
};
