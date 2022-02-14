package observer

import "fmt"

type People interface {
	Lunch()
}

type Chinese struct {
	Name string
}
type Japanese struct {
	Name string
}

func (c Chinese) Lunch() {
	fmt.Println(c.Name, "eat chicken")
}

func (j Japanese) Lunch() {
	fmt.Println(j.Name, "eat sushi")
}

type Company struct {
	Colleague []People
}

func (c *Company) Join(p People) {
	c.Colleague = append(c.Colleague, p)
}

func (c *Company) Afternoon() {
	for _, v := range c.Colleague {
		v.Lunch()
	}
}
