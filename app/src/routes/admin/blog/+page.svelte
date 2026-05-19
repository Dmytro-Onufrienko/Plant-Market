<script lang="ts">
	import * as m from '$lib/paraglide/messages';
	import { enhance } from '$app/forms';
	import type { PageData } from './$types';
	let { data }: { data: PageData } = $props();
</script>

<div class="p-8">
	<div class="flex items-center justify-between mb-6">
		<div>
			<h1 class="text-2xl font-bold text-stone-900">Blog Posts</h1>
			<p class="text-stone-500 text-sm">{data.posts.length} total</p>
		</div>
		<a href="/admin/blog/new"
			class="px-4 py-2.5 rounded-xl bg-bark-700 text-white text-sm font-semibold hover:bg-bark-600 transition-colors">
			+ New Post
		</a>
	</div>

	<div class="bg-white rounded-2xl border border-stone-100 shadow-sm overflow-hidden overflow-x-auto">
		{#if data.posts.length === 0}
			<p class="p-8 text-center text-stone-400">No posts yet.</p>
		{:else}
			<table class="w-full text-sm">
				<thead class="bg-stone-50 border-b border-stone-100">
					<tr>
						{#each ['Title', 'Status', 'Date', 'Actions'] as h}
							<th class="px-5 py-3 text-left text-xs font-semibold text-stone-500 uppercase tracking-wide">{h}</th>
						{/each}
					</tr>
				</thead>
				<tbody class="divide-y divide-stone-100">
					{#each data.posts as post}
						<tr class="hover:bg-stone-50/50">
							<td class="px-5 py-3.5">
								<div class="font-medium text-stone-900">{post.title}</div>
								<div class="text-xs text-stone-400 font-mono">{post.slug}</div>
							</td>
							<td class="px-5 py-3.5">
								<form method="POST" action="?/togglePublish" use:enhance>
									<input type="hidden" name="id"        value={post.id} />
									<input type="hidden" name="published" value={post.published} />
									<button class="px-2.5 py-1 rounded-full text-xs font-semibold transition-colors
										{post.published
											? 'bg-green-100 text-green-700 hover:bg-green-200'
											: 'bg-stone-100 text-stone-500 hover:bg-stone-200'}">
										{post.published ? '● Published' : '○ Draft'}
									</button>
								</form>
							</td>
							<td class="px-5 py-3.5 text-stone-400 text-xs">
								{post.createdAt ? new Date(post.createdAt).toLocaleDateString('en-GB') : '—'}
							</td>
							<td class="px-5 py-3.5 flex items-center gap-2">
								<a href="/admin/blog/{post.id}"
									class="px-3 py-1.5 rounded-lg border border-stone-200 text-xs font-medium text-stone-600 hover:border-green-400 hover:text-green-700 transition-colors">
									Edit
								</a>
								<form method="POST" action="?/delete" use:enhance
									onsubmit={(e) => { if (!confirm('Delete this post?')) e.preventDefault(); }}>
									<input type="hidden" name="id" value={post.id} />
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
