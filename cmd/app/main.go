package main

import (
	"flag"
	"fmt"
	"go-project-template/internal/config"
	"go-project-template/internal/handler"
	"go-project-template/internal/service"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var f = flag.String("f", "etc/application.yaml", "config file")
var version = "development"

func main() {
	fmt.Println("version:\t", version)
	flag.Parse()
	var c config.Config
	conf.MustLoad(*f, &c)

	ctx := service.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)

	defer server.Stop()

	handler.RegisterHandles(server, ctx)

	fmt.Printf("Starting server at %s:%d......\n", c.Host, c.Port)
	fmt.Println("this is new version-----------")
	server.Start()
	//envErr := godotenv.Load("etc/.env")
	//if envErr != nil {
	//	log.Fatal(envErr)
	//}
	//fmt.Println(os.Getenv("username"))
	//fmt.Println(os.Getenv("password"))
	//service.Test()
	////-----------------------------------------------
}
