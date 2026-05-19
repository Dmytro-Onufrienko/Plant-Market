import { fail, redirect } from '@sveltejs/kit';
import { ApiError } from '$lib/modules/common/api/fetcher';
import { plants as plantsApi } from '$lib/modules/plants/api';
import type { Actions, PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ locals }) => {
	if (locals.user?.role === 'admin') redirect(302, '/admin');
	return {};
};

export const actions: Actions = {
	add: async ({ request, locals }) => {
		if (!locals.user) return fail(401, { error: 'Login required' });

		const cookie  = request.headers.get('cookie') ?? undefined;
		const form    = await request.formData();
		const plantId = Number(form.get('plantId'));
		const qty     = Math.max(1, Number(form.get('quantity') ?? 1));

		if (!plantId) return fail(400, { error: 'Invalid plant' });

		let plantData;
		try {
			plantData = await plantsApi.getByID(plantId, cookie);
		} catch (e) {
			if (e instanceof ApiError && e.status === 404) return fail(404, { error: 'Plant not found' });
			return fail(500, { error: 'Failed to fetch plant' });
		}

		const plant = plantData.plant;
		if (plant.stock < qty) return fail(400, { error: 'Not enough stock' });

		const image = plant.images[0] ?? '';

		return {
			added: {
				plantId:  plant.id,
				slug:     plant.slug,
				name:     plant.name,
				price:    plant.price,
				image,
				quantity: qty
			}
		};
	}
};
