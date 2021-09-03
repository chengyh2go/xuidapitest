package main


import (
	"apiservertest/defs"
	"apiservertest/handler"
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)

var (
	Help bool
	Port string
	XuidLength int
)

func init() {
	flag.BoolVar(&Help, "help", false, "xuid查询使用帮助")
	flag.StringVar(&Port, "port", "8080","指定监听端口")
	flag.IntVar(&XuidLength, "xuidLength", 10,"指定生成xuid的长度")
	flag.Usage = usage
}

func usage() {
	_, _ = fmt.Fprintf(os.Stdout, `xuid query api server: v1.0
Options:
`)
	flag.PrintDefaults()
}

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/xuid",handler.XuidHandler)
	router.GET("/",handler.HomeHandler)
	return router
}

func main() {

	//解析命令行参数
	flag.Parse()

	if Help {
		flag.Usage()
	} else {
		err := defs.GenerateGlobalConfig()
		defs.GlobalConfig.Port = Port
		defs.GlobalConfig.XuidLength = XuidLength
		if err != nil {
			fmt.Println("创建配置失败，程序退出!")
			os.Exit(-1)
		}
		r := RegisterHandlers()
		log.Fatal(http.ListenAndServe(":"+defs.GlobalConfig.Port ,r))
	}

}