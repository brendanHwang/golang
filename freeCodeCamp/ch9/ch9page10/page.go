package påge10

import "fmt"

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)

	for _, name := range names {
		firstLetter := rune(name[0])
		if _, ok := counts[firstLetter]; !ok {
			counts[firstLetter] = make(map[string]int)
		}
		if _, ok := counts[firstLetter][name]; !ok {
			counts[firstLetter][name] = 0
		}
		counts[firstLetter][name]++
	}

	return counts
}

func Nested() {
	names := []string{
		"brendan",
		"leonard",
		"park",
		"seong",
	}

	counts := getNameCounts(names)
	fmt.Println(counts)
}
