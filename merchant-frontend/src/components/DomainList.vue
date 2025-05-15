<template>
  <el-table
    :data="domains"
    stripe
    border
    class="mt-4"
  >
    <el-table-column prop="name" label="域名" width="200" />
    <el-table-column prop="status" label="状态" width="120">
      <template #default="scope">
        <el-badge
          :type="getStatusType(scope.row.status)"
          :text="getStatusText(scope.row.status)"
        />
      </template>
    </el-table-column>
    <el-table-column label="操作" width="180">
      <template #default="scope">
        <el-button
          type="primary"
          size="small"
          @click="editDomain(scope.row.id)"
        >
          编辑
        </el-button>
        <el-button
          type="danger"
          size="small"
          @click="deleteDomain(scope.row.id)"
          :disabled="scope.row.status !== 'pending'"
        >
          删除
        </el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { ElTable, ElButton, ElBadge } from 'element-plus'
import { useRouter } from 'vue-router'

const props = defineProps({
  domains: {
    type: Array,
    default: () => []
  }
})

const router = useRouter()

const getStatusType = (status: string) => {
  switch (status) {
    case 'active': return 'success'
    case 'pending': return 'warning'
    case 'disabled': return 'danger'
    default: return 'info'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'active': return '已激活'
    case 'pending': return '待审核'
    case 'disabled': return '已禁用'
    default: return '未知'
  }
}

const editDomain = (id: string) => {
  router.push(`/domains/${id}/edit`)
}

const deleteDomain = (id: string) => {
  // 调用API删除
}
</script>