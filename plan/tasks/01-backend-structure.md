# Task 1: Go后端项目结构与数据库模型

## 上下文概要
- **目标**: 搭建Go后端项目结构，使用Gin+GORM+SQLite，实现6张表的数据模型
- **前置条件**: 无
- **依赖步骤**: 无

## 任务清单
1. 创建 `backend/` 目录结构
2. 初始化 Go module (`backend/go.mod`)
3. 创建配置文件 `backend/config/config.go`
4. 创建数据库连接 `backend/database/database.go`
5. 创建数据模型:
   - `backend/models/warehouse.go`
   - `backend/models/location.go`
   - `backend/models/product.go`
   - `backend/models/inventory.go`
   - `backend/models/record.go`
6. 创建响应工具 `backend/utils/response.go`
7. 创建路由结构 `backend/routes/routes.go` (仅骨架)
8. 创建入口 `backend/main.go`

## 验证命令
```bash
cd backend && go mod tidy && go build -o wms .
```

## 成功标准
- `go build` 成功无报错
- 生成 `backend/wms` 可执行文件

## 备注
- 数据库使用SQLite，文件路径: `backend/wms.db`
- 所有模型需有GORM标签
- 启用CORS支持前端跨域
