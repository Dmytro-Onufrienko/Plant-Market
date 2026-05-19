<script lang="ts">
	import * as m from '$lib/paraglide/messages';
	import { enhance } from '$app/forms';
	import type { PageData } from './$types';

	let { data }: { data: PageData } = $props();

	function stars(n: number) {
		return '★'.repeat(n) + '☆'.repeat(5 - n);
	}
</script>

<div class="p-8">
	<div class="flex items-center justify-between mb-6">
		<div>
			<h1 class="text-2xl font-bold text-stone-900">Reviews</h1>
			<p class="text-stone-500 text-sm">{data.reviews.length} total</p>
		</div>
	</div>

	<div class="bg-white rounded-2xl border border-stone-100 shadow-sm overflow-hidden overflow-x-auto">
		{#if data.reviews.length === 0}
			<p class="p-8 text-center text-stone-400">No reviews yet.</p>
		{:else}
			<table class="w-full text-sm">
				<thead class="bg-stone-50 border-b border-stone-100">
					<tr>
						{#each ['Plant', 'Customer', 'Rating', 'Review', 'Date', 'Actions'] as h}
							<th class="px-5 py-3 text-left text-xs font-semibold text-stone-500 uppercase tracking-wide">{h}</th>
						{/each}
					</tr>
				</thead>
				<tbody class="divide-y divide-stone-100">
					{#each data.reviews as review}
						<tr class="hover:bg-stone-50/50 align-top">
							<td class="px-5 py-3.5">
								<a href="/shop/{review.plantSlug}"
									class="font-medium text-green-700 hover:underline text-xs">
									{review.plantName ?? '—'}
								</a>
							</td>
							<td class="px-5 py-3.5">
								<div class="text-xs text-stone-700">{review.userName ?? '—'}</div>
								<div class="text-xs text-stone-400">{review.userEmail ?? ''}</div>
							</td>
							<td class="px-5 py-3.5">
								<span class="text-amber-500 tracking-tighter text-sm">{stars(review.rating)}</span>
								<span class="text-xs text-stone-400 ml-1">{review.rating}/5</span>
							</td>
							<td class="px-5 py-3.5 max-w-xs">
								<p class="text-stone-700 text-xs leading-relaxed line-clamp-3">{review.text}</p>
							</td>
							<td class="px-5 py-3.5 text-stone-400 text-xs whitespace-nowrap">
								{review.createdAt ? new Date(review.createdAt).toLocaleDateString('en-GB') : '—'}
							</td>
							<td class="px-5 py-3.5">
								<form method="POST" action="?/delete" use:enhance
									onsubmit={(e) => { if (!confirm('Delete this review?')) e.preventDefault(); }}>
									<input type="hidden" name="id" value={review.id} />
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
