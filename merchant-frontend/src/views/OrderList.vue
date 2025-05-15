<template>
  <div class="container mx-auto p-6">
    <h2 class="text-xl font-bold mb-6">订单管理</h2>
    
    <el-search
      v-model="searchQuery"
      placeholder="搜索订单号/客户名称"
      @search="fetchOrders"
      class="mb-6"
    />
    
    <el-table
      :data="orders"
      stripe
      border
      :loading="loading"
      @sort-change="handleSort"
    >
      <el-table-column prop="orderNo" label="订单号" width="200" />
      <el-table-column prop="customerName" label="客户名称" width="150" />
      <el-table-column prop="totalAmount" label="金额" width="120">
        <template #default="scope">¥{{ scope.row.totalAmount.toFixed(2) }}</template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="120">
        <template #default="scope">
          <el-badge :type="getStatusType(scope.row.status)" :text="getStatusText(scope.row.status)" />
        </template>
      </el-table-column>
      <el-table-column label="操作" width="180">
        <template #default="scope">
          <el-button size="small" @click="viewOrder(scope.row.id)">查看</el-button>
          <el-button
            size="small"
            type="primary"
            @click="editOrder(scope.row.id)"
            v-if="scope.row.status === 'pending'"
          >
            编辑
          </el-button>
          <el-button
            size="small"
            type="danger"
            @click="deleteOrder(scope.row.id)"
            v-if="scope.row.status === 'pending'"
          >
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>
    
    <el-pagination
      :total="total"
      :current-page="currentPage"
      :page-size="pageSize"
      @current-change="handlePageChange"
      class="mt-6"
    />
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'
import { ElTable, ElSearch, ElPagination, ElBadge, ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { useOrderStore } from '@/stores/orderStore'

const router = useRouter()
const orderStore = useOrderStore()

const searchQuery = ref('')
const currentPage = ref(1)
const pageSize = ref(10)
const total = ref(0)
const loading = ref(false)

watch([searchQuery, currentPage], () => {
  fetchOrders()
})

const fetchOrders = async () => {
  loading.value = true
  try {
    const response = await orderStore.getOrders({
      page: currentPage.value,
      pageSize: pageSize.value,
      search: searchQuery.value
    })
    orderStore.setOrders(response.records)
    total.value = response.total
  } catch (error) {
    ElMessage.error('获取订单列表失败')
  } finally {
    loading.value = false
  }
}

const getStatusType = (status: string) => {
  switch (status) {
    case 'pending': return 'warning'
    case 'paid': return 'success'
    case 'cancelled': return 'danger'
    default: return 'info'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'pending': return '待支付'
    case 'paid': return '已支付'
    case 'cancelled': return '已取消'
    default: return '未知'
  }
}

const viewOrder = (id: string) => {
  router.push(`/orders/${id}`)
}
</script>