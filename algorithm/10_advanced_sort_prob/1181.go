package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Sort1181() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 첫 줄 읽기 (단어 개수, 여기서는 사용하지 않음)

	wordMap := make(map[string]struct{})
	words := []string{}

	for scanner.Scan() {
		word := scanner.Text()
		if _, exists := wordMap[word]; !exists {
			wordMap[word] = struct{}{}
			words = append(words, word)
		}
	}

	sort.SliceStable(words, func(i, j int) bool {
		if len(words[i]) < len(words[j]) {
			return true
		} else if len(words[i]) == len(words[j]) && words[i] < words[j] {
			return true
		}
		return false
	},
	)

	for _, word := range words {
		fmt.Println(word)
	}

}
