import { fail }   from '@sveltejs/kit';
import { admin } from '$lib/modules/admin/api';
import type { PageServerLoad, Actions } from './$types';

export const load: PageServerLoad = async ({ request }) => {
	const cookie = request.headers.get('cookie') ?? undefined;
	const data   = await admin.blog.list(cookie);

	const posts = data.posts.map(p => ({
		id:        p.id,
		title:     p.title,
		slug:      p.slug,
		published: p.published,
		createdAt: p.created_at
	}));

	return { posts };
};

export const actions: Actions = {
	delete: async ({ request }) => {
		const cookie = request.headers.get('cookie') ?? undefined;
		const form   = await request.formData();
		const id     = Number(form.get('id'));
		if (!id) return fail(400, { error: 'Invalid id' });

		try {
			await admin.blog.delete(id, cookie);
		} catch {
			return fail(500, { error: 'Failed to delete post' });
		}

		return { success: true };
	},

	togglePublish: async ({ request }) => {
		const cookie    = request.headers.get('cookie') ?? undefined;
		const form      = await request.formData();
		const id        = Number(form.get('id'));
		const published = form.get('published') === 'true';
		if (!id) return fail(400);

		try {
			await admin.blog.update(id, { published: !published }, cookie);
		} catch {
			return fail(500, { error: 'Failed to toggle publish' });
		}

		return { success: true };
	}
};
