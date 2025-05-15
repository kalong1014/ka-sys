<template>
  <div class="flex">
    <!-- 左侧元素面板 -->
    <div class="w-64 bg-white p-4">
      <h3 class="text-lg font-bold mb-4">页面元素</h3>
      
      <el-card v-for="element in elements" :key="element.id" class="mb-4">
        <el-button
          type="primary"
          icon="fa fa-plus"
          @click="addElement(element)"
          class="w-full"
        >
          {{ element.name }}
        </el-button>
      </el-card>
    </div>
    
    <!-- 中间画布 -->
    <div class="flex-1 p-4">
      <div
        v-for="(element, index) in canvasElements"
        :key="index"
        :class="element.type"
        draggable
        @dragstart="handleDragStart(index)"
      >
        {{ element.content }}
      </div>
    </div>
    
    <!-- 右侧属性面板 -->
    <div class="w-48 bg-white p-4">
      <h3 class="text-lg font-bold mb-4">元素属性</h3>
      
      <el-input
        v-model="selectedElement.content"
        label="内容"
        v-if="selectedElement"
      />
      
      <el-select
        v-model="selectedElement.type"
        label="类型"
        v-if="selectedElement"
      >
        <el-option label="标题" value="title" />
        <el-option label="文本" value="text" />
        <el-option label="图片" value="image" />
      </el-select>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const elements = ref([
  { id: 1, name: "标题", type: "title" },
  { id: 2, name: "文本", type: "text" },
  { id: 3, name: "图片", type: "image" }
])

const canvasElements = ref<ElementData[]>([])
const selectedElement = ref<ElementData | null>(null)

const addElement = (element: { type: string }) => {
  canvasElements.value.push({
    id: Date.now(),
    type: element.type,
    content: "",
    position: canvasElements.value.length
  })
}

const handleDragStart = (index: number) => {
  selectedElement.value = canvasElements.value[index]
}

interface ElementData {
  id: number
  type: string
  content: string
  position: number
}

const savePage = async () => {
  try {
    const pageData = {
      name: "首页",
      elements: canvasElements.value.map(element => ({
        type: element.type,
        content: element.content,
        position: element.position
      }))
    }
    
    await axios.post('/api/v1/pages', pageData)
    ElMessage.success('页面保存成功')
  } catch (error) {
    ElMessage.error('保存失败，请重试')
  }
}
</script>