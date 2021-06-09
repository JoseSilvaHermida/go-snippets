package main

import (
	"fmt"
	"time"

	"github.com/josesilvahermida/go-snippets/jshmodule"
	"github.com/josesilvahermida/go-snippets/workers"
	"gopkg.in/yaml.v2"
)

type Yaml struct {
	Name string `yaml:"name"`
	Age  int    `yaml:"age"`
}

var WORKERS workers.WorkerList

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

	st := jshmodule.SuperType{
		Id:  1,
		In:  make(chan interface{}),
		Out: make(chan interface{}),
	}
	fmt.Printf("ID: %d\n", st.Id)

	WORKERS = workers.NewWorkerList()

	tasks := make(chan string, 1)

	for _, t := range []string{"sequential", "parallel"} {
		WORKERS.List[t] = make(map[int]workers.Worker)

		for id := 1; id <= 5; id++ {
			w := workers.NewWorker(id, t, tasks)
			WORKERS.Add(w)

			go w.Run()
		}
	}

	time.Sleep(time.Duration(5 * time.Second))

	for _, t := range []string{"sequential", "parallel"} {
		for id := 1; id <= 5; id++ {
			WORKERS.Get(t, id).Stop()
		}
	}

	time.Sleep(time.Duration(time.Second))
	fmt.Println("END")
}
