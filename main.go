package main

import (
	"net/http"

	"github.com/adhish002/webservice/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
