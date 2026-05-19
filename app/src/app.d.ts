declare global {
	namespace App {
		interface Locals {
			user: UserPublic | null;
		}

		interface UserPublic {
			id: string;
			email: string;
			name: string;
			role: 'user' | 'admin';
			created_at: string;
		}

		interface Category {
			id: number;
			name: string;
			slug: string;
		}

		interface Plant {
			id: number;
			name: string;
			slug: string;
			description: string;
			price: number;
			images: string[];
			category_id: number | null;
			stock: number;
			featured: boolean;
			created_at: string;
		}

		interface ShippingAddress {
			full_name: string;
			street: string;
			city: string;
			zip: string;
			country: string;
		}

		interface Order {
			id: number;
			user_id: string;
			status: 'pending' | 'paid' | 'shipped' | 'delivered';
			total_price: number;
			shipping_address: ShippingAddress;
			created_at: string;
		}

		interface OrderItem {
			id: number;
			order_id: number;
			plant_id: number;
			quantity: number;
			price_at_purchase: number;
		}

		interface OrderItemWithPlant extends OrderItem {
			plant_name: string;
			plant_slug: string;
			plant_images: string[];
		}

		interface OrderWithItems extends Order {
			items: OrderItemWithPlant[];
		}

		interface OrderWithUser extends Order {
			user_email: string;
			user_name: string;
		}

		interface Review {
			id: number;
			user_id: string;
			plant_id: number;
			rating: number;
			text: string;
			created_at: string;
		}

		interface ReviewWithUser extends Review {
			user_name: string;
		}

		interface ReviewWithDetails extends Review {
			user_name: string;
			user_email: string;
			plant_name: string;
			plant_slug: string;
		}

		interface BlogPost {
			id: number;
			title: string;
			slug: string;
			content: string;
			excerpt: string;
			cover_image: string | null;
			published: boolean;
			created_at: string;
		}

		interface AdminStats {
			plants: number;
			orders: number;
			reviews: number;
			users: number;
			recent_orders: OrderWithUser[];
		}
	}
}

export {};
