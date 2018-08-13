package config

import(
	env "github.com/joho/godotenv"
	"os"
	"log"
)

func Env() map[string]interface{}{
	
	set := make(map[string]interface{})
	
	err := env.Load()
	if err != nil{
		log.Fatal("Error Load env File")
	} 
	//set map
	set["host"] = os.Getenv("HOST")
	set["database"] = os.Getenv("DB")
	set["user"] = os.Getenv("USER")
	set["password"] = os.Getenv("PASSWORD")
	
	return set
}

