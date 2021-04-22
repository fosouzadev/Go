package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   rune
	Name string
}

func main() {
	tipoStruct()
}

func tipoStruct() {
	user := User{1, "noe"}
	valueof := reflect.ValueOf(&user).Elem()
	typeof := valueof.Type()

	for i := 0; i < valueof.NumField(); i++ {
		field := valueof.Field(i)

		fmt.Println(
			"struct", typeof.Name(),
			"field type", field.Type().Name(),
			"field name", typeof.Field(i).Name,
			"field value", field)
	}

	valueof.Field(0).SetInt(3)
	valueof.Field(1).SetString("jose")
	fmt.Println(user)
}

func tipoPrimitivo() {
	var x float64 = 3.4
	typeof := reflect.TypeOf(x)
	valueof := reflect.ValueOf(x)

	fmt.Println("type x", typeof, "value x", valueof)

	if valueof.Kind() == reflect.Float64 {
		fmt.Println("type valueof", valueof.Type(), "value valueof", valueof.Float())
	}

	// editar valor
	pValueOf := reflect.ValueOf(&x)
	fmt.Println("type pValueOf", pValueOf.Type(), "canset", pValueOf.CanSet())

	pValueOfSet := pValueOf.Elem()
	fmt.Println("type pValueOfSet", pValueOfSet.Type(), "canset", pValueOfSet.CanSet())

	pValueOfSet.SetFloat(5.2)
	fmt.Println("value pValueOfSet", pValueOfSet, "value x", x)
}
