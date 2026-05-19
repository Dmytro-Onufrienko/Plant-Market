<script lang="ts">
	import * as m from '$lib/paraglide/messages';
	import { enhance } from '$app/forms';
	import { formatPrice } from '$lib/modules/common/utils/format';
	import type { PageData } from './$types';
	let { data }: { data: PageData } = $props();
</script>

<div class="p-8">
	<div class="flex items-center justify-between mb-6">
		<div>
			<h1 class="text-2xl font-bold text-stone-900">Plants</h1>
			<p class="text-stone-500 text-sm">{data.plants.length} total</p>
		</div>
		<a href="/admin/plants/new"
			class="px-4 py-2.5 rounded-xl bg-green-700 text-white text-sm font-semibold hover:bg-green-600 transition-colors">
			+ Add Plant
		</a>
	</div>

	<div class="bg-white rounded-2xl border border-stone-100 shadow-sm overflow-hidden overflow-x-auto">
		{#if data.plants.length === 0}
			<p class="p-8 text-center text-stone-400">No plants yet. <a href="/admin/plants/new" class="text-green-700 font-medium hover:underline">Add the first one →</a></p>
		{:else}
			<table class="w-full text-sm">
				<thead class="bg-stone-50 border-b border-stone-100">
					<tr>
						{#each ['Name', 'Category', 'Price', 'Stock', 'Featured', 'Actions'] as h}
							<th class="px-5 py-3 text-left text-xs font-semibold text-stone-500 uppercase tracking-wide">{h}</th>
						{/each}
					</tr>
				</thead>
				<tbody class="divide-y divide-stone-100">
					{#each data.plants as plant}
						<tr class="hover:bg-stone-50/50">
							<td class="px-5 py-3.5">
								<div class="font-medium text-stone-900">{plant.name}</div>
								<div class="text-xs text-stone-400 font-mono">{plant.slug}</div>
							</td>
							<td class="px-5 py-3.5 text-stone-500">{plant.category ?? '—'}</td>
							<td class="px-5 py-3.5 font-semibold text-green-700">{formatPrice(plant.price)}</td>
							<td class="px-5 py-3.5">
								<span class="px-2 py-0.5 rounded-full text-xs font-medium
									{plant.stock > 0 ? 'bg-green-100 text-green-700' : 'bg-red-100 text-red-600'}">
									{plant.stock}
								</span>
							</td>
							<td class="px-5 py-3.5">
								{#if plant.featured}
									<span class="text-yellow-500">★</span>
								{:else}
									<span class="text-stone-300">☆</span>
								{/if}
							</td>
							<td class="px-5 py-3.5 flex items-center gap-2">
								<a href="/admin/plants/{plant.id}"
									class="px-3 py-1.5 rounded-lg border border-stone-200 text-xs font-medium text-stone-600 hover:border-green-400 hover:text-green-700 transition-colors">
									Edit
								</a>
								<form method="POST" action="?/delete" use:enhance
									onsubmit={(e) => { if (!confirm(`Delete "${plant.name}"?`)) e.preventDefault(); }}>
									<input type="hidden" name="id" value={plant.id} />
									<button class="px-3 py-1.5 rounded-lg border border-stone-200 text-xs font-medium text-red-500 hover:border-red-300 hover:bg-red-50 transition-colors">
										Delete
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
