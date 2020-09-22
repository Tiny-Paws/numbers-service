package numbers

import "errors"

type NumbersService interface {
	Add(a, b int) (int, error)
	Sub(a, b int) (int, error)
}

type numbersService struct{}

func Add(a, b int) (int, error) {
	// Must add checks to prevent int overflow
	return a + b, nil
}

func Sub(a, b int) (int, error) {
	// Must add checks to prevent int underflow
	return a - b, nil
}

var ErrOverflow = errors.New("Integer overflow detected")
