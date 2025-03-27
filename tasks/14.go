package main

import (
	"fmt"
)

func checkType(t interface{}) string {
	switch t.(type) {
	case int:
		return "int"
	case float64:
		return "float"
	case string:
		return "string"
	case bool:
		return "bool"
	case uint:
		return "uint"
	}
	return "non type"
}
func main() {
	fmt.Println(checkType(132.124))
}

//func getType(value interface{}) string {
//	// Получаем тип значения интерфейса
//	typ := reflect.TypeOf(value)
//
//	return typ.String()
//}
//
//func main() {
//	intValue := 42
//	strValue := "Hello"
//	boolValue := true
//	chValue := make(chan int)
//
//	fmt.Println("Type intValue:", getType(intValue))
//	fmt.Println("Type strValue:", getType(strValue))
//	fmt.Println("Type boolValue:", getType(boolValue))
//	fmt.Println("Type chValue:", getType(chValue))
//}
