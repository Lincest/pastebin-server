## deploy

```shell
set CGO_ENABLED=0 // powershell使用$env:CGO_ENABLED=0, 下同
set GOARCH=amd64
set GOOS=linux
go build -o main main.go
docker login -u .. -p ..
docker build -t pastebin-server .
docker tag xxx moreality/pastebin-server
docker push moreality/pastebin-server
```

测试: `docker run -p 9989:9989 -d moreality/pastebin-server`