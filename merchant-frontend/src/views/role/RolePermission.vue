<template>
  <div class="container mx-auto p-6">
    <h2 class="text-xl font-bold mb-6">角色权限设置</h2>
    
    <el-card class="mb-6">
      <el-row>
        <el-col :span="6">
          <el-form-item label="角色名称">
            <el-input v-model="role.name" disabled />
          </el-form-item>
        </el-col>
      </el-row>
    </el-card>
    
    <el-tree
      :data="permissionsTree"
      show-checkbox
      node-key="id"
      ref="treeRef"
      :props="treeProps"
      @check-change="handleCheckChange"
    />
    
    <el-button type="primary" @click="savePermissions" class="mt-6">
      保存设置
    </el-button>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { ElTree, ElButton, ElInput, ElMessage } from 'element-plus'
import { useRoleStore } from '@/stores/roleStore'

const roleStore = useRoleStore()
const treeRef = ref(null)
const role = ref({ id: '', name: '', permissions: [] })
const permissionsTree = ref([])

const treeProps = {
  label: 'name',
  children: 'children'
}

onMounted(() => {
  fetchRolePermissions()
})

const fetchRolePermissions = async () => {
  try {
    // 获取角色信息
    const roleId = '1' // 从路由参数获取
    const roleInfo = await roleStore.getRole(roleId)
    role.value = roleInfo
    
    // 获取权限树
    const permissions = await roleStore.getPermissionsTree()
    permissionsTree.value = permissions
    
    // 设置已选中的权限
    setTimeout(() => {
      if (treeRef.value) {
        treeRef.value.setCheckedKeys(role.value.permissions)
      }
    }, 100)
  } catch (error) {
    ElMessage.error('获取权限数据失败')
  }
}

const handleCheckChange = (node, checked, indeterminate) => {
  // 权限选择变化处理
}

const savePermissions = async () => {
  try {
    const checkedKeys = treeRef.value.getCheckedKeys()
    const halfCheckedKeys = treeRef.value.getHalfCheckedKeys()
    const allPermissions = [...checkedKeys, ...halfCheckedKeys]
    
    await roleStore.updateRolePermissions(role.value.id, allPermissions)
    ElMessage.success('权限设置保存成功')
  } catch (error) {
    ElMessage.error('权限设置保存失败')
  }
}
</script>