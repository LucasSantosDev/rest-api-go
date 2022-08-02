package routes

import (
	"go-rest-api/controllers"
	"go-rest-api/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)

	r.Use(middleware.SetContentTypeMiddleware)
	r.HandleFunc("/api/personalities", controllers.ListPersonalities).Methods("Get")
	r.HandleFunc("/api/personalities/{id}", controllers.ShowPersonalitie).Methods("Get")
	r.HandleFunc("/api/personalities", controllers.CreatePersonalitie).Methods("Post")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonalitie).Methods("Delete")
	r.HandleFunc("/api/personalities/{id}", controllers.UpdatePersonalitie).Methods("Put")

	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
