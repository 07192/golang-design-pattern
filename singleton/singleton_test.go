package singleton

import (
	"fmt"
	"testing"
)

func TestGetInstance(t *testing.T) {
	p1 := GetInstance("Tom")
	p2 := GetInstance("jerry")
	fmt.Println(p1)
	fmt.Println(p2)
}
