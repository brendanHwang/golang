package main

import (
	"bufio"
	"fmt"
	"os"
)

// 계수 정렬
//func Sort10989() {
//	reader := bufio.NewReader(os.Stdin)
//	writer := bufio.NewWriter(os.Stdout)
//	defer writer.Flush()
//
//	var n int
//	fmt.Fscanln(reader, &n)
//
//	// 10,000 이하의 자연수이므로, 각 숫자의 빈도를 저장할 배열을 선언합니다.
//	counts := make([]int, 10001)
//
//	for i := 0; i < n; i++ {
//		num, _ := reader.ReadString('\n')
//		numInt, _ := strconv.Atoi(num[:len(num)-1]) // 줄바꿈 문자 제거 후 정수 변환
//		counts[numInt]++                            // 해당 숫자의 빈도를 증가시킵니다.
//	}
//
//	// 빈도 배열을 사용하여 결과를 출력합니다.
//	for i, count := range counts {
//		for j := 0; j < count; j++ {
//			fmt.Fprintln(writer, i)
//		}
//	}
//}

func Sort10989() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)

	counts := make(map[int]int)

	for i := 0; i < n; i++ {
		var num int
		fmt.Fscanln(reader, &num)
		if _, ok := counts[num]; !ok {
			counts[num] = 1
		} else {
			counts[num]++
		}
	}

	for i := 0; i < 10001; i++ {
		for j := 0; j < counts[i]; j++ {
			fmt.Fprintln(writer, i)
		}
	}

}
