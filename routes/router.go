package router

import(
	handler "APIs/controller"
	"github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
)

var router = echo.New()

func Router(){
	router.GET("/", handler.Show("APIs", "v1"))
}	

func Middleware(){
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
}

func Start(logger bool){
	verbose := router.Start(":9121")		
	router.Logger.Fatal(verbose)
}