<template>
  <div v-if="loaded">
    <component :is="LazyComponent" />
  </div>
  <div v-else class="loading">加载中...</div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'

const loaded = ref(false)
const LazyComponent = ref(null)

onMounted(async () => {
  // 模拟延迟加载
  await new Promise(resolve => setTimeout(resolve, 500))
  LazyComponent.value = (await import('@/components/HeavyComponent.vue')).default
  loaded.value = true
})
</script>