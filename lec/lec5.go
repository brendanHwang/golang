package lec

import "fmt"

// lec5는 배열과 슬라이스의 사용법을 보여주는 Go 함수입니다.
//
// 이 함수는 매개변수를 받지 않습니다.
// 반환값이 없습니다.
func lec5() {
	// var ages [3]int = [3]int{20, 30, 40}
	var ages = [3]int{20, 30, 40}

	names := [4]string{"Kim", "Lee", "Park", "Choi"}

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// slices (use arrays under the hood)
	var scores = []int{100, 90, 80}
	scores[2] = 25
	scores = append(scores, 85)
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	fmt.Println(rangeOne)
	rangeTwo := scores[1:]
	rangeThree := scores[:3]

	fmt.Println(rangeTwo)
	fmt.Println(rangeOne, rangeTwo, rangeThree)

}
