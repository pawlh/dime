import { getCookie, setCookie } from 'typescript-cookie';
import { browser } from '$app/environment';
import type { CookieAttributes } from 'typescript-cookie/dist/types';

const get = (name: string): string | null => {
	if (browser) return getCookie(name) || null;

	return null;
};

const set = (name: string, value: string, options?: CookieAttributes): void => {
	if (browser) setCookie(name, value, options || {});
};

export default { get, set };
