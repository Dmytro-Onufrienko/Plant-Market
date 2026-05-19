import { plants as plantsApi, categories as categoriesApi } from '$lib/modules/plants/api';
import type { PageServerLoad } from './$types';

const LIMIT = 12;

export const load: PageServerLoad = async ({ url, request }) => {
	const cookie       = request.headers.get('cookie') ?? undefined;
	const categorySlug = url.searchParams.get('category') ?? null;
	const page         = Math.max(1, Number(url.searchParams.get('page') ?? 1));

	const [plantsData, catsData] = await Promise.all([
		plantsApi.list({ category: categorySlug ?? undefined, page, limit: LIMIT }, cookie),
		categoriesApi.list(cookie)
	]);

	const categoryMap = new Map(catsData.categories.map(c => [c.id, c.name]));

	const mappedPlants = plantsData.plants.map(p => ({
		id:       p.id,
		name:     p.name,
		slug:     p.slug,
		price:    p.price,
		images:   p.images,
		stock:    p.stock,
		featured: p.featured,
		category: p.category_id ? (categoryMap.get(p.category_id) ?? null) : null
	}));

	return {
		plants:         mappedPlants,
		categories:     catsData.categories,
		activeCategory: categorySlug,
		page,
		hasMore:        plantsData.total > page * LIMIT
	};
};
