<template>
  <div class="bg-white rounded-xl shadow-sm p-6 mb-8">
    <div class="flex justify-between items-center mb-6">
      <h3 class="font-bold text-neutral-700">域名管理</h3>
      <button class="btn-primary flex items-center gap-2" @click="openAddDomainModal">
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
          <tr v-for="domain in domains" :key="domain.id" class="border-b border-neutral-100 hover:bg-neutral-50">
            <td class="py-4 px-4">{{ domain.name }}</td>
            <td class="py-4 px-4">
              <span :class="getStatusClass(domain.status)">{{ getStatusText(domain.status) }}</span>
            </td>
            <td class="py-4 px-4">{{ domain.page_template }}</td>
            <td class="py-4 px-4">{{ formatDate(domain.expire_date) }}</td>
            <td class="py-4 px-4">
              <div class="flex gap-2">
                <button class="text-primary hover:text-primary/80" @click="openEditDomainModal(domain)">
                  <i class="fa fa-edit"></i>
                </button>
                <button class="text-danger hover:text-danger/80" @click="deleteDomain(domain.id)">
                  <i class="fa fa-trash"></i>
                </button>
              </div>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    
    <!-- 添加域名模态框 -->
    <div v-if="addDomainModalVisible" class="fixed inset-0 bg-black/50 flex items-center justify-center z-50">
      <div class="bg-white rounded-xl shadow-lg w-full max-w-md p-6">
        <div class="flex justify-between items-center mb-4">
          <h3 class="font-bold text-lg">添加域名</h3>
          <button @click="addDomainModalVisible = false">
            <i class="fa fa-times"></i>
          </button>
        </div>
        
        <form @submit.prevent="handleAddDomain">
          <div class="mb-4">
            <label class="block text-sm font-medium text-neutral-700 mb-1">域名</label>
            <input 
              type="text" 
              v-model="newDomain.name" 
              class="w-full px-3 py-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary/20 focus:border-primary"
              placeholder="example.com"
            >
          </div>
          
          <div class="mb-4">
            <label class="block text-sm font-medium text-neutral-700 mb-1">页面模板</label>
            <select 
              v-model="newDomain.page_template" 
              class="w-full px-3 py-2 border border-neutral-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-primary/20 focus:border-primary"
            >
              <option value="home">首页模板</option>
              <option value="product">产品列表</option>
              <option value="promotion">促销页面</option>
              <option value="custom">自定义模板</option>
            </select>
          </div>
          
          <div class="flex justify-end gap-2 mt-6">
            <button type="button" class="btn-secondary" @click="addDomainModalVisible = false">取消</button>
            <button type="submit" class="btn-primary">保存</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import { getDomains, addDomain, updateDomain, deleteDomain } from '@/services/domainService'

export default {
  data() {
    return {
      domains: [],
      addDomainModalVisible: false,
      newDomain: {
        name: '',
        page_template: 'home',
        status: 'pending',
        expire_date: new Date().toISOString().split('T')[0]
      }
    }
  },
  mounted() {
    this.fetchDomains()
  },
  methods: {
    async fetchDomains() {
      try {
        this.domains = await getDomains()
      } catch (error) {
        console.error('获取域名列表失败', error)
        // 显示错误提示
      }
    },
    
    openAddDomainModal() {
      this.addDomainModalVisible = true
      this.newDomain = {
        name: '',
        page_template: 'home',
        status: 'pending',
        expire_date: new Date().toISOString().split('T')[0]
      }
    },
    
    async handleAddDomain() {
      try {
        await addDomain(this.newDomain)
        this.addDomainModalVisible = false
        this.fetchDomains()
        // 显示成功提示
      } catch (error) {
        console.error('添加域名失败', error)
        // 显示错误提示
      }
    },
    
    openEditDomainModal(domain) {
      // 实现编辑功能
    },
    
    async deleteDomain(id) {
      if (confirm('确定要删除这个域名吗？')) {
        try {
          await deleteDomain(id)
          this.fetchDomains()
          // 显示成功提示
        } catch (error) {
          console.error('删除域名失败', error)
          // 显示错误提示
        }
      }
    },
    
    getStatusClass(status) {
      switch (status) {
        case 'active':
          return 'bg-success/10 text-success text-xs px-2 py-1 rounded-full'
        case 'pending':
          return 'bg-warning/10 text-warning text-xs px-2 py-1 rounded-full'
        case 'disabled':
          return 'bg-danger/10 text-danger text-xs px-2 py-1 rounded-full'
        default:
          return 'bg-neutral-100 text-neutral-500 text-xs px-2 py-1 rounded-full'
      }
    },
    
    getStatusText(status) {
      switch (status) {
        case 'active':
          return '已激活'
        case 'pending':
          return '待审核'
        case 'disabled':
          return '已禁用'
        default:
          return '未知'
      }
    },
    
    formatDate(dateStr) {
      if (!dateStr) return ''
      const date = new Date(dateStr)
      return `${date.getFullYear()}-${(date.getMonth() + 1).toString().padStart(2, '0')}-${date.getDate().toString().padStart(2, '0')}`
    }
  }
}
</script>