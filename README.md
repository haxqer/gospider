# Go-Spider

爬取 芒果TV 剧集

---
每分钟自动注入一次 seed, 根据 seed 爬取相关网页数据(广度优先)。

通过控制 url 去重时间(seed 也是 url), 控制采集频次。

---
基于 go-micro 微服务框架
(引用包中适配 freebsd 时使用了一些 C 代码，CGo 和 跨平台编译不共存，暂时只能提供 linux 版本的二进制文件)

## 功能描述
+ spider: 负责采集数据，从 etcd 获取 itemsave 服务列表，将采集的数据投递给 itemsave (grpc)
+ itemsave: 注册服务至 etcd，接收 spider 投递的数据存入数据库(mysql)中
+ micro: 第三方工具包，etcd 可视化，查看 itemsave 服务状态 (https://github.com/micro/micro/releases 建议 v2.4.0)
+ etcd: 注册中心 (github.com/etcd-io/etcd 建议 v3.4.x)
+ spiderhttp: 手工采集接口，提供 http 接口(后续可能提供 grpc 接口)，解决紧急采集剧集的需求(自动采集有一定的延迟)
+ prometheus: 应用程序状态监控(https://github.com/prometheus/prometheus/releases 建议 v2.18.0)
+ grafana: 可视化 (https://github.com/grafana/grafana/releases 建议 v6.7.3)


数据库 schema 在 `docs/sql` 中

grafana 配置在 `docs/grafana` 中

默认情况下，日志文件在 `runtime` 下 


---
## 编译
### 爬虫编译
+ dev: `go build`
+ linux: `env GOOS=linux go build`

生成二进制文件 `spider` 在工作目录

### ItemSave 编译
+ dev: `go build ./services/itemsave`
+ linux: `env GOOS=linux go build ./services/itemsave`

生成二进制文件 `itemsave` 在工作目录

### SpiderHttp 编译
+ dev: `go build ./services/spiderhttp`
+ linux: `env GOOS=linux go build ./services/spiderhttp`

生成二进制文件 `spiderhttp` 在工作目录

---
## 配置文件
配置文件位置: `conf/app.ini`

生成配置文件: `cp ./conf/app.dist.ini ./conf/app.ini`

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
+ MetricsPort: expose Prometheus metrics (注意端口不要和其他服务相同)

#### spider
spider 无需配置数据库

server
+ RegistryAddr: 注册中心地址
+ UrlExpire: 控制采集频次，单位为 分钟
+ MetricsPort: expose Prometheus metrics (注意端口不要和其他服务相同)


#### spiderhttp
spiderhttp 无需配置数据库

server
+ RegistryAddr: 注册中心地址
+ RunMode: 正式环境改为 `release`
+ HttpPort: http 服务的端口
+ MetricsPort: expose Prometheus metrics (注意端口不要和其他服务相同)

---
## 部署
1. 启动 `etcd`
2. 启动 `micro` 可视化服务
3. 启动 `itemsave` (可多台，支持水平扩展)
4. 启动 `spider` (目前建议一台，支持水平扩展)
5. 启动 `spiderhttp` (目前建议一台，支持水平扩展)
6. 启动 `prometheus` 采集 go 应用程序数据
7. 启动 `grafana` 展示 prometheus 采集的数据

### etcd

参考: https://github.com/etcd-io/etcd/releases

### micro
将 YOURIP 替换成 etcd 的 ip
`micro --registry=etcd --registry_address=YOURIP:2379 web`

### itemsave
修改配置文件的 `server`(只需配置注册中心地址即可) 和 `database`

启动成功后，在 `micro` 提供的 web 界面中能看到此服务

### spider
修改配置文件的 `server`

启动成功后，会向 `itemsave` 投递数据，每投递1万条，会写一次 log(log 无需保存)

### spiderhttp
修改配置文件的 `server` (只需配置注册中心地址和http配置)

启动成功后，接收http请求，采集指定芒果剧集，向 `itemsave` 投递数据

### prometheus

采集配置

prometheus.yml

```yaml
scrape_configs:
- job_name: itemsave-01
  scrape_interval: 10s
  static_configs:
  - targets:
    - localhost:12112

- job_name: itemsave-02
  scrape_interval: 10s
  static_configs:
  - targets:
    - localhost:12113

- job_name: spider-01
  scrape_interval: 10s
  static_configs:
  - targets:
    - localhost:12121

- job_name: spiderhttp-01
  scrape_interval: 10s
  static_configs:
  - targets:
    - localhost:12131

```

### grafana
连接 prometheus 后，导入 `docs/grafana` 中的配置 



---
## 数据流向(用途描述)
数据流向:
1. spider 采集，投递给 itemsave
2. itemsave 写入 mysql
3. cpanel 上的脚本定时读取 mysql 数据生成索引 (用于定向剧集), 写入 redis
4. dsp 读取 redis 中的剧集索引，指导广告投放

http服务:
1. 接收请求中的 mgtv_url
2. 采集剧集信息，投递给 itemsave

---
## micro web
micro
![micro](docs/images/micro-services.png)

itemsave
![itemsave](docs/images/itemsave.png)


 
---
## todo
- [x] etcd 服务注册/发现
- [x] micro 健康检查
- [x] prometheus 监控
- [ ] jaeger/zipkin 链路追踪
- [x] hystrix 熔断器 (spider 使用)
- [x] uber.ratelimiter 限流 (itemsave spiderhttp 使用)
- [ ] Testing: Unit Testing, Behavior Testing, Integration Testing
- [x] 提供手工采集的接口
- [x] itemsave 接口 (grpc => api)
- [ ] error 日志收集
- [ ] channel 数据长度监控


## 异步消费 
(已尝试，暂时不启用)
- [x] rabbitmq/redis broker 


