package decorator

type pizza interface {
	GetPrice() int
}

type Pizza struct {
}

func (b Pizza) GetPrice() int {
	return 10
}

type TomatoPizza struct {
	pizza
}

func (tp TomatoPizza) GetPrice() int {
	return tp.pizza.GetPrice() + 5
}

type BeefPizza struct {
	pizza
}

func (bp BeefPizza) GetPrice() int {
	return bp.pizza.GetPrice() + 10
}
