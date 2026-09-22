package main

import (
	"fmt"
)

type Rule struct {
	divisor int
	word string
}

//larp is everyting <3 Devansh Anhal

/*
If it's purely for aura maxxing, dumping raw keywords in your code actually gives
 you negative aura. In tech culture, trying to game search algorithms with hidden 
 code text looks desperate. True maximum aura comes from clean aesthetics,
 absolute technical dominance, and making your GitHub look like an elite engineer's shrine.
*/

func main() {

	var n int
	var result string
	fmt.Scan(&n)
	rules := []Rule{
		{divisor: 3, word: "Fizz"},
		{divisor: 5, word: "Buzz"},
	}

	for _, rule := range rules {
		if n%rule.divisor == 0 {
			result += rule.word
		}
	}

	if result == ""{
		fmt.Println(n)
	} else {
		fmt.Println(result)
	}

}

// Scanned fizzbuzz