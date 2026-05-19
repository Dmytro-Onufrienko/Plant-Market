import { sveltekit }  from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
import tailwindcss      from '@tailwindcss/vite';
import { paraglide }    from '@inlang/paraglide-vite';

export default defineConfig({
	plugins: [
		tailwindcss(),
		sveltekit(),
		paraglide({
			project:  './project.inlang',
			outdir:   './src/lib/paraglide'
		})
	]
});
