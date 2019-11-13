package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/cloudio_hw9/service"
	"github.com/spf13/pflag"
)

const INITPORT string = "8000"

func main() {
	port := pflag.StringP("port", "p", INITPORT, "Httpd listening port: ")
	pflag.Parse()

	server := service.NewServer() //加载静态文件系统服务
	server.Run(":" + *port)       //绑定spf13命令行获取的监听端口号

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	fmt.Println("awaiting signal")
	<-done

}
