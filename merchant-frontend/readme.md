merchant-frontend/
├── src/
│   ├── api/                  # API请求
│   ├── assets/               # 静态资源
│   ├── components/           # 组件
│   │   ├── editor/           # 编辑器组件
│   │   ├── layout/           # 布局组件
│   │   └── upload/           # 上传组件
│   ├── stores/               # 状态管理
│   ├── router/               # 路由配置
│   ├── views/                # 页面视图
│   │   ├── dashboard/        # 仪表盘
│   │   ├── domains/          # 域名管理
│   │   ├── orders/           # 订单管理
│   │   ├── pages/            # 页面构建器
│   │   └── role/             # 权限管理
│   ├── styles/               # 样式文件
│   ├── utils/                # 工具函数
│   └── App.vue               # 根组件
├── tests/                    # 测试代码
│   ├── unit/                 # 单元测试
│   └── e2e/                  # E2E测试
├── cypress/                  # Cypress配置
├── .env                      # 环境变量
├── package.json              # 依赖配置
└── vite.config.ts            # Vite配置