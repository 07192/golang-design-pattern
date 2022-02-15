package decorator

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	pizza := Pizza{}
	tomatoPizza := TomatoPizza{pizza: pizza}
	beefPizza := BeefPizza{pizza: pizza}
	fmt.Println(pizza.GetPrice())
	fmt.Println(tomatoPizza.GetPrice())
	fmt.Println(beefPizza.GetPrice())
}
