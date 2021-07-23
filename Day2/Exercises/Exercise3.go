package main

import (
	"errors"
	"fmt"
	"sync"
)

type Account struct {
	mutex   sync.Mutex
	balance float64
}

func (a *Account) Deposit(amount float64) Response {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.balance += amount
	return Response{a.balance, nil}
}

func (a *Account) Withdraw(amount float64) Response {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if amount > a.balance {
		return Response{a.balance, errors.New("insufficient balance")}
	}
	a.balance -= amount
	return Response{a.balance, nil}
}

func (a *Account) GetBalance() Response {
	return Response{a.balance, nil}
}

type Transaction struct {
	account         *Account
	transactionType string
	amount          float64
}

type Response struct {
	balance float64
	error   error
}

func ReceiveTransaction(txnChan chan Transaction, respChan chan Response) {
	defer close(respChan)
	for txn := range txnChan {
		switch txn.transactionType {
		case "Deposit":
			respChan <- txn.account.Deposit(txn.amount)
		case "Withdraw":
			respChan <- txn.account.Withdraw(txn.amount)
		case "Balance":
			respChan <- txn.account.GetBalance()
		}
	}
}

func main() {
	acc := Account{balance: 500}
	txnChan := make(chan Transaction)
	respChan := make(chan Response)

	transactions := []Transaction{
		{&acc, "Deposit", 22},
		{&acc, "Deposit", 42},
		{&acc, "Balance", 0},
		{&acc, "Withdraw", 22},
		{&acc, "Withdraw", 1000},
	}

	go ReceiveTransaction(txnChan, respChan)
	var responses []Response
	done := make(chan struct{})
	go func() {
		for resp := range respChan {
			responses = append(responses, resp)
		}
		done <- struct{}{}
	}()

	for _, txn := range transactions {
		txnChan <- txn
	}
	close(txnChan)

	<-done
	fmt.Println(responses)
}
