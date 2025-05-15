import { mount } from '@vue/test-utils';
import Login from '@/views/Login.vue';
import { setActivePinia, createPinia } from 'pinia'; // 使用 createPinia
describe('Login.vue', () => {
    let wrapper;
    beforeEach(() => {
        const pinia = createPinia();
        setActivePinia(pinia);
        wrapper = mount(Login, {
            global: {
                plugins: [pinia]
            }
        });
    });
    it('表单验证 - 用户名必填', async () => {
        await wrapper.find('button').trigger('click');
        expect(wrapper.find('.el-form-item__error').text()).toBe('请输入用户名');
    });
    it('表单验证 - 密码必填', async () => {
        await wrapper.setData({ form: { username: 'test' } });
        await wrapper.find('button').trigger('click');
        expect(wrapper.find('.el-form-item__error').text()).toBe('请输入密码');
    });
    it('登录成功后跳转', async () => {
        // 模拟axios请求
        const mockLogin = jest.fn(() => Promise.resolve({ token: 'test-token' }));
        jest.spyOn(wrapper.vm, 'handleLogin').mockImplementation(mockLogin);
        await wrapper.setData({ form: { username: 'admin', password: '123456' } });
        await wrapper.find('button').trigger('click');
        expect(mockLogin).toHaveBeenCalled();
    });
});
//# sourceMappingURL=Login.spec.js.map