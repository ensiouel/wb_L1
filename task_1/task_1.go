package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (human Human) String() string {
	return fmt.Sprintf("human %s is %d years old", human.Name, human.Age)
}

// Action Встраивание структуры Human в структуру Action
type Action struct {
	Human
}

func (action Action) String() string {
	return fmt.Sprintf("action %s is %d years old", action.Name, action.Age)
}

func main() {
	action := Action{Human: Human{
		Name: "Ivan",
		Age:  18,
	}}

	// Вызовется String дочерней структуры, т.е. Action
	fmt.Println(action.String())
}
