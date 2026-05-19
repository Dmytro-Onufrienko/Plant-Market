import { PUBLIC_API_URL } from "$env/static/public";

/** Convert integer cents to "€ 34.99" */
export function formatPrice(cents: number): string {
  return `€ ${(cents / 100).toFixed(2)}`;
}

/** Get full absolute URL for an image path */
export function getImageUrl(path: string | undefined | null): string {
  if (!path) return "";
  if (
    path.startsWith("http://") ||
    path.startsWith("https://") ||
    path.startsWith("data:")
  )
    return path;
  return `${PUBLIC_API_URL}${path.startsWith("/") ? "" : "/"}${path}`;
}

/** Parse a JSON images field safely — handles both JSON string and plain array */
export function parseImages(raw: string | string[]): string[] {
  if (Array.isArray(raw)) return raw;
  try {
    const arr = JSON.parse(raw);
    return Array.isArray(arr) ? arr : [];
  } catch {
    return [];
  }
}

/** Truncate text to n chars */
export function truncate(str: string, n: number): string {
  return str.length > n ? str.slice(0, n - 1) + "…" : str;
}
