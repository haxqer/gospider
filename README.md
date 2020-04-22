# Go-Spider
爬取 芒果TV 剧集
---
每分钟自动注入一次 seed, 根据 seed 爬取相关网页数据(广度优先)。

通过控制 url 去重时间(seed 也是 url), 控制采集频次。

## 功能描述
+ spider: 负责采集数据，从 etcd 获取 itemsave 服务列表，将采集的数据投递给 itemsave
+ itemsave: 注册服务至 etcd，接收 spider 投递的数据存入数据库中
+ basmicro: 工具包，etcd 可视化，查看 itemsave 服务状态
+ etcd: 注册中心


数据库 schema 在 `docs/sql` 中

默认情况下，日志文件在 `runtime` 下 


---
## 编译
### 爬虫编译
+ dev: go build
+ linux: env GOOS=linux go build

生成二进制文件 `spider` 在工作目录

### ItemSave 编译
+ dev: go build ./services/itemsave
+ linux: env GOOS=linux go build ./services/itemsave

生成二进制文件 `itemsave` 在工作目录

---
## 配置文件
配置文件位置: `conf/app.ini`

生成配置文件: `cp ./conf/app.dist.ini ./conf/app.init`

### 最小修改:

#### itemsave
只有 itemsave 需要数据库

database
+ User: 数据库用户名
+ Password: 数据库密码
+ Host: 数据库host+port
+ Name: 数据库名称
+ TablePrefix: 数据库表名前缀

server
+ RegistryAddr: 注册中心地址

#### spider
spider 无需配置数据库

server
+ RegistryAddr: 注册中心地址
+ UrlExpire: 控制采集频次，单位为 分钟

---
## 部署
1. 启动 `etcd`
2. 启动 `basmicro` 可视化服务
3. 启动 `itemsave` (可多台，支持水平扩展)
4. 启动 `spider` (目前建议一台，支持水平扩展)

### etcd

docker 启动方式
```bash
rm -rf /tmp/etcd-data.tmp && mkdir -p /tmp/etcd-data.tmp && \
  docker pull haxqer/etcd:v3.4.5 || true && \
  docker run \
  -p 2379:2379 \
  -p 2380:2380 \
  --mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \
  --name etcd-v3.4.5 \
  haxqer/etcd:v3.4.5 \
  /usr/local/bin/etcd \
  --name s1 \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-advertise-peer-urls http://0.0.0.0:2380 \
  --initial-cluster s1=http://0.0.0.0:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --log-level info \
  --logger zap \
  --log-outputs stderr

```

### basmicro
将 YOURIP 替换成 etcd 的 ip
`basmicro --registry=etcd --registry_address=YOURIP:2379 web`

### itemsave
修改配置文件的 `server`(只需配置注册中心地址即可) 和 `database`

启动成功后，在 `basmicro` 提供的 web 界面中能看到此服务

### spider
修改配置文件的 `server`

启动成功后，会向 `itemsave` 投递数据，每投递1万条，会写一次 log


