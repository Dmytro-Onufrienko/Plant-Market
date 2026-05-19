<script lang="ts">
    import { enhance } from "$app/forms";
    import { goto } from "$app/navigation";
    import { cart, cartTotal } from "$lib/modules/orders/stores/cart";
    import { formatPrice, getImageUrl } from "$lib/modules/common/utils/format";
    import * as m from '$lib/paraglide/messages';
    import type { ActionData } from "./$types";

    let { form }: { form: ActionData } = $props();

    // After successful order, clear cart and let redirect handle it
    function handleSubmit() {
        return ({ result }: { result: any }) => {
            if (result.type === "redirect") {
                cart.clear();
                goto(result.location);
            }
        };
    }
</script>

<svelte:head>
    <title>{m.checkout_title()} — Monstera Shop</title>
</svelte:head>

<div class="bg-green-900 py-10">
    <div class="container">
        <p
            class="text-green-400 text-xs font-semibold tracking-widest uppercase mb-1"
        >
            {m.checkout_almost()}
        </p>
        <h1 class="text-3xl font-bold text-white">{m.checkout_title()}</h1>
    </div>
</div>

<div class="container py-10">
    {#if $cart.length === 0}
        <div class="py-24 text-center">
            <p class="text-stone-500 mb-4">{m.checkout_empty()}</p>
            <a
                href="/shop"
                class="inline-block px-6 py-3 rounded-xl bg-green-700 text-white font-semibold hover:bg-green-600 transition-colors"
                >{m.checkout_browse()}</a
            >
        </div>
    {:else}
        <form method="POST" use:enhance={handleSubmit}>
            <!-- Hidden cart snapshot -->
            <input
                type="hidden"
                name="cartItems"
                value={JSON.stringify(
                    $cart.map((i) => ({
                        plantId: i.plantId,
                        quantity: i.quantity,
                        price: i.price,
                    })),
                )}
            />

            <div class="grid grid-cols-1 lg:grid-cols-3 gap-8">
                <!-- Left: Shipping + Payment -->
                <div class="lg:col-span-2 space-y-6">
                    <!-- Shipping -->
                    <div
                        class="bg-white rounded-2xl border border-stone-100 shadow-sm p-6"
                    >
                        <h2
                            class="text-base font-bold text-stone-900 mb-5 flex items-center gap-2"
                        >
                            <span
                                class="w-6 h-6 rounded-full bg-green-700 text-white text-xs font-bold flex items-center justify-center"
                                >1</span
                            >
                            {m.checkout_shipping()}
                        </h2>
                        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                            <div class="sm:col-span-2">
                                <label
                                    for="name"
                                    class="block text-sm font-medium text-stone-700 mb-1"
                                    >{m.checkout_name()} *</label
                                >
                                <input
                                    id="name"
                                    name="name"
                                    type="text"
                                    value={form?.values?.name ?? ""}
                                    placeholder="Sofia Kovalenko"
                                    class="w-full px-4 py-2.5 rounded-xl border text-sm outline-none
										{form?.errors?.name
                                        ? 'border-red-400 bg-red-50'
                                        : 'border-stone-200 focus:border-green-500'}"
                                />
                                {#if form?.errors?.name}<p
                                        class="mt-1 text-xs text-red-500"
                                    >
                                        {form.errors.name[0]}
                                    </p>{/if}
                            </div>
                            <div class="sm:col-span-2">
                                <label
                                    for="address"
                                    class="block text-sm font-medium text-stone-700 mb-1"
                                    >{m.checkout_street()} *</label
                                >
                                <input
                                    id="address"
                                    name="address"
                                    type="text"
                                    value={form?.values?.address ?? ""}
                                    placeholder="Khreshchatyk St, 1"
                                    class="w-full px-4 py-2.5 rounded-xl border text-sm outline-none
										{form?.errors?.address
                                        ? 'border-red-400 bg-red-50'
                                        : 'border-stone-200 focus:border-green-500'}"
                                />
                                {#if form?.errors?.address}<p
                                        class="mt-1 text-xs text-red-500"
                                    >
                                        {form.errors.address[0]}
                                    </p>{/if}
                            </div>
                            <div>
                                <label
                                    for="city"
                                    class="block text-sm font-medium text-stone-700 mb-1"
                                    >{m.checkout_city()} *</label
                                >
                                <input
                                    id="city"
                                    name="city"
                                    type="text"
                                    value={form?.values?.city ?? ""}
                                    placeholder="Kyiv"
                                    class="w-full px-4 py-2.5 rounded-xl border text-sm outline-none
										{form?.errors?.city
                                        ? 'border-red-400 bg-red-50'
                                        : 'border-stone-200 focus:border-green-500'}"
                                />
                            </div>
                            <div>
                                <label
                                    for="postalCode"
                                    class="block text-sm font-medium text-stone-700 mb-1"
                                    >{m.checkout_postal()} *</label
                                >
                                <input
                                    id="postalCode"
                                    name="postalCode"
                                    type="text"
                                    value={form?.values?.postalCode ?? ""}
                                    placeholder="01001"
                                    class="w-full px-4 py-2.5 rounded-xl border text-sm outline-none
										{form?.errors?.postalCode
                                        ? 'border-red-400 bg-red-50'
                                        : 'border-stone-200 focus:border-green-500'}"
                                />
                            </div>
                            <div class="sm:col-span-2">
                                <label
                                    for="country"
                                    class="block text-sm font-medium text-stone-700 mb-1"
                                    >{m.checkout_country()} *</label
                                >
                                <input
                                    id="country"
                                    name="country"
                                    type="text"
                                    value={form?.values?.country ?? ""}
                                    placeholder="Ukraine"
                                    class="w-full px-4 py-2.5 rounded-xl border text-sm outline-none
										{form?.errors?.country
                                        ? 'border-red-400 bg-red-50'
                                        : 'border-stone-200 focus:border-green-500'}"
                                />
                            </div>
                        </div>
                    </div>

                    <!-- Payment (mock) -->
                    <div
                        class="bg-white rounded-2xl border border-stone-100 shadow-sm p-6"
                    >
                        <h2
                            class="text-base font-bold text-stone-900 mb-1 flex items-center gap-2"
                        >
                            <span
                                class="w-6 h-6 rounded-full bg-green-700 text-white text-xs font-bold flex items-center justify-center"
                                >2</span
                            >
                            {m.checkout_payment()}
                        </h2>
                        <p class="text-xs text-stone-400 mb-5 ml-8">
                            {m.checkout_demo_mode()}
                        </p>

                        <div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
                            <div class="sm:col-span-2">
                                <label
                                    for="card"
                                    class="block text-sm font-medium text-stone-700 mb-1"
                                    >{m.checkout_card_number()}</label
                                >
                                <input
                                    id="card"
                                    type="text"
                                    maxlength="19"
                                    placeholder="4242 4242 4242 4242"
                                    class="w-full px-4 py-2.5 rounded-xl border border-stone-200 text-sm outline-none focus:border-green-500 font-mono"
                                    oninput={(e) => {
                                        const t = e.target as HTMLInputElement;
                                        t.value = t.value
                                            .replace(/\D/g, "")
                                            .replace(/(.{4})/g, "$1 ")
                                            .trim()
                                            .slice(0, 19);
                                    }}
                                />
                            </div>
                            <div>
                                <label
                                    for="expiry"
                                    class="block text-sm font-medium text-stone-700 mb-1"
                                    >{m.checkout_expiry()}</label
                                >
                                <input
                                    id="expiry"
                                    type="text"
                                    maxlength="5"
                                    placeholder="MM/YY"
                                    class="w-full px-4 py-2.5 rounded-xl border border-stone-200 text-sm outline-none focus:border-green-500 font-mono"
                                />
                            </div>
                            <div>
                                <label
                                    for="cvc"
                                    class="block text-sm font-medium text-stone-700 mb-1"
                                    >{m.checkout_cvc()}</label
                                >
                                <input
                                    id="cvc"
                                    type="text"
                                    maxlength="3"
                                    placeholder="123"
                                    class="w-full px-4 py-2.5 rounded-xl border border-stone-200 text-sm outline-none focus:border-green-500 font-mono"
                                />
                            </div>
                        </div>

                        <div
                            class="mt-4 p-3 rounded-xl bg-amber-50 border border-amber-100 text-xs text-amber-700 flex items-start gap-2"
                        >
                            <span>⚠️</span>
                            <span>{m.checkout_card_mock()}</span>
                        </div>
                    </div>
                </div>

                <!-- Right: Order summary -->
                <div class="lg:col-span-1">
                    <div
                        class="bg-white rounded-2xl border border-stone-100 shadow-sm p-6 sticky top-24"
                    >
                        <h2 class="text-base font-bold text-stone-900 mb-4">
                            {m.checkout_your_order()}
                        </h2>

                        <div class="space-y-3 mb-5">
                            {#each $cart as item}
                                <div class="flex items-center gap-3">
                                    {#if item.image}
                                        <img
                                            src={getImageUrl(item.image)}
                                            alt={item.name}
                                            class="w-10 h-10 rounded-lg object-cover bg-green-100 shrink-0"
                                        />
                                    {:else}
                                        <div
                                            class="w-10 h-10 rounded-lg bg-green-100 flex items-center justify-center shrink-0 text-sm"
                                        >
                                            🌿
                                        </div>
                                    {/if}
                                    <div class="flex-1 min-w-0">
                                        <p
                                            class="text-xs font-medium text-stone-800 truncate"
                                        >
                                            {item.name}
                                        </p>
                                        <p class="text-xs text-stone-400">
                                            × {item.quantity}
                                        </p>
                                    </div>
                                    <p
                                        class="text-xs font-semibold text-stone-800 shrink-0"
                                    >
                                        {formatPrice(
                                            item.price * item.quantity,
                                        )}
                                    </p>
                                </div>
                            {/each}
                        </div>

                        <div
                            class="border-t border-stone-100 pt-4 mb-6 space-y-1.5"
                        >
                            <div
                                class="flex justify-between text-sm text-stone-600"
                            >
                                <span>{m.checkout_subtotal()}</span>
                                <span>{formatPrice($cartTotal)}</span>
                            </div>
                            <div
                                class="flex justify-between text-sm text-stone-400"
                            >
                                <span>{m.checkout_ship_cost()}</span>
                                <span>{m.checkout_free()}</span>
                            </div>
                            <div
                                class="flex justify-between font-bold text-stone-900 text-base pt-1"
                            >
                                <span>{m.checkout_total()}</span>
                                <span class="text-green-700"
                                    >{formatPrice($cartTotal)}</span
                                >
                            </div>
                        </div>

                        {#if form?.error}
                            <p class="text-xs text-red-500 mb-3">
                                {form.error}
                            </p>
                        {/if}

                        <button
                            type="submit"
                            class="w-full py-3.5 rounded-xl bg-green-700 text-white font-semibold text-sm
								hover:bg-green-600 active:scale-[.98] transition-all"
                        >
                            {m.checkout_place()}
                        </button>

                        <a
                            href="/cart"
                            class="block w-full mt-3 py-2 text-xs text-center text-stone-400 hover:text-stone-600 transition-colors"
                        >
                            {m.checkout_back_cart()}
                        </a>
                    </div>
                </div>
            </div>
        </form>
    {/if}
</div>
