## gin+gorm restful demo
### 创建表结构
```
go run migrate.go
```
### 运行项目
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