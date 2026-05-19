import { error, fail }  from '@sveltejs/kit';
import { ApiError } from '$lib/modules/common/api/fetcher';
import { plants as plantsApi, categories as categoriesApi } from '$lib/modules/plants/api';
import { reviews as reviewsApi } from '$lib/modules/reviews/api';
import { reviewSchema }  from '$lib/server/validations/schemas';
import type { PageServerLoad, Actions } from './$types';

export const load: PageServerLoad = async ({ params, locals, request }) => {
	const cookie = request.headers.get('cookie') ?? undefined;

	let data;
	try {
		data = await plantsApi.getBySlug(params.slug, cookie);
	} catch (e) {
		if (e instanceof ApiError && e.status === 404) error(404, 'Plant not found');
		throw e;
	}

	const catsData   = await categoriesApi.list(cookie);
	const categoryMap = new Map(catsData.categories.map(c => [c.id, c.slug]));
	const categoryNameMap = new Map(catsData.categories.map(c => [c.id, c.name]));

	const p = data.plant;

	return {
		plant: {
			id:           p.id,
			name:         p.name,
			slug:         p.slug,
			description:  p.description,
			price:        p.price,
			images:       p.images,
			stock:        p.stock,
			featured:     p.featured,
			category:     p.category_id ? (categoryNameMap.get(p.category_id) ?? null) : null,
			categorySlug: p.category_id ? (categoryMap.get(p.category_id) ?? null) : null
		},
		reviews: data.reviews.map(r => ({
			id:         r.id,
			rating:     r.rating,
			text:       r.text,
			createdAt:  r.created_at,
			authorName: r.user_name
		})),
		avgRating: data.avg_rating > 0 ? data.avg_rating : null,
		user:      locals.user ?? null
	};
};

export const actions: Actions = {
	review: async ({ request, locals, params }) => {
		if (!locals.user) return fail(401, { error: 'Login required' });

		const cookie = request.headers.get('cookie') ?? undefined;

		let plantData;
		try {
			plantData = await plantsApi.getBySlug(params.slug, cookie);
		} catch {
			return fail(404, { error: 'Plant not found' });
		}

		const form   = await request.formData();
		const parsed = reviewSchema.safeParse({
			rating: Number(form.get('rating')),
			text:   String(form.get('text') ?? '')
		});

		if (!parsed.success) {
			return fail(400, { errors: parsed.error.flatten().fieldErrors });
		}

		try {
			await reviewsApi.create({
				plant_id: plantData.plant.id,
				rating:   parsed.data.rating,
				text:     parsed.data.text
			}, cookie);
		} catch (e) {
			if (e instanceof ApiError && e.status === 409) {
				return fail(400, { errors: { rating: ['You already reviewed this plant'], text: [] as string[] } });
			}
			return fail(500, { error: 'Failed to create review' });
		}

		return { success: true };
	}
};
