<script lang="ts">
    import Navbar from "$lib/modules/layout/components/Navbar.svelte";
    import Footer from "$lib/modules/layout/components/Footer.svelte";
    import type { LayoutData } from "./$types";
    import { fly } from "svelte/transition";
    import { page } from "$app/stores";

    let {
        data,
        children,
    }: { data: LayoutData; children: import("svelte").Snippet } = $props();
</script>

<div class="flex flex-col min-h-screen">
    <Navbar user={data.user} />
    <main class="flex-1">
        {#key $page.url.pathname}
            <div
                in:fly={{ y: 16, duration: 250, delay: 60 }}
                out:fly={{ y: -8, duration: 150 }}
            >
                {@render children()}
            </div>
        {/key}
    </main>
    <Footer />
</div>
