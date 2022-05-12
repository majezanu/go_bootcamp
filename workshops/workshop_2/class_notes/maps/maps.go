package main

import "fmt"

var myMap map[string]string

func printMap(myMap map[string]string, includeLenght bool = true) {
	fmt.Printf("myMap = %#v\n", myMap)
	if includeLenght {
		fmt.Println("myMap.len = ", len(myMap))
	}
}

func main() {
	printMap(myMap, true)

	myMap = map[string]string{
		"greeting": "Hello, world",
	}
	printMap(myMap, true)
	myMap["greeting"] = "Hello, Zavala"
	printMap(myMap, true)
	myMap["greeting2"] = "Hello, Zavala2"
	printMap(myMap, true)
	delete(myMap, "greeting2")
	printMap(myMap, true)
}
