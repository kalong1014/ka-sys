<template>
  <div>
    <editor
      :init="editorInit"
      v-model="content"
      @change="handleChange"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Editor from '@tinymce/tinymce-vue'

const props = defineProps({
  modelValue: {
    type: String,
    default: ''
  }
})

const emit = defineEmits(['update:modelValue', 'change'])
const content = ref(props.modelValue)

const editorInit = {
  height: 500,
  menubar: false,
  plugins: [
    'advlist', 'autolink', 'lists', 'link', 'image', 'charmap', 'preview',
    'anchor', 'searchreplace', 'visualblocks', 'code', 'fullscreen',
    'insertdatetime', 'media', 'table', 'code', 'help', 'wordcount'
  ],
  toolbar: 'undo redo | blocks | bold italic backcolor | ' +
    'alignleft aligncenter alignright alignjustify | ' +
    'bullist numlist outdent indent | removeformat | help',
  images_upload_url: '/api/v1/upload/image',
  images_upload_handler: (blobInfo, success, failure) => {
    uploadImage(blobInfo.blob(), success, failure)
  }
}

const uploadImage = (file: File, success: Function, failure: Function) => {
  const formData = new FormData()
  formData.append('file', file)
  
  axios.post('/api/v1/upload/image', formData, {
    headers: { 'Content-Type': 'multipart/form-data' }
  })
    .then(response => {
      success(response.data.url)
    })
    .catch(error => {
      failure('上传失败')
    })
}

const handleChange = (e) => {
  content.value = e.target.getContent()
  emit('update:modelValue', content.value)
  emit('change', content.value)
}
</script>

<style scoped>
.tinymce-container {
  border: 1px solid #ccc;
  border-radius: 4px;
  margin-top: 10px;
}
</style>