package page7

func threeStringArray() [3]string {
	return [3]string{
		"hello",
		"nice",
		"to meet you",
	}
}

func SliceAndArrayDiff() []string {

	threeStringArray := threeStringArray()

	return threeStringArray[:]

}
