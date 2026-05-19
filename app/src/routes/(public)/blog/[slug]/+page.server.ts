import { error }   from '@sveltejs/kit';
import { ApiError } from '$lib/modules/common/api/fetcher';
import { blog } from '$lib/modules/blog/api';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async ({ params, request }) => {
	const cookie = request.headers.get('cookie') ?? undefined;

	let data;
	try {
		data = await blog.getBySlug(params.slug, cookie);
	} catch (e) {
		if (e instanceof ApiError && e.status === 404) error(404, 'Post not found');
		throw e;
	}

	const p = data.post;
	return {
		post: {
			id:         p.id,
			title:      p.title,
			slug:       p.slug,
			content:    p.content,
			excerpt:    p.excerpt,
			coverImage: p.cover_image,
			published:  p.published,
			createdAt:  p.created_at
		}
	};
};
