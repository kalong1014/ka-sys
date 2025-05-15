<template>
  <div>
    <el-upload
      class="upload-demo"
      ref="uploadRef"
      :action="uploadUrl"
      :headers="headers"
      :on-success="handleSuccess"
      :on-error="handleError"
      :before-upload="beforeUpload"
      :show-file-list="false"
    >
      <el-button size="small" type="primary">点击上传</el-button>
      <div slot="tip" class="el-upload__tip">
        只能上传jpg/png文件，且不超过2MB
      </div>
    </el-upload>
    
    <img v-if="imageUrl" :src="imageUrl" class="mt-2" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { ElUpload, ElButton } from 'element-plus'

const props = defineProps({
  value: {
    type: String,
    default: ''
  },
  uploadUrl: {
    type: String,
    default: '/api/v1/upload/image'
  }
})

const emit = defineEmits(['update:modelValue'])
const uploadRef = ref(null)
const imageUrl = computed({
  get() { return props.value },
  set(val) { emit('update:modelValue', val) }
})

const headers = computed(() => {
  const token = localStorage.getItem('token')
  return { Authorization: `Bearer ${token}` }
})

const beforeUpload = (file: File) => {
  const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png'
  if (!isJpgOrPng) {
    ElMessage.error('只能上传JPG/PNG文件!')
    return false
  }
  
  const isLt2M = file.size / 1024 / 1024 < 2
  if (!isLt2M) {
    ElMessage.error('图片大小不能超过2MB!')
    return false
  }
  
  return true
}

const handleSuccess = (response, file, fileList) => {
  if (response.code === 200) {
    imageUrl.value = response.data.url
    ElMessage.success('上传成功')
  } else {
    ElMessage.error(response.message || '上传失败')
  }
}

const handleError = (error) => {
  ElMessage.error('上传失败，请重试')
}
</script>