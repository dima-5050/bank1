package stats

import (
	"fmt"
	"github.com/dima-5050/bank/pkg/bank/types"
)

func ExampleAvg() {
	
	payment := []types.Payment{
		{
			ID: 1,
			Amount: 1_000_00,
		},
		{
			ID: 2,
			Amount: 3_000_00,
		},
		{
			ID: 3,
			Amount: 4_000_00,
		},
	}

	fmt.Println(Avg(payment))
	//Output: 400000
	
}