package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/mysql"
	"waysgallery/repositories"

	"github.com/gorilla/mux"
)

func UserRoutes(r *mux.Router) {
	userRepository := repositories.RepositoryUser(mysql.DB)
	h := handlers.HandlerUser(userRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.Auth(h.GetUser)).Methods("GET")
	r.HandleFunc("/user/{id}", middleware.Auth(middleware.UploadBestArt(middleware.UploadImage(h.UpdateUser)))).Methods("PATCH")
}
