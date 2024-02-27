package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name  string `required:"true" max:"10"`
	Email string `required:"true" max:"10"`
}

type Person struct {
	Name string
}

func readField(value any) {
	var valuetype reflect.Type = reflect.TypeOf(value)
	fmt.Println("Name Is", valuetype.Name())
	for i := 0; i < valuetype.NumField(); i++ {
		var valuefield reflect.StructField = valuetype.Field(i)
		fmt.Println(valuefield.Name, "With Type ", valuefield.Type)
		fmt.Println(valuefield.Tag.Get("required"))
		fmt.Println(valuefield.Tag.Get("max"))
	}
}
func IsValid(value any) (result bool) {
	t := reflect.TypeOf(value)          //Typeof mengexek data bertipe apa dari strutc
	for i := 0; i < t.NumField(); i++ { //Numfield menmbuat field di strutc menjadi Number
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface() //Interface membuat data
			result = data != ""
			if result == false {
				return result

			}
		}
	}
	return result
}

func main() {

	readField(Sample{"Chandra", "Chandra@gmail.com"})
	readField(Person{"Ekoo"})

	sample := Sample{"Wanderer", "Wanderer@gmail.com"}
	fmt.Println(IsValid(sample))

}
