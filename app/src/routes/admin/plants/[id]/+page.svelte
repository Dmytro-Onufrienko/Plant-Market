<script lang="ts">
    import { enhance } from "$app/forms";
    import { parseImages, getImageUrl } from "$lib/modules/common/utils/format";
    import type { PageData, ActionData } from "./$types";

    let { data, form }: { data: PageData; form: ActionData } = $props();
    const p = data.plant;

    // Track which existing images to keep
    const existingImgs = $derived(parseImages(p.images));
    let keepImgs = $state<string[]>(existingImgs);

    function removeImg(src: string) {
        keepImgs = keepImgs.filter((i) => i !== src);
    }
</script>

<div class="p-8 max-w-2xl">
    <div class="flex items-center gap-3 mb-6">
        <a
            href="/admin/plants"
            class="text-stone-400 hover:text-stone-600 text-sm">← Plants</a
        >
        <span class="text-stone-200">/</span>
        <h1 class="text-2xl font-bold text-stone-900">Edit: {p.name}</h1>
    </div>

    <form
        method="POST"
        enctype="multipart/form-data"
        use:enhance
        class="space-y-5 bg-white rounded-2xl border border-stone-100 shadow-sm p-6"
    >
        <div>
            <label
                for="name"
                class="block text-sm font-medium text-stone-700 mb-1"
                >Name *</label
            >
            <input
                id="name"
                name="name"
                type="text"
                value={form?.values?.name ?? p.name}
                class="w-full px-4 py-2.5 rounded-xl border {form?.errors?.name
                    ? 'border-red-400'
                    : 'border-stone-200 focus:border-green-500'} text-stone-900 text-sm outline-none"
            />
            {#if form?.errors?.name}<p class="mt-1 text-xs text-red-500">
                    {form.errors.name[0]}
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
                class="w-full px-4 py-2.5 rounded-xl border border-stone-200 focus:border-green-500 text-stone-900 text-sm font-mono outline-none"
            />
        </div>

        <div>
            <label
                for="desc"
                class="block text-sm font-medium text-stone-700 mb-1"
                >Description</label
            >
            <textarea
                id="desc"
                name="description"
                rows="4"
                class="w-full px-4 py-2.5 rounded-xl border border-stone-200 focus:border-green-500 text-stone-900 text-sm resize-none outline-none"
                >{form?.values?.description ?? p.description}</textarea
            >
        </div>

        <div class="grid grid-cols-2 gap-4">
            <div>
                <label
                    for="price"
                    class="block text-sm font-medium text-stone-700 mb-1"
                    >Price (€) *</label
                >
                <input
                    id="price"
                    name="price"
                    type="number"
                    min="0"
                    step="0.01"
                    value={form?.values?.price
                        ? (form.values.price / 100).toFixed(2)
                        : (p.price / 100).toFixed(2)}
                    class="w-full px-4 py-2.5 rounded-xl border border-stone-200 focus:border-green-500 text-stone-900 text-sm outline-none"
                />
            </div>
            <div>
                <label
                    for="stock"
                    class="block text-sm font-medium text-stone-700 mb-1"
                    >Stock *</label
                >
                <input
                    id="stock"
                    name="stock"
                    type="number"
                    min="0"
                    value={form?.values?.stock ?? p.stock}
                    class="w-full px-4 py-2.5 rounded-xl border border-stone-200 focus:border-green-500 text-stone-900 text-sm outline-none"
                />
            </div>
        </div>

        <div>
            <label
                for="cat"
                class="block text-sm font-medium text-stone-700 mb-1"
                >Category</label
            >
            <select
                id="cat"
                name="categoryId"
                class="w-full px-4 py-2.5 rounded-xl border border-stone-200 focus:border-green-500 text-stone-900 text-sm outline-none bg-white"
            >
                <option value="">— No category —</option>
                {#each data.categories as cat}
                    <option value={cat.id} selected={cat.id === p.category_id}
                        >{cat.name}</option
                    >
                {/each}
            </select>
        </div>

        <!-- Current images -->
        {#if existingImgs.length > 0}
            <div>
                <label class="block text-sm font-medium text-stone-700 mb-2"
                    >Current Photos</label
                >
                <div class="flex flex-wrap gap-3">
                    {#each existingImgs as src}
                        <div class="relative group">
                            <img
                                src={getImageUrl(src)}
                                alt=""
                                class="w-20 h-20 object-cover rounded-xl border-2
                				{keepImgs.includes(src)
                                    ? 'border-green-500'
                                    : 'border-stone-200 opacity-40'}"
                            />
                            {#if keepImgs.includes(src)}
                                <!-- Hidden field to keep this image -->
                                <input
                                    type="hidden"
                                    name="keepImage"
                                    value={src}
                                />
                                <button
                                    type="button"
                                    onclick={() => removeImg(src)}
                                    class="absolute -top-1.5 -right-1.5 w-5 h-5 rounded-full bg-red-500 text-white text-xs
										flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity"
                                >
                                    ×
                                </button>
                            {:else}
                                <div
                                    class="absolute inset-0 flex items-center justify-center"
                                >
                                    <span
                                        class="text-xs text-stone-400 bg-white/80 px-1 rounded"
                                        >removed</span
                                    >
                                </div>
                                <button
                                    type="button"
                                    onclick={() =>
                                        (keepImgs = [...keepImgs, src])}
                                    class="absolute -top-1.5 -right-1.5 w-5 h-5 rounded-full bg-green-500 text-white text-xs
										flex items-center justify-center opacity-0 group-hover:opacity-100 transition-opacity"
                                >
                                    +
                                </button>
                            {/if}
                        </div>
                    {/each}
                </div>
                <p class="text-xs text-stone-400 mt-1">
                    Hover over a photo and click × to remove it
                </p>
            </div>
        {/if}

        <!-- Upload new images -->
        <div>
            <label class="block text-sm font-medium text-stone-700 mb-1">
                Add New Photos <span class="text-stone-400 font-normal"
                    >(WebP, max 800px)</span
                >
            </label>
            <input
                type="file"
                name="images"
                accept="image/*"
                multiple
                class="w-full text-sm text-stone-600 file:mr-3 file:py-2 file:px-4 file:rounded-lg
					file:border-0 file:text-sm file:font-semibold file:bg-green-50 file:text-green-700
					hover:file:bg-green-100 file:cursor-pointer"
            />
        </div>

        <div class="flex items-center gap-3">
            <input
                id="featured"
                name="featured"
                type="checkbox"
                checked={p.featured}
                class="w-4 h-4 rounded border-stone-300 accent-green-700 cursor-pointer"
            />
            <label
                for="featured"
                class="text-sm font-medium text-stone-700 cursor-pointer"
                >Featured</label
            >
        </div>

        <div class="flex gap-3 pt-2">
            <button
                type="submit"
                class="px-6 py-2.5 rounded-xl bg-green-700 text-white text-sm font-semibold hover:bg-green-600 active:scale-95 transition-all"
            >
                Save Changes
            </button>
            <a
                href="/admin/plants"
                class="px-6 py-2.5 rounded-xl border border-stone-200 text-stone-600 text-sm font-medium hover:border-stone-300 transition-colors"
            >
                Cancel
            </a>
        </div>
    </form>
</div>
