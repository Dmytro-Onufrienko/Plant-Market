import { fail }   from '@sveltejs/kit';
import { categories as categoriesApi } from '$lib/modules/plants/api';
import { admin } from '$lib/modules/admin/api';
import type { PageServerLoad, Actions } from './$types';

export const load: PageServerLoad = async ({ request }) => {
	const cookie = request.headers.get('cookie') ?? undefined;

	const [plantsData, catsData] = await Promise.all([
		admin.plants.list(cookie),
		categoriesApi.list(cookie)
	]);

	const categoryMap = new Map(catsData.categories.map(c => [c.id, c.name]));

	const plants = plantsData.plants.map(p => ({
		id:       p.id,
		name:     p.name,
		slug:     p.slug,
		price:    p.price,
		stock:    p.stock,
		featured: p.featured,
		category: p.category_id ? (categoryMap.get(p.category_id) ?? null) : null
	}));

	return { plants };
};

export const actions: Actions = {
	delete: async ({ request }) => {
		const cookie = request.headers.get('cookie') ?? undefined;
		const form   = await request.formData();
		const id     = Number(form.get('id'));
		if (!id) return fail(400, { error: 'Invalid id' });

		try {
			await admin.plants.delete(id, cookie);
		} catch {
			return fail(500, { error: 'Failed to delete plant' });
		}

		return { success: true };
	}
};
