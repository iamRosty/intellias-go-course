package main

import "fmt"

const (
	applePrice = 5.99
	pearPrice  = 7.
	myBudget   = 23.
)

func main() {
	fmt.Printf("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?\n")
	necessaryAmount := applePrice*9 + pearPrice*8
	fmt.Printf("Щоб купити 9 яблук та 8 груш, нам потрiбно %.2f\n", necessaryAmount)
	fmt.Printf("Скільки груш ми можемо купити?\n")
	numberPearsWeCanBuy := myBudget / pearPrice
	fmt.Printf("Ми можемо купити груш: %v шт.\n", int(numberPearsWeCanBuy))
	fmt.Printf("Скільки яблук ми можемо купити?\n")
	numberApplesWeCanBuy := myBudget / pearPrice
	fmt.Printf("Ми можемо купити яблук: %v шт.\n", int(numberApplesWeCanBuy))
	canBuy := myBudget >= 2*(pearPrice+applePrice) //Чи можемо ми купити 2 груші та 2 яблука?
	fmt.Printf("Можемо ми купити 2 груші та 2 яблука?: %t", canBuy)
}
