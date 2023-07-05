import { goto } from '$app/navigation';
import { browser } from '$app/environment';
import { page as pageStore } from '$app/stores';
import type { Page } from '@sveltejs/kit';

/**
 * Navigate to an internal page
 * @param path 
 */
export function navigateTo(path: string) {
	if (browser) {
		const page = getPage();

		if (page.url.pathname !== path) goto(path);
	}
}

export function getPage(): Page {
	let currentPage = null;
	const unsubscribe = pageStore.subscribe((value) => {
		currentPage = value;
	});
	unsubscribe();

	if (currentPage == null) {
		throw Error('something went wrong, page returned null');
	}

	return currentPage;
}
