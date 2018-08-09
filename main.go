package main

import(
 	"github.com/labstack/echo"
	handler "APIs/controller"
	middleware"github.com/labstack/echo/middleware"
)

func main(){
	engine := echo.New()

	//middleware
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())

	//Stack Routing
	engine.GET("/", handler.Show("APIs", "v1"))

	//Logger
	engine.Logger.Fatal(engine.Start(":1234"))
}