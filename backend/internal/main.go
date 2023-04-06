// Code generated by hertz generator.

package main

import (
	"context"
	"github.com/wego2023/weflow/internal/biz/handler/hello"
	"github.com/wego2023/weflow/pkg/parser"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	hertzlogrus "github.com/hertz-contrib/logger/logrus"
	"github.com/hertz-contrib/pprof"
	"github.com/wego2023/weflow/internal/biz/router"
	"github.com/wego2023/weflow/internal/conf"
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	// init dal
	// dal.Init()
	address := conf.GetConf().Hertz.Address
	h := server.New(server.WithHostPorts(address))
	h.Use(hello.GlobalErrorHandler)
	// add a ping route to test
	h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
		parser.FlowParserServiceImpl.Test()
		ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
	})

	router.GeneratedRegister(h)
	// do what you wanted
	// add some render data: <no value>

	registerMiddleware(h)

	h.Spin()
}

func registerMiddleware(h *server.Hertz) {
	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}
	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// log
	logger := hertzlogrus.NewLogger()
	hlog.SetLogger(logger)
	hlog.SetLevel(conf.LogLevel())
	hlog.SetOutput(&lumberjack.Logger{
		Filename:   conf.GetConf().Hertz.LogFileName,
		MaxSize:    conf.GetConf().Hertz.LogMaxSize,
		MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
		MaxAge:     conf.GetConf().Hertz.LogMaxAge,
	})

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())
}