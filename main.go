package httpServer

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"

)

func main(){
	e := echo.New()
	err := e.Start(fmt.Sprintf(":%d", 8080))
	if err != nil{
		log.Fatal(err)
	}
}