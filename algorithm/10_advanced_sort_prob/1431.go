package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

// totalSum 함수는 숫자들의 합을 계산합니다.
func totalSum(s string) int {
	total := 0
	for _, v := range s {
		if '0' <= v && v <= '9' {
			total += int(v - '0')
		}
	}
	return total
}

func Sort1413() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // 첫 번째 줄 (N) 읽기 - 현재 구현에서는 사용하지 않음

	var serials []string

	// 입력 받은 시리얼 번호 저장
	for scanner.Scan() {
		serial := scanner.Text()
		serials = append(serials, serial)
	}

	// 조건에 따라 시리얼 번호 정렬
	sort.Slice(serials, func(i, j int) bool {
		if len(serials[i]) != len(serials[j]) {
			return len(serials[i]) < len(serials[j])
		}

		sumI := totalSum(serials[i])
		sumJ := totalSum(serials[j])
		if sumI != sumJ {
			return sumI < sumJ
		}

		return serials[i] < serials[j]
	})

	// 정렬된 시리얼 번호 출력
	for _, serial := range serials {
		fmt.Println(serial)
	}
}
