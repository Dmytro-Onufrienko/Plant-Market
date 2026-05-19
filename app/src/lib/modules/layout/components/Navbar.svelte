<script lang="ts">
    import { page } from "$app/stores";
    import { enhance } from "$app/forms";
    import { cartCount } from "$lib/modules/orders/stores/cart";
    import { slide } from "svelte/transition";
    import { i18n } from "$lib/i18n";
    import * as runtime from "$lib/paraglide/runtime";
    import * as m from "$lib/paraglide/messages";

    const availableLangs = runtime.availableLanguageTags;
    const langLabels: Record<string, string> = { uk: "UA", en: "EN" };

    let {
        user,
    }: { user: { name: string; email: string; role: string } | null } =
        $props();

    let mobileOpen = $state(false);

    function switchLanguage(lang: (typeof availableLangs)[number]) {
        runtime.setLanguageTag(lang);
        if (typeof localStorage !== "undefined") {
            localStorage.setItem("lang", lang);
        }
        // Refresh the page to apply language change
        window.location.href = i18n.route($page.url.pathname);
    }

    const links = [
        { href: "/shop", label: m.nav_shop },
        { href: "/blog", label: m.nav_blog },
    ];
</script>

<header class="sticky top-0 z-50 bg-green-900 shadow-lg">
    <nav class="container flex items-center justify-between h-16">
        <!-- Logo -->
        <a href="/" class="flex items-center gap-2.5 group shrink-0">
            <span
                class="w-8 h-8 rounded-full bg-green-300 flex items-center justify-center
				text-green-900 font-bold text-sm transition-transform group-hover:scale-110"
            >
                M
            </span>
            <span class="text-white font-bold text-lg leading-none"
                >Monstera Shop</span
            >
        </a>

        <!-- Desktop nav -->
        <div class="hidden md:flex items-center gap-8">
            {#each links as link}
                <a
                    href={link.href}
                    class="text-sm font-medium transition-colors
						{$page.url.pathname.startsWith(link.href)
                        ? 'text-green-300'
                        : 'text-green-100/80 hover:text-white'}"
                >
                    {link.label()}
                </a>
            {/each}
        </div>

        <!-- Language switcher (desktop) -->
        <div class="hidden md:flex items-center gap-1 mr-1">
            {#each availableLangs as lang}
                <button
                    onclick={() => switchLanguage(lang)}
                    class="px-2 py-1 rounded text-xs font-semibold transition-colors
						{runtime.languageTag() === lang
                        ? 'bg-green-600 text-white'
                        : 'text-green-300 hover:text-white'}"
                >
                    {langLabels[lang] ?? lang.toUpperCase()}
                </button>
            {/each}
        </div>

        <!-- Desktop actions -->
        <div class="hidden md:flex items-center gap-3">
            {#if user}
                {#if user.role !== "admin"}
                    <a
                        href="/cart"
                        class="relative flex items-center gap-1.5 px-4 py-2 rounded-full
						bg-green-700 text-white text-sm font-semibold hover:bg-green-600 transition-colors"
                    >
                        <svg
                            xmlns="http://www.w3.org/2000/svg"
                            class="w-4 h-4"
                            fill="none"
                            viewBox="0 0 24 24"
                            stroke="currentColor"
                            stroke-width="2"
                        >
                            <path
                                stroke-linecap="round"
                                stroke-linejoin="round"
                                d="M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z"
                            />
                        </svg>
                        {m.nav_cart()}
                        {#if $cartCount > 0}
                            <span
                                class="absolute -top-1.5 -right-1.5 w-4 h-4 rounded-full bg-bark-500 text-white text-[10px] font-bold flex items-center justify-center"
                            >
                                {$cartCount > 9 ? "9+" : $cartCount}
                            </span>
                        {/if}
                    </a>
                {/if}

                <!-- User menu -->
                <div class="relative group">
                    <button
                        class="flex items-center gap-2 px-3 py-2 rounded-full
						border border-green-600 text-green-100 text-sm hover:border-green-400 transition-colors"
                    >
                        <span
                            class="w-6 h-6 rounded-full bg-green-600 flex items-center justify-center text-xs font-bold"
                        >
                            {user.name[0].toUpperCase()}
                        </span>
                        {user.name.split(" ")[0]}
                        {#if user.role === "admin"}
                            <span
                                class="text-xs px-1.5 py-0.5 rounded bg-bark-700 text-bark-200 font-medium"
                                >Admin</span
                            >
                        {/if}
                    </button>

                    <!-- Dropdown -->
                    <div
                        class="absolute right-0 top-full mt-1 w-48 bg-white rounded-xl shadow-xl border border-stone-100
						opacity-0 invisible group-hover:opacity-100 group-hover:visible transition-all translate-y-1 group-hover:translate-y-0"
                    >
                        <div class="p-1">
                            {#if user.role !== "admin"}
                                <a
                                    href="/orders"
                                    class="flex items-center gap-2 px-3 py-2 rounded-lg text-sm text-stone-700 hover:bg-stone-50"
                                >
                                    📦 {m.nav_orders()}
                                </a>
                            {/if}
                            <a
                                href="/settings"
                                class="flex items-center gap-2 px-3 py-2 rounded-lg text-sm text-stone-700 hover:bg-stone-50"
                            >
                                ⚙️ {m.nav_settings()}
                            </a>
                            {#if user.role === "admin"}
                                <hr class="my-1 border-stone-100" />
                                <a
                                    href="/admin"
                                    class="flex items-center gap-2 px-3 py-2 rounded-lg text-sm text-green-700 font-medium hover:bg-green-50"
                                >
                                    🛠 {m.nav_admin()}
                                </a>
                            {/if}
                            <hr class="my-1 border-stone-100" />
                            <form
                                method="POST"
                                action="/auth/logout"
                                use:enhance
                            >
                                <button
                                    class="w-full flex items-center gap-2 px-3 py-2 rounded-lg text-sm text-red-500 hover:bg-red-50 text-left"
                                >
                                    ← {m.nav_signout()}
                                </button>
                            </form>
                        </div>
                    </div>
                </div>
            {:else}
                <a
                    href="/auth/login"
                    class="px-4 py-2 rounded-full border border-green-500 text-green-100
						text-sm font-medium hover:border-green-300 hover:text-white transition-colors"
                >
                    {m.nav_signin()}
                </a>
                <a
                    href="/auth/register"
                    class="px-4 py-2 rounded-full bg-green-300 text-green-900
						text-sm font-semibold hover:bg-green-200 transition-colors"
                >
                    {m.nav_getstart()}
                </a>
            {/if}
        </div>

        <!-- Mobile burger -->
        <button
            class="md:hidden p-2 text-green-100 hover:text-white"
            onclick={() => (mobileOpen = !mobileOpen)}
            aria-label="Toggle menu"
        >
            {#if mobileOpen}
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="w-6 h-6"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                    stroke-width="2"
                >
                    <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        d="M6 18L18 6M6 6l12 12"
                    />
                </svg>
            {:else}
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="w-6 h-6"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke="currentColor"
                    stroke-width="2"
                >
                    <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        d="M4 6h16M4 12h16M4 18h16"
                    />
                </svg>
            {/if}
        </button>
    </nav>

    <!-- Mobile menu -->
    {#if mobileOpen}
        <div
            transition:slide={{ duration: 200 }}
            class="md:hidden bg-green-950 border-t border-green-800 px-4 pb-4 space-y-1"
        >
            {#each links as link}
                <a
                    href={link.href}
                    class="block px-3 py-2.5 rounded-lg text-sm font-medium
						{$page.url.pathname.startsWith(link.href)
                        ? 'bg-green-800 text-white'
                        : 'text-green-100 hover:bg-green-800'}"
                    onclick={() => (mobileOpen = false)}
                >
                    {link.label()}
                </a>
            {/each}

            {#if user}
                {#if user.role !== "admin"}
                    <a
                        href="/cart"
                        class="block px-3 py-2.5 rounded-lg text-sm text-green-100 hover:bg-green-800"
                        >🛒 {m.nav_cart()}</a
                    >
                    <a
                        href="/orders"
                        class="block px-3 py-2.5 rounded-lg text-sm text-green-100 hover:bg-green-800"
                        >📦 {m.nav_orders()}</a
                    >
                {/if}
                <a
                    href="/settings"
                    class="block px-3 py-2.5 rounded-lg text-sm text-green-100 hover:bg-green-800"
                    >⚙️ {m.nav_settings()}</a
                >
                {#if user.role === "admin"}
                    <a
                        href="/admin"
                        class="block px-3 py-2.5 rounded-lg text-sm text-green-300 font-medium hover:bg-green-800"
                        >🛠 {m.nav_admin()}</a
                    >
                {/if}
                <form method="POST" action="/auth/logout" use:enhance>
                    <button
                        class="w-full text-left px-3 py-2.5 rounded-lg text-sm text-red-400 hover:bg-green-800"
                        >← {m.nav_signout()}</button
                    >
                </form>
            {:else}
                <a
                    href="/auth/login"
                    class="block px-3 py-2.5 rounded-lg text-sm text-green-100 hover:bg-green-800"
                    >{m.nav_signin()}</a
                >
                <a
                    href="/auth/register"
                    class="block px-3 py-2.5 rounded-lg text-sm text-green-300 font-medium hover:bg-green-800"
                    >{m.nav_getstart()}</a
                >
            {/if}
        </div>
    {/if}
</header>
