package routes

import (
	"go-rest-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonas).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.OnePersona).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreateNewPersona).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersona).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditPersona).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
