<template>
  <div class="container mx-auto p-6">
    <h2 class="text-xl font-bold mb-6">角色管理</h2>
    
    <el-card class="mb-6">
      <el-row>
        <el-col :span="12">
          <el-button type="primary" @click="openCreateRoleDialog">
            创建角色
          </el-button>
        </el-col>
      </el-row>
    </el-card>
    
    <el-table
      :data="roles"
      stripe
      border
      @row-click="handleRowClick"
    >
      <el-table-column prop="name" label="角色名称" />
      <el-table-column prop="description" label="角色描述" />
      <el-table-column label="操作">
        <template #default="scope">
          <el-button size="small" @click="editRole(scope.row.id)">编辑</el-button>
          <el-button size="small" type="danger" @click="deleteRole(scope.row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElTable, ElButton, ElCard, ElMessage } from 'element-plus'
import { useRoleStore } from '@/stores/roleStore'

const roleStore = useRoleStore()
const roles = ref([])
const currentRoleId = ref('')

onMounted(() => {
  fetchRoles()
})

const fetchRoles = async () => {
  try {
    const response = await roleStore.getRoles()
    roles.value = response
  } catch (error) {
    ElMessage.error('获取角色列表失败')
  }
}

const openCreateRoleDialog = () => {
  // 打开创建角色对话框
}

const editRole = (id: string) => {
  currentRoleId.value = id
  // 打开编辑角色对话框
}

const deleteRole = async (id: string) => {
  try {
    await roleStore.deleteRole(id)
    ElMessage.success('角色删除成功')
    fetchRoles()
  } catch (error) {
    ElMessage.error('角色删除失败')
  }
}
</script>