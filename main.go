package main

import (
	"go-gin-template/app/router"
)

func main() {
	//挂载路由
	router.InitRouter()

	//todo 对接nsq
	//middlewares.Consumer()
	//middlewares.Producer()
}
