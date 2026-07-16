# 实战：记录和管理API日志

## 查找文件命令

```bash
# 查找某个包 -r recursive 递归搜索
grep -r "lexkong" .
# 只查 Go 文件
grep -r "lexkong" --include="*.go" .

# 删除旧依赖
go mod edit -droprequire github.com/lexkong/log
# 添加新的 log
go get github.com/zxmrlc/log
```

## go项目如何替换第三方包版本，该第三方包被多个go文件引用

- 例如将 log包：【github.com/lexkong/log】 -> 【github.com/zxmrlc/log】 

1. 修改 【go.mod】文件, 使用 replace 命令

```lua
module apiserver

go 1.24

require (
    github.com/lexkong/log v1.0.0
)

replace github.com/lexkong/log => github.com/zxmrlc/log
```

2. 执行命令

```bash
go mod tidy
```

3. 或者Linux/Mac命令

   
```bash
# -r (recursive) [递归查找]
# -l (list files) [只输出文件名]
# | 【管道】: 把前一个命令输出交给后一个命令
# xargs : 把输入内容变成命令参数,如把前面 grep 查找到的： ./main.go  ./config/log.go =>  sed xxx ./main.go ./config/log.go [一次修改多个文件]
# sed : 文本替换工具  -i [in-place]:直接修改文件,没有-i,只查找，不修改
# -i '' : 这是 macOS BSD sed 格式, Linux 通常: -i 
# 's#github.com/lexkong/log#github.com/zxmrlc/log#g': 替换表达式,  s#旧内容#新内容#g
#  s: (substitute) [替换]， g: (global) [全部替换], 如果没有g,只替换第一个
grep -rl "github.com/lexkong/log" . | xargs sed -i '' 's#github.com/lexkong/log#github.com/zxmrlc/log#g'
```
