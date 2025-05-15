import { defineStore } from 'pinia';
import axios from 'axios';
export const useOrderStore = defineStore('order', {
    state: () => ({
        orders: [],
        total: 0
    }),
    actions: {
        async getOrders(params) {
            const response = await axios.get('/api/v1/orders', { params });
            this.orders = response.data.records;
            this.total = response.data.total;
            return response.data;
        },
        setOrders(orders) {
            this.orders = orders;
        }
    }
});
//# sourceMappingURL=orderStore.js.map