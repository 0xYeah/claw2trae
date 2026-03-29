# Task 4: Vue3 前端页面与组件

## 上下文概要
- **目标**: 实现所有前端页面和组件
- **前置条件**: Task 3 完成
- **依赖步骤**: 03-frontend-setup.md

## 任务清单
1. 创建布局组件 `frontend/src/components/Layout.vue`
2. 创建侧边导航 `frontend/src/components/Sidebar.vue`
3. 创建 Dashboard 页面 `frontend/src/views/Dashboard.vue`
   - 统计卡片: 仓库数、商品数、库存总数、入库数、出库数
4. 创建仓库管理页面 `frontend/src/views/Warehouse.vue`
   - 列表展示、新增、编辑、删除
5. 创建库位管理页面 `frontend/src/views/Location.vue`
   - 列表展示、新增、编辑、删除
6. 创建商品管理页面 `frontend/src/views/Product.vue`
   - 列表展示、新增、编辑、删除
7. 创建库存查询页面 `frontend/src/views/Inventory.vue`
   - 筛选功能、列表展示
8. 创建入库管理页面 `frontend/src/views/Inbound.vue`
   - 入库记录列表、入库操作
9. 创建出库管理页面 `frontend/src/views/Outbound.vue`
   - 出库记录列表、出库操作
10. 更新路由配置 `frontend/src/router/index.ts`

## 验证命令
```bash
cd frontend && npm run dev -- --host 0.0.0.0
# 各页面应正常渲染，路由切换正常
```

## 成功标准
- 所有页面路由正常
- 基本CRUD操作可执行
- 前后端API联调成功

## 备注
- 使用 Element Plus 或 TailwindCSS 做UI
- API请求统一封装在 api/index.ts
- 使用 Pinia 管理全局状态
