package jshmodule

import (
	"fmt"
)

type SuperType struct {
	Id  int
	In  chan interface{}
	Out chan interface{}
}

func Hello(name string) {
	fmt.Println("Hello " + name)
}
