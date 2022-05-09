package heroes

import "fmt"

// Define one list of heroes from Marvel and DC comics (at least 4 heroes), and separate it on 2
// slices, one slice for the Marvel heroes and other slices for the DC comics heroes

// Input: heroes – [“Batman”, “Spiderman”, “Superman”, “Dr. Strange”]
//Output:  DC – [“Batman”,”Superman”]
//Marvel – [“Spiderman”, “Dr. Strange”]

func isMarvelHero(hero string) bool {
	return hero == "Ironman" || hero == "Spiderman"
}

func isDCHero(hero string) bool {
	return hero == "Superman" || hero == "Batman"
}

func ProcessHeroes() {
	fmt.Println("Challenge 1: Heroes")
	heroes := []string{"Superman", "Batman", "Ironman", "Spiderman"}

	dcHeroes := []string{}
	marvelHeroes := []string{}

	for _, hero := range heroes {
		if isDCHero(hero) {
			dcHeroes = append(dcHeroes, hero)
		}
		if isMarvelHero(hero) {
			marvelHeroes = append(marvelHeroes, hero)
		}
	}
	fmt.Println("DC -", dcHeroes)
	fmt.Println("Marvel -", marvelHeroes)
}
