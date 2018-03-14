package main

import "errors"

import "fmt"

func f1(arg int) (int, error) {
	if 42 == arg {
		return -1, errors.New("Can't work with 42")
	}

	return arg + 3, nil
}

// ArgError initial
type ArgError struct {
	arg  int
	prob string
}

func (e *ArgError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if 42 == arg {
		return -1, &ArgError{arg, "Can't work with it!"}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		r, e := f1(i)
		if e != nil {
			fmt.Println("f1 failed: ", e)
		} else {
			fmt.Println("f1 worked: ", r)
		}
	}
}
