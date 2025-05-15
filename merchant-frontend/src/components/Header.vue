<template>
  <header class="bg-white shadow-sm">
    <div class="container mx-auto px-4 py-3 flex justify-between items-center">
      <div class="flex items-center">
        <div class="text-xl font-bold">{{ $t('common.welcome') }}</div>
      </div>
      
      <div class="flex items-center space-x-4">
        <el-select v-model="currentLocale" @change="changeLocale">
          <el-option label="中文" value="zh"></el-option>
          <el-option label="English" value="en"></el-option>
        </el-select>
      </div>
    </div>
  </header>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'

const { t, locale } = useI18n()
const currentLocale = ref(locale.value)

const changeLocale = (val: string) => {
  locale.value = val
  localStorage.setItem('locale', val)
}
</script>