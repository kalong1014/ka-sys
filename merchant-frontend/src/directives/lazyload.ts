import type { App } from 'vue'  // 添加 type 关键字

export const lazyload = {
  install(app: App) {
    app.directive('lazyload', {
      mounted(el: HTMLImageElement, binding) {
        const observer = new IntersectionObserver((entries) => {
          entries.forEach(entry => {
            if (entry.isIntersecting) {
              el.src = binding.value
              observer.unobserve(el)
            }
          })
        })
        observer.observe(el)
      },
    })
  },
}