module git.trac.cn/nv/spider

go 1.14

require (
	contrib.go.opencensus.io/exporter/jaeger v0.2.0
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/gin-gonic/gin v1.6.2
	github.com/go-ini/ini v1.55.0
	github.com/go-playground/validator/v10 v10.2.0
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/golang/protobuf v1.4.1
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro/v2 v2.5.0
	github.com/micro/go-plugins/wrapper/breaker/hystrix/v2 v2.5.0
	github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2 v2.5.0
	github.com/micro/go-plugins/wrapper/trace/opencensus/v2 v2.5.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pquerna/ffjson v0.0.0-20190930134022-aa0246cd15f7
	github.com/prometheus/client_golang v1.6.0
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.5.1
	go.opencensus.io v0.22.3
	golang.org/x/net v0.0.0-20200506145744-7e3656a0809f
	golang.org/x/sys v0.0.0-20200501145240-bc7a7d42d5c3 // indirect
	golang.org/x/text v0.3.2
)
