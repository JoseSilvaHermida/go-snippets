package main

import (
	"fmt"

	"github.com/josesilvahermida/go-snippets/jshmodule"
	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

func main() {
	fmt.Println("Hi there!")
	fmt.Println("General Kenobi")

	jshmodule.Hello("Jose")

	var data = `
name: Hello
age: 43
`
	t := Yaml{}
	yaml.Unmarshal([]byte(data), &t)
	fmt.Printf("name: %s, age: %d\n", t.Name, t.Age)
}
