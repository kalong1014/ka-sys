<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-100">
    <div class="bg-white p-8 rounded-lg shadow-lg">
      <h2 class="text-2xl font-bold mb-6 text-gray-800">商户管理平台</h2>
<!--       <el-form ref="formRef" :model="form" :rules="rules" label-width="80px">
        <el-form-item prop="username" :rules="{ required: true, message: '请输入用户名' }">
          <el-input
            v-model="form.username"
            placeholder="请输入用户名"
            prefix-icon="fa fa-user"
          />
        </el-form-item>
        <el-form-item prop="password" :rules="{ required: true, message: '请输入密码' }">
          <el-input
            v-model="form.password"
            type="password"
            placeholder="请输入密码"
            prefix-icon="fa fa-lock"
          />
        </el-form-item>
        <el-form-item>
          <el-button
            type="primary"
            @click="handleLogin"
            class="w-full"
          >
            登录
          </el-button>
        </el-form-item>
      </el-form> -->
<el-form :model="form" :rules="rules">
  <el-form-item prop="username" :rules="{ 
    required: true, 
    message: '请输入用户名', 
    trigger: 'blur',
    pattern: /^[a-zA-Z0-9_]{3,20}$/,
    message: '用户名格式不正确'
  }">
    <el-input v-model="form.username" />
  </el-form-item>
  <el-form-item prop="password" :rules="{ 
    required: true, 
    message: '请输入密码', 
    min: 6, 
    max: 20,
    message: '密码长度6-20位'
  }">
    <el-input type="password" v-model="form.password" />
  </el-form-item>
</el-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import axios from 'axios'
import { ElMessage } from 'element-plus'

const form = ref({
  username: '',
  password: ''
})

const router = useRouter()

const handleLogin = async () => {
  try {
    const response = await axios.post('/api/v1/auth/login', form.value)
    localStorage.setItem('token', response.token)
    router.push('/dashboard')
  } catch (error: any) {
    ElMessage.error(error.response?.data?.message || '登录失败')
  }
}


</script>