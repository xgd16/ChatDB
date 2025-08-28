# AI Chat SQL 前端

这是一个基于 Vue 3 + TypeScript + Naive UI 的 AI 对话系统前端应用。

## 功能特性

- 🤖 **AI 对话**: 支持与 AI 助手进行自然语言对话
- 📝 **Markdown 渲染**: AI 回复支持完整的 Markdown 格式渲染
- 🌓 **主题切换**: 支持明暗主题切换，自动保存用户偏好
- 🔄 **流式响应**: 实时显示 AI 回复，提供更好的用户体验
- 🗄️ **数据库集成**: 支持多种数据库配置和查询
- 📱 **响应式设计**: 适配桌面和移动设备

## 技术栈

- **前端框架**: Vue 3 + TypeScript
- **UI 组件库**: Naive UI
- **构建工具**: Vite
- **包管理器**: pnpm
- **Markdown 渲染**: marked
- **HTTP 客户端**: Axios

## 快速开始

### 安装依赖

```bash
pnpm install
```

### 开发环境运行

```bash
pnpm dev
```

应用将在 `http://localhost:5173` 启动。

### 构建生产版本

```bash
pnpm build
```

## 使用说明

### 登录

使用以下账号登录系统：
- 用户名: `admin`
- 密码: `486213`

### AI 对话

1. 登录后进入主页，可以看到 AI 对话框界面
2. 在顶部工具栏选择：
   - AI 模型 (OpenAI 或 DeepSeek)
   - 具体模型版本
   - 目标数据库
3. 在底部输入框输入问题
4. 按回车键或点击发送按钮开始对话
5. AI 回复会以 Markdown 格式实时显示

### 主题切换

点击右上角的主题切换按钮可以在明暗主题之间切换，设置会自动保存。

## API 接口

### 认证接口

- `POST /api/v1/user/login` - 用户登录

### AI 对话接口

- `POST /api/v1/chats` - AI 聊天 (流式响应)

### 数据库配置接口

- `GET /api/v1/config/db/listConfig` - 获取数据库配置列表
- `POST /api/v1/config/db/setConfig` - 设置数据库配置
- `PUT /api/v1/config/db/updateConfig` - 更新数据库配置
- `DELETE /api/v1/config/db/deleteConfig` - 删除数据库配置

### AI 模型接口

- `GET /api/v1/config/ai/model/list` - 获取 AI 模型列表

## 项目结构

```
src/
├── api/           # API 接口定义
├── components/    # 公共组件
├── router/        # 路由配置
├── stores/        # 状态管理
├── types/         # TypeScript 类型定义
├── utils/         # 工具函数
├── views/         # 页面组件
│   ├── Home.vue   # 主页 (AI 对话框)
│   └── user/      # 用户相关页面
└── App.vue        # 根组件
```

## 开发说明

### 环境配置

开发环境使用 Vite 代理将 `/api` 请求转发到后端服务 `http://127.0.0.1:18200`。

### 主题系统

应用使用 Naive UI 的主题系统，支持：
- 自动检测系统主题
- 手动切换明暗主题
- 主题偏好本地持久化

### Markdown 渲染

AI 回复使用 `marked` 库进行 Markdown 渲染，支持：
- 代码高亮
- 表格
- 列表
- 引用
- 链接
- 图片等

## 注意事项

1. 确保后端服务在 `http://127.0.0.1:18200` 运行
2. 首次使用需要先登录获取 token
3. 数据库配置需要提前在后端设置
4. AI 模型需要配置相应的 API Key

## 许可证

MIT License
