import { blog } from '$lib/modules/blog/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ request }) => {
	const cookie = request.headers.get('cookie') ?? undefined;
	const data   = await blog.list(cookie);

	const posts = data.posts.map(p => ({
		id:         p.id,
		title:      p.title,
		slug:       p.slug,
		excerpt:    p.excerpt,
		coverImage: p.cover_image,
		createdAt:  p.created_at
	}));

	return { posts };
};
