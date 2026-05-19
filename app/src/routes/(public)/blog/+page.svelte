<script lang="ts">
    import { getImageUrl } from "$lib/modules/common/utils/format";
    import type { PageData } from "./$types";
    let { data }: { data: PageData } = $props();

    function formatDate(d: Date | string | null): string {
        if (!d) return "";
        const date = typeof d === "string" ? new Date(d) : d;
        return date.toLocaleDateString("en-GB", {
            day: "numeric",
            month: "long",
            year: "numeric",
        });
    }
</script>

<svelte:head>
    <title>Blog — Monstera Shop</title>
    <meta
        name="description"
        content="Plant care tips, tropical inspiration and green living guides."
    />
</svelte:head>

<!-- Header -->
<div class="bg-green-900 py-10">
    <div class="container">
        <p
            class="text-green-400 text-xs font-semibold tracking-widest uppercase mb-1"
        >
            Plant care & inspiration
        </p>
        <h1 class="text-3xl font-bold text-white">Blog</h1>
    </div>
</div>

<div class="container py-12">
    {#if data.posts.length === 0}
        <div class="py-20 text-center">
            <p class="text-4xl mb-3">✍️</p>
            <p class="text-stone-500">No posts yet. Check back soon!</p>
        </div>
    {:else}
        <!-- Featured first post -->
        {@const [first, ...rest] = data.posts}
        <a
            href="/blog/{first.slug}"
            class="group grid grid-cols-1 lg:grid-cols-2 gap-0 rounded-2xl overflow-hidden bg-white
				shadow-sm hover:shadow-lg transition-all duration-300 mb-10"
        >
            <!-- Cover -->
            <div
                class="aspect-video lg:aspect-auto bg-gradient-to-br from-green-200 to-green-400 relative overflow-hidden"
            >
                {#if first.coverImage}
                    <img
                        src={getImageUrl(first.coverImage)}
                        alt={first.title}
                        class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105"
                    />
                {:else}
                    <div
                        class="w-full h-full flex items-center justify-center opacity-30"
                    >
                        <svg
                            viewBox="0 0 200 200"
                            class="w-32 h-32 text-green-700"
                            fill="currentColor"
                        >
                            <ellipse
                                cx="100"
                                cy="90"
                                rx="55"
                                ry="75"
                                transform="rotate(-15 100 100)"
                            />
                            <ellipse
                                cx="88"
                                cy="72"
                                rx="16"
                                ry="13"
                                transform="rotate(-15 100 100)"
                                fill="white"
                                opacity="0.6"
                            />
                            <ellipse
                                cx="118"
                                cy="60"
                                rx="13"
                                ry="10"
                                transform="rotate(-15 100 100)"
                                fill="white"
                                opacity="0.6"
                            />
                        </svg>
                    </div>
                {/if}
            </div>
            <!-- Content -->
            <div class="p-8 lg:p-10 flex flex-col justify-center">
                <span
                    class="text-xs text-green-600 font-semibold uppercase tracking-widest mb-3"
                    >Latest post</span
                >
                <h2
                    class="text-2xl font-bold text-stone-900 group-hover:text-green-700 transition-colors mb-3 leading-snug"
                >
                    {first.title}
                </h2>
                <p class="text-stone-500 text-sm leading-relaxed mb-5">
                    {first.excerpt}
                </p>
                <div class="flex items-center justify-between">
                    <span class="text-xs text-stone-400"
                        >{formatDate(first.createdAt)}</span
                    >
                    <span
                        class="text-sm font-semibold text-green-700 group-hover:gap-2 transition-all"
                        >Read more →</span
                    >
                </div>
            </div>
        </a>

        <!-- Rest of posts -->
        {#if rest.length > 0}
            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
                {#each rest as post}
                    <a
                        href="/blog/{post.slug}"
                        class="group bg-white rounded-2xl overflow-hidden shadow-sm hover:shadow-md transition-all duration-300 flex flex-col"
                    >
                        <div
                            class="aspect-video bg-gradient-to-br from-green-100 to-green-200 overflow-hidden"
                        >
                            {#if post.coverImage}
                                <img
                                    src={getImageUrl(post.coverImage)}
                                    alt={post.title}
                                    class="w-full h-full object-cover transition-transform duration-500 group-hover:scale-105"
                                />
                            {:else}
                                <div
                                    class="w-full h-full flex items-center justify-center opacity-20"
                                >
                                    <svg
                                        viewBox="0 0 100 100"
                                        class="w-16 h-16 text-green-700"
                                        fill="currentColor"
                                    >
                                        <ellipse
                                            cx="50"
                                            cy="45"
                                            rx="28"
                                            ry="38"
                                            transform="rotate(-15 50 50)"
                                        />
                                    </svg>
                                </div>
                            {/if}
                        </div>
                        <div class="p-5 flex flex-col flex-1">
                            <h3
                                class="font-semibold text-stone-900 group-hover:text-green-700 transition-colors mb-2 leading-snug"
                            >
                                {post.title}
                            </h3>
                            <p
                                class="text-sm text-stone-500 leading-relaxed flex-1 mb-4"
                            >
                                {post.excerpt}
                            </p>
                            <div
                                class="flex items-center justify-between mt-auto"
                            >
                                <span class="text-xs text-stone-400"
                                    >{formatDate(post.createdAt)}</span
                                >
                                <span
                                    class="text-xs font-semibold text-green-700"
                                    >Read →</span
                                >
                            </div>
                        </div>
                    </a>
                {/each}
            </div>
        {/if}
    {/if}
</div>
