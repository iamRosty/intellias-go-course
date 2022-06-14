package main

import "fmt"

const (
	applePrice   = 5.99
	pearPrice    = 7.
	myBudget     = 23.
	numberApples = 9
	numberPears  = 8
)

func main() {
	fmt.Println("Скільки грошей треба витратити, щоб купити 9 яблук та 8 груш?")
	necessaryAmount := applePrice*numberApples + pearPrice*numberPears
	fmt.Println("Щоб купити 9 яблук та 8 груш, нам потрiбно", necessaryAmount)
	fmt.Println("Скільки груш ми можемо купити?")
	numberPearsWeCanBuy := myBudget / pearPrice
	fmt.Printf("Ми можемо купити груш: %v шт.\n", int(numberPearsWeCanBuy))
	fmt.Println("Скільки яблук ми можемо купити?")
	numberApplesWeCanBuy := myBudget / pearPrice
	fmt.Printf("Ми можемо купити яблук: %v шт.\n", int(numberApplesWeCanBuy))
	canBuy := myBudget >= 2*(pearPrice+applePrice) //Чи можемо ми купити 2 груші та 2 яблука?
	fmt.Printf("Можемо ми купити 2 груші та 2 яблука?: %t", canBuy)
}
