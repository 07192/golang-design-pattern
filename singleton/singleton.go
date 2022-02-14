package singleton

import "sync"

var instance *People
var once sync.Once

type People struct {
	Name string
}

func GetInstance(name string) *People {
	once.Do(func() {
		instance = &People{Name: name}
	})
	return instance
}
