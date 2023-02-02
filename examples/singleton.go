package main

import "sync"

type Singleton struct{}

var (
	ins  *Singleton
	once sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		ins = &Singleton{}
	})
	return ins
}
