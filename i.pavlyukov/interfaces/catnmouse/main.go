package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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

func (c *Cat) String() string {
	return c.Name + " " + c.Surname
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

func (m *Mouse) Say(word ...string) {
	fmt.Println(word)
}

func (m *Mouse) String() string {
	return "mouse has no name"
}

// ------ //

type Animal interface {
	Greeting(int) error
	Walk()
	String() string
}

func Welcome(an Animal) {
	fmt.Print("Welcome, introduce yourself... ")
	_ = an.Greeting(0)
}

func main() {

	//ca2 := &Cat{
	//	Name:    "Rosa",
	//	Surname: "Sibitova",
	//}

	//mo := &Mouse{}

	resp, _ := http.Get("https://ya.ru")
	defer resp.Body.Close()

	full, _ := ioutil.ReadAll(resp.Body)

	var an interface{} = int32(100)
	// type casting
	switch an.(type) {
	case int:
	case int32:
	case Animal:

	}

	// type conversion
	fullStr := []byte(full)

	fmt.Println(fullStr)
}
