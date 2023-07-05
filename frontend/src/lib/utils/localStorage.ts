import { browser } from '$app/environment';

const get = (name: string): string | null => {
    return browser ? localStorage.getItem(name) : null;
};

const set = (name: string, value: string) => {
    if (browser) {
        localStorage.setItem(name, value);
    }
};

const remove = (name: string) => {
    if (browser) {
        localStorage.removeItem(name);
    }
};

export default {
    get,
    set,
    remove
}