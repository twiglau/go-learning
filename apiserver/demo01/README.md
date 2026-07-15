# 实战：启动一个最简单的RESTful API服务器

## 命令

```bash
# 初始化， 生成 go.mod
go mod init demo01

# 清理mod缓存
go clean -modcache

# 如果调整module名称，重新检查，生成依赖
go mod tidy

```

## 编译前检查

```bash
# 格式化
gofmt -w .
# 或
go fmt ./...

# 检查
go vet ./...

# 整理依赖
go mod tidy

# 编译
go build -v .

# 运行
go run .
```

## curl 使用

- 参数
  
```lua
-X/--request [GET|POST] 指定请求的HTTP方法
-H/--header 指定请求的 HTTP Header
-d/--date 指定请求的 HTTP 消息体 (Body)
-v/--verbose 输出详细的返回信息
-u/--user 指定账号，密码
-b/--cookie 读取 cookie
```

- 示例
  
```bash
curl -v -XPOST -H "Content-Type:application/json" http://127.0.0.1:8080/user -d '{"username": "admin", "password": "admin1234"}'
```
