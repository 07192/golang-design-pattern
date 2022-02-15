package simplefactory

import "fmt"

type Greater interface {
	Say(string)
}

func NewGreater(isShort bool) Greater {
	if isShort {
		return &hi{}
	}
	return &hello{}
}

type hi struct {
}

func (h hi) Say(name string) {
	fmt.Println("hi,", name, "!")
}

type hello struct {
}

func (h hello) Say(name string) {
	fmt.Println("hello,", name, "!")
}
