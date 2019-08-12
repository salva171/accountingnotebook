package transaction


import (
	"os"
	"fmt"
	"log"
	"time"
	"sync"
	"bufio"
	"encoding/json"
	"github.com/salva171/accountingnotebook/pkg/common"
)

type transaction struct {
	Id string
	Type string
	Amount float64
	AccountAmount float64
	EffectiveDate string
}

type account struct {
	Amount float64
}

var currentAccountState account
var syncAccessMemory sync.WaitGroup


var validTypesTransaction = []string{"credit","debit"}


func init() {

	accountData,err := getCurrentStateAccount();

	if err != nil {
		log.Fatal(err)
	}

	currentAccountState = accountData

}



func validateTransaction(t *transaction,accountData account) (bool,error) {


	if t.Type == "" || !common.FindElementInArray(&validTypesTransaction,t.Type) {
		return false,fmt.Errorf("Invalid transaction type %s\n",t.Type)
	}
	

	if !validAccountOperation(t.Type,accountData.Amount,t.Amount) {
		return false,fmt.Errorf("Invalid account operation");		
	}

	return true,nil

}



func getCurrentStateAccount() (account,error) {

	accountState := account{}

	exists,err := common.ExistFile("pkg/transaction/storage/account_state.json")

	if err != nil {
		return accountState,err
	}

	if !exists {
		return accountState,nil
	}


	accountFile, err := os.Open("pkg/transaction/storage/account_state.json")
	
	defer accountFile.Close()
	
	if err != nil {
		return accountState,err
	}

	jsonParser := json.NewDecoder(accountFile)

	err = jsonParser.Decode(&accountState)

	if err != nil {
		return accountState,err
	}

	return accountState,nil
}


func saveTransaction(t *transaction) error {
	//Block
	syncAccessMemory.Wait()
	syncAccessMemory.Add(1)

	defer func() {
		//Unblock
		syncAccessMemory.Done()
	}()



	validt, err := validateTransaction(t,currentAccountState)

	if !validt {
		return err
	}

	if t.Type == "debit" {
		currentAccountState.Amount = currentAccountState.Amount-t.Amount
	} else {
		currentAccountState.Amount = currentAccountState.Amount+t.Amount
	}

	t.Id = common.RandStringBytes(10)
	currentTime  := time.Now() 
	t.EffectiveDate = currentTime.Format("2006-01-02T15:04:05-0700")
	t.AccountAmount = currentAccountState.Amount


	f , err := os.OpenFile("pkg/transaction/storage/account_state.json",os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)

	defer f.Close()

	if err != nil {
		return err
	}

	data, _ := json.Marshal(t)
	adata, _ := json.Marshal(currentAccountState)

	_, err = f.WriteString(string(adata))
	
	if err != nil {
        return err
	}

	err  = appendToHistory(string(data));
	
	if err != nil {
		return err
	}

	return nil
}


func appendToHistory(data string) error {

	f, err := os.OpenFile("pkg/transaction/storage/transaction_history",os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)


	defer f.Close()

	if err != nil {
		return err
	}

	f.WriteString(data+"\n")

	return nil
}


func validAccountOperation(typeOperation string,currentAmount float64,amount float64) bool {
	if typeOperation == "debit" {
		newAmount := currentAmount-amount
		return newAmount >= 0
	} 
	return true
}



func getHistory() ([]transaction,error) {

	syncAccessMemory.Wait()
	f, err := os.OpenFile("pkg/transaction/storage/transaction_history",os.O_CREATE|os.O_RDONLY, 0644)
	
	defer f.Close()

	if err != nil {
		return nil,err
	}
	
	scanner := bufio.NewScanner(f)

	history := make([]transaction,0,100)

	for scanner.Scan() {

		t := transaction{}
		tData := scanner.Text();
		_ = json.Unmarshal([]byte(tData),&t)
		
		history = append(history,t)
	}

	return history,nil
}



func getTransactionFromHistory(transactionID string) (transaction,error) {

	syncAccessMemory.Wait()
	f, err := os.OpenFile("pkg/transaction/storage/transaction_history",os.O_CREATE|os.O_RDONLY, 0644)
	defer f.Close()

	if err != nil {
		return transaction{},err
	}
	
	scanner := bufio.NewScanner(f)


	for scanner.Scan() {
		t := transaction{}
		tData := scanner.Text();
		_ = json.Unmarshal([]byte(tData),&t)
		
		if t.Id == transactionID {
			return t,nil
		}
	}

	return transaction{},nil
}


func getAccountState() (account) {
	return currentAccountState
}