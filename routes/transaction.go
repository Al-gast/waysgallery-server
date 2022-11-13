package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/mysql"
	"waysgallery/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/my-transactions", middleware.Auth(h.ShowTransaction)).Methods("GET")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.GetTransactionByID)).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	r.HandleFunc("/transaction/{id}", middleware.Auth(middleware.UploadPost(middleware.UploadPost2(middleware.UploadPost3(middleware.UploadPost4(middleware.UploadPost5(h.UpdateTransaction))))))).Methods("PATCH")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.DeleteTransaction)).Methods("DELETE")
}
