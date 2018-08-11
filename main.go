package main

import(
	"github.com/labstack/echo"
	handler "APIs/controller"
	middleware"github.com/labstack/echo/middleware"
	"APIs/core"
)


func main(){	
	engine := echo.New()

	//middleware
	engine.Use(middleware.Logger())
	engine.Use(middleware.Recover())

	//Stack Routing
	engine.GET("/", handler.Show("APIs", "v1"))
	

	//recover
	verbose := engine.Start(":9121")		
	engine.Logger.Fatal(verbose)
	core.NetListen("tcp", ":9121")
}