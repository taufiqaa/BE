package routes

import (
	"waysbuck-API/handlers"
	"waysbuck-API/pkg/middleware"
	"waysbuck-API/pkg/mysql"
	"waysbuck-API/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", middleware.Auth(h.FindTransactions)).Methods("GET")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.GetTransaction)).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth((h.CreateTransaction))).Methods("POST")
	r.HandleFunc("/transaction_id", middleware.Auth(h.UpdateTransaction)).Methods("PATCH")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.DeleteTransaction)).Methods("DELETE")
	r.HandleFunc("/transaction-id", middleware.Auth(h.GetUserTransaction)).Methods("GET")
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}
