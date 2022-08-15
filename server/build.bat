
set CGO_ENABLED=0
set GOARCH=amd64
set GOOS=linux
go build -o main main.go
docker build -t server .
docker tag server moreality/pastebin-server
docker push moreality/pastebin-server
