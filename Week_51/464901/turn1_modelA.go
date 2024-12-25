package main

import (
	"fmt"
)

// Define the types for the callback and rollback functions
type CallbackFunc func(err error)
type RollbackFunc func()

// Deposit money into an account
func Deposit(amount float64, callback CallbackFunc, rollback RollbackFunc) {
	fmt.Println("Depositing $", amount)
	// Simulate a transaction that could fail
	if amount < 0 {
		callback(fmt.Errorf("Invalid deposit amount: %f", amount))
	} else {
		callback(nil)
	}
}

// Withdraw money from an account
func Withdraw(amount float64, callback CallbackFunc, rollback RollbackFunc) {
	fmt.Println("Withdrawing $", amount)
	// Simulate a transaction that could fail
	if amount < 0 {
		callback(fmt.Errorf("Invalid withdrawal amount: %f", amount))
	} else {
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
		fmt.Println("Rolling back deposit")
	}
	rollbackWithdraw := func() {
		fmt.Println("Rolling back withdrawal")
	}

	Deposit(depositAmount, func(err error) {
		if err != nil {
			RollbackTransactions(err, rollbackDeposit)
			return
		}

		Withdraw(withdrawAmount, func(err error) {
			if err != nil {
				RollbackTransactions(err, rollbackWithdraw, rollbackDeposit)
				return
			}

			complete(nil)
		}, rollbackWithdraw)
	}, rollbackDeposit)
}

func main() {
	doChainedTransactions(100.0, 50.0, func(err error) {
		if err != nil {
			fmt.Println("Final Error:", err)
		} else {
			fmt.Println("All transactions completed successfully")
		}
	})
}
