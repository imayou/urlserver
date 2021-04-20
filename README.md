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

### key生成参考
[https://zhuanlan.zhihu.com/p/59168096](https://zhuanlan.zhihu.com/p/59168096)

### Dockerfile
```
FROM  busybox:stable-glibc
RUN mkdir /app
WORKDIR /app
COPY main /app
# COPY conf.yaml /app
ENTRYPOINT ["./main"]
```
### 部署
```
docker rm -f urlserver
docker rmi ysy/urlserver
rm -f main
mv urlserver main
chmod +x main

docker build -t ysy/urlserver:latest .

docker run -it \
-v /etc/localtime:/etc/localtime \
--net nets \
--name=urlserver \
--restart=unless-stopped \
-p 8811:8811 \
 ysy/urlserver:latest;
```
