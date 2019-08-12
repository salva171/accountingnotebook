package transaction

import (
	"github.com/gorilla/mux"
	"net/http"
)


func Router() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/transaction/add", addTransaction).Methods("POST")
	r.HandleFunc("/transaction/get/{id}", getTransaction).Methods("GET")
	r.HandleFunc("/transaction/history", historyTransaction).Methods("GET")
	r.HandleFunc("/account/state", accountState).Methods("GET")
	return r
}



