package main

import (
	"./heroes"
	"./montecarlo"
	"./statistics"
)

func main() {
	heroes.ProcessHeroes()
	statistics.ProccesNumbers()
	montecarlo.EstimatePi()
}
