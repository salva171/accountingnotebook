# Accounting notebook API


## Getting Started

```
git clone https://github.com/salva171/accountingnotebook.git
```



To run the project in your local machine first you have to make sure you have the port 8080 available.


Go to the root directory of the project and run the binary with the next command 

```
cd /home/user/accountingnotebook/
./accountingnotebook
```

IMPORTANT:if you dont find the binary 'accountingnotebook' you have to install go in your machine and after that follow the next steps.

Go to the root directory of the project and run go build command

Example.

```
cd /home/user/accountingnotebook/
go build
```

Then run the binary generated in the last step like the follow example.


```
./accountingnotebook
```


#Endpoints

Here is a list with the available endpoints.


* /transaction/add  
* /transaction/get/{id}
* /transaction/history
* /account/state



With '/transaction/add' you can add a new transaction (debit or credit) to the client account like the following example.

Example
```
path: http://localhost:8080/transaction/add
method: "POST"
Data: 
{
    "type": "credit",
    "amount": 200
}


```

With '/transaction/history' you can get the complete history of transactions for the user account

Example
```
path: http://localhost:8080/transaction/history
method: "GET"

```


With '/transaction/get/{id}' you can get the complete data from a specific transaction.

Example
```
path: http://localhost:8080/transaction/get/MRAjWwhTHc
method: "GET"

```


With '/account/state' you can get the current amount from the client account.

Example
```
path: http://localhost:8080/account/state
method: "GET"

```




