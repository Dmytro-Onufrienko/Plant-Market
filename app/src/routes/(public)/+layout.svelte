<script lang="ts">
	import Navbar from '$lib/modules/layout/components/Navbar.svelte';
	import Footer from '$lib/modules/layout/components/Footer.svelte';
	import type { LayoutData } from './$types';

	let { data, children }: { data: LayoutData; children: import('svelte').Snippet } = $props();
</script>

<div class="flex flex-col min-h-screen">
	<Navbar user={data.user} />
	<main class="flex-1">
		{@render children()}
	</main>
	<Footer />
</div>
