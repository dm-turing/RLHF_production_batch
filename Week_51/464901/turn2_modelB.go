package main

import (
	"fmt"
	"sync"
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

// Global balance variable (shared state)
var balance float64
var balanceMutex sync.Mutex

func concurrentDeposit(amount float64, wg *sync.WaitGroup) {
	defer wg.Done()
	// Protect concurrent access to balance using mutex
	balanceMutex.Lock()
	defer balanceMutex.Unlock()

	Deposit(amount, func(err error) {
		if err != nil {
			fmt.Println("Deposit failed:", err)
			// Rollback deposit if it fails
			balanceMutex.Lock()
			defer balanceMutex.Unlock()
			balance -= amount
			return
		}
		// Simulate success
		balance += amount
	}, func() {
		// Rollback deposit if transaction rollback is triggered
		balanceMutex.Lock()
		defer balanceMutex.Unlock()
		balance -= amount
	})
}

func concurrentWithdraw(amount float64, wg *sync.WaitGroup) {
	defer wg.Done()
	// Protect concurrent access to balance using mutex
	balanceMutex.Lock()
	defer balanceMutex.Unlock()

	Withdraw(amount, func(err error) {
		if err != nil {
			fmt.Println("Withdrawal failed:", err)
			// Rollback withdrawal if it fails
			balanceMutex.Lock()
			defer balanceMutex.Unlock()
			balance += amount
			return
		}
		// Simulate success
		balance -= amount
	}, func() {
		// Rollback withdrawal if transaction rollback is triggered
		balanceMutex.Lock()
		defer balanceMutex.Unlock()
		balance += amount
	})
}

func main() {
	var wg sync.WaitGroup
	depositAmount := 100.0
	withdrawAmount := 50.0

	// Add goroutines to the wait group to track completion
	wg.Add(2)
	go concurrentDeposit(depositAmount, &wg)
	go concurrentWithdraw(withdrawAmount, &wg)

	wg.Wait() // Wait for all transactions to complete

	// After all goroutines have finished, read and print the balance
	balanceMutex.Lock()
	fmt.Println("Final Balance: $", balance)
	balanceMutex.Unlock()
}
