package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/mysql"
	"waysgallery/repositories"

	"github.com/gorilla/mux"
)

func PostRoutes(r *mux.Router) {
	postRepository := repositories.RepositoryPost(mysql.DB)
	h := handlers.HandlerPost(postRepository)

	r.HandleFunc("/posts", middleware.Auth(h.FindPosts)).Methods("GET")
	r.HandleFunc("/posts/{id}", middleware.Auth(h.FindPostsByUserId)).Methods("GET")
	r.HandleFunc("/post/{id}", middleware.Auth(h.GetPost)).Methods("GET")
	r.HandleFunc("/post", middleware.Auth(middleware.UploadPost(middleware.UploadPost2(middleware.UploadPost3(middleware.UploadPost4(middleware.UploadPost5(h.CreatePost))))))).Methods("POST")
}
