import { defineStore } from 'pinia'
import axios from 'axios'

export const useOrderStore = defineStore('order', {
  state: () => ({
    orders: [] as Order[],
    total: 0
  }),
  actions: {
    async getOrders(params: { page: number; pageSize: number; search?: string }) {
      const response = await axios.get('/api/v1/orders', { params })
      this.orders = response.data.records
      this.total = response.data.total
      return response.data
    },
    setOrders(orders: Order[]) {
      this.orders = orders
    }
  }
})

interface Order {
  id: string
  orderNo: string
  customerName: string
  totalAmount: number
  status: string
  createdAt: string
  paymentType: string
  items: OrderItem[]
}

interface OrderItem {
  name: string
  quantity: number
  price: number
  total: number
}