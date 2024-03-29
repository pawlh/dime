import { derived, writable } from 'svelte/store';
import cookies from '$lib/utils/cookies';
import localStorage from '$lib/utils/localStorage';

interface User {
	firstName: string | null;
	lastName: string | null;
}

function getUserFromLocalStorage(): User {
	const firstName = localStorage.get('firstName');
	const lastName = localStorage.get('lastName');

	if (firstName && lastName) {
		return {
			firstName,
			lastName
		} as User;
	}

	return {
		firstName: null,
		lastName: null
	};
}

const createUser = () => {
	const { subscribe, set } = writable<User>(getUserFromLocalStorage());

	function refresh() {
		const loggedIn = cookies.get('loggedIn') === 'true';
		if (!loggedIn) {
			localStorage.remove('firstName');
			localStorage.remove('lastName');

			set({
				firstName: null,
				lastName: null
			});
		} else {
			set(getUserFromLocalStorage());
		}
	}

	function logout() {
		set({
			firstName: null,
			lastName: null
		});
	}

	function login(user: User) {
		set(user);
	}

	refresh();

	return {
		subscribe,
		refresh,
		logout,
		login
	};
};

export const user = createUser();
user.subscribe((value) => {
	if (value?.firstName && value?.lastName) {
		localStorage.set('firstName', value.firstName);
		localStorage.set('lastName', value.lastName);
	} else {
		localStorage.remove('firstName');
		localStorage.remove('lastName');
	}
});

export const isLoggedIn = derived(
	user,
	($user) => $user.firstName != null && $user.lastName != null
);
