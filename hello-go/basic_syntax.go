package main

import "fmt"

// Variables
var name string = "Gopher"
var age int = 5
var isCool bool = true

// Multiple variable declaration
var (
	country string  = "USA"
	weight  float32 = 180.5
	hasPet  bool    = true
)

// Constants
const Pi = 3.14
const Language = "Go"

// Multiple constants
const (
	E       = 2.71
	Version = "1.16"
)

// Functions
func greet(name string) string {
	return "Hello, " + name + "!"
}

func add(a int, b int) int {
	return a + b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func functionTestings() {
	// Short-hand declaration
	city := "San Francisco"
	height := 6.1
	isEmployed := false

	// Conditionals
	if age > 18 {
		fmt.Println("Adult")
	}
	if city == "San Francisco" {
		fmt.Println("You live in SF!")
	}
	if height > 6.0 && isEmployed {
		fmt.Println("Tall and employed")
	}

	if isCool {
		fmt.Println("You're cool!")
	} else {
		fmt.Println("You're not cool.")
	}
}

// Loops
func loopTestings() {
	// For loop
	// var i int
	i := 0
	for i = 0; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}

	// While loop (using for)
	j := 0
	for j < 3 {
		fmt.Println("While loop iteration:", j)
		j++
	}

	// Infinite loop with break
	k := 0
	for {
		if k >= 2 {
			break
		}
		fmt.Println("Infinite loop iteration:", k)
		k++
	}
}

// Structs
type Person struct {
	Name string
	Age  int
}

// Methods instead of class mentods
func (p Person) Greet() string {
	return "hi, my name is " + p.Name
}

// Compositions instead of Mixins
type SoftDelete struct {
	DeletedAt string
}

func (sd SoftDelete) IsDeleted() bool {
	return sd.DeletedAt != ""
}

// Employee struct embedding Person and SoftDelete explicitly, not extending
type Employee struct {
	Person
	SoftDelete
	Position string
}

func showBasicSyntax() {
	fmt.Println("Basic Syntax Demonstration:")
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Cool:", isCool)
	fmt.Println("Country:", country)
	fmt.Println("Pi:", Pi)
	fmt.Println("Language:", Language)
	fmt.Println("Greet:", greet("Go"))
	fmt.Println("Add:", add(2, 3))
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Divide:", result)
	}
	functionTestings()
	loopTestings()
	p := Person{Name: "John", Age: 30}
	fmt.Println(p.Greet())
	emp := Employee{Person: Person{Name: "Jane", Age: 25}, Position: "Developer"}
	fmt.Println("Employee:", emp.Name, emp.Position)
}
