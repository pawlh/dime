import localStorage from '$lib/utils/localStorage';
import { writable } from 'svelte/store';

type Theme = 'light' | 'dark';

const createTheme = () => {
	const { subscribe, set } = writable<Theme>('light');

	if (localStorage.get('theme') === 'dark') {
		set('dark');
	}

	function setLight() {
		localStorage.set('theme', 'light');
		set('light');
	}

	function setDark() {
		localStorage.set('theme', 'dark');
		set('dark');
	}

	function toggle() {
		if (localStorage.get('theme') === 'dark') {
			setLight();
		} else {
			setDark();
		}
	}

	return {
		subscribe,
		toggle
	};
};

export const theme = createTheme();
