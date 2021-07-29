package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	/* strings */
	tweet := "I study Golang. Golang is Beautiful!!"
	fmt.Println(strings.Contains(tweet, "Golang"))                 // true
	fmt.Println(strings.ReplaceAll(tweet, "Golang", "JavaScript")) // I study JavaScript. JavaScript is Beautiful!!
	fmt.Println(strings.ToUpper(tweet))                            // I STUDY GOLANG. GOLANG IS BEAUTIFUL!!
	fmt.Println(strings.ToLower(tweet))                            // i study golang. golang is beautiful!!
	fmt.Println(strings.Split(tweet, " "))                         // ["I", "study", "Golang", "Golang", "is", "Beautiful", "!!"])
	fmt.Println(strings.Fields(tweet))                             // ["I", "study", "Golang", "Golang", "is", "Beautiful", "!!"]
	fmt.Println(strings.Index(tweet, "Golang"))                    // 6)
	// これらは非破壊的
	fmt.Println(tweet) // I study Golang. Golang is Beautiful!!

	/* sort */
	numbers := []int{51, 22, 38, 47, 56, 96, 79, 18, 9, 101}

	index := sort.SearchInts(numbers, 47)
	fmt.Println(index)          // 4
	fmt.Println(numbers[index]) // 47

	sort.Ints(numbers) // 戻り値はなし
	// sortNumbers := sort.Ints(numbers) // Error!!
	fmt.Println(numbers) // [9 18 22 38 47 51 56 79 96 101]

	words := strings.Split(tweet, " ")
	sort.Strings(words)
	fmt.Println(words) // ["I", "study", "Golang", "Golang", "is", "Beautiful", "!!"]
}
