import { fail }   from '@sveltejs/kit';
import { admin } from '$lib/modules/admin/api';
import type { PageServerLoad, Actions } from './$types';

export const load: PageServerLoad = async ({ request }) => {
	const cookie = request.headers.get('cookie') ?? undefined;
	const data   = await admin.orders.list(cookie);

	const orders = data.orders.map(o => ({
		id:         o.id,
		status:     o.status,
		totalPrice: o.total_price,
		createdAt:  o.created_at,
		userName:   o.user_name,
		userEmail:  o.user_email
	}));

	return { orders };
};

export const actions: Actions = {
	setStatus: async ({ request }) => {
		const cookie = request.headers.get('cookie') ?? undefined;
		const form   = await request.formData();
		const id     = Number(form.get('id'));
		const status = String(form.get('status') ?? '');

		const validStatuses = ['pending', 'paid', 'shipped', 'delivered'];
		if (!id || !validStatuses.includes(status)) {
			return fail(400, { error: 'Invalid data' });
		}

		try {
			await admin.orders.updateStatus(id, status, cookie);
		} catch {
			return fail(500, { error: 'Failed to update status' });
		}

		return { success: true };
	}
};
