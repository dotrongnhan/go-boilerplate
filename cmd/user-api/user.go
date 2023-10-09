package main

import (
	"flag"
	"fmt"

	"git.savvycom.vn/go-services/go-boilerplate/cmd/user-api/internal/config"
	"git.savvycom.vn/go-services/go-boilerplate/cmd/user-api/internal/handler"
	"git.savvycom.vn/go-services/go-boilerplate/cmd/user-api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
