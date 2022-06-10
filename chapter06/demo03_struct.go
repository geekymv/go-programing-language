package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Say(content string) {
	fmt.Println(p.Name + " say " + content)
}

type Student struct {
	Person
	No int
}

func main() {

	s1 := Student{
		No: 1000,
		Person: Person{
			Name: "tom",
		},
	}

	s1.Say("hello")
	s1.Person.Say("hello")
}
