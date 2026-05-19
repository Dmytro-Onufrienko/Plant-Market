<script lang="ts">
	import * as m from '$lib/paraglide/messages';
	import { enhance } from '$app/forms';
	import type { ActionData } from './$types';

	let { form }: { form: ActionData } = $props();
</script>

<svelte:head>
	<title>Register — Monstera Shop</title>
</svelte:head>

<div class="min-h-screen bg-green-950 flex items-center justify-center px-4 py-16">
	<!-- Decorative blobs -->
	<div class="absolute inset-0 overflow-hidden pointer-events-none">
		<div class="absolute -top-32 -right-32 w-96 h-96 rounded-full bg-green-700 opacity-20"></div>
		<div class="absolute -bottom-24 -left-24 w-72 h-72 rounded-full bg-green-800 opacity-15"></div>
	</div>

	<div class="relative w-full max-w-md">
		<!-- Logo -->
		<a href="/" class="flex items-center gap-2 justify-center mb-8 group">
			<span class="w-8 h-8 rounded-full bg-green-300 flex items-center justify-center text-green-900 font-bold text-sm group-hover:scale-110 transition-transform">M</span>
			<span class="text-white font-bold text-xl">Monstera Shop</span>
		</a>

		<!-- Card -->
		<div class="bg-white rounded-2xl p-8 shadow-2xl">
			<h1 class="text-2xl font-bold text-stone-950 mb-1">Create account</h1>
			<p class="text-stone-500 text-sm mb-6">Start your plant journey today</p>

			<form method="POST" use:enhance class="space-y-4">
				<!-- Name -->
				<div>
					<label for="name" class="block text-sm font-medium text-stone-700 mb-1">Full name</label>
					<input
						id="name" name="name" type="text"
						value={form?.values?.name ?? ''}
						placeholder="Sofia Kovalenko"
						autocomplete="name"
						class="w-full px-4 py-2.5 rounded-xl border text-stone-950 text-sm transition
							{form?.errors?.name ? 'border-red-400 bg-red-50' : 'border-stone-200 bg-stone-50 focus:border-green-500 focus:bg-white'}
							outline-none"
					/>
					{#if form?.errors?.name}
						<p class="mt-1 text-xs text-red-500">{form.errors.name[0]}</p>
					{/if}
				</div>

				<!-- Email -->
				<div>
					<label for="email" class="block text-sm font-medium text-stone-700 mb-1">Email</label>
					<input
						id="email" name="email" type="email"
						value={form?.values?.email ?? ''}
						placeholder="you@example.com"
						autocomplete="email"
						class="w-full px-4 py-2.5 rounded-xl border text-stone-950 text-sm transition
							{form?.errors?.email ? 'border-red-400 bg-red-50' : 'border-stone-200 bg-stone-50 focus:border-green-500 focus:bg-white'}
							outline-none"
					/>
					{#if form?.errors?.email}
						<p class="mt-1 text-xs text-red-500">{form.errors.email[0]}</p>
					{/if}
				</div>

				<!-- Password -->
				<div>
					<label for="password" class="block text-sm font-medium text-stone-700 mb-1">Password</label>
					<input
						id="password" name="password" type="password"
						placeholder="Min. 8 characters"
						autocomplete="new-password"
						class="w-full px-4 py-2.5 rounded-xl border text-stone-950 text-sm transition
							{form?.errors?.password ? 'border-red-400 bg-red-50' : 'border-stone-200 bg-stone-50 focus:border-green-500 focus:bg-white'}
							outline-none"
					/>
					{#if form?.errors?.password}
						<p class="mt-1 text-xs text-red-500">{form.errors.password[0]}</p>
					{/if}
				</div>

				<button
					type="submit"
					class="w-full py-3 rounded-xl bg-green-700 text-white font-semibold text-sm
						hover:bg-green-600 active:scale-[.98] transition-all"
				>
					Create account
				</button>
			</form>

			<p class="mt-6 text-center text-sm text-stone-500">
				Already have an account?
				<a href="/auth/login" class="text-green-700 font-semibold hover:underline">Sign in</a>
			</p>
		</div>
	</div>
</div>
