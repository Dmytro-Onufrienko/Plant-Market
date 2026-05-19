import { sequence }    from '@sveltejs/kit/hooks';
import type { Handle } from '@sveltejs/kit';
import { i18n }        from '$lib/i18n';
import { auth } from '$lib/modules/auth/api';

const authHandle: Handle = async ({ event, resolve }) => {
	const cookie = event.request.headers.get('cookie') ?? undefined;

	try {
		event.locals.user = await auth.me(cookie);
	} catch {
		event.locals.user = null;
	}

	return resolve(event);
};

export const handle: Handle = sequence(i18n.handle(), authHandle);
