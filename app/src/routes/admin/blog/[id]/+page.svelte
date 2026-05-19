<script lang="ts">
    import { enhance } from "$app/forms";
    import TiptapEditor from "$lib/modules/admin/components/TiptapEditor.svelte";
    import { getImageUrl } from "$lib/modules/common/utils/format";
    import type { PageData, ActionData } from "./$types";

    let { data, form }: { data: PageData; form: ActionData } = $props();
    const p = data.post;

    let editorContent = $state(form?.values?.content ?? p.content);
    let removeCoverImage = $state(false);
</script>

<div class="p-8 max-w-3xl">
    <div class="flex items-center gap-3 mb-6">
        <a
            href="/admin/blog"
            class="text-stone-400 hover:text-stone-600 text-sm">← Blog</a
        >
        <span class="text-stone-200">/</span>
        <h1 class="text-2xl font-bold text-stone-900">Edit: {p.title}</h1>
    </div>

    <form
        method="POST"
        enctype="multipart/form-data"
        use:enhance
        class="space-y-5"
    >
        <!-- Meta fields -->
        <div
            class="bg-white rounded-2xl border border-stone-100 shadow-sm p-6 space-y-5"
        >
            <div>
                <label
                    for="title"
                    class="block text-sm font-medium text-stone-700 mb-1"
                    >Title *</label
                >
                <input
                    id="title"
                    name="title"
                    type="text"
                    value={form?.values?.title ?? p.title}
                    class="w-full px-4 py-2.5 rounded-xl border text-stone-900 text-sm outline-none
						{form?.errors?.title
                        ? 'border-red-400 bg-red-50'
                        : 'border-stone-200 focus:border-green-500'}"
                />
                {#if form?.errors?.title}<p class="mt-1 text-xs text-red-500">
                        {form.errors.title[0]}
                    </p>{/if}
            </div>

            <div>
                <label
                    for="slug"
                    class="block text-sm font-medium text-stone-700 mb-1"
                    >Slug *</label
                >
                <input
                    id="slug"
                    name="slug"
                    type="text"
                    value={form?.values?.slug ?? p.slug}
                    class="w-full px-4 py-2.5 rounded-xl border border-stone-200 focus:border-green-500 text-stone-900 text-sm font-mono outline-none
						{form?.errors?.slug ? 'border-red-400 bg-red-50' : ''}"
                />
                {#if form?.errors?.slug}<p class="mt-1 text-xs text-red-500">
                        {form.errors.slug[0]}
                    </p>{/if}
            </div>

            <div>
                <label
                    for="excerpt"
                    class="block text-sm font-medium text-stone-700 mb-1"
                >
                    Excerpt <span class="text-stone-400 font-normal"
                        >(max 300 chars)</span
                    >
                </label>
                <textarea
                    id="excerpt"
                    name="excerpt"
                    rows="2"
                    maxlength="300"
                    class="w-full px-4 py-2.5 rounded-xl border border-stone-200 focus:border-green-500 text-stone-900 text-sm resize-none outline-none"
                    >{form?.values?.excerpt ?? p.excerpt}</textarea
                >
            </div>

            <div>
                <label class="block text-sm font-medium text-stone-700 mb-2"
                    >Cover Image</label
                >

                {#if p.coverImage && !removeCoverImage}
                    <div class="relative w-48 h-32 mb-3 group">
                        <img
                            src={getImageUrl(p.coverImage)}
                            alt="Cover"
                            class="w-full h-full object-cover rounded-xl border border-stone-200"
                        />
                        <input
                            type="hidden"
                            name="existingCoverImage"
                            value={p.coverImage}
                        />
                        <button
                            type="button"
                            onclick={() => (removeCoverImage = true)}
                            class="absolute -top-2 -right-2 w-6 h-6 rounded-full bg-red-500 text-white text-xs flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity shadow-sm"
                        >
                            ×
                        </button>
                    </div>
                {/if}

                {#if removeCoverImage}
                    <input type="hidden" name="removeCoverImage" value="true" />
                {/if}

                <input
                    type="file"
                    name="coverImage"
                    accept="image/*"
                    class="w-full text-sm text-stone-600 file:mr-3 file:py-2 file:px-4 file:rounded-lg
            			file:border-0 file:text-sm file:font-semibold file:bg-green-50 file:text-green-700
            			hover:file:bg-green-100 file:cursor-pointer"
                />
            </div>
        </div>

        <!-- Tiptap Editor -->
        <div>
            <label class="block text-sm font-medium text-stone-700 mb-2"
                >Content</label
            >
            <TiptapEditor bind:content={editorContent} name="content" />
        </div>

        <!-- Footer -->
        <div
            class="flex items-center justify-between bg-white rounded-2xl border border-stone-100 shadow-sm px-6 py-4"
        >
            <div class="flex items-center gap-3">
                <input
                    id="published"
                    name="published"
                    type="checkbox"
                    checked={form?.values?.published ?? p.published}
                    class="w-4 h-4 rounded border-stone-300 accent-green-700 cursor-pointer"
                />
                <label
                    for="published"
                    class="text-sm font-medium text-stone-700 cursor-pointer"
                    >Published</label
                >
            </div>
            <div class="flex gap-3">
                <a
                    href="/admin/blog"
                    class="px-5 py-2.5 rounded-xl border border-stone-200 text-stone-600 text-sm font-medium hover:border-stone-300 transition-colors"
                >
                    Cancel
                </a>
                <button
                    type="submit"
                    class="px-6 py-2.5 rounded-xl bg-bark-700 text-white text-sm font-semibold hover:bg-bark-600 active:scale-95 transition-all"
                >
                    Save Changes
                </button>
            </div>
        </div>
    </form>
</div>
