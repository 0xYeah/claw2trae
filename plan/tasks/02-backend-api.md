# Task 2: Go后端 CRUD API 实现

## 上下文概要
- **目标**: 实现所有模块的CRUD API handlers和路由
- **前置条件**: Task 1 完成
- **依赖步骤**: 01-backend-structure.md

## 任务清单
1. 实现仓库管理 handlers (`backend/handlers/warehouse.go`)
   - GET /api/warehouses 列表
   - GET /api/warehouses/:id 详情
   - POST /api/warehouses 创建
   - PUT /api/warehouses/:id 更新
   - DELETE /api/warehouses/:id 删除

2. 实现库位管理 handlers (`backend/handlers/location.go`)
   - CRUD 完整操作

3. 实现商品管理 handlers (`backend/handlers/product.go`)
   - CRUD 完整操作

4. 实现库存管理 handlers (`backend/handlers/inventory.go`)
   - GET /api/inventory 列表(带仓库/库位/商品筛选)
   - GET /api/inventory/summary 汇总

5. 实现入库出库 handlers (`backend/handlers/record.go`)
   - GET /api/inbounds 列表
   - POST /api/inbounds 创建入库
   - GET /api/outbounds 列表
   - POST /api/outbounds 创建出库

6. 更新路由文件注册所有路由

## 验证命令
```bash
cd backend && go build -o wms .
curl http://localhost:8080/api/warehouses  # 应返回空数组
```

## 成功标准
- 所有API路由注册成功
- CRUD操作基本验证通过
- 创建入库/出库时自动更新库存数量

## 备注
- 入库时: 库存数量增加
- 出库时: 库存数量减少，需校验库存充足
