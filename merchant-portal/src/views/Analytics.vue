<template>
  <div class="bg-white rounded-xl shadow-sm p-6 mb-8">
    <div class="mb-6">
      <h2 class="text-xl font-bold text-neutral-700">数据统计</h2>
      <p class="text-neutral-500">查看您的业务数据和分析</p>
    </div>
    
    <!-- 数据概览卡片 -->
    <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
      <div class="bg-white rounded-xl shadow-sm p-6 card-hover">
        <div class="flex justify-between items-start">
          <div>
            <p class="text-neutral-500 text-sm">本月销售额</p>
            <h3 class="text-3xl font-bold text-neutral-700 mt-1">¥{{ overview.total_sales.toFixed(2) }}</h3>
            <p class="text-success text-sm mt-2 flex items-center">
              <i class="fa fa-arrow-up mr-1"></i> {{ overview.sales_growth }}% 较上月
            </p>
          </div>
          <div class="bg-primary/10 p-3 rounded-lg">
            <i class="fa fa-line-chart text-primary text-xl"></i>
          </div>
        </div>
      </div>
      
      <div class="bg-white rounded-xl shadow-sm p-6 card-hover">
        <div class="flex justify-between items-start">
          <div>
            <p class="text-neutral-500 text-sm">订单数量</p>
            <h3 class="text-3xl font-bold text-neutral-700 mt-1">{{ overview.total_orders }}</h3>
            <p class="text-success text-sm mt-2 flex items-center">
              <i class="fa fa-arrow-up mr-1"></i> {{ overview.orders_growth }}% 较上月
            </p>
          </div>
          <div class="bg-secondary/10 p-3 rounded-lg">
            <i class="fa fa-shopping-cart text-secondary text-xl"></i>
          </div>
        </div>
      </div>
      
      <div class="bg-white rounded-xl shadow-sm p-6 card-hover">
        <div class="flex justify-between items-start">
          <div>
            <p class="text-neutral-500 text-sm">访客数</p>
            <h3 class="text-3xl font-bold text-neutral-700 mt-1">{{ overview.total_visitors }}</h3>
            <p class="text-success text-sm mt-2 flex items-center">
              <i class="fa fa-arrow-up mr-1"></i> {{ overview.visitors_growth }}% 较上月
            </p>
          </div>
          <div class="bg-warning/10 p-3 rounded-lg">
            <i class="fa fa-users text-warning text-xl"></i>
          </div>
        </div>
      </div>
      
      <div class="bg-white rounded-xl shadow-sm p-6 card-hover">
        <div class="flex justify-between items-start">
          <div>
            <p class="text-neutral-500 text-sm">转化率</p>
            <h3 class="text-3xl font-bold text-neutral-700 mt-1">{{ overview.conversion_rate.toFixed(1) }}%</h3>
            <p class="text-success text-sm mt-2 flex items-center">
              <i class="fa fa-arrow-up mr-1"></i> {{ overview.conversion_growth }}% 较上月
            </p>
          </div>
          <div class="bg-success/10 p-3 rounded-lg">
            <i class="fa fa-percent text-success text-xl"></i>
          </div>
        </div>
      </div>
    </div>
    
    <!-- 图表 -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mb-8">
      <div class="bg-white rounded-xl shadow-sm p-6">
        <div class="flex justify-between items-center mb-6">
          <h3 class="font-bold text-neutral-700">销售额趋势</h3>
          <div class="flex gap-2">
            <button class="btn-secondary text-sm">周</button>
            <button class="btn-primary text-sm">月</button>
            <button class="btn-secondary text-sm">年</button>
          </div>
        </div>
        <div>
          <canvas id="salesChart" height="250"></canvas>
        </div>
      </div>
      
      <div class="bg-white rounded-xl shadow-sm p-6">
        <div class="flex justify-between items-center mb-6">
          <h3 class="font-bold text-neutral-700">流量来源</h3>
          <select class="text-sm border border-neutral-200 rounded-lg px-3 py-1 focus:outline-none focus:ring-2 focus:ring-primary/20 focus:border-primary">
            <option>全部域名</option>
            <option>shop1.example.com</option>
            <option>store2.example.com</option>
          </select>
        </div>
        <div>
          <canvas id="trafficSourceChart" height="250"></canvas>
        </div>
      </div>
    </div>
    
    <!-- 热门产品 -->
    <div class="bg-white rounded-xl shadow-sm p-6 mb-8">
      <div class="flex justify-between items-center mb-6">
        <h3 class="font-bold text-neutral-700">热门产品</h3>
        <a href="#" class="text-primary text-sm">查看全部</a>
      </div>
      
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead>
            <tr class="border-b border-neutral-200">
              <th class="text-left py-3 px-4 font-medium text-neutral-500">产品名称</th>
              <th class="text-left py-3 px-4 font-medium text-neutral-500">销量</th>
              <th class="text-left py-3 px-4 font-medium text-neutral-500">销售额</th>
              <th class="text-left py-3 px-4 font-medium text-neutral-500">转化率</th>
              <th class="text-left py-3 px-4 font-medium text-neutral-500">操作</th>
            </tr>
          </thead>
          <tbody>
            <tr class="border-b border-neutral-100 hover:bg-neutral-50">
              <td class="py-4 px-4">
                <div class="flex items-center gap-3">
                  <img src="https://picsum.photos/100/100" alt="产品图片" class="w-10 h-10 object-cover rounded-lg">
                  <span>高级智能手机</span>
                </div>
              </td>
              <td class="py-4 px-4">128</td>
              <td class="py-4 px-4">¥64,000</td>
              <td class="py-4 px-4">3.2%</td>
              <td class="py-4 px-4">
                <button class="text-primary hover:text-primary/80">查看详情</button>
              </td>
            </tr>
            <tr class="border-b border-neutral-100 hover:bg-neutral-50">
              <td class="py-4 px-4">
                <div class="flex items-center gap-3">
                  <img src="https://picsum.photos/100/101" alt="产品图片" class="w-10 h-10 object-cover rounded-lg">
                  <span>无线耳机</span>
                </div>
              </td>
              <td class="py-4 px-4">96</td>
              <td class="py-4 px-4">¥24,000</td>
              <td class="py-4 px-4">2.8%</td>
              <td class="py-4 px-4">
                <button class="text-primary hover:text-primary/80">查看详情</button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
import { getMerchantOverview, getDomainTraffic } from '@/services/analyticsService'

export default {
  data() {
    return {
      overview: {
        total_sales: 0,
        sales_growth: 0,
        total_orders: 0,
        orders_growth: 0,
        total_visitors: 0,
        visitors_growth: 0,
        conversion_rate: 0,
        conversion_growth: 0
      },
      trafficData: [],
      loading: false
    }
  },
  mounted() {
    this.fetchOverviewData()
    this.fetchTrafficData()
    this.initCharts()
  },
  methods: {
    async fetchOverviewData() {
      this.loading = true
      try {
        const merchantId = localStorage.getItem('merchant_id')
        this.overview = await getMerchantOverview(merchantId)
      } catch (error) {
        console.error('获取概览数据失败', error)
        // 显示错误提示
      } finally {
        this.loading = false
      }
    },
    
    async fetchTrafficData() {
      try {
        const domainId = localStorage.getItem('default_domain_id')
        this.trafficData = await getDomainTraffic(domainId, 30)
        this.updateSalesChart()
      } catch (error) {
        console.error('获取流量数据失败', error)
      }
    },
    
    initCharts() {
      // 初始化图表（详细代码略）
    },
    
    updateSalesChart() {
      // 更新图表数据（详细代码略）
    }
  }
}
</script>