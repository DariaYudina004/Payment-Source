package card 

import (
	"bank/pkg/bank/types"
	"fmt"
	
)


func ExamplePaymentSource()  {
	cards := []types.Card{
		{
			Balance: 99909,
			PAN: "1",
			Active: true,
		},
		{
			Balance: 99909,
			PAN: "2",
			Active: false,
		},
		{
			Balance: -99909,
			PAN: "3",
			Active: true,
		},
		{
			Balance: 1037909,
			PAN: "4",
			Active: true,
		},
	}
	card :=PaymentSource(cards)
	fmt.Println(card)
	//Output: 
	// [{card 1 99909} {card 4 1037909}]
	
}