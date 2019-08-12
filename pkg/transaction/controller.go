package transaction

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
)

func addTransaction(w http.ResponseWriter, r *http.Request) {

	newTransaction := transaction{}
	
	reqBody,_ := ioutil.ReadAll(r.Body);
	
	err := json.Unmarshal(reqBody, &newTransaction)

	if err != nil  {
		http.Error(w,"Invalid input", http.StatusBadRequest)
		return
	}
	
	err = saveTransaction(&newTransaction)

	if err != nil {
		log.Printf("%v",err)
		err = fmt.Errorf("Something goes wrong tring to save the transaction: %v",err)
		http.Error(w,err.Error(), http.StatusBadRequest)
		return
	}

	response,_:= json.Marshal(newTransaction)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(response)
	return 
}

func getTransaction(w http.ResponseWriter, r *http.Request) {

	transactionID := mux.Vars(r)["id"]
	t,err := getTransactionFromHistory(transactionID)
	if err != nil  {
		log.Printf("%v",err)
		http.Error(w,"Something goes getting transaction", http.StatusBadRequest)
		return 
	}

	if t.Id == "" {
		w.Write([]byte("Transaction not found"))
		return
	}

	response,_:= json.Marshal(t)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(response)
	return 
}

func historyTransaction(w http.ResponseWriter, r *http.Request) {

	transactions, err := getHistory()

	if err != nil {
		log.Printf("%v",err)
		http.Error(w,"Something goes wrong trying to get the transaction history", http.StatusInternalServerError)
		return 
	}

	response,_:= json.Marshal(transactions)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	w.Write(response)
	return 

}


func accountState(w http.ResponseWriter, r *http.Request) {

	account := getAccountState()

	response,_:= json.Marshal(account)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(response)
	return 

}