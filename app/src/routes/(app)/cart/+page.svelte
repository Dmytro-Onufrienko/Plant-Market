<script lang="ts">
    import * as m from '$lib/paraglide/messages';
    import { cart, cartTotal } from "$lib/modules/orders/stores/cart";
    import { formatPrice, getImageUrl } from "$lib/modules/common/utils/format";
    import { enhance } from "$app/forms";
    import { fly, fade } from "svelte/transition";

    function handleAdd() {
        return ({ result }: { result: any }) => {
            if (result.type === "success" && result.data?.added) {
                cart.add(result.data.added, result.data.added.quantity);
            }
        };
    }
</script>

<svelte:head>
    <title>{m.cart_title()} — Monstera Shop</title>
</svelte:head>

<div class="bg-green-900 py-10">
    <div class="container">
        <p
            class="text-green-400 text-xs font-semibold tracking-widest uppercase mb-1"
        >
            {m.cart_your()}
        </p>
        <h1 class="text-3xl font-bold text-white">{m.cart_title()}</h1>
    </div>
</div>

<div class="container py-10">
    {#if $cart.length === 0}
        <div class="py-24 text-center">
            <p class="text-5xl mb-4">🛒</p>
            <p class="text-stone-500 text-lg mb-2">{m.cart_empty()}</p>
            <p class="text-stone-400 text-sm mb-6">
                {m.cart_empty_msg()}
            </p>
            <a
                href="/shop"
                class="inline-block px-6 py-3 rounded-xl bg-green-700 text-white font-semibold hover:bg-green-600 transition-colors"
            >
                {m.cart_browse()}
            </a>
        </div>
    {:else}
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
            <!-- Items list -->
            <div class="lg:col-span-2 space-y-3">
                {#each $cart as item (item.plantId)}
                    <div
                        in:fly={{ x: -20, duration: 250 }}
                        out:fly={{ x: 20, duration: 200 }}
                        class="bg-white rounded-2xl border border-stone-100 shadow-sm p-4 flex items-center gap-4"
                    >
                        <!-- Image -->
                        <a href="/shop/{item.slug}" class="shrink-0">
                            {#if item.image}
                                <img
                                    src={getImageUrl(item.image)}
                                    alt={item.name}
                                    class="w-20 h-20 rounded-xl object-cover bg-green-100"
                                />
                            {:else}
                                <div
                                    class="w-20 h-20 rounded-xl bg-green-100 flex items-center justify-center"
                                >
                                    <span class="text-2xl">🌿</span>
                                </div>
                            {/if}
                        </a>

                        <!-- Info -->
                        <div class="flex-1 min-w-0">
                            <a
                                href="/shop/{item.slug}"
                                class="font-semibold text-stone-900 hover:text-green-700 transition-colors text-sm leading-snug block truncate"
                            >
                                {item.name}
                            </a>
                            <p class="text-sm text-green-700 font-bold mt-0.5">
                                {formatPrice(item.price)}
                            </p>
                        </div>

                        <!-- Qty controls -->
                        <div class="flex items-center gap-2 shrink-0">
                            <button
                                onclick={() =>
                                    cart.setQty(
                                        item.plantId,
                                        item.quantity - 1,
                                    )}
                                class="w-7 h-7 rounded-full border border-stone-200 text-stone-600 flex items-center justify-center
									hover:border-stone-400 transition-colors text-sm font-bold"
                            >
                                −
                            </button>
                            <span
                                class="w-6 text-center text-sm font-semibold text-stone-900"
                                >{item.quantity}</span
                            >
                            <button
                                onclick={() =>
                                    cart.setQty(
                                        item.plantId,
                                        item.quantity + 1,
                                    )}
                                class="w-7 h-7 rounded-full border border-stone-200 text-stone-600 flex items-center justify-center
									hover:border-stone-400 transition-colors text-sm font-bold"
                            >
                                +
                            </button>
                        </div>

                        <!-- Line total -->
                        <p
                            class="text-sm font-bold text-stone-800 w-16 text-right shrink-0"
                        >
                            {formatPrice(item.price * item.quantity)}
                        </p>

                        <!-- Remove -->
                        <button
                            onclick={() => cart.remove(item.plantId)}
                            class="w-7 h-7 rounded-full text-stone-300 hover:text-red-400 hover:bg-red-50
								flex items-center justify-center transition-colors text-base shrink-0"
                        >
                            ×
                        </button>
                    </div>
                {/each}
            </div>

            <!-- Order summary -->
            <div class="lg:col-span-1">
                <div
                    class="bg-white rounded-2xl border border-stone-100 shadow-sm p-6 sticky top-24"
                >
                    <h2 class="text-lg font-bold text-stone-900 mb-5">
                        {m.cart_order_summary()}
                    </h2>

                    <div class="space-y-2 mb-5">
                        {#each $cart as item}
                            <div
                                class="flex justify-between text-sm text-stone-600"
                            >
                                <span class="truncate mr-2"
                                    >{item.name} × {item.quantity}</span
                                >
                                <span class="shrink-0 font-medium"
                                    >{formatPrice(
                                        item.price * item.quantity,
                                    )}</span
                                >
                            </div>
                        {/each}
                    </div>

                    <div class="border-t border-stone-100 pt-4 mb-6">
                        <div class="flex justify-between">
                            <span class="font-semibold text-stone-700"
                                >{m.cart_total()}</span
                            >
                            <span class="text-xl font-bold text-green-700"
                                >{formatPrice($cartTotal)}</span
                            >
                        </div>
                        <p class="text-xs text-stone-400 mt-1">
                            {m.cart_shipping_calc()}
                        </p>
                    </div>

                    <a
                        href="/checkout"
                        class="block w-full py-3.5 rounded-xl bg-green-700 text-white font-semibold text-sm text-center
							hover:bg-green-600 active:scale-[.98] transition-all"
                    >
                        {m.cart_proceed()}
                    </a>

                    <button
                        onclick={() => cart.clear()}
                        class="block w-full mt-3 py-2 text-xs text-stone-400 hover:text-red-400 transition-colors"
                    >
                        {m.cart_clear()}
                    </button>
                </div>
            </div>
        </div>
    {/if}
</div>
