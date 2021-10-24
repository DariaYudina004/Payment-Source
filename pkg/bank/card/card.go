package payment

import (
	"bank/pkg/bank/types"
)

func PaymentSource(cards []types.Card)[]types.PaymentSource{

	var source  []types.PaymentSource
	for _, card := range cards {
		
		if card.Balance<0 {
			continue
		}
		if !card.Active {
			continue
		}
		
		source=append(source ,types.PaymentSource{"card" , string(card.PAN),types.Money(card.Balance)})
	}
	
	return source 
}