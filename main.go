package main

import (
	"bk_analysis/app"
	arango "bk_analysis/arangodb"
	"bk_analysis/config"
	"crypto/tls"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/yangioc/bk_pack/log"
)

var configPath = flag.String("config", "./env.yaml", "specific config to processing")

func main() {
	if err := config.Init(*configPath); err != nil {
		panic(err)
	}

	arango.LaunchInstans(config.EnvInfo.Arango.Addr, config.EnvInfo.Arango.Username, config.EnvInfo.Arango.Password, config.EnvInfo.Arango.DataBase)

	// 設定 log
	log.Level = log.Level_Info // 預設
	if logLevel, ok := log.LevelToStringMap[config.EnvInfo.Log.Level]; ok {
		log.Level = logLevel
	}

	// 測試用設定
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	// messageq 介面
	// type1
	// handle_messageq := msg_nats.New(context.TODO(), *config.EnvInfo, nil)

	// 核心服務
	// handle_app := app.New(*config.EnvInfo, handle_messageq)
	// handle_app.Launch()
	handle_app := app.New(*config.EnvInfo)

	//////// test //////////////
	handle_app.AddStocId("3481")
	handle_app.AddStocId("3711")
	handle_app.AddStocId("5283")
	handle_app.AddStocId("6505")
	handle_app.AddStocId("8046")
	handle_app.RunCloseData()
	////////////////////

	log.Info("Service Up.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	<-c

	log.Info("Service Down.")
}
