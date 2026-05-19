<script lang="ts">
    import { formatPrice, parseImages, getImageUrl } from "$lib/modules/common/utils/format";
    import * as m from '$lib/paraglide/messages';
    import type { PageData } from "./$types";

    let { data }: { data: PageData } = $props();
    const { order, items, shippingAddress, justPlaced } = data;

    const STATUS_STEPS = ["pending", "paid", "shipped", "delivered"] as const;

    function statusLabel(step: string): string {
        const map: Record<string, () => string> = {
            pending:   m.status_pending,
            paid:      m.status_paid,
            shipped:   m.status_shipped,
            delivered: m.status_delivered
        };
        return map[step]?.() ?? step;
    }

    const currentStep = $derived(
        STATUS_STEPS.indexOf(order.status as (typeof STATUS_STEPS)[number]),
    );
</script>

<svelte:head>
    <title>{m.orders_order()} #{order.id} — Monstera Shop</title>
</svelte:head>

<div class="bg-green-900 py-10">
    <div class="container">
        <a
            href="/orders"
            class="text-green-400 text-xs hover:text-green-200 transition-colors mb-2 inline-block"
            >{m.orders_all()}</a
        >
        <h1 class="text-3xl font-bold text-white">{m.orders_order()} #{order.id}</h1>
    </div>
</div>

<div class="container py-10 max-w-3xl">
    <!-- Success banner -->
    {#if justPlaced}
        <div
            class="mb-6 p-5 rounded-2xl bg-green-50 border border-green-200 flex items-start gap-3"
        >
            <span class="text-2xl">🎉</span>
            <div>
                <p class="font-semibold text-green-800">
                    {m.orders_placed_ok()}
                </p>
                <p class="text-sm text-green-600 mt-0.5">
                    {m.orders_preparing()}
                </p>
            </div>
        </div>
    {/if}

    <!-- Status tracker -->
    <div
        class="bg-white rounded-2xl border border-stone-100 shadow-sm p-6 mb-5"
    >
        <h2 class="text-sm font-semibold text-stone-700 mb-5">{m.orders_status()}</h2>
        <div class="flex items-center">
            {#each STATUS_STEPS as step, i}
                <div class="flex flex-col items-center flex-1">
                    <div
                        class="w-8 h-8 rounded-full flex items-center justify-center text-sm font-bold mb-1.5
						{i <= currentStep ? 'bg-green-700 text-white' : 'bg-stone-100 text-stone-400'}"
                    >
                        {i < currentStep ? "✓" : i + 1}
                    </div>
                    <span
                        class="text-[10px] font-medium text-center leading-tight
						{i <= currentStep ? 'text-green-700' : 'text-stone-400'}"
                    >
                        {statusLabel(step)}
                    </span>
                </div>
                {#if i < STATUS_STEPS.length - 1}
                    <div
                        class="flex-1 h-0.5 mb-6 {i < currentStep
                            ? 'bg-green-400'
                            : 'bg-stone-100'}"
                    ></div>
                {/if}
            {/each}
        </div>
    </div>

    <div class="grid grid-cols-1 sm:grid-cols-2 gap-5 mb-5">
        <!-- Order items -->
        <div
            class="bg-white rounded-2xl border border-stone-100 shadow-sm p-5 sm:col-span-2"
        >
            <h2 class="text-sm font-semibold text-stone-700 mb-4">{m.orders_items()}</h2>
            <div class="space-y-3">
                {#each items as item}
                    {@const img = parseImages(item.plantImages ?? "[]")[0]}
                    <div class="flex items-center gap-3">
                        {#if img}
                            <img
                                src={getImageUrl(img)}
                                alt={item.plantName}
                                class="w-12 h-12 rounded-xl object-cover bg-green-100 shrink-0"
                            />
                        {:else}
                            <div
                                class="w-12 h-12 rounded-xl bg-green-100 flex items-center justify-center text-lg shrink-0"
                            >
                                🌿
                            </div>
                        {/if}
                        <div class="flex-1 min-w-0">
                            <a
                                href="/shop/{item.plantSlug}"
                                class="text-sm font-medium text-stone-900 hover:text-green-700 transition-colors"
                            >
                                {item.plantName}
                            </a>
                            <p class="text-xs text-stone-400">
                                × {item.quantity} · {formatPrice(
                                    item.priceAtPurchase,
                                )} each
                            </p>
                        </div>
                        <p class="text-sm font-bold text-stone-800">
                            {formatPrice(item.priceAtPurchase * item.quantity)}
                        </p>
                    </div>
                {/each}
            </div>
            <div
                class="border-t border-stone-100 pt-3 mt-4 flex justify-between"
            >
                <span class="font-semibold text-stone-700">{m.checkout_total()}</span>
                <span class="font-bold text-green-700 text-base"
                    >{formatPrice(order.totalPrice)}</span
                >
            </div>
        </div>

        <!-- Shipping address -->
        <div class="bg-white rounded-2xl border border-stone-100 shadow-sm p-5">
            <h2 class="text-sm font-semibold text-stone-700 mb-3">
                {m.orders_shipping_to()}
            </h2>
            {#if shippingAddress.name}
                <div class="text-sm text-stone-600 space-y-0.5">
                    <p class="font-medium text-stone-900">
                        {shippingAddress.name}
                    </p>
                    <p>{shippingAddress.address}</p>
                    <p>{shippingAddress.postalCode} {shippingAddress.city}</p>
                    <p>{shippingAddress.country}</p>
                </div>
            {:else}
                <p class="text-sm text-stone-400">{m.orders_no_address()}</p>
            {/if}
        </div>

        <!-- Order info -->
        <div class="bg-white rounded-2xl border border-stone-100 shadow-sm p-5">
            <h2 class="text-sm font-semibold text-stone-700 mb-3">
                {m.orders_info()}
            </h2>
            <div class="text-sm text-stone-600 space-y-1.5">
                <div class="flex justify-between">
                    <span>{m.orders_id()}</span>
                    <span class="font-mono text-stone-800">#{order.id}</span>
                </div>
                <div class="flex justify-between">
                    <span>{m.orders_date()}</span>
                    <span
                        >{order.createdAt
                            ? new Date(order.createdAt).toLocaleDateString(
                                  "en-GB",
                              )
                            : "—"}</span
                    >
                </div>
                <div class="flex justify-between">
                    <span>{m.orders_payment()}</span>
                    <span class="text-green-600 font-medium">{m.orders_payment_mock()}</span>
                </div>
            </div>
        </div>
    </div>

    <div class="flex gap-3">
        <a
            href="/orders"
            class="px-5 py-2.5 rounded-xl border border-stone-200 text-stone-600 text-sm font-medium hover:border-stone-300 transition-colors"
        >
            {m.orders_all()}
        </a>
        <a
            href="/shop"
            class="px-5 py-2.5 rounded-xl bg-green-700 text-white text-sm font-semibold hover:bg-green-600 transition-colors"
        >
            {m.orders_continue()}
        </a>
    </div>
</div>
