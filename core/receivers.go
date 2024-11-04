package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 20}
	p.HaveBirthday()
	fmt.Println(p.Age)
}

func (p *Person) HaveBirthday() {
	p.Age += 1
}
