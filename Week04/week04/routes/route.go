package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"week04/controllers"

)

var App *iris.Application

func InitRoute() {
	App = iris.New()

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposedHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	})

	App.Get("/kpi/ping", controllers.TestPing)
	kpiSns := App.Party("/kpi/", crs).AllowMethods(iris.MethodOptions) // <- important for the preflight.
	{
		kpiSns.Post("/ping1", controllers.TestPing)
	}



}
