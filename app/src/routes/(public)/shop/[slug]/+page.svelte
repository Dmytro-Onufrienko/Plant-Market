<script lang="ts">
    import { enhance } from "$app/forms";
    import { goto } from "$app/navigation";
    import { formatPrice, parseImages, getImageUrl } from "$lib/modules/common/utils/format";
    import { cart } from "$lib/modules/orders/stores/cart";
    import * as m from '$lib/paraglide/messages';
    import type { PageData, ActionData } from "./$types";

    let { data, form }: { data: PageData; form: ActionData } = $props();

    const imgs = parseImages(data.plant.images).map(getImageUrl);
    let activeImg = $state(0);
    function setActive(i: number) { activeImg = i; }

    let reviewRating = $state(5);
    let addedToCart = $state(false);

    function handleAddToCart() {
        return ({ result }: { result: any }) => {
            if (result.type === "success" && result.data?.added) {
                cart.add(result.data.added, result.data.added.quantity);
                addedToCart = true;
                setTimeout(() => { addedToCart = false; }, 2000);
            }
        };
    }
</script>

<svelte:head>
    <title>{data.plant.name} — Monstera Shop</title>
    <meta name="description" content={data.plant.description.slice(0, 160)} />
</svelte:head>

<div class="bg-stone-100 border-b border-stone-200">
    <div class="container py-2.5 text-xs text-stone-400 flex items-center gap-1.5">
        <a href="/" class="hover:text-stone-600">{m.breadcrumb_home()}</a>
        <span>/</span>
        <a href="/shop" class="hover:text-stone-600">{m.nav_shop()}</a>
        {#if data.plant.category}
            <span>/</span>
            <a href="/shop?category={data.plant.categorySlug}" class="hover:text-stone-600">{data.plant.category}</a>
        {/if}
        <span>/</span>
        <span class="text-stone-600">{data.plant.name}</span>
    </div>
</div>

<div class="container py-10 lg:py-14">
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-10 lg:gap-16">
        <div class="space-y-3">
            <div class="aspect-square rounded-2xl overflow-hidden bg-green-100">
                {#if imgs.length > 0 && imgs[activeImg]}
                    <img src={imgs[activeImg]} alt={data.plant.name} class="w-full h-full object-cover" />
                {:else}
                    <div class="w-full h-full flex items-center justify-center">
                        <svg viewBox="0 0 200 200" class="w-40 h-40 text-green-400 opacity-50" fill="currentColor">
                            <ellipse cx="100" cy="90" rx="55" ry="75" transform="rotate(-15 100 100)"/>
                            <ellipse cx="124" cy="84" rx="36" ry="55" transform="rotate(10 124 84)" fill="white" opacity="0.5"/>
                            <ellipse cx="88" cy="72" rx="16" ry="13" transform="rotate(-15 100 100)" fill="#1c3d1c" opacity="0.7"/>
                            <ellipse cx="118" cy="60" rx="13" ry="10" transform="rotate(-15 100 100)" fill="#1c3d1c" opacity="0.7"/>
                            <rect x="97" y="162" width="6" height="30" rx="3" fill="#6b421f"/>
                            <rect x="82" y="182" width="36" height="22" rx="6" fill="#6b421f"/>
                        </svg>
                    </div>
                {/if}
            </div>
            {#if imgs.length > 1}
                <div class="flex gap-2">
                    {#each imgs as img, i}
                        <button onclick={() => setActive(i)}
                            class="w-16 h-16 rounded-xl overflow-hidden border-2 transition-colors
                                {activeImg === i ? 'border-green-600' : 'border-transparent hover:border-green-300'}">
                            <img src={img} alt="" class="w-full h-full object-cover" />
                        </button>
                    {/each}
                </div>
            {/if}
        </div>

        <div class="flex flex-col gap-5">
            {#if data.plant.category}
                <p class="text-xs font-medium text-stone-400 uppercase tracking-widest">{data.plant.category}</p>
            {/if}

            <h1 class="text-3xl lg:text-4xl font-bold text-stone-900 leading-tight">{data.plant.name}</h1>

            {#if data.avgRating !== null}
                <div class="flex items-center gap-2">
                    <div class="flex gap-0.5">
                        {#each Array(5) as _, i}
                            <span class="text-base {i < Math.round(data.avgRating!) ? 'text-yellow-400' : 'text-stone-200'}">★</span>
                        {/each}
                    </div>
                    <span class="text-sm text-stone-500">{data.avgRating!.toFixed(1)} ({data.reviews.length} {data.reviews.length === 1 ? m.review_singular() : m.review_plural()})</span>
                </div>
            {/if}

            <div class="flex items-baseline gap-4">
                <span class="text-3xl font-bold text-green-700">{formatPrice(data.plant.price)}</span>
                <span class="text-sm font-medium {data.plant.stock > 0 ? 'text-green-600' : 'text-red-500'}">
                    {data.plant.stock > 0 ? m.plant_instock({ count: data.plant.stock }) : m.plant_outstock()}
                </span>
            </div>

            <p class="text-stone-600 leading-relaxed">{data.plant.description}</p>

            <div class="flex flex-col sm:flex-row gap-3 pt-2">
                {#if data.plant.stock > 0}
                    {#if data.user}
                        <form method="POST" action="/cart?/add" use:enhance={handleAddToCart} class="flex gap-3 flex-1">
                            <input type="hidden" name="plantId" value={data.plant.id} />
                            <input type="hidden" name="slug" value={data.plant.slug} />
                            <select name="quantity" class="px-3 py-3 rounded-xl border border-stone-200 text-sm bg-white text-stone-700 outline-none focus:border-green-500">
                                {#each Array(Math.min(data.plant.stock, 10)) as _, i}
                                    <option value={i + 1}>{i + 1}</option>
                                {/each}
                            </select>
                            <button class="flex-1 py-3 rounded-xl font-semibold text-sm transition-all active:scale-[.98]
                                {addedToCart ? 'bg-bark-700 text-white' : 'bg-green-700 text-white hover:bg-green-600'}">
                                {addedToCart ? m.plant_added() : m.plant_addcart()}
                            </button>
                        </form>
                    {:else}
                        <a href="/auth/login" class="flex-1 py-3 rounded-xl bg-green-700 text-white font-semibold text-sm text-center hover:bg-green-600 transition-colors">
                            {m.plant_signin_buy()}
                        </a>
                    {/if}
                {:else}
                    <button disabled class="flex-1 py-3 rounded-xl bg-stone-100 text-stone-400 font-semibold text-sm cursor-not-allowed">
                        {m.plant_outstock()}
                    </button>
                {/if}
            </div>

            <div class="grid grid-cols-3 gap-3 pt-2">
                {#each [m.plant_handpicked(), m.plant_eco(), m.plant_guarantee()] as label}
                    <div class="flex flex-col items-center gap-1 p-3 rounded-xl bg-stone-50 text-center">
                        <span class="text-xs text-stone-500 font-medium leading-tight">{label}</span>
                    </div>
                {/each}
            </div>
        </div>
    </div>

    <div class="mt-16 border-t border-stone-200 pt-12">
        <h2 class="text-2xl font-bold text-stone-900 mb-8">
            {m.review_title()}
            {#if data.reviews.length > 0}
                <span class="text-base font-normal text-stone-400 ml-2">({data.reviews.length})</span>
            {/if}
        </h2>

        <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
            <div class="space-y-4">
                {#if data.reviews.length === 0}
                    <p class="text-stone-400 text-sm">{m.review_empty()}</p>
                {:else}
                    {#each data.reviews as review}
                        <div class="bg-stone-50 rounded-2xl p-5">
                            <div class="flex items-center justify-between mb-2">
                                <div class="flex items-center gap-2">
                                    <div class="w-7 h-7 rounded-full bg-green-600 flex items-center justify-center text-white text-xs font-bold">
                                        {review.authorName[0].toUpperCase()}
                                    </div>
                                    <span class="text-sm font-medium text-stone-700">{review.authorName}</span>
                                </div>
                                <div class="flex gap-0.5">
                                    {#each Array(5) as _, i}
                                        <span class="text-sm {i < review.rating ? 'text-yellow-400' : 'text-stone-200'}">★</span>
                                    {/each}
                                </div>
                            </div>
                            <p class="text-sm text-stone-600 leading-relaxed">{review.text}</p>
                        </div>
                    {/each}
                {/if}
            </div>

            <div>
                {#if data.user}
                    <div class="bg-green-50 rounded-2xl p-6 border border-green-100">
                        <h3 class="font-semibold text-stone-900 mb-4">{m.review_share()}</h3>
                        {#if form?.success}
                            <p class="text-green-700 text-sm font-medium">✓ {m.review_submit()}</p>
                        {:else}
                            <form method="POST" action="?/review" use:enhance class="space-y-4">
                                <div>
                                    <label class="block text-sm font-medium text-stone-700 mb-2">{m.review_rating()}</label>
                                    <div class="flex gap-1">
                                        {#each [1, 2, 3, 4, 5] as star}
                                            <button type="button" onclick={() => (reviewRating = star)}
                                                class="text-2xl transition-transform hover:scale-110
                                                    {star <= reviewRating ? 'text-yellow-400' : 'text-stone-200'}">★</button>
                                        {/each}
                                    </div>
                                    <input type="hidden" name="rating" value={reviewRating} />
                                </div>
                                <div>
                                    <textarea name="text" rows="4" placeholder={m.review_placeholder()}
                                        class="w-full px-4 py-3 rounded-xl border border-stone-200 text-sm resize-none outline-none focus:border-green-500 transition-colors bg-white"></textarea>
                                    {#if form?.errors?.text}
                                        <p class="mt-1 text-xs text-red-500">{form.errors.text[0]}</p>
                                    {/if}
                                </div>
                                <button type="submit" class="w-full py-2.5 rounded-xl bg-green-700 text-white text-sm font-semibold hover:bg-green-600 active:scale-[.98] transition-all">
                                    {m.review_submit()}
                                </button>
                            </form>
                        {/if}
                    </div>
                {:else}
                    <div class="bg-stone-50 rounded-2xl p-6 border border-stone-200 text-center">
                        <p class="text-stone-500 text-sm mb-3">{m.review_no_auth()}</p>
                        <a href="/auth/login" class="inline-block px-5 py-2.5 rounded-xl bg-green-700 text-white text-sm font-semibold hover:bg-green-600 transition-colors">
                            {m.nav_signin()}
                        </a>
                    </div>
                {/if}
            </div>
        </div>
    </div>
</div>
