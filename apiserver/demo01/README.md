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
