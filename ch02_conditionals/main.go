package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	// don't edit above this line

	finalCost = bulkMessageCost

	if isPremiumUser {
		finalCost = finalCost - (finalCost * discountRate)

		if (accountBalance < finalCost) {
			fmt.Println(insufficientFundMessage)
		} else {
			accountBalance = accountBalance - finalCost
			fmt.Println(purchaseSuccessMessage)
		}
	} else {
		if (accountBalance < finalCost) {
			fmt.Println(insufficientFundMessage)
		} else {
			accountBalance = accountBalance - finalCost
			fmt.Println(purchaseSuccessMessage)
		}
	}

	

	// don't edit below this line

	fmt.Println("Account balance:", accountBalance)
}
