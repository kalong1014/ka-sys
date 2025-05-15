<template>
  <div class="container mx-auto p-6">
    <div class="mb-8">
      <h1 class="text-3xl font-bold text-gray-800">欢迎回来，管理员</h1>
      <p class="text-gray-500 mt-2">今日数据概览</p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6">
      <el-card class="h-32">
        <div class="flex justify-between items-center">
          <div>
            <p class="text-sm text-gray-500">总域名数</p>
            <h2 class="text-2xl font-bold text-primary">{{ domainCount }}</h2>
          </div>
          <i class="fa fa-globe text-2xl text-primary"></i>
        </div>
      </el-card>

      <el-card class="h-32">
        <div class="flex justify-between items-center">
          <div>
            <p class="text-sm text-gray-500">今日订单</p>
            <h2 class="text-2xl font-bold text-success">{{ orderCount }}</h2>
          </div>
          <i class="fa fa-shopping-cart text-2xl text-success"></i>
        </div>
      </el-card>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { ElCard } from 'element-plus'

const domainCount = ref(0)
const orderCount = ref(0)

onMounted(() => {
  fetchDashboardData()
})

const fetchDashboardData = async () => {
  try {
    const response = await axios.get('/api/v1/analytics/overview')
    domainCount.value = response.domainCount
    orderCount.value = response.orderCount
  } catch (error) {
    console.error('获取数据失败', error)
  }
}
</script>