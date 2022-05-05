package main

import (
	"log"

	"github.com/terloo/myadmission/server"
)

func main() {
	err := server.RootCmd.Execute()
	if err != nil {
		log.Fatalln("服务启动失败：", err)
	}
}
