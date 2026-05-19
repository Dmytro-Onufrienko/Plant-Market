<script lang="ts">
	import PlantCard       from '$lib/modules/plants/components/PlantCard.svelte';
	import { formatPrice } from '$lib/modules/common/utils/format';
	import { scrollReveal } from '$lib/modules/common/actions/scrollReveal';
	import * as m          from '$lib/paraglide/messages';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();
</script>

<svelte:head>
	<title>Monstera Shop</title>
</svelte:head>

<!-- ── HERO ─────────────────────────────────────────────── -->
<section class="relative bg-green-900 overflow-hidden min-h-[600px] flex items-center">
	<div class="absolute inset-0 pointer-events-none">
		<div class="absolute top-0 right-0 w-[600px] h-[600px] rounded-full bg-green-700 opacity-25 translate-x-1/3 -translate-y-1/4"></div>
		<div class="absolute bottom-0 left-0 w-80 h-80 rounded-full bg-green-800 opacity-20 -translate-x-1/3 translate-y-1/4"></div>
	</div>

	<div class="absolute right-[5%] top-1/2 -translate-y-1/2 w-[420px] h-[420px] pointer-events-none opacity-40 hidden lg:block leaf-float">
		<svg viewBox="0 0 400 400" fill="none" xmlns="http://www.w3.org/2000/svg">
			<ellipse cx="200" cy="185" rx="130" ry="170" fill="#4A8C4D" transform="rotate(-12 200 200)"/>
			<ellipse cx="255" cy="165" rx="90" ry="130" fill="#3d823a" transform="rotate(15 255 165)" opacity="0.7"/>
			<ellipse cx="170" cy="150" rx="28" ry="22" fill="#1c3d1c" transform="rotate(-12 200 200)"/>
			<ellipse cx="230" cy="130" rx="22" ry="18" fill="#1c3d1c" transform="rotate(-12 200 200)"/>
			<ellipse cx="200" cy="195" rx="20" ry="16" fill="#1c3d1c" transform="rotate(-12 200 200)"/>
			<rect x="195" y="340" width="10" height="55" rx="5" fill="#6b421f"/>
			<rect x="170" y="362" width="60" height="38" rx="8" fill="#6b421f"/>
			<rect x="164" y="356" width="72" height="12" rx="4" fill="#9c6130"/>
		</svg>
	</div>

	<div class="container relative py-20 lg:py-28">
		<div class="max-w-xl">
			<div class="inline-flex items-center gap-2 px-3 py-1.5 rounded-full border border-green-500 text-green-300 text-xs font-medium mb-6">
				<span class="w-1.5 h-1.5 rounded-full bg-green-400 animate-pulse"></span>
				{m.hero_eyebrow()}
			</div>

			<h1 class="text-5xl lg:text-6xl font-bold text-white leading-[1.1] mb-6">
				{@html m.hero_title().replace(/\n/g, '<br><span class="text-green-300">') + '</span>'}
			</h1>

			<p class="text-green-100/80 text-lg leading-relaxed mb-8 max-w-md">
				{m.hero_desc()}
			</p>

			<div class="flex flex-wrap gap-3">
				<a href="/shop"
					class="px-7 py-3.5 rounded-full bg-green-300 text-green-900 font-semibold
						hover:bg-green-200 active:scale-95 transition-all text-sm">
					{m.hero_cta_shop()}
				</a>
				<a href="/blog"
					class="px-7 py-3.5 rounded-full border border-green-500 text-green-200 font-medium
						hover:border-green-300 hover:text-white transition-all text-sm">
					{m.hero_cta_blog()}
				</a>
			</div>

			<div class="flex gap-8 mt-10 pt-8 border-t border-green-800/60">
				{#each [['200+', m.home_stat_varieties()], ['4.9★', m.home_stat_rating()], [m.home_stat_free(), m.home_stat_delivery()]] as [num, label]}
					<div>
						<p class="text-xl font-bold text-white">{num}</p>
						<p class="text-xs text-green-300/70 mt-0.5">{label}</p>
					</div>
				{/each}
			</div>
		</div>
	</div>
</section>

<!-- ── FEATURED PLANTS ───────────────────────────────────── -->
{#if data.featured.length > 0}
<section class="py-16 lg:py-20">
	<div class="container">
		<div use:scrollReveal class="flex items-end justify-between mb-10">
			<div>
				<p class="text-xs font-semibold text-green-600 tracking-widest uppercase mb-1">{m.home_handpicked()}</p>
				<h2 class="text-3xl font-bold text-stone-900">{m.home_featured_plants()}</h2>
			</div>
			<a href="/shop" class="hidden sm:inline-flex items-center gap-1 text-sm font-semibold text-green-700
				border border-green-700 px-4 py-2 rounded-full hover:bg-green-50 transition-colors">
				{m.home_see_all()}
			</a>
		</div>

		<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
			{#each data.featured as plant, i}
				<PlantCard
					id={plant.id}
					name={plant.name}
					slug={plant.slug}
					price={plant.price}
					images={plant.images}
					stock={plant.stock}
					category={plant.category ?? null}
					badge={m.common_featured()}
					delay={i * 80}
				/>
			{/each}
		</div>

		<div class="mt-8 text-center sm:hidden">
			<a href="/shop" class="inline-flex items-center gap-1 text-sm font-semibold text-green-700
				border border-green-700 px-6 py-2.5 rounded-full hover:bg-green-50 transition-colors">
				{m.home_see_all_plants()}
			</a>
		</div>
	</div>
</section>
{/if}

<!-- ── WHY US ─────────────────────────────────────────────── -->
<section class="py-16 bg-bark-50">
	<div class="container">
		<h2 class="text-2xl font-bold text-stone-900 mb-8 text-center">{m.home_why_us()}</h2>
		<div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-5">
			{#each [
				{ icon: '🌱', title: m.home_carefully_selected(), desc: m.home_selected_desc() },
				{ icon: '📦', title: m.home_safe_packaging(),     desc: m.home_packaging_desc() },
				{ icon: '🚚', title: m.home_fast_delivery(),      desc: m.home_delivery_desc() },
				{ icon: '💚', title: m.home_guarantee(),          desc: m.home_guarantee_desc() },
			] as perk}
				<div class="bg-white rounded-2xl p-5 shadow-sm hover:shadow-md transition-shadow">
					<span class="text-2xl block mb-3">{perk.icon}</span>
					<h3 class="font-semibold text-stone-900 mb-1.5">{perk.title}</h3>
					<p class="text-sm text-stone-500 leading-relaxed">{perk.desc}</p>
				</div>
			{/each}
		</div>
	</div>
</section>

<!-- ── REVIEWS ───────────────────────────────────────────── -->
{#if data.latestReviews.length > 0}
<section class="py-16 lg:py-20 bg-green-900 relative overflow-hidden">
	<div class="absolute -top-20 -left-20 w-72 h-72 rounded-full bg-green-800 opacity-30"></div>
	<div class="container relative">
		<div class="mb-10">
			<p class="text-xs font-semibold text-green-400 tracking-widest uppercase mb-1">{m.home_what_they_say()}</p>
			<h2 class="text-3xl font-bold text-white">{m.home_customer_reviews()}</h2>
		</div>

		<div class="grid grid-cols-1 md:grid-cols-3 gap-5">
			{#each data.latestReviews as review}
				<div class="bg-white/[0.07] border border-white/10 rounded-2xl p-6">
					<div class="flex gap-0.5 mb-3">
						{#each Array(review.rating) as _}
							<span class="text-yellow-400 text-sm">★</span>
						{/each}
					</div>
					<p class="text-green-100/90 text-sm leading-relaxed mb-4">"{review.text}"</p>
					<div class="flex items-center gap-2.5">
						<div class="w-8 h-8 rounded-full bg-green-600 flex items-center justify-center text-white text-xs font-bold">
							{review.authorName[0].toUpperCase()}
						</div>
						<div>
							<p class="text-white text-sm font-medium">{review.authorName}</p>
							<p class="text-green-400 text-xs">{m.home_review_on()} {review.plantName}</p>
						</div>
					</div>
				</div>
			{/each}
		</div>
	</div>
</section>
{/if}

<!-- ── CTA BANNER ─────────────────────────────────────────── -->
<section class="py-14 bg-bark-700 text-white text-center">
	<div class="container">
		<h2 class="text-2xl font-bold mb-2">{m.home_ready_grow()}</h2>
		<p class="text-bark-200 mb-6 text-sm">{m.home_ready_desc()}</p>
		<a href="/shop"
			class="inline-flex px-8 py-3.5 rounded-full bg-green-300 text-green-900 font-semibold text-sm
				hover:bg-green-200 active:scale-95 transition-all">
			{m.hero_cta_shop()}
		</a>
	</div>
</section>
