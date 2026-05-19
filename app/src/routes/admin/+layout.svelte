<script lang="ts">
	import { page } from '$app/stores';
	import { enhance } from '$app/forms';
	import type { LayoutData } from './$types';

	let { data, children }: { data: LayoutData; children: import('svelte').Snippet } = $props();

	const navItems = [
		{ href: '/admin',         icon: '📊', label: 'Dashboard'  },
		{ href: '/admin/plants',  icon: '🌿', label: 'Plants'     },
		{ href: '/admin/blog',    icon: '✍️',  label: 'Blog'       },
		{ href: '/admin/orders',  icon: '📦', label: 'Orders'     },
		{ href: '/admin/reviews', icon: '⭐', label: 'Reviews'    },
	];
</script>

<div class="min-h-screen bg-stone-50 flex">
	<!-- Sidebar -->
	<aside class="w-60 shrink-0 bg-green-900 min-h-screen flex flex-col">
		<!-- Logo -->
		<div class="px-5 py-5 border-b border-green-800">
			<a href="/" class="flex items-center gap-2 group">
				<span class="w-7 h-7 rounded-full bg-green-300 flex items-center justify-center text-green-900 font-bold text-xs">M</span>
				<div>
					<p class="text-white font-bold text-sm leading-none">Monstera Shop</p>
					<p class="text-green-400 text-xs mt-0.5">Admin Panel</p>
				</div>
			</a>
		</div>

		<!-- Nav -->
		<nav class="flex-1 px-3 py-4 space-y-0.5">
			{#each navItems as item}
				<a
					href={item.href}
					class="flex items-center gap-3 px-3 py-2.5 rounded-xl text-sm font-medium transition-colors
						{$page.url.pathname === item.href || ($page.url.pathname.startsWith(item.href + '/') && item.href !== '/admin')
							? 'bg-green-700 text-white'
							: 'text-green-100/70 hover:bg-green-800 hover:text-white'}"
				>
					<span class="text-base">{item.icon}</span>
					{item.label}
				</a>
			{/each}
		</nav>

		<!-- User + logout -->
		<div class="px-4 py-4 border-t border-green-800">
			<p class="text-xs text-green-400 mb-2 truncate">{data.user.email}</p>
			<form method="POST" action="/auth/logout" use:enhance>
				<button class="w-full text-left flex items-center gap-2 px-3 py-2 rounded-xl text-sm
					text-green-200/60 hover:bg-green-800 hover:text-red-300 transition-colors">
					← Sign out
				</button>
			</form>
			<a href="/" class="flex items-center gap-2 px-3 py-2 rounded-xl text-sm text-green-200/60 hover:bg-green-800 hover:text-white transition-colors">
				↗ View site
			</a>
		</div>
	</aside>

	<!-- Main content -->
	<div class="flex-1 overflow-auto">
		{@render children()}
	</div>
</div>
