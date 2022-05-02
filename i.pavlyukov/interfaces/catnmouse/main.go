package main

import "fmt"

type Cat struct {
	Name    string
	Surname string
	Age     int
}

func (c *Cat) Greeting(c8 int) error {
	fmt.Printf("Hi, I'm %s %s \n", c.Name, c.Surname)

	return nil
}

func (c *Cat) Walk() {
}

// ------------- //

type Mouse struct {
}

func (m *Mouse) Greeting(p2 int) error {
	fmt.Printf("Hi, mouse have no name\n")

	return nil
}

func (m *Mouse) Walk() {
}

func (m *Mouse) Say(word string) {
	fmt.Println(word)
}

// ------ //

type Animal interface {
	Greeting(int) error
	Walk()
}

func Welcome(an Animal) {
	fmt.Print("Welcome, introduce yourself... ")
	_ = an.Greeting(0)
}

func main() {
	ca := &Cat{
		Name:    "Larisa",
		Surname: "Guzeeva",
	}

	ca2 := &Cat{
		Name:    "Rosa",
		Surname: "Sibitova",
	}

	mo := &Mouse{}

	Welcome(ca)

	Welcome(ca2)

	Welcome(mo)
}
