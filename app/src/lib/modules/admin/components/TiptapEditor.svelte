<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { Editor } from '@tiptap/core';
	import StarterKit from '@tiptap/starter-kit';
	import Placeholder from '@tiptap/extension-placeholder';

	let {
		content = $bindable(''),
		placeholder = 'Start writing…',
		name = 'content'
	}: {
		content?: string;
		placeholder?: string;
		name?: string;
	} = $props();

	let element: HTMLDivElement;
	let editor: Editor | undefined;
	let htmlValue = $state(content);

	onMount(() => {
		editor = new Editor({
			element,
			extensions: [
				StarterKit,
				Placeholder.configure({ placeholder })
			],
			content,
			editorProps: {
				attributes: {
					class: 'prose prose-stone max-w-none min-h-[300px] focus:outline-none px-5 py-4'
				}
			},
			onTransaction() {
				htmlValue = editor?.getHTML() ?? '';
				content   = htmlValue;
			}
		});
	});

	onDestroy(() => editor?.destroy());
</script>

<!-- Toolbar -->
<div class="flex flex-wrap gap-1 px-3 py-2 border-b border-stone-200 bg-stone-50 rounded-t-xl">
	{#each [
		{ label: 'B',   title: 'Bold',          action: () => editor?.chain().focus().toggleBold().run(),        active: () => editor?.isActive('bold') },
		{ label: 'I',   title: 'Italic',        action: () => editor?.chain().focus().toggleItalic().run(),      active: () => editor?.isActive('italic') },
		{ label: 'S',   title: 'Strike',        action: () => editor?.chain().focus().toggleStrike().run(),      active: () => editor?.isActive('strike') },
		{ label: 'H1',  title: 'Heading 1',     action: () => editor?.chain().focus().toggleHeading({ level: 1 }).run(), active: () => editor?.isActive('heading', { level: 1 }) },
		{ label: 'H2',  title: 'Heading 2',     action: () => editor?.chain().focus().toggleHeading({ level: 2 }).run(), active: () => editor?.isActive('heading', { level: 2 }) },
		{ label: 'H3',  title: 'Heading 3',     action: () => editor?.chain().focus().toggleHeading({ level: 3 }).run(), active: () => editor?.isActive('heading', { level: 3 }) },
		{ label: '≡',   title: 'Bullet list',   action: () => editor?.chain().focus().toggleBulletList().run(),  active: () => editor?.isActive('bulletList') },
		{ label: '1.',  title: 'Ordered list',  action: () => editor?.chain().focus().toggleOrderedList().run(), active: () => editor?.isActive('orderedList') },
		{ label: '❝',   title: 'Blockquote',    action: () => editor?.chain().focus().toggleBlockquote().run(),  active: () => editor?.isActive('blockquote') },
		{ label: '</>',  title: 'Code block',   action: () => editor?.chain().focus().toggleCodeBlock().run(),   active: () => editor?.isActive('codeBlock') },
		{ label: '—',   title: 'Divider',       action: () => editor?.chain().focus().setHorizontalRule().run(), active: () => false },
		{ label: '↩',   title: 'Undo',          action: () => editor?.chain().focus().undo().run(),              active: () => false },
		{ label: '↪',   title: 'Redo',          action: () => editor?.chain().focus().redo().run(),              active: () => false },
	] as btn}
		<button
			type="button"
			title={btn.title}
			onclick={btn.action}
			class="px-2 py-1 text-xs font-mono rounded hover:bg-stone-200 transition-colors
				{btn.active() ? 'bg-stone-200 text-stone-900 font-bold' : 'text-stone-600'}">
			{btn.label}
		</button>
	{/each}
</div>

<!-- Editor area -->
<div bind:this={element} class="rounded-b-xl border-x border-b border-stone-200 bg-white min-h-[300px]"></div>

<!-- Hidden field for form submission -->
<input type="hidden" {name} value={htmlValue} />

<style>
	:global(.tiptap p.is-editor-empty:first-child::before) {
		color: #a8a29e;
		content: attr(data-placeholder);
		float: left;
		height: 0;
		pointer-events: none;
	}
	:global(.tiptap h1) { font-size: 1.75rem; font-weight: 700; margin: 1rem 0 0.5rem; }
	:global(.tiptap h2) { font-size: 1.4rem; font-weight: 700; margin: 0.75rem 0 0.5rem; }
	:global(.tiptap h3) { font-size: 1.15rem; font-weight: 600; margin: 0.5rem 0 0.25rem; }
	:global(.tiptap ul) { list-style: disc; padding-left: 1.5rem; margin: 0.5rem 0; }
	:global(.tiptap ol) { list-style: decimal; padding-left: 1.5rem; margin: 0.5rem 0; }
	:global(.tiptap blockquote) { border-left: 3px solid #d6d3d1; padding-left: 1rem; color: #78716c; margin: 0.5rem 0; }
	:global(.tiptap pre) { background: #1c1917; color: #e7e5e4; padding: 0.75rem 1rem; border-radius: 0.5rem; font-size: 0.85rem; overflow-x: auto; margin: 0.5rem 0; }
	:global(.tiptap code) { background: #f5f5f4; color: #0c4a6e; padding: 0.1em 0.3em; border-radius: 0.2em; font-size: 0.9em; }
	:global(.tiptap pre code) { background: none; color: inherit; padding: 0; }
	:global(.tiptap hr) { border: none; border-top: 2px solid #e7e5e4; margin: 1rem 0; }
	:global(.tiptap strong) { font-weight: 700; }
	:global(.tiptap em) { font-style: italic; }
	:global(.tiptap s) { text-decoration: line-through; }
</style>
