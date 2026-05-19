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
    <title>{data.post.title} — Monstera Shop Blog</title>
    <meta name="description" content={data.post.excerpt} />
</svelte:head>

<!-- Breadcrumb -->
<div class="bg-stone-100 border-b border-stone-200">
    <div
        class="container py-2.5 text-xs text-stone-400 flex items-center gap-1.5"
    >
        <a href="/" class="hover:text-stone-600">Home</a>
        <span>/</span>
        <a href="/blog" class="hover:text-stone-600">Blog</a>
        <span>/</span>
        <span class="text-stone-600 truncate max-w-[200px]"
            >{data.post.title}</span
        >
    </div>
</div>

<!-- Cover image -->
{#if data.post.coverImage}
    <div class="w-full h-64 lg:h-96 bg-green-100 overflow-hidden">
        <img
            src={getImageUrl(data.post.coverImage)}
            alt={data.post.title}
            class="w-full h-full object-cover"
        />
    </div>
{:else}
    <div
        class="w-full h-48 bg-gradient-to-br from-green-800 to-green-950 flex items-center justify-center"
    >
        <svg
            viewBox="0 0 200 120"
            class="w-32 h-auto text-green-600 opacity-30"
            fill="currentColor"
        >
            <ellipse
                cx="100"
                cy="60"
                rx="55"
                ry="55"
                transform="rotate(-15 100 60)"
            />
            <ellipse
                cx="88"
                cy="48"
                rx="14"
                ry="11"
                transform="rotate(-15 100 60)"
                fill="white"
                opacity="0.6"
            />
            <ellipse
                cx="118"
                cy="40"
                rx="11"
                ry="9"
                transform="rotate(-15 100 60)"
                fill="white"
                opacity="0.6"
            />
        </svg>
    </div>
{/if}

<div class="container py-10 lg:py-14">
    <div class="max-w-2xl mx-auto">
        <!-- Post header -->
        <div class="mb-8">
            <a
                href="/blog"
                class="text-xs text-green-600 font-semibold uppercase tracking-widest hover:underline mb-4 inline-block"
            >
                ← Back to Blog
            </a>
            <h1
                class="text-3xl lg:text-4xl font-bold text-stone-900 leading-tight mb-3"
            >
                {data.post.title}
            </h1>
            <div class="flex items-center gap-3 text-sm text-stone-400">
                <span>Monstera Shop</span>
                <span>·</span>
                <span>{formatDate(data.post.createdAt)}</span>
            </div>
        </div>

        <!-- Post content — rendered HTML from Tiptap -->
        <div
            class="prose prose-stone prose-headings:font-bold prose-headings:text-stone-900
			prose-p:text-stone-600 prose-p:leading-relaxed prose-a:text-green-700
			prose-strong:text-stone-900 prose-li:text-stone-600 max-w-none"
        >
            {@html data.post.content}
        </div>

        <!-- Back link -->
        <div class="mt-12 pt-8 border-t border-stone-200">
            <a
                href="/blog"
                class="inline-flex items-center gap-2 text-sm font-semibold text-green-700
					hover:text-green-600 transition-colors"
            >
                ← All posts
            </a>
        </div>
    </div>
</div>
