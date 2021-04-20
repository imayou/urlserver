### 说明
go+gin+redis短链接生成

### 编译
```
set GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

go mod tidy
go mod vendor

SET GOARCH=amd64
SET GOOS=linux
go build ./
```

### key生成
[https://zhuanlan.zhihu.com/p/59168096](https://zhuanlan.zhihu.com/p/59168096)