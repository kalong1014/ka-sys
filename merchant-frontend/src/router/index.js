import { createRouter, createWebHistory } from 'vue-router';
const routes = [
    {
        path: '/',
        redirect: '/login',
    },
    {
        path: '/login',
        name: 'Login',
        component: () => import('@/views/Login.vue'), // 懒加载
    },
    {
        path: '/dashboard',
        name: 'Dashboard',
        component: () => import('@/views/Dashboard.vue'), // 懒加载
        meta: { requiresAuth: true },
    },
    {
        path: '/domains',
        name: 'Domains',
        component: () => import('@/views/domains/DomainList.vue'), // 懒加载
        meta: { requiresAuth: true },
    },
    // 其他路由...
];
const router = createRouter({
    history: createWebHistory(),
    routes,
});
export default router;
//# sourceMappingURL=index.js.map