import { PUBLIC_API_URL } from '$env/static/public';

export class ApiError extends Error {
	constructor(
		public status: number,
		message: string
	) {
		super(message);
	}
}

// Базовий fetcher — передає cookies із запиту (server-side)
export async function fetcher<T>(
	path: string,
	options: RequestInit = {},
	cookieHeader?: string
): Promise<T> {
	const headers: Record<string, string> = {
		'Content-Type': 'application/json',
		...(options.headers as Record<string, string>)
	};

	if (cookieHeader) {
		headers['Cookie'] = cookieHeader;
	}

	const res = await fetch(`${PUBLIC_API_URL}${path}`, {
		...options,
		headers,
		credentials: 'include'
	});

	const body = await res.json();

	if (!res.ok) {
		throw new ApiError(res.status, body?.error ?? 'Unknown error');
	}

	return body.data as T;
}
