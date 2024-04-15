package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"

	"github.com/ch3nnn/goview-gozero/restful/internal/config"
	"github.com/ch3nnn/goview-gozero/restful/internal/handler"
	"github.com/ch3nnn/goview-gozero/restful/internal/svc"
)

var configFile = flag.String("f", "etc/restful.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// static file
	// https://github.com/zeromicro/zero-examples/blob/main/http/static/main.go#L22
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/static/:file",
		Handler: http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))).ServeHTTP,
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
