package main

import (
	"github.com/kataras/iris/v12"
	"week04/routes"
	"week04/utils"
)

func init() {
	utils.InitMongo("mongo")
	routes.InitRoute()
}

func main() {

	routes.App.Run(
		iris.Addr(utils.Conf.App.Port),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	)
}
