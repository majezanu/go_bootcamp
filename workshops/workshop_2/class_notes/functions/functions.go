package main

import (
	"fmt"
	"math/rand"
)

func buildFullName(firstName string, lastName string) string {
	return fmt.Sprintf("%s %s", firstName, lastName)
}

func buildRGB() (R, G, B int) {
	R = rand.Int()
	G = rand.Int()
	B = rand.Int()
	return
}

func Sum(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	fullName := buildFullName("Manuel", "Zavala")
	echo(fullName)

	r, g, b := buildRGB()
	echo(r)
	echo(g)
	echo(b)

	sum := Sum(1, 2, 3)
	echo(sum)

	step := counter
	step(echo)      // ix = 1
	step(echo)      // ix = 2
	step(increment) // ix = 4
	echo(ix)
	step(increment) // ix = 6
	echo(ix)
}

var ix int

func counter(callables ...func(interface{})) {
	ix++
	for _, f := range callables {
		f(ix)
	}
}

func echo(s interface{}) {
	fmt.Println(s)
}

func increment(i interface{}) {
	temp, ok := i.(int)
	if ok {
		temp++
		i = temp
	}

}
