<template>
  <div class="bg-white rounded-xl shadow-sm p-6 mb-8">
    <div class="flex justify-between items-center mb-6">
      <h3 class="font-bold text-neutral-700">订单管理</h3>
      <div class="flex gap-2">
        <div class="relative">
          <input type="text" placeholder="搜索订单号" class="pl-10 pr-4 py-2 rounded-lg border border-neutral-200 focus:outline-none focus:ring-2 focus:ring-primary/20 focus:border-primary">
          <i class="fa fa-search absolute left-3 top-1/2 -translate-y-1/2 text-neutral-400"></i>
        </div>
        <button class="btn-secondary flex items-center gap-2">
          <i class="fa fa-filter"></i>
          <span>筛选</span>
        </button>
      </div>
    </div>
    
    <div class="overflow-x-auto">
      <table class="w-full">
        <thead>
          <tr class="border-b border-neutral-200">
            <th class="text-left py-3 px-4 font-medium text-neutral-500">订单号</th>
            <th class="text-left py-3 px-4 font-medium text-neutral-500">客户</th>
            <th class="text-left py-3 px-4 font-medium text-neutral-500">金额</th>
            <th class="text-left py-3 px-4 font-medium text-neutral-500">状态</th>
            <th class="text-left py-3 px-4 font-medium text-neutral-500">创建时间</th>
            <th class="text-left py-3 px-4 font-medium text-neutral-500">操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="order in orders" :key="order.id" class="border-b border-neutral-100 hover:bg-neutral-50">
            <td class="py-4 px-4">{{ order.order_no }}</td>
            <td class="py-4 px-4">{{ order.customer_name }}</td>
            <td class="py-4 px-4">¥{{ order.total_amount.toFixed(2) }}</td>
            <td class="py-4 px-4">
              <span :class="getStatusClass(order.status)">{{ getStatusText(order.status) }}</span>
            </td>
            <td class="py-4 px-4">{{ formatDate(order.created_at) }}</td>
            <td class="py-4 px-4">
              <button class="text-primary hover:text-primary/80" @click="viewOrderDetails(order.id)">
                查看详情
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    
    <!-- 分页 -->
    <div class="flex justify-between items-center mt-6">
      <div class="text-sm text-neutral-500">
        显示 {{ (currentPage - 1) * pageSize + 1 }} 到 {{ Math.min(currentPage * pageSize, total) }} 条，共 {{ total }} 条
      </div>
      <div class="flex gap-1">
        <button class="w-9 h-9 flex items-center justify-center rounded-lg border border-neutral-200 hover:bg-neutral-100" :disabled="currentPage === 1" @click="currentPage--">
          <i class="fa fa-angle-left"></i>
        </button>
        <button class="w-9 h-9 flex items-center justify-center rounded-lg border border-neutral-200 bg-primary text-white" v-for="page in totalPages" :key="page" :class="{ 'bg-primary text-white': page === currentPage }" @click="currentPage = page">
          {{ page }}
        </button>
        <button class="w-9 h-9 flex items-center justify-center rounded-lg border border-neutral-200 hover:bg-neutral-100" :disabled="currentPage === totalPages" @click="currentPage++">
          <i class="fa fa-angle-right"></i>
        </button>
      </div>
    </div>
  </div>
</template>

<script>
import { listOrders } from '@/services/orderService'

export default {
  data() {
    return {
      orders: [],
      currentPage: 1,
      pageSize: 10,
      total: 0,
      loading: false
    }
  },
  computed: {
    totalPages() {
      return Math.ceil(this.total / this.pageSize)
    }
  },
  mounted() {
    this.fetchOrders()
  },
  methods: {
    async fetchOrders() {
      this.loading = true
      try {
        const merchantId = localStorage.getItem('merchant_id')
        const response = await listOrders(merchantId, this.currentPage, this.pageSize)
        this.orders = response.items
        this.total = response.total
      } catch (error) {
        console.error('获取订单列表失败', error)
        // 显示错误提示
      } finally {
        this.loading = false
      }
    },
    
    viewOrderDetails(orderId) {
      // 跳转到订单详情页
      this.$router.push(`/orders/${orderId}`)
    },
    
    getStatusClass(status) {
      switch (status) {
        case 'pending':
          return 'bg-neutral-100 text-neutral-500 text-xs px-2 py-1 rounded-full'
        case 'paid':
          return 'bg-success/10 text-success text-xs px-2 py-1 rounded-full'
        case 'shipped':
          return 'bg-primary/10 text-primary text-xs px-2 py-1 rounded-full'
        case 'completed':
          return 'bg-secondary/10 text-secondary text-xs px-2 py-1 rounded-full'
        case 'cancelled':
          return 'bg-danger/10 text-danger text-xs px-2 py-1 rounded-full'
        default:
          return 'bg-neutral-100 text-neutral-500 text-xs px-2 py-1 rounded-full'
      }
    },
    
    getStatusText(status) {
      switch (status) {
        case 'pending':
          return '待支付'
        case 'paid':
          return '已支付'
        case 'shipped':
          return '已发货'
        case 'completed':
          return '已完成'
        case 'cancelled':
          return '已取消'
        default:
          return '未知'
      }
    },
    
    formatDate(dateStr) {
      if (!dateStr) return ''
      const date = new Date(dateStr)
      return `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, '0')}-${date.getDate().toString().padStart(2, '0')} ${date.getHours().toString().padStart(2, '0')}:${date.getMinutes().toString().padStart(2, '0')}`
    }
  }
}
</script>