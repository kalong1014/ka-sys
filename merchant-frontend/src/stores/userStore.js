import { defineStore } from 'pinia';
import axios from 'axios';
export const useUserStore = defineStore('user', {
    state: () => ({
        token: localStorage.getItem('token') || '',
        userInfo: null
    }),
    actions: {
        setToken(token) {
            this.token = token;
            localStorage.setItem('token', token);
        },
        async fetchUserInfo() {
            try {
                const response = await axios.get('/api/v1/auth/me', {
                    headers: { Authorization: `Bearer ${this.token}` }
                });
                this.userInfo = response.data;
            }
            catch (error) {
                console.error('获取用户信息失败', error);
            }
        },
        logout() {
            this.token = '';
            localStorage.removeItem('token');
            this.userInfo = null;
        }
    }
});
//# sourceMappingURL=userStore.js.map