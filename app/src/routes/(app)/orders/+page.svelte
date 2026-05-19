<script lang="ts">
	import { formatPrice } from '$lib/modules/common/utils/format';
	import * as m from '$lib/paraglide/messages';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	const STATUS_CLASS: Record<string, string> = {
		pending:   'bg-yellow-100 text-yellow-700',
		paid:      'bg-blue-100 text-blue-700',
		shipped:   'bg-purple-100 text-purple-700',
		delivered: 'bg-green-100 text-green-700'
	};

	function statusLabel(status: string): string {
		const map: Record<string, () => string> = {
			pending:   m.status_pending,
			paid:      m.status_paid,
			shipped:   m.status_shipped,
			delivered: m.status_delivered
		};
		return map[status]?.() ?? status;
	}
</script>

<svelte:head>
	<title>{m.orders_title()} — Monstera Shop</title>
</svelte:head>

<div class="bg-green-900 py-10">
	<div class="container">
		<p class="text-green-400 text-xs font-semibold tracking-widest uppercase mb-1">{m.orders_your()}</p>
		<h1 class="text-3xl font-bold text-white">{m.orders_title()}</h1>
	</div>
</div>

<div class="container py-10 max-w-3xl">
	{#if data.orders.length === 0}
		<div class="py-24 text-center">
			<p class="text-4xl mb-4">📦</p>
			<p class="text-stone-500 mb-2">{m.orders_empty()}</p>
			<a href="/shop" class="inline-block mt-2 px-6 py-3 rounded-xl bg-green-700 text-white font-semibold hover:bg-green-600 transition-colors">
				{m.orders_start()}
			</a>
		</div>
	{:else}
		<div class="space-y-3">
			{#each data.orders as order}
				<a href="/orders/{order.id}"
					class="bg-white rounded-2xl border border-stone-100 shadow-sm p-5
						flex items-center justify-between gap-4 hover:shadow-md transition-shadow group">
					<div class="flex items-center gap-4">
						<div class="w-10 h-10 rounded-xl bg-green-50 flex items-center justify-center text-lg">📦</div>
						<div>
							<p class="font-semibold text-stone-900 text-sm group-hover:text-green-700 transition-colors">
								{m.orders_order()} #{order.id}
							</p>
							<p class="text-xs text-stone-400 mt-0.5">
								{order.createdAt ? new Date(order.createdAt).toLocaleDateString('en-GB', { day: 'numeric', month: 'long', year: 'numeric' }) : '—'}
							</p>
						</div>
					</div>
					<div class="flex items-center gap-4">
						<span class="px-2.5 py-1 rounded-full text-xs font-semibold {STATUS_CLASS[order.status ?? 'pending']}">
							{statusLabel(order.status ?? 'pending')}
						</span>
						<span class="font-bold text-stone-800">{formatPrice(order.totalPrice)}</span>
						<span class="text-stone-300 group-hover:text-green-500 transition-colors">→</span>
					</div>
				</a>
			{/each}
		</div>
	{/if}
</div>
