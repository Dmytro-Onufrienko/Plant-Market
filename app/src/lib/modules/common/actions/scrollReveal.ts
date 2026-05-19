/**
 * Svelte action: reveals element when it enters the viewport.
 * Usage: <div use:scrollReveal>...</div>
 *        <div use:scrollReveal={{ delay: 100 }}>...</div>
 */
export interface ScrollRevealOptions {
	delay?:     number;   // ms, default 0
	threshold?: number;   // 0–1, default 0.15
	duration?:  number;   // ms CSS transition, default 500
	translateY?: number;  // px, default 20
}

export function scrollReveal(node: HTMLElement, opts: ScrollRevealOptions = {}) {
	const {
		delay     = 0,
		threshold = 0.15,
		duration  = 500,
		translateY = 20
	} = opts;

	node.style.opacity   = '0';
	node.style.transform = `translateY(${translateY}px)`;
	node.style.transition = `opacity ${duration}ms ease, transform ${duration}ms ease`;
	if (delay) node.style.transitionDelay = `${delay}ms`;

	const observer = new IntersectionObserver(
		([entry]) => {
			if (entry.isIntersecting) {
				node.style.opacity   = '1';
				node.style.transform = 'translateY(0)';
				observer.disconnect();
			}
		},
		{ threshold }
	);

	observer.observe(node);

	return {
		destroy() {
			observer.disconnect();
		}
	};
}
