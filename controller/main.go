package controller
import(
	"time"
	_"strconv"
	"net/http"
	"log"
	echo "github.com/labstack/echo"
	core "APIs/core"

)

func Show(title, version string) echo.HandlerFunc{
	core.Config()
	return func(ctx echo.Context) error{
		collect := make(map[string]string)
		collect["title"] = title
		collect["version"] = version
		t := time.Now()
		intoFiles := []byte("Testing Bug " + t.String() + "\n")		
		
		data, err := core.DEBUG("DEBUG_TEST" , intoFiles)
		
		if err != nil{
			log.Fatalf("Error Write Log")
		}

		log.Println(data)

		return ctx.JSON(http.StatusOK, collect)
		
	}
}