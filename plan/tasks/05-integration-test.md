# Task 5: 前后端联调与自测

## 上下文概要
- **目标**: 前后端联调，完整功能验证
- **前置条件**: Task 4 完成
- **依赖步骤**: 04-frontend-pages.md

## 任务清单
1. 启动后端服务: `cd backend && ./wms`
2. 启动前端服务: `cd frontend && npm run dev`
3. 功能测试:
   - [ ] 仓库CRUD测试
   - [ ] 库位CRUD测试
   - [ ] 商品CRUD测试
   - [ ] 入库操作测试(库存应增加)
   - [ ] 出库操作测试(库存应减少)
   - [ ] 库存查询测试
4. 跨域问题排查
5. 错误处理验证

## 验证命令
```bash
# 后端
cd backend && ./wms &
curl http://localhost:8080/api/warehouses

# 前端
cd frontend && npm run dev
# 浏览器访问 http://localhost:5173
```

## 成功标准
- 完整流程测试通过:
  1. 创建仓库 -> 创建库位 -> 创建商品
  2. 执行入库 -> 库存数量验证
  3. 执行出库 -> 库存数量验证
- 无明显前端报错
- 无明显后端panic

## 备注
- 后端端口: 8080
- 前端端口: 5173
- 数据库: backend/wms.db
