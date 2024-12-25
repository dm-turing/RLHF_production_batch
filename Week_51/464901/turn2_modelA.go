package main

import (
	"fmt"
	"sync"
	"time"
)

// Define the types for the callback and rollback functions
type CallbackFunc func(err error)
type RollbackFunc func()

var sharedData float64
var mu sync.Mutex

// Deposit money into an account
func Deposit(amount float64, callback CallbackFunc, rollback RollbackFunc) {
	mu.Lock()
	defer mu.Unlock()

	sharedData += amount
	fmt.Println("Depositing $", amount, "; New balance:", sharedData)

	// Simulate a transaction that could fail
	if amount < 0 {
		callback(fmt.Errorf("Invalid deposit amount: %f", amount))
	} else {
		callback(nil)
	}
}

// Withdraw money from an account
func Withdraw(amount float64, callback CallbackFunc, rollback RollbackFunc) {
	mu.Lock()
	defer mu.Unlock()

	if sharedData < amount {
		callback(fmt.Errorf("Insufficient funds"))
	} else {
		sharedData -= amount
		fmt.Println("Withdrawing $", amount, "; New balance:", sharedData)
		callback(nil)
	}
}

// Rollback function to revert transactions
func RollbackTransactions(previousError error, rollbacks ...RollbackFunc) {
	fmt.Println("Rolling back transactions due to:", previousError)
	for _, rollback := range rollbacks {
		rollback()
	}
}

// Do chained transactions with potential rollback
func doChainedTransactions(depositAmount, withdrawAmount float64, complete CallbackFunc) {
	rollbackDeposit := func() {
		mu.Lock()
		defer mu.Unlock()
		sharedData -= depositAmount
		fmt.Println("Rolling back deposit; New balance:", sharedData)
	}
	rollbackWithdraw := func() {
		mu.Lock()
		defer mu.Unlock()
		sharedData += withdrawAmount
		fmt.Println("Rolling back withdrawal; New balance:", sharedData)
	}

	go Deposit(depositAmount, func(err error) {
		if err != nil {
			RollbackTransactions(err, rollbackDeposit)
			return
		}

		go Withdraw(withdrawAmount, func(err error) {
			if err != nil {
				RollbackTransactions(err, rollbackWithdraw, rollbackDeposit)
				return
			}

			complete(nil)
		}, rollbackWithdraw)
	}, rollbackDeposit)
}

func main() {
	time.Sleep(500 * time.Millisecond)

	doChainedTransactions(100.0, 50.0, func(err error) {
		if err != nil {
			fmt.Println("Final Error:", err)
		} else {
			fmt.Println("All transactions completed successfully")
		}
	})

	doChainedTransactions(200.0, 300.0, func(err error) {
		if err != nil {
			fmt.Println("Final Error:", err)
		} else {
			fmt.Println("All transactions completed successfully")
		}
	})

	time.Sleep(2 * time.Second) // Wait for transactions to finish
	fmt.Println("Final shared data balance:", sharedData)
}
