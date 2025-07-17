package main

func CountEvens(list []int) int {
	count := 0
	for _, num := range list {
		if num%2 == 0 {
			count++
		}
	}
	return count
}

func GetWordCounts(words []string) map[string]int {
	result := make(map[string]int)
	for _, word := range words {
		result[word]++
	}
	return result
}

func Anagrams(f string, s string) bool {
	if len(f) != len(s) {
		return false
	}

	letterCount := make(map[rune]int)

	for _, ch := range f {
		letterCount[ch]++
	}

	for _, ch := range s {
		letterCount[ch]--
		if letterCount[ch] < 0 {
			return false
		}
	}

	return true
}

/*
func main() {
	fmt.Println(CountEvens([]int{1, 2, 3, 4, 5, 6}))

	words := []string{"go", "is", "fun", "go", "go", "fun"}
	fmt.Println(GetWordCounts(words))

	fmt.Println(Anagrams("listen", "silent"))
}
*/
