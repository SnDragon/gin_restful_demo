# gin+gorm restful demo
## 目录结构
```
.
├── README.md
├── api                     // 对外的API接口(controller)
│   ├── error_code.go
│   └── user.go
├── dao                     // 数据库操作
│   ├── db.go
│   └── user.go
├── go.mod
├── go.sum
├── main.go                 // 主程序
├── migrate.go              // 迁移工具(自动创建表）
├── model                   // 模型
│   └── user.go
└── utils                   // 工具包
    └── utils.go

```
## 修改数据库配置
修改`dao/db.go`的`connection`变量
## 创建表结构
```
go run migrate.go
```
## 运行项目
```
go run main.go
```
1. 创建用户
```curl
curl 'http://127.0.0.1:8080/users' -XPOST -H 'Content-Type: application/json' -d '{"name":"longerwu","age":23}'
```
2. 更新用户
```curl
 curl 'http://127.0.0.1:8080/users/5' -XPUT -H 'Content-Type: application/json' -d '{"name":"longerwu","age":24}'
```
3. 获取单个用户
```curl
curl 'http://127.0.0.1:8080/users/5'
```
4. 获取多个用户(分页)
```curl
 curl 'http://127.0.0.1:8080/users?limit=1&page=2'
```