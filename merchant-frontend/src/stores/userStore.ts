import { defineStore } from 'pinia'
import axios from 'axios'

export const useUserStore = defineStore('user', {
  state: () => ({
    token: localStorage.getItem('token') || '',
    userInfo: null as UserInfo | null
  }),
  actions: {
    setToken(token: string) {
      this.token = token
      localStorage.setItem('token', token)
    },
    async fetchUserInfo() {
      try {
        const response = await axios.get('/api/v1/auth/me', {
          headers: { Authorization: `Bearer ${this.token}` }
        })
        this.userInfo = response.data
      } catch (error) {
        console.error('获取用户信息失败', error)
      }
    },
    logout() {
      this.token = ''
      localStorage.removeItem('token')
      this.userInfo = null
    }
  }
})

interface UserInfo {
  id: string
  username: string
  role: 'admin' | 'merchant'
}