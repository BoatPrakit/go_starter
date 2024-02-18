package main

import (
	"fmt"
)

func main() {
	x := "Hello, 世界 😀😆🥰"
	xs := []rune(x)
	for i := 0; i < len(x); i++ {
		a := x[i]
		fmt.Println(string(a))
		fmt.Printf("%T\n", a)
	}
	for i := 0; i < len(xs); i++ {
		a := xs[i]
		fmt.Println(string(a), a)
		fmt.Printf("%T\n", a)
	}
	// arr := make(map[string]string)
	// arr["A"] = "Test"
	// strArr := strings.Split(arr["A"], "")
	// for e, i := range strArr {
	// 	fmt.Println(e, i)
	// }
	// p := createPerson("A", 10)
	// fmt.Println(p)
	// myMap := map[string]any{"A": "test"}
	// myMap["B"] = 2
	// if myMap["A"] != nil {
	// 	fmt.Println(myMap)
	// }
}

type Person struct {
	Name string
	Age  int
}

func createPerson(name string, age int) Person {
	return Person{
		Name: name,
		Age:  age,
	}
}
