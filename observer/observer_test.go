package observer

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	var company Company
	tom := Chinese{"tom"}
	jerry := Japanese{"jerry"}

	company.Join(tom)
	company.Join(jerry)
	fmt.Println(company)
	company.Afternoon()
}
