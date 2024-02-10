package page6

import (
	"errors"
	"fmt"
)

func Deposit(krw float64) error {
	if krw <= 0 {
		return fmt.Errorf("Invalid KRW amount: %.2f", krw)
	} else {
		fmt.Println("Deposited KRW: ", krw)
		return nil
	}
}

func Division(dividend, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("division by zero")
	}
	return dividend / divisor, nil
}
