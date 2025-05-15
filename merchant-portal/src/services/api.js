import axios from 'axios'

const API_BASE_URL = 'http://localhost:8080/api/v1'

const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 添加请求拦截器
api.interceptors.request.use(config => {
  // 从本地存储获取token
  const token = localStorage.getItem('token')
  if (token) {
    config.headers['Authorization'] = `Bearer ${token}`
  }
  return config
}, error => {
  return Promise.reject(error)
})

// 添加响应拦截器
api.interceptors.response.use(response => {
  return response.data
}, error => {
  if (error.response) {
    // 处理HTTP错误
    if (error.response.status === 401) {
      // 未授权，跳转到登录页
      window.location.href = '/login'
    }
  }
  return Promise.reject(error)
})

export default api