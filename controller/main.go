package controller
import(
	echo "github.com/labstack/echo"
	"net/http"
)

func Show(title, version string) echo.HandlerFunc{
	return func(ctx echo.Context) error{
		collect := make(map[string]string)
		collect["title"] = title
		collect["version"] = version
	
		return ctx.JSON(http.StatusOK, collect)
	
	}
}