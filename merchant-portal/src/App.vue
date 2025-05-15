<template>
  <div class="flex h-screen overflow-hidden">
    <!-- 侧边栏 -->
    <aside class="w-64 bg-white shadow-lg flex-shrink-0 hidden md:block transition-all duration-300" :class="sidebarOpen ? 'translate-x-0' : '-translate-x-full'">
      <div class="p-4 border-b">
        <div class="flex items-center gap-3">
          <div class="bg-primary text-white p-2 rounded-lg">
            <i class="fa fa-shopping-bag text-xl"></i>
          </div>
          <h1 class="text-xl font-bold text-neutral-700">商户管理平台</h1>
        </div>
      </div>
      
      <nav class="p-4">
        <ul class="space-y-1">
          <li>
            <a href="#" class="sidebar-item sidebar-item-active">
              <i class="fa fa-tachometer"></i>
              <span>控制台</span>
            </a>
          </li>
          <li>
            <a href="#" class="sidebar-item hover:bg-neutral-100">
              <i class="fa fa-globe"></i>
              <span>域名管理</span>
            </a>
          </li>
          <li>
            <a href="#" class="sidebar-item hover:bg-neutral-100">
              <i class="fa fa-file-text-o"></i>
              <span>页面管理</span>
            </a>
          </li>
          <li>
            <a href="#" class="sidebar-item hover:bg-neutral-100">
              <i class="fa fa-shopping-cart"></i>
              <span>订单管理</span>
            </a>
          </li>
          <li>
            <a href="#" class="sidebar-item hover:bg-neutral-100">
              <i class="fa fa-line-chart"></i>
              <span>数据统计</span>
            </a>
          </li>
          <li>
            <a href="#" class="sidebar-item hover:bg-neutral-100">
              <i class="fa fa-cog"></i>
              <span>账户设置</span>
            </a>
          </li>
        </ul>
      </nav>
    </aside>

    <!-- 主内容区 -->
    <div class="flex-1 flex flex-col overflow-hidden">
      <!-- 顶部导航栏 -->
      <header class="bg-white shadow-sm h-16 flex items-center justify-between px-6">
        <button @click="toggleSidebar" class="md:hidden text-neutral-700">
          <i class="fa fa-bars text-xl"></i>
        </button>
        
        <div class="flex items-center gap-4">
          <div class="relative">
            <input type="text" placeholder="搜索..." class="pl-10 pr-4 py-2 rounded-lg border border-neutral-200 focus:outline-none focus:ring-2 focus:ring-primary/20 focus:border-primary">
            <i class="fa fa-search absolute left-3 top-1/2 -translate-y-1/2 text-neutral-400"></i>
          </div>
          
          <div class="relative">
            <button class="relative text-neutral-700">
              <i class="fa fa-bell text-xl"></i>
              <span class="absolute -top-1 -right-1 bg-danger text-white text-xs rounded-full h-4 w-4 flex items-center justify-center">3</span>
            </button>
          </div>
          
          <div class="flex items-center gap-3">
            <img src="https://picsum.photos/200/200" alt="用户头像" class="w-8 h-8 rounded-full object-cover">
            <span class="hidden md:inline">管理员</span>
          </div>
        </div>
      </header>
      
      <!-- 内容滚动区 -->
      <main class="flex-1 overflow-y-auto p-6">
        <div class="mb-6">
          <h2 class="text-2xl font-bold text-neutral-700">商户控制台</h2>
          <p class="text-neutral-500">欢迎回来，管理员！这是您的商户管理中心。</p>
        </div>
        
        <!-- 数据概览卡片 -->
        <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
          <div class="bg-white rounded-xl shadow-sm p-6 card-hover">
            <div class="flex justify-between items-start">
              <div>
                <p class="text-neutral-500 text-sm">总域名数</p>
                <h3 class="text-3xl font-bold text-neutral-700 mt-1">12</h3>
                <p class="text-success text-sm mt-2 flex items-center">
                  <i class="fa fa-arrow-up mr-1"></i> 较上月增长 15%
                </p>
              </div>
              <div class="bg-primary/10 p-3 rounded-lg">
                <i class="fa fa-globe text-primary text-xl"></i>
              </div>
            </div>
          </div>
          
          <div class="bg-white rounded-xl shadow-sm p-6 card-hover">
            <div class="flex justify-between items-start">
              <div>
                <p class="text-neutral-500 text-sm">活跃页面</p>
                <h3 class="text-3xl font-bold text-neutral-700 mt-1">48</h3>
                <p class="text-success text-sm mt-2 flex items-center">
                  <i class="fa fa-arrow-up mr-1"></i> 较上月增长 8%
                </p>
              </div>
              <div class="bg-secondary/10 p-3 rounded-lg">
                <i class="fa fa-file-text-o text-secondary text-xl"></i>
              </div>
            </div>
          </div>
          
          <div class="bg-white rounded-xl shadow-sm p-6 card-hover">
            <div class="flex justify-between items-start">
              <div>
                <p class="text-neutral-500 text-sm">本月订单</p>
                <h3 class="text-3xl font-bold text-neutral-700 mt-1">1,248</h3>
                <p class="text-success text-sm mt-2 flex items-center">
                  <i class="fa fa-arrow-up mr-1"></i> 较上月增长 23%
                </p>
              </div>
              <div class="bg-warning/10 p-3 rounded-lg">
                <i class="fa fa-shopping-cart text-warning text-xl"></i>
              </div>
            </div>
          </div>
          
          <div class="bg-white rounded-xl shadow-sm p-6 card-hover">
            <div class="flex justify-between items-start">
              <div>
                <p class="text-neutral-500 text-sm">总交易额</p>
                <h3 class="text-3xl font-bold text-neutral-700 mt-1">¥128,590</h3>
                <p class="text-success text-sm mt-2 flex items-center">
                  <i class="fa fa-arrow-up mr-1"></i> 较上月增长 18%
                </p>
              </div>
              <div class="bg-success/10 p-3 rounded-lg">
                <i class="fa fa-line-chart text-success text-xl"></i>
              </div>
            </div>
          </div>
        </div>
        
        <!-- 图表和表格 -->
        <div class="grid grid-cols-1 lg:grid-cols-3 gap-6 mb-8">
          <div class="bg-white rounded-xl shadow-sm p-6 lg:col-span-2">
            <div class="flex justify-between items-center mb-6">
              <h3 class="font-bold text-neutral-700">流量趋势</h3>
              <div class="flex gap-2">
                <button class="btn-secondary text-sm">日</button>
                <button class="btn-primary text-sm">周</button>
                <button class="btn-secondary text-sm">月</button>
              </div>
            </div>
            <div>
              <canvas id="trafficChart" height="250"></canvas>
            </div>
          </div>
          
          <div class="bg-white rounded-xl shadow-sm p-6">
            <div class="flex justify-between items-center mb-6">
              <h3 class="font-bold text-neutral-700">热门域名</h3>
              <a href="#" class="text-primary text-sm">查看全部</a>
            </div>
            <div class="space-y-4">
              <div class="flex items-center justify-between">
                <div class="flex items-center gap-3">
                  <div class="bg-primary/10 p-2 rounded-lg">
                    <i class="fa fa-globe text-primary"></i>
                  </div>
                  <div>
                    <p class="font-medium">shop1.example.com</p>
                    <p class="text-xs text-neutral-500">访问量: 2,847</p>
                  </div>
                </div>
                <div class="text-success text-xs">+12%</div>
              </div>
              
              <div class="flex items-center justify-between">
                <div class="flex items-center gap-3">
                  <div class="bg-secondary/10 p-2 rounded-lg">
                    <i class="fa fa-globe text-secondary"></i>
                  </div>
                  <div>
                    <p class="font-medium">store2.example.com</p>
                    <p class="text-xs text-neutral-500">访问量: 1,932</p>
                  </div>
                </div>
                <div class="text-success text-xs">+8%</div>
              </div>
              
              <div class="flex items-center justify-between">
                <div class="flex items-center gap-3">
                  <div class="bg-warning/10 p-2 rounded-lg">
                    <i class="fa fa-globe text-warning"></i>
                  </div>
                  <div>
                    <p class="font-medium">mall3.example.com</p>
                    <p class="text-xs text-neutral-500">访问量: 1,568</p>
                  </div>
                </div>
                <div class="text-danger text-xs">-3%</div>
              </div>
              
              <div class="flex items-center justify-between">
                <div class="flex items-center gap-3">
                  <div class="bg-success/10 p-2 rounded-lg">
                    <i class="fa fa-globe text-success"></i>
                  </div>
                  <div>
                    <p class="font-medium">market4.example.com</p>
                    <p class="text-xs text-neutral-500">访问量: 1,205</p>
                  </div>
                </div>
                <div class="text-success text-xs">+15%</div>
              </div>
            </div>
          </div>
        </div>
        
        <!-- 域名管理 -->
        <div class="bg-white rounded-xl shadow-sm p-6 mb-8">
          <div class="flex justify-between items-center mb-6">
            <h3 class="font-bold text-neutral-700">域名管理</h3>
            <button class="btn-primary flex items-center gap-2">
              <i class="fa fa-plus"></i>
              <span>添加域名</span>
            </button>
          </div>
          
          <div class="overflow-x-auto">
            <table class="w-full">
              <thead>
                <tr class="border-b border-neutral-200">
                  <th class="text-left py-3 px-4 font-medium text-neutral-500">域名</th>
                  <th class="text-left py-3 px-4 font-medium text-neutral-500">状态</th>
                  <th class="text-left py-3 px-4 font-medium text-neutral-500">关联页面</th>
                  <th class="text-left py-3 px-4 font-medium text-neutral-500">到期时间</th>
                  <th class="text-left py-3 px-4 font-medium text-neutral-500">操作</th>
                </tr>
              </thead>
              <tbody>
                <tr class="border-b border-neutral-100 hover:bg-neutral-50">
                  <td class="py-4 px-4">shop1.example.com</td>
                  <td class="py-4 px-4">
                    <span class="bg-success/10 text-success text-xs px-2 py-1 rounded-full">已激活</span>
                  </td>
                  <td class="py-4 px-4">首页模板</td>
                  <td class="py-4 px-4">2025-08-15</td>
                  <td class="py-4 px-4">
                    <div class="flex gap-2">
                      <button class="text-primary hover:text-primary/80"><i class="fa fa-edit"></i></button>
                      <button class="text-danger hover:text-danger/80"><i class="fa fa-trash"></i></button>
                    </div>
                  </td>
                </tr>
                <tr class="border-b border-neutral-100 hover:bg-neutral-50">
                  <td class="py-4 px-4">store2.example.com</td>
                  <td class="py-4 px-4">
                    <span class="bg-success/10 text-success text-xs px-2 py-1 rounded-full">已激活</span>
                  </td>
                  <td class="py-4 px-4">产品列表</td>
                  <td class="py-4 px-4">2025-07-22</td>
                  <td class="py-4 px-4">
                    <div class="flex gap-2">
                      <button class="text-primary hover:text-primary/80"><i class="fa fa-edit"></i></button>
                      <button class="text-danger hover:text-danger/80"><i class="fa fa-trash"></i></button>
                    </div>
                  </td>
                </tr>
                <tr class="border-b border-neutral-100 hover:bg-neutral-50">
                  <td class="py-4 px-4">mall3.example.com</td>
                  <td class="py-4 px-4">
                    <span class="bg-warning/10 text-warning text-xs px-2 py-1 rounded-full">待审核</span>
                  </td>
                  <td class="py-4 px-4">促销页面</td>
                  <td class="py-4 px-4">2025-06-30</td>
                  <td class="py-4 px-4">
                    <div class="flex gap-2">
                      <button class="text-primary hover:text-primary/80"><i class="fa fa-edit"></i></button>
                      <button class="text-danger hover:text-danger/80"><i class="fa fa-trash"></i></button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </main>
    </div>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      sidebarOpen: false,
    }
  },
  methods: {
    toggleSidebar() {
      this.sidebarOpen = !this.sidebarOpen
    }
  },
  mounted() {
    // 初始化流量趋势图表
    const ctx = document.getElementById('trafficChart').getContext('2d');
    new Chart(ctx, {
      type: 'line',
      data: {
        labels: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
        datasets: [{
          label: '访问量',
          data: [1200, 1900, 1700, 2100, 1600, 2500, 2300],
          borderColor: '#165DFF',
          backgroundColor: 'rgba(22, 93, 255, 0.1)',
          tension: 0.3,
          fill: true
        }, {
          label: '独立用户',
          data: [800, 1200, 1100, 1400, 1000, 1800, 1600],
          borderColor: '#36CFC9',
          backgroundColor: 'rgba(54, 207, 201, 0.1)',
          tension: 0.3,
          fill: true
        }]
      },
      options: {
        responsive: true,
        plugins: {
          legend: {
            position: 'top',
          }
        },
        scales: {
          y: {
            beginAtZero: true
          }
        }
      }
    });
  }
}
</script>