# 仓储管理系统 (WMS) - 思维导图

## 系统概述
- **项目名称**: Warehouse Management System (WMS)
- **技术栈**:
  - 前端: Vue3 + Vite + TypeScript + Pinia + Vue Router + Axios
  - 后端: Go + Gin + GORM + SQLite
  - 端口: 前端:5173, 后端:8080

## 功能模块

### 1. 基础数据管理
- [ ] 仓库管理 (Warehouse)
  - 仓库编号、名称、地址、面积、负责人
  - CRUD 操作
- [ ] 库位管理 (Storage Location)
  - 所属仓库、库位编号、容量、状态
  - CRUD 操作

### 2. 商品管理
- [ ] 商品信息 (Product)
  - 商品编号、名称、规格、单位、分类
  - CRUD 操作

### 3. 库存管理
- [ ] 库存查询 (Inventory Query)
  - 按仓库/库位/商品筛选
  - 库存数量展示
- [ ] 入库管理 (Inbound)
  - 入库单号、商品、数量、来源仓库/库位、入库时间
  - 入库确认操作
- [ ] 出库管理 (Outbound)
  - 出库单号、商品、数量、目的仓库/库位、出库时间
  - 出库确认操作

### 4. 记录查询
- [ ] 入库记录查询
- [ ] 出库记录查询
- [ ] 库存变动记录

## 数据模型

### Warehouse (仓库)
```
- Id: uint (PK)
- Code: string (unique)
- Name: string
- Address: string
- Area: float64 (m²)
- Manager: string
- Status: int (1=active, 0=inactive)
- CreatedAt/UpdatedAt: timestamp
```

### StorageLocation (库位)
```
- Id: uint (PK)
- WarehouseId: uint (FK)
- Code: string
- Capacity: float64
- Status: int (1=available, 0=occupied)
- CreatedAt/UpdatedAt: timestamp
```

### Product (商品)
```
- Id: uint (PK)
- Code: string (unique)
- Name: string
- Spec: string
- Unit: string
- Category: string
- CreatedAt/UpdatedAt: timestamp
```

### Inventory (库存)
```
- Id: uint (PK)
- WarehouseId: uint (FK)
- StorageLocationId: uint (FK)
- ProductId: uint (FK)
- Quantity: float64
- UpdatedAt: timestamp
```

### InboundRecord (入库记录)
```
- Id: uint (PK)
- InboundNo: string (unique)
- ProductId: uint (FK)
- WarehouseId: uint (FK)
- StorageLocationId: uint (FK)
- Quantity: float64
- SourceInfo: string
- Operator: string
- CreatedAt: timestamp
```

### OutboundRecord (出库记录)
```
- Id: uint (PK)
- OutboundNo: string (unique)
- ProductId: uint (FK)
- WarehouseId: uint (FK)
- StorageLocationId: uint (FK)
- Quantity: float64
- DestinationInfo: string
- Operator: string
- CreatedAt: timestamp
```

## API 设计

### 仓库管理
- GET    /api/warehouses          - 获取仓库列表
- GET    /api/warehouses/:id      - 获取仓库详情
- POST   /api/warehouses          - 创建仓库
- PUT    /api/warehouses/:id      - 更新仓库
- DELETE /api/warehouses/:id      - 删除仓库

### 库位管理
- GET    /api/locations           - 获取库位列表
- GET    /api/locations/:id       - 获取库位详情
- POST   /api/locations           - 创建库位
- PUT    /api/locations/:id       - 更新库位
- DELETE /api/locations/:id       - 删除库位

### 商品管理
- GET    /api/products            - 获取商品列表
- GET    /api/products/:id        - 获取商品详情
- POST   /api/products            - 创建商品
- PUT    /api/products/:id        - 更新商品
- DELETE /api/products/:id       - 删除商品

### 库存管理
- GET    /api/inventory           - 获取库存列表
- GET    /api/inventory/summary   - 库存汇总

### 入库管理
- GET    /api/inbounds            - 获取入库记录
- POST   /api/inbounds            - 创建入库记录

### 出库管理
- GET    /api/outbounds           - 获取出库记录
- POST   /api/outbounds           - 创建出库记录

## 项目结构

### Go 后端 (backend/)
```
backend/
├── main.go
├── go.mod
├── config/
│   └── config.go
├── models/
│   ├── warehouse.go
│   ├── location.go
│   ├── product.go
│   ├── inventory.go
│   └── record.go
├── handlers/
│   ├── warehouse.go
│   ├── location.go
│   ├── product.go
│   ├── inventory.go
│   └── record.go
├── routes/
│   └── routes.go
├── database/
│   └── database.go
└── utils/
    └── response.go
```

### Vue3 前端 (frontend/)
```
frontend/
├── src/
│   ├── main.ts
│   ├── App.vue
│   ├── api/
│   │   └── index.ts
│   ├── stores/
│   │   └── app.ts
│   ├── views/
│   │   ├── Dashboard.vue
│   │   ├── Warehouse.vue
│   │   ├── Location.vue
│   │   ├── Product.vue
│   │   ├── Inventory.vue
│   │   ├── Inbound.vue
│   │   └── Outbound.vue
│   ├── components/
│   │   └── ...
│   └── router/
│       └── index.ts
├── index.html
├── vite.config.ts
├── tsconfig.json
└── package.json
```

## 依赖清单

### Go 依赖
- github.com/gin-gonic/gin
- github.com/gin-contrib/cors
- gorm.io/gorm
- gorm.io/driver/sqlite
- github.com/google/uuid

### 前端依赖
- vue@3
- vue-router@4
- pinia
- axios
- @vitejs/plugin-vue
- typescript
- vite
