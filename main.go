package main

import (
	"fmt"
	"net/http"
	"github.com/salva171/accountingnotebook/pkg/transaction"
)


func main() {
	fmt.Println("Server start, listening on port 8080")
	fmt.Println(http.ListenAndServe(":8080",transaction.Router()))
}