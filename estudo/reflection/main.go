package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int64  `json:id`
	Name string `json:name`
}

func main() {
	//tipoPrimitivo()
	//tipoStruct()

	tp := reflect.TypeOf(User{})
	userInterface := createObjet(tp, 13, "{ \"name\": \"felipe\" }")
	user := userInterface.(User)
	user.Name = "felipe"
	fmt.Println(user)

	/* sliceType := reflect.TypeOf([]User{})
	itemType := reflect.TypeOf(User{})
	usersInterface := createArray(sliceType, itemType)
	users := usersInterface.([]User)
	users[0].Id = 1
	users[0].Name = "felipe"
	fmt.Println(users) */
}

func createArray(sliceType reflect.Type, itemType reflect.Type) interface{} {
	items := reflect.MakeSlice(sliceType, 0, 3)

	items = reflect.Append(items, reflect.New(itemType).Elem())
	items = reflect.Append(items, reflect.New(itemType).Elem())

	//fmt.Println("value    : ", items)
	return items.Interface()
}

func createObjet(tp reflect.Type, id int64, jsonValue string) interface{} {
	item := reflect.New(tp)

	valueOfSet := item.Elem()
	valueOfSet.FieldByName("Id").SetInt(id)

	fmt.Println("type     :", valueOfSet.Type().Name())
	fmt.Println("can edit : ", item.Elem().CanSet())
	fmt.Println("value    : ", item)

	return valueOfSet.Interface()
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
