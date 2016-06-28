package main

import (
	"fmt"
	"github.com/Arthelon/n10n/controllers"
	"github.com/Arthelon/n10n/models"
	"log"
	"net/http"
)

func main() {
	err := models.Init()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Listening on port %v\n", models.Conf.Port[1:])
	http.ListenAndServe(models.Conf.Port, controllers.GetRoutes())
}
