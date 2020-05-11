package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RateLimit      time.Duration
	WorkerCount    int
	UrlExpire      time.Duration
	SaveItemExpire time.Duration
	IsFull         bool
	RegistryAddr   string

	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	MetricsPort         int
	JaegerAgentAddr     string
	JaegerCollectorAddr string
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
	MaxIdle     int
	MaxOpen     int
}

var DatabaseSetting = &Database{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)

	ServerSetting.RateLimit = ServerSetting.RateLimit * time.Millisecond
	ServerSetting.UrlExpire = ServerSetting.UrlExpire * time.Minute
	ServerSetting.SaveItemExpire = ServerSetting.SaveItemExpire * time.Minute

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second

}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}
