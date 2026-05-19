import { fail }   from '@sveltejs/kit';
import { admin } from '$lib/modules/admin/api';
import type { PageServerLoad, Actions } from './$types';

export const load: PageServerLoad = async ({ request }) => {
	const cookie = request.headers.get('cookie') ?? undefined;
	const data   = await admin.reviews.list(cookie);

	const reviews = data.reviews.map(r => ({
		id:        r.id,
		rating:    r.rating,
		text:      r.text,
		createdAt: r.created_at,
		userName:  r.user_name,
		userEmail: r.user_email,
		plantName: r.plant_name,
		plantSlug: r.plant_slug
	}));

	return { reviews };
};

export const actions: Actions = {
	delete: async ({ request }) => {
		const cookie = request.headers.get('cookie') ?? undefined;
		const form   = await request.formData();
		const id     = Number(form.get('id'));
		if (!id) return fail(400, { error: 'Invalid id' });

		try {
			await admin.reviews.delete(id, cookie);
		} catch {
			return fail(500, { error: 'Failed to delete review' });
		}

		return { success: true };
	}
};
