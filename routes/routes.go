package routes

import (
	"go-rest-api/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
