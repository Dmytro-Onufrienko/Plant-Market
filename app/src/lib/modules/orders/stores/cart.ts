import { writable, derived, get } from 'svelte/store';
import { browser } from '$app/environment';

export interface CartItem {
	plantId:  number;
	slug:     string;
	name:     string;
	price:    number;   // cents
	image:    string;
	quantity: number;
}

const STORAGE_KEY = 'monstera_cart';

function loadFromStorage(): CartItem[] {
	if (!browser) return [];
	try {
		const raw = localStorage.getItem(STORAGE_KEY);
		return raw ? JSON.parse(raw) : [];
	} catch {
		return [];
	}
}

function createCartStore() {
	const { subscribe, set, update } = writable<CartItem[]>(loadFromStorage());

	// Persist to localStorage on every change
	if (browser) {
		subscribe(items => {
			localStorage.setItem(STORAGE_KEY, JSON.stringify(items));
		});
	}

	return {
		subscribe,

		add(item: Omit<CartItem, 'quantity'>, qty = 1) {
			update(items => {
				const existing = items.find(i => i.plantId === item.plantId);
				if (existing) {
					return items.map(i =>
						i.plantId === item.plantId
							? { ...i, quantity: i.quantity + qty }
							: i
					);
				}
				return [...items, { ...item, quantity: qty }];
			});
		},

		remove(plantId: number) {
			update(items => items.filter(i => i.plantId !== plantId));
		},

		setQty(plantId: number, qty: number) {
			if (qty <= 0) {
				update(items => items.filter(i => i.plantId !== plantId));
			} else {
				update(items => items.map(i =>
					i.plantId === plantId ? { ...i, quantity: qty } : i
				));
			}
		},

		clear() {
			set([]);
		},

		// Snapshot for form submission
		snapshot(): CartItem[] {
			return get({ subscribe });
		}
	};
}

export const cart = createCartStore();

/** Total item count (sum of quantities) */
export const cartCount = derived(cart, $cart =>
	$cart.reduce((sum, i) => sum + i.quantity, 0)
);

/** Total price in cents */
export const cartTotal = derived(cart, $cart =>
	$cart.reduce((sum, i) => sum + i.price * i.quantity, 0)
);
