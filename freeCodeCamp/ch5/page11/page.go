package page11

import "fmt"

func PrintNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Println("int : ", v)
	case float64:
		fmt.Println("float64 : ", v)
	case string:
		fmt.Println("string : ", v)
	default:
		fmt.Printf("unknown type: %v\n", v)
	}

}

func GetExpenseReport(e expense) (string, float64) {
	switch v := e.(type) {
	case Email:
		return v.ToAddress, e.cost()
	case Sms:
		return v.ToPhoneNum, e.cost()
	default:
		return "", 0.0
	}
}

type expense interface {
	cost() float64
}

type Email struct {
	IsSubscribed bool
	Body         string
	ToAddress    string
}

func (e Email) cost() float64 {
	if !e.IsSubscribed {
		return float64(len(e.Body))
	}

	return float64(len(e.Body))
}

type Sms struct {
	IsSubscribed bool
	Body         string
	ToPhoneNum   string
}

func (s Sms) cost() float64 {
	if !s.IsSubscribed {
		return float64(len(s.Body))
	}

	return float64(len(s.Body))
}
