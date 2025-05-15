<template>
  <header class="bg-white shadow-sm">
    <div class="container mx-auto px-4 py-3 flex justify-between items-center">
      <div class="flex items-center">
        <div class="text-xl font-bold">商户管理平台</div>
      </div>
      
      <div class="hidden md:flex items-center space-x-6">
        <router-link to="/dashboard">仪表盘</router-link>
        <router-link to="/domains">域名管理</router-link>
        <router-link to="/orders">订单管理</router-link>
        <router-link to="/pages">页面构建</router-link>
      </div>
      
      <div class="md:hidden">
        <el-button @click="toggleSidebar" icon="fa fa-bars"></el-button>
      </div>
    </div>
    
    <el-drawer
      :visible.sync="sidebarVisible"
      direction="rtl"
      :size="200"
      @close="sidebarVisible = false"
    >
      <template #content>
        <div class="p-4">
          <router-link to="/dashboard" class="block py-2">仪表盘</router-link>
          <router-link to="/domains" class="block py-2">域名管理</router-link>
          <router-link to="/orders" class="block py-2">订单管理</router-link>
          <router-link to="/pages" class="block py-2">页面构建</router-link>
        </div>
      </template>
    </el-drawer>
  </header>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElButton, ElDrawer } from 'element-plus'

const sidebarVisible = ref(false)

const toggleSidebar = () => {
  sidebarVisible.value = !sidebarVisible.value
}
</script>