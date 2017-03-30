package main

import "fmt"

type Person struct{
	name string
	age int
}

func main(){

	myMap := make(map[string]Person)
	myMap["marcio"] = Person{"Marcio",33}
	myMap["claudia"] = Person{"Claudia", 46}


	for key, val := range myMap {
		fmt.Println(key, " - ", val)
	}

}
