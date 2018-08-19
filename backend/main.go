package main

import (
	"flag"

	log "github.com/golang/glog"
	"github.com/lincolnzhou/check-in/backend/app"
	"github.com/lincolnzhou/check-in/backend/conf"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "c", "./config.toml", "set directory config file path")
}

func main() {
	flag.Parse()
	defer log.Flush()

	log.Info("check in system start")
	// 初始化配置
	err := conf.InitConfig(configFile)
	if err != nil {
		log.Errorf("NewConfig(\"%s\") error(%v)", configFile, err)
	}
	log.Infof("Version: %s", conf.ConfigData.Version)

	// 初始化HTTP服务
	app.InitRouter()
}
