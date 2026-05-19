<script lang="ts">
	import PlantCard from '$lib/modules/plants/components/PlantCard.svelte';
	import * as m from '$lib/paraglide/messages';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<svelte:head>
	<title>{m.nav_shop()} — Monstera Shop</title>
</svelte:head>

<div class="bg-green-900 py-10 relative overflow-hidden">
	<div class="absolute right-8 top-1/2 -translate-y-1/2 opacity-10 pointer-events-none hidden lg:block leaf-sway">
		<svg viewBox="0 0 200 260" class="w-36 h-auto" fill="white">
			<ellipse cx="100" cy="120" rx="70" ry="110" transform="rotate(-10 100 130)"/>
			<ellipse cx="120" cy="105" rx="45" ry="80" transform="rotate(15 120 105)" fill="#1c3d1c" opacity="0.6"/>
			<ellipse cx="75" cy="90" rx="22" ry="18" transform="rotate(-10 100 130)" fill="#1c3d1c" opacity="0.7"/>
			<ellipse cx="115" cy="70" rx="18" ry="14" transform="rotate(-10 100 130)" fill="#1c3d1c" opacity="0.7"/>
			<ellipse cx="95" cy="130" rx="16" ry="13" transform="rotate(-10 100 130)" fill="#1c3d1c" opacity="0.6"/>
			<rect x="97" y="220" width="6" height="35" rx="3" fill="#6b421f"/>
		</svg>
	</div>
	<div class="container relative">
		<p class="text-green-400 text-xs font-semibold tracking-widest uppercase mb-1">{m.shop_our()}</p>
		<h1 class="text-3xl font-bold text-white">{m.shop_title()}</h1>
	</div>
</div>

<div class="container py-8 relative">
	<div class="flex flex-wrap gap-2 mb-8">
		<a
			href="/shop"
			class="px-4 py-2 rounded-full text-sm font-medium transition-colors
				{!data.activeCategory ? 'bg-green-700 text-white' : 'bg-white text-stone-600 border border-stone-200 hover:border-green-400 hover:text-green-700'}"
		>
			{m.shop_all()}
		</a>
		{#each data.categories as cat}
			<a
				href="/shop?category={cat.slug}"
				class="px-4 py-2 rounded-full text-sm font-medium transition-colors capitalize
					{data.activeCategory === cat.slug ? 'bg-green-700 text-white' : 'bg-white text-stone-600 border border-stone-200 hover:border-green-400 hover:text-green-700'}"
			>
				{cat.name}
			</a>
		{/each}
	</div>

	{#if data.plants.length === 0}
		<div class="py-20 text-center">
			<p class="text-4xl mb-3">🌿</p>
			<p class="text-stone-500">{m.shop_empty()}</p>
			<a href="/shop" class="mt-4 inline-block text-sm text-green-700 font-medium hover:underline">{m.shop_browse_all()}</a>
		</div>
	{:else}
		<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
			{#each data.plants as plant, i}
				<PlantCard
					id={plant.id}
					name={plant.name}
					slug={plant.slug}
					price={plant.price}
					images={plant.images}
					stock={plant.stock}
					category={plant.category ?? null}
					badge={plant.featured ? m.common_featured() : undefined}
					delay={i * 60}
				/>
			{/each}
		</div>

		<div class="flex justify-center gap-3 mt-12">
			{#if data.page > 1}
				<a href="/shop?page={data.page - 1}{data.activeCategory ? `&category=${data.activeCategory}` : ''}"
					class="px-5 py-2.5 rounded-full border border-stone-200 text-sm font-medium hover:border-green-400 transition-colors">
					{m.shop_prev()}
				</a>
			{/if}
			{#if data.hasMore}
				<a href="/shop?page={data.page + 1}{data.activeCategory ? `&category=${data.activeCategory}` : ''}"
					class="px-5 py-2.5 rounded-full bg-green-700 text-white text-sm font-medium hover:bg-green-600 transition-colors">
					{m.shop_next()}
				</a>
			{/if}
		</div>
	{/if}
</div>
