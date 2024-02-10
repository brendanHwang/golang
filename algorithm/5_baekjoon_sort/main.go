package main

import "fmt"

//// 2750
//func main() {
//
//
//	scanner := bufio.NewScanner(os.Stdin)
//	scanner.Scan()                       // 첫 번째 줄 읽기 (숫자의 개수)
//	n, _ := strconv.Atoi(scanner.Text()) // 에러 처리 생략
//
//	data := make([]int, n) // n 크기의 int 슬라이스 생성
//	for i := 0; i < n; i++ {
//		scanner.Scan()                            // 다음 숫자 읽기
//		data[i], _ = strconv.Atoi(scanner.Text()) // 에러 처리 생략
//	}
//	selectionSort(data)
//
//	ans := ""
//	for _, v := range data {
//		ans += fmt.Sprintf("%v\n", v)
//	}
//	fmt.Println(ans)
//
//}

//2752
//func main() {
//
//	var a, b, c int
//	fmt.Scan(&a, &b, &c)
//
//	numbers := []int{a, b, c}
//
//	three_num_sort(numbers)
//
//	ans := ""
//	for _, v := range numbers {
//		ans += fmt.Sprintf("%v ", v)
//	}
//	fmt.Println(ans)
//
//}

// 2751
func main() {
	var number int
	fmt.Scan(&number)

	data := make([]int, number)
	for i := 0; i < number; i++ {
		fmt.Scan(&data[i])
	}

	quickSort(data, 0, number-1)

	ans := ""
	for _, v := range data {
		ans += fmt.Sprintf("%v\n", v)
	}

	fmt.Println(ans)
}
