package simplefactory

import "testing"

func TestSimpleFactory(t *testing.T) {
	greater := NewGreater(true)
	greater.Say("Tom")
	greater2 := NewGreater(false)
	greater2.Say("Jerry")
}
