package main

import(
	route "APIs/routes"
)


func main(){	
	route.Router()
	route.Middleware()
	route.Start(true)
}