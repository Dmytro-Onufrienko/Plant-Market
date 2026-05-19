<script lang="ts">
    import {
        formatPrice,
        parseImages,
        getImageUrl,
    } from "$lib/modules/common/utils/format";
    import { scrollReveal } from "$lib/modules/common/actions/scrollReveal";
    import type { ScrollRevealOptions } from "$lib/modules/common/actions/scrollReveal";
    import * as m from "$lib/paraglide/messages";

    let {
        id,
        name,
        slug,
        price,
        images,
        stock,
        category,
        badge,
        delay = 0,
    }: {
        id: number;
        name: string;
        slug: string;
        price: number;
        images: string | string[];
        stock: number;
        category: string | null;
        badge?: string;
        delay?: number;
    } = $props();

    const imgs = $derived(parseImages(images).map(getImageUrl));
    const img = $derived(imgs[0] ?? null);
    const inStock = $derived(stock > 0);
</script>

<article
    use:scrollReveal={{ delay, duration: 450 }}
    class="group bg-white rounded-2xl overflow-hidden shadow-sm hover:shadow-xl
	transition-all duration-300 hover:-translate-y-1 flex flex-col"
>
    <a
        href="/shop/{slug}"
        class="block relative aspect-[4/3] bg-green-100 overflow-hidden"
    >
        {#if img}
            <img
                src={img}
                alt={name}
                class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105"
                loading="lazy"
            />
        {:else}
            <div
                class="w-full h-full flex items-center justify-center bg-gradient-to-br from-green-100 to-green-200"
            >
                <svg
                    viewBox="0 0 100 100"
                    class="w-24 h-24 text-green-400 opacity-60"
                    fill="currentColor"
                >
                    <ellipse
                        cx="50"
                        cy="45"
                        rx="28"
                        ry="38"
                        transform="rotate(-15 50 50)"
                    />
                    <ellipse
                        cx="62"
                        cy="42"
                        rx="18"
                        ry="28"
                        transform="rotate(10 62 42)"
                        fill="white"
                        opacity="0.5"
                    />
                    <line
                        x1="50"
                        y1="83"
                        x2="50"
                        y2="100"
                        stroke="currentColor"
                        stroke-width="3"
                        class="text-bark-500"
                    />
                </svg>
            </div>
        {/if}

        <div class="absolute top-3 left-3 flex flex-col gap-1.5">
            {#if badge}
                <span
                    class="px-2.5 py-1 rounded-full text-xs font-semibold bg-green-700 text-white"
                >
                    {badge}
                </span>
            {/if}
            {#if !inStock}
                <span
                    class="px-2.5 py-1 rounded-full text-xs font-semibold bg-stone-500 text-white"
                >
                    {m.plant_soldout()}
                </span>
            {/if}
        </div>
    </a>

    <div class="p-4 flex flex-col flex-1">
        {#if category}
            <p class="text-xs text-stone-400 mb-0.5">{category}</p>
        {/if}

        <a
            href="/shop/{slug}"
            class="font-semibold text-stone-900 hover:text-green-700 transition-colors leading-tight mb-3 flex-1"
        >
            {name}
        </a>

        <div class="flex items-center justify-between gap-2 mt-auto">
            <span class="text-xl font-bold text-green-700"
                >{formatPrice(price)}</span
            >

            <a
                href="/shop/{slug}"
                class="px-4 py-2 rounded-xl text-sm font-semibold transition-all
					{inStock
                    ? 'bg-green-700 text-white hover:bg-green-600 active:scale-95'
                    : 'bg-stone-100 text-stone-400 cursor-not-allowed pointer-events-none'}"
            >
                {inStock ? m.plant_view() : m.plant_soldout()}
            </a>
        </div>
    </div>
</article>
