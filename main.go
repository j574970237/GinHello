package main

import (
	_ "GinHello/docs"
	"GinHello/initRouter"
)

// @title Gin swagger
// @version 1.0
// @description Gin Swagger Learning

// @contact.name jack
// @contact.email 574970237@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()
}
