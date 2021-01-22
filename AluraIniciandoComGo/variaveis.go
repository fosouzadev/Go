package main

import (
	"fmt"
	"reflect"
)

func main() {
	var name string = "Fulano" // valor default ""
	var nameEmpty string
	name2 := "Fulano 2"
	var typeString = "abc"

	var age int = 30 // valor default 0
	var ageEmpty int
	age2 := 49
	var typeInt = 30

	var value float32 = 10000.99 // valor default 0.0
	var valueEmpty float32
	value2 := 453.44
	var typeFloat = 4.222

	var isValid bool = true // valor default false
	var isValidEmpty bool
	isValid2 := false
	var typeBool = true

	fmt.Println("string: ", name, " ", nameEmpty, " ", name2, " ", reflect.TypeOf(typeString))
	fmt.Println("int   : ", age, " ", ageEmpty, " ", age2, " ", reflect.TypeOf(typeInt))
	fmt.Println("float : ", value, " ", valueEmpty, " ", value2, " ", reflect.TypeOf(typeFloat))
	fmt.Println("bool  : ", isValid, " ", isValidEmpty, " ", isValid2, " ", reflect.TypeOf(typeBool))
}
