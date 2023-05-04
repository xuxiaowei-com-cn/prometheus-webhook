package main

import (
	"flag"
	"fmt"
	"github.com/xuxiaowei-com-cn/prometheus-webhook/hello"
	"log"
	"net/http"
	"os"
)

var (
	Name        = "prometheus-webhook"
	Version     = "v0.0.1"
	Author      = "xuxiaowei-com-cn/prometheus-webhook: https://github.com/xuxiaowei-com-cn/prometheus-webhook"
	Email       = "徐晓伟 <xuxiaowei@xuxiaowei.com.cn>"
	Copyright   = "徐晓伟工作室 <xuxiaowei@xuxiaowei.com.cn>"
	Description = "普罗米修斯 Webhook"
)

func main() {

	// 信息
	info()

	var version bool
	var port string
	flag.BoolVar(&version, "version", false, "查看版本号")
	flag.StringVar(&port, "port", "", "参数中的端口")

	flag.Parse()

	if version {
		return
	}

	http.HandleFunc("/", hello.IndexHandler)

	if port == "" {
		port = os.Getenv("PORT")
		if port == "" {
			port = "8080"
			log.Printf("默认端口 %s", port)
		} else {
			log.Printf("环境中的端口 %s\n", port)
		}
	} else {
		log.Printf("参数中的端口 %s\n", port)
	}

	log.Printf("侦听端口 %s", port)
	log.Printf("http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

// 信息
func info() {
	fmt.Printf("Name：%s\n", Name)
	fmt.Printf("Version：%s\n", Version)
	fmt.Printf("Author：%s\n", Author)
	fmt.Printf("Email：%s\n", Email)
	fmt.Printf("Copyright：%s\n", Copyright)
	fmt.Printf("Description：%s\n", Description)
}
