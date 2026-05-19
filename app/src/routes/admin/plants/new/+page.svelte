<script lang="ts">
	import { enhance } from '$app/forms';
	import { slugify } from '$lib/modules/common/utils/slugify';
	import type { PageData, ActionData } from './$types';

	let { data, form }: { data: PageData; form: ActionData } = $props();

	let nameVal = $state('');
	let slugVal = $state('');

	function onNameInput(e: Event) {
		nameVal = (e.target as HTMLInputElement).value;
		if (!slugVal || slugVal === slugify(nameVal.slice(0, -1))) {
			slugVal = slugify(nameVal);
		}
	}
</script>

<div class="p-8 max-w-2xl">
	<div class="flex items-center gap-3 mb-6">
		<a href="/admin/plants" class="text-stone-400 hover:text-stone-600 text-sm">← Plants</a>
		<span class="text-stone-200">/</span>
		<h1 class="text-2xl font-bold text-stone-900">Add Plant</h1>
	</div>

	<form method="POST" enctype="multipart/form-data" use:enhance class="space-y-5 bg-white rounded-2xl border border-stone-100 shadow-sm p-6">
		<!-- Name -->
		<div>
			<label for="name" class="block text-sm font-medium text-stone-700 mb-1">Name *</label>
			<input id="name" name="name" type="text" value={form?.values?.name ?? nameVal}
				oninput={onNameInput}
				placeholder="Monstera Deliciosa"
				class="w-full px-4 py-2.5 rounded-xl border text-stone-900 text-sm outline-none
					{form?.errors?.name ? 'border-red-400 bg-red-50' : 'border-stone-200 focus:border-green-500'}" />
			{#if form?.errors?.name}<p class="mt-1 text-xs text-red-500">{form.errors.name[0]}</p>{/if}
		</div>

		<!-- Slug -->
		<div>
			<label for="slug" class="block text-sm font-medium text-stone-700 mb-1">Slug *</label>
			<input id="slug" name="slug" type="text" value={form?.values?.slug ?? slugVal}
				oninput={(e) => slugVal = (e.target as HTMLInputElement).value}
				placeholder="monstera-deliciosa"
				class="w-full px-4 py-2.5 rounded-xl border text-stone-900 text-sm font-mono outline-none
					{form?.errors?.slug ? 'border-red-400 bg-red-50' : 'border-stone-200 focus:border-green-500'}" />
			{#if form?.errors?.slug}<p class="mt-1 text-xs text-red-500">{form.errors.slug[0]}</p>{/if}
		</div>

		<!-- Description -->
		<div>
			<label for="desc" class="block text-sm font-medium text-stone-700 mb-1">Description</label>
			<textarea id="desc" name="description" rows="4"
				placeholder="Describe the plant…"
				class="w-full px-4 py-2.5 rounded-xl border border-stone-200 text-stone-900 text-sm resize-none outline-none focus:border-green-500"
			>{form?.values?.description ?? ''}</textarea>
		</div>

		<!-- Price + Stock row -->
		<div class="grid grid-cols-2 gap-4">
			<div>
				<label for="price" class="block text-sm font-medium text-stone-700 mb-1">Price (€) *</label>
				<input id="price" name="price" type="number" min="0" step="0.01"
					value={form?.values?.price ? (form.values.price / 100).toFixed(2) : ''}
					placeholder="34.99"
					class="w-full px-4 py-2.5 rounded-xl border border-stone-200 text-stone-900 text-sm outline-none focus:border-green-500" />
				{#if form?.errors?.price}<p class="mt-1 text-xs text-red-500">{form.errors.price[0]}</p>{/if}
			</div>
			<div>
				<label for="stock" class="block text-sm font-medium text-stone-700 mb-1">Stock *</label>
				<input id="stock" name="stock" type="number" min="0"
					value={form?.values?.stock ?? 0}
					class="w-full px-4 py-2.5 rounded-xl border border-stone-200 text-stone-900 text-sm outline-none focus:border-green-500" />
			</div>
		</div>

		<!-- Category -->
		<div>
			<label for="cat" class="block text-sm font-medium text-stone-700 mb-1">Category</label>
			<select id="cat" name="categoryId"
				class="w-full px-4 py-2.5 rounded-xl border border-stone-200 text-stone-900 text-sm outline-none focus:border-green-500 bg-white">
				<option value="">— No category —</option>
				{#each data.categories as cat}
					<option value={cat.id}>{cat.name}</option>
				{/each}
			</select>
		</div>

		<!-- Images -->
		<div>
			<label class="block text-sm font-medium text-stone-700 mb-1">
				Photos <span class="text-stone-400 font-normal">(WebP, max 800px, multiple allowed)</span>
			</label>
			<input type="file" name="images" accept="image/*" multiple
				class="w-full text-sm text-stone-600 file:mr-3 file:py-2 file:px-4 file:rounded-lg
					file:border-0 file:text-sm file:font-semibold file:bg-green-50 file:text-green-700
					hover:file:bg-green-100 file:cursor-pointer" />
		</div>

		<!-- Featured -->
		<div class="flex items-center gap-3">
			<input id="featured" name="featured" type="checkbox"
				class="w-4 h-4 rounded border-stone-300 accent-green-700 cursor-pointer" />
			<label for="featured" class="text-sm font-medium text-stone-700 cursor-pointer">
				Featured (shown on homepage)
			</label>
		</div>

		<!-- Actions -->
		<div class="flex gap-3 pt-2">
			<button type="submit"
				class="px-6 py-2.5 rounded-xl bg-green-700 text-white text-sm font-semibold hover:bg-green-600 active:scale-95 transition-all">
				Save Plant
			</button>
			<a href="/admin/plants"
				class="px-6 py-2.5 rounded-xl border border-stone-200 text-stone-600 text-sm font-medium hover:border-stone-300 transition-colors">
				Cancel
			</a>
		</div>
	</form>
</div>
