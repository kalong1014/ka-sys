export const lazyload = {
    install(app) {
        app.directive('lazyload', {
            mounted(el, binding) {
                const observer = new IntersectionObserver((entries) => {
                    entries.forEach(entry => {
                        if (entry.isIntersecting) {
                            el.src = binding.value;
                            observer.unobserve(el);
                        }
                    });
                });
                observer.observe(el);
            },
        });
    },
};
//# sourceMappingURL=lazyload.js.map