import { z } from 'zod';

// ─── Auth ─────────────────────────────────────────────────
export const registerSchema = z.object({
	name:     z.string().min(2, 'Name must be at least 2 characters'),
	email:    z.email('Invalid email address'),
	password: z.string().min(8, 'Password must be at least 8 characters')
});

export const loginSchema = z.object({
	email:    z.email('Invalid email address'),
	password: z.string().min(1, 'Password is required')
});

export const changeEmailSchema = z.object({
	newEmail:        z.email('Invalid email address'),
	currentPassword: z.string().min(1, 'Current password is required')
});

export const changePasswordSchema = z.object({
	currentPassword: z.string().min(1, 'Current password is required'),
	newPassword:     z.string().min(8, 'New password must be at least 8 characters'),
	confirmPassword: z.string()
}).refine(d => d.newPassword === d.confirmPassword, {
	message: 'Passwords do not match',
	path:    ['confirmPassword']
});

// ─── Plants ───────────────────────────────────────────────
export const plantSchema = z.object({
	name:        z.string().min(1, 'Name is required'),
	slug:        z.string().min(1).regex(/^[a-z0-9-]+$/, 'Slug must be lowercase with hyphens'),
	description: z.string().default(''),
	price:       z.number().int().positive('Price must be positive'),  // cents
	categoryId:  z.number().int().nullable().optional(),
	stock:       z.number().int().min(0),
	featured:    z.boolean().default(false)
});

// ─── Orders ───────────────────────────────────────────────
export const checkoutSchema = z.object({
	name:       z.string().min(2, 'Full name is required'),
	address:    z.string().min(5, 'Address is required'),
	city:       z.string().min(2, 'City is required'),
	country:    z.string().min(2, 'Country is required'),
	postalCode: z.string().min(4, 'Postal code is required'),
	// Mock card (not stored in DB)
	cardNumber: z.string().optional(),
	cardExpiry: z.string().optional(),
	cardCvc:    z.string().optional()
});

// ─── Reviews ──────────────────────────────────────────────
export const reviewSchema = z.object({
	rating: z.number().int().min(1).max(5),
	text:   z.string().min(10, 'Review must be at least 10 characters').max(1000)
});

// ─── Blog ─────────────────────────────────────────────────
export const blogPostSchema = z.object({
	title:      z.string().min(1, 'Title is required'),
	slug:       z.string().min(1).regex(/^[a-z0-9-]+$/),
	content:    z.string().default(''),
	excerpt:    z.string().max(300).default(''),
	published:  z.boolean().default(false)
});

// ─── Types ────────────────────────────────────────────────
export type RegisterInput   = z.infer<typeof registerSchema>;
export type LoginInput      = z.infer<typeof loginSchema>;
export type PlantInput      = z.infer<typeof plantSchema>;
export type CheckoutInput   = z.infer<typeof checkoutSchema>;
export type ReviewInput     = z.infer<typeof reviewSchema>;
export type BlogPostInput   = z.infer<typeof blogPostSchema>;
