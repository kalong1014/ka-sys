import api from './api'

// 用户登录
export const login = async (username, password) => {
  const response = await api.post('/auth/login', {
    username,
    password
  })
  
  // 保存token到本地存储
  localStorage.setItem('token', response.token)
  return response
}

// 用户注册
export const register = async (userData) => {
  return api.post('/auth/register', userData)
}

// 检查用户是否已登录
export const checkAuth = () => {
  const token = localStorage.getItem('token')
  return token !== null
}

// 用户退出登录
export const logout = () => {
  localStorage.removeItem('token')
  window.location.href = '/login'
}