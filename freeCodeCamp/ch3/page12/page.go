package page12

import (
	"errors"
)

func Divide(dividend, divsisor int) (int, error) {
	if divsisor == 0 {
		return 0, errors.New("division by zero")
	}
	return dividend / divsisor, nil
}
