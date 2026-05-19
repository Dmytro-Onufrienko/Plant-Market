<script lang="ts">
	import * as m from '$lib/paraglide/messages';
	import { formatPrice } from '$lib/modules/common/utils/format';
	import type { PageData } from './$types';
	let { data }: { data: PageData } = $props();

	const statusColor: Record<string, string> = {
		pending:   'bg-yellow-100 text-yellow-700',
		paid:      'bg-blue-100 text-blue-700',
		shipped:   'bg-purple-100 text-purple-700',
		delivered: 'bg-green-100 text-green-700'
	};
</script>

<div class="p-8">
	<h1 class="text-2xl font-bold text-stone-900 mb-1">Dashboard</h1>
	<p class="text-stone-500 text-sm mb-8">Welcome back. Here's what's happening.</p>

	<!-- Stats grid -->
	<div class="grid grid-cols-2 lg:grid-cols-4 gap-4 mb-10">
		{#each [
			{ label: 'Plants',   value: data.stats.plants,  icon: '🌿', href: '/admin/plants'  },
			{ label: 'Orders',   value: data.stats.orders,  icon: '📦', href: '/admin/orders'  },
			{ label: 'Reviews',  value: data.stats.reviews, icon: '⭐', href: '/admin/reviews' },
			{ label: 'Users',    value: data.stats.users,   icon: '👤', href: '#'              },
		] as stat}
			<a href={stat.href}
				class="bg-white rounded-2xl p-5 border border-stone-100 shadow-sm hover:shadow-md transition-all hover:-translate-y-0.5 group">
				<div class="flex items-start justify-between mb-3">
					<span class="text-2xl">{stat.icon}</span>
					<span class="text-xs text-green-600 font-medium group-hover:underline">View →</span>
				</div>
				<p class="text-3xl font-bold text-stone-900">{stat.value}</p>
				<p class="text-sm text-stone-400 mt-0.5">{stat.label}</p>
			</a>
		{/each}
	</div>

	<!-- Quick actions -->
	<div class="flex flex-wrap gap-3 mb-10">
		<a href="/admin/plants/new"
			class="flex items-center gap-2 px-4 py-2.5 rounded-xl bg-green-700 text-white text-sm font-semibold hover:bg-green-600 transition-colors">
			+ Add Plant
		</a>
		<a href="/admin/blog/new"
			class="flex items-center gap-2 px-4 py-2.5 rounded-xl bg-bark-700 text-white text-sm font-semibold hover:bg-bark-600 transition-colors">
			+ New Post
		</a>
	</div>

	<!-- Recent orders -->
	<div class="bg-white rounded-2xl border border-stone-100 shadow-sm overflow-hidden">
		<div class="px-6 py-4 border-b border-stone-100 flex items-center justify-between">
			<h2 class="font-semibold text-stone-900">Recent Orders</h2>
			<a href="/admin/orders" class="text-xs text-green-700 font-medium hover:underline">See all →</a>
		</div>

		{#if data.recentOrders.length === 0}
			<p class="p-6 text-sm text-stone-400">No orders yet.</p>
		{:else}
			<table class="w-full text-sm">
				<thead class="bg-stone-50">
					<tr>
						{#each ['Order', 'Customer', 'Amount', 'Status', 'Date'] as h}
							<th class="px-6 py-3 text-left text-xs font-semibold text-stone-500 uppercase tracking-wide">{h}</th>
						{/each}
					</tr>
				</thead>
				<tbody class="divide-y divide-stone-100">
					{#each data.recentOrders as order}
						<tr class="hover:bg-stone-50/50">
							<td class="px-6 py-3.5 font-mono text-stone-600">#{order.id}</td>
							<td class="px-6 py-3.5 text-stone-700">{order.userName}</td>
							<td class="px-6 py-3.5 font-semibold text-stone-900">{formatPrice(order.totalPrice)}</td>
							<td class="px-6 py-3.5">
								<span class="px-2.5 py-1 rounded-full text-xs font-semibold capitalize {statusColor[order.status] ?? 'bg-stone-100 text-stone-600'}">
									{order.status}
								</span>
							</td>
							<td class="px-6 py-3.5 text-stone-400 text-xs">
								{order.createdAt ? new Date(order.createdAt).toLocaleDateString('en-GB') : '—'}
							</td>
						</tr>
					{/each}
				</tbody>
			</table>
		{/if}
	</div>
</div>
