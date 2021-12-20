package main

import (
	"flag"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/qingants/gin-skeleton/cache"
	"github.com/qingants/gin-skeleton/model"
	"github.com/qingants/gin-skeleton/pkg/bininfo"
	"github.com/qingants/gin-skeleton/pkg/log"
	"github.com/qingants/gin-skeleton/pkg/utils"
	"github.com/qingants/gin-skeleton/router"
	"github.com/qingants/gin-skeleton/setting"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"os"
	"time"
)

var (
	version  bool
	filename string
)

func Usage() {
	flag.BoolVar(&version, "v", false, "version")
	flag.BoolVar(&version, "version", false, "version")
	flag.StringVar(&filename, "f", "./conf.ini", "configuration file")
	flag.StringVar(&filename, "conf", "./conf.ini", "configuration file")

	flag.Parse()
	if version || filename == "" {
		fmt.Printf("%s", bininfo.PrettyVersion())
		os.Exit(0)
	}
}

// @title gin-skeleton
// @version 0.0.1
// @description  Go Gin Skeleton
func main() {
	// Usage
	Usage()

	// Parser Conf
	if err := setting.Open(filename); err != nil {
		fmt.Printf("Parser conf: %s err: %v", filename, err)
		return
	}

	// Open Error monitoring
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:         setting.SentryDsn,
		Environment: utils.GetEnvironment(),
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
	defer sentry.Flush(1 * time.Second)

	// Setup log
	log.Open(setting.LoggerCfg.Filename, zapcore.Level(setting.LoggerCfg.Level),
		setting.LoggerCfg.MaxSize, setting.LoggerCfg.MaxBackups,
		setting.LoggerCfg.MaxAge, setting.LoggerCfg.Compress)
	defer log.Close()

	// Connect all cache
	cache.Open()
	defer cache.Close()

	// Connect all database
	model.Open()
	defer model.Close()

	r := router.SetupRouter()

	s := &http.Server{
		Addr:           setting.HTTPAddr,
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	zap.S().Infof("listening on %v", setting.HTTPAddr)

	if err := s.ListenAndServe(); err != nil {
		zap.L().Error(err.Error())
	}
}
