import axios from 'axios';
import { ElMessage } from 'element-plus';
// 创建axios实例
const service = axios.create({
    baseURL: import.meta.env.VITE_API_BASE_URL,
    timeout: 10000,
});
// 请求拦截器
service.interceptors.request.use(config => {
    // 添加CSRF token
    const csrfToken = localStorage.getItem('csrfToken');
    if (csrfToken) {
        config.headers['X-CSRF-Token'] = csrfToken;
    }
    // 添加认证token
    const authToken = localStorage.getItem('token');
    if (authToken) {
        config.headers['Authorization'] = `Bearer ${authToken}`;
    }
    return config;
}, error => {
    console.error(error);
    return Promise.reject(error);
});
// 响应拦截器
service.interceptors.response.use(response => {
    // 提取新的CSRF token
    const newCsrfToken = response.headers['x-csrf-token'];
    if (newCsrfToken) {
        localStorage.setItem('csrfToken', newCsrfToken);
    }
    return response;
}, error => {
    if (error.response) {
        const { status, data } = error.response;
        switch (status) {
            case 401:
                ElMessage.error('认证失败，请重新登录');
                // 跳转到登录页
                break;
            case 403:
                ElMessage.error('权限不足');
                break;
            case 404:
                ElMessage.error('资源不存在');
                break;
            case 500:
                ElMessage.error('服务器内部错误');
                break;
        }
    }
    return Promise.reject(error);
});
export default service;
//# sourceMappingURL=axios.js.map