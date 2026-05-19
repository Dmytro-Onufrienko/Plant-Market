<script>
	import * as m from '$lib/paraglide/messages';
</script>

<footer class="bg-green-900 text-green-100/70 mt-auto">
	<div class="container py-12">
		<div class="grid grid-cols-1 md:grid-cols-4 gap-8 mb-10">
			<div class="md:col-span-2">
				<div class="flex items-center gap-2 mb-3">
					<span class="w-7 h-7 rounded-full bg-green-300 flex items-center justify-center text-green-900 font-bold text-xs">M</span>
					<span class="text-white font-bold text-base">Monstera Shop</span>
				</div>
				<p class="text-sm leading-relaxed max-w-xs">{m.footer_brand_desc()}</p>
			</div>

			<div>
				<h4 class="text-white font-semibold text-sm mb-3">{m.nav_shop()}</h4>
				<ul class="space-y-2 text-sm">
					<li><a href="/shop" class="hover:text-white transition-colors">{m.footer_all_plants()}</a></li>
					<li><a href="/shop?featured=true" class="hover:text-white transition-colors">{m.common_featured()}</a></li>
				</ul>
			</div>

			<div>
				<h4 class="text-white font-semibold text-sm mb-3">{m.footer_info()}</h4>
				<ul class="space-y-2 text-sm">
					<li><a href="/blog" class="hover:text-white transition-colors">{m.nav_blog()}</a></li>
					<li><a href="#" class="hover:text-white transition-colors">{m.footer_shipping_info()}</a></li>
					<li><a href="#" class="hover:text-white transition-colors">{m.footer_returns()}</a></li>
					<li><a href="/auth/register" class="hover:text-white transition-colors">{m.footer_create_account()}</a></li>
				</ul>
			</div>
		</div>

		<div class="border-t border-green-800 pt-6 flex flex-col sm:flex-row justify-between items-center gap-3 text-xs">
			<p>{m.footer_copyright({ year: new Date().getFullYear() })}</p>
			<div class="flex gap-4">
				<a href="#" class="hover:text-white transition-colors">{m.footer_privacy()}</a>
				<a href="#" class="hover:text-white transition-colors">{m.footer_terms()}</a>
			</div>
		</div>
	</div>
</footer>
