<template>
  <div class="container mx-auto p-6">
    <h2 class="text-xl font-bold mb-6">订单详情 #{{ order.orderNo }}</h2>
    
    <el-card class="mb-6">
      <el-row :gutter="20">
        <el-col :span="12">
          <el-row>
            <el-col :span="6" class="font-medium">订单状态：</el-col>
            <el-col :span="18">{{ getStatusText(order.status) }}</el-col>
          </el-row>
          <el-row>
            <el-col :span="6" class="font-medium">下单时间：</el-col>
            <el-col :span="18">{{ formatDate(order.createdAt) }}</el-col>
          </el-row>
        </el-col>
        <el-col :span="12">
          <el-row>
            <el-col :span="6" class="font-medium">总金额：</el-col>
            <el-col :span="18">¥{{ order.totalAmount.toFixed(2) }}</el-col>
          </el-row>
          <el-row>
            <el-col :span="6" class="font-medium">支付方式：</el-col>
            <el-col :span="18">{{ order.paymentType }}</el-col>
          </el-row>
        </el-col>
      </el-row>
    </el-card>
    
    <h3 class="text-lg font-bold mb-4">商品列表</h3>
    <el-table :data="order.items" border stripe>
      <el-table-column prop="name" label="商品名称" />
      <el-table-column prop="quantity" label="数量" width="80" />
      <el-table-column prop="price" label="单价" width="100">¥{{ scope.row.price.toFixed(2) }}</el-table-column>
      <el-table-column prop="total" label="小计" width="100">¥{{ scope.row.total.toFixed(2) }}</el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRoute } from 'vue-router'
import { useOrderStore } from '@/stores/orderStore'

const route = useRoute()
const orderStore = useOrderStore()

const order = computed(() => 
  orderStore.orders.find(order => order.id === route.params.id) || null
)

const formatDate = (dateStr: string) => {
  return new Date(dateStr).toLocaleString()
}
</script>