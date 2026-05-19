import { plants as plantsApi } from '$lib/modules/plants/api';
import { reviews as reviewsApi } from '$lib/modules/reviews/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ request }) => {
	const cookie = request.headers.get('cookie') ?? undefined;

	const [featuredData, reviewsData] = await Promise.all([
		plantsApi.list({ featured: true, limit: 4 }, cookie),
		reviewsApi.latest(3, cookie)
	]);

	const featured = featuredData.plants.map(p => ({
		id:       p.id,
		name:     p.name,
		slug:     p.slug,
		price:    p.price,
		images:   p.images,
		stock:    p.stock,
		category: null as string | null
	}));

	const latestReviews = reviewsData.reviews.map(r => ({
		id:         r.id,
		rating:     r.rating,
		text:       r.text,
		createdAt:  r.created_at,
		authorName: r.user_name,
		plantName:  r.plant_name
	}));

	return { featured, latestReviews };
};
