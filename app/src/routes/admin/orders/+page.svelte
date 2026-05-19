<script lang="ts">
	import * as m from '$lib/paraglide/messages';
	import { enhance }   from '$app/forms';
	import { formatPrice } from '$lib/modules/common/utils/format';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	const STATUS_LABELS: Record<string, string> = {
		pending:   '⏳ Pending',
		paid:      '💳 Paid',
		shipped:   '🚚 Shipped',
		delivered: '✅ Delivered'
	};

	const STATUS_CLASSES: Record<string, string> = {
		pending:   'bg-yellow-100 text-yellow-700',
		paid:      'bg-blue-100 text-blue-700',
		shipped:   'bg-purple-100 text-purple-700',
		delivered: 'bg-green-100 text-green-700'
	};

	const STATUSES = ['pending', 'paid', 'shipped', 'delivered'] as const;
</script>

<div class="p-8">
	<div class="flex items-center justify-between mb-6">
		<div>
			<h1 class="text-2xl font-bold text-stone-900">Orders</h1>
			<p class="text-stone-500 text-sm">{data.orders.length} total</p>
		</div>
	</div>

	<div class="bg-white rounded-2xl border border-stone-100 shadow-sm overflow-hidden overflow-x-auto">
		{#if data.orders.length === 0}
			<p class="p-8 text-center text-stone-400">No orders yet.</p>
		{:else}
			<table class="w-full text-sm">
				<thead class="bg-stone-50 border-b border-stone-100">
					<tr>
						{#each ['#', 'Customer', 'Total', 'Status', 'Date', 'Change Status'] as h}
							<th class="px-5 py-3 text-left text-xs font-semibold text-stone-500 uppercase tracking-wide">{h}</th>
						{/each}
					</tr>
				</thead>
				<tbody class="divide-y divide-stone-100">
					{#each data.orders as order}
						<tr class="hover:bg-stone-50/50">
							<td class="px-5 py-3.5 text-stone-400 font-mono text-xs">#{order.id}</td>
							<td class="px-5 py-3.5">
								<div class="font-medium text-stone-900">{order.userName ?? '—'}</div>
								<div class="text-xs text-stone-400">{order.userEmail ?? ''}</div>
							</td>
							<td class="px-5 py-3.5 font-semibold text-stone-800">
								{formatPrice(order.totalPrice)}
							</td>
							<td class="px-5 py-3.5">
								<span class="px-2.5 py-1 rounded-full text-xs font-semibold {STATUS_CLASSES[order.status ?? 'pending']}">
									{STATUS_LABELS[order.status ?? 'pending']}
								</span>
							</td>
							<td class="px-5 py-3.5 text-stone-400 text-xs">
								{order.createdAt ? new Date(order.createdAt).toLocaleDateString('en-GB') : '—'}
							</td>
							<td class="px-5 py-3.5">
								<form method="POST" action="?/setStatus" use:enhance class="flex items-center gap-2">
									<input type="hidden" name="id" value={order.id} />
									<select name="status"
										class="px-3 py-1.5 rounded-lg border border-stone-200 text-xs text-stone-700 outline-none focus:border-green-500 bg-white">
										{#each STATUSES as s}
											<option value={s} selected={s === order.status}>{STATUS_LABELS[s]}</option>
										{/each}
									</select>
									<button type="submit"
										class="px-3 py-1.5 rounded-lg bg-green-700 text-white text-xs font-medium hover:bg-green-600 transition-colors">
										Update
									</button>
								</form>
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		{/if}
	</div>
</div>
