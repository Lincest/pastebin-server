## 服务器部署

`git clone https://github.com/Lincest/pastebin-server.git`

修改`./docker/`下的: 

- `config.example.yml` =>`config.yml`
- `redis.conf.example` => `redis.conf`

之后: 

```shell
cd docker
docker-compose up
```

配置开放`config.yml`配置的端口

## api

```yml
GET: /pastebin/:id

RESP: 
{
    "code": -1, # 0成功, -1失败
    "msg": "", # 信息
    "data": null # string / null
}
```

```yaml
POST: /pastebin

BODY: 
{
    "data": "hello this is me", # 字符串
    "expired": 1 # 过期天数
}

RESP: 
{
    "code": 0,
    "msg": "保存成功",
    "data": "779b647c" # key值
}
```

