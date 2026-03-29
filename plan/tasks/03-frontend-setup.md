# Task 3: Vue3 前端项目初始化

## 上下文概要
- **目标**: 初始化Vue3前端项目，配置Vite/TypeScript/Pinia/VueRouter
- **前置条件**: Task 2 完成
- **依赖步骤**: 02-backend-api.md

## 任务清单
1. 创建 `frontend/` 目录
2. 创建 `frontend/package.json`
3. 创建 `frontend/vite.config.ts`
4. 创建 `frontend/tsconfig.json`
5. 创建 `frontend/index.html`
6. 创建 `frontend/src/main.ts`
7. 创建 `frontend/src/App.vue`
8. 创建 `frontend/src/router/index.ts`
9. 创建 `frontend/src/api/index.ts` (Axios配置)
10. 创建 `frontend/src/stores/app.ts` (Pinia store)
11. 安装依赖: `cd frontend && npm install`

## 验证命令
```bash
cd frontend && npm install && npm run dev -- --host 0.0.0.0
# 访问 http://localhost:5173 应看到页面
```

## 成功标准
- `npm install` 成功无报错
- `npm run dev` 启动成功
- 浏览器可访问

## 依赖版本
- vue@3
- vue-router@4
- pinia
- axios
- @vitejs/plugin-vue
- typescript
- vite
- vue-tsc
