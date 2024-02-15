package main

import "fmt"

func divide(dividend int, divisor int) (int, error) {
	if divisor == 0 {
		return 0, fmt.Errorf("cannot divide by 0")
	}
	return dividend / divisor, nil
}

func main() {
	result, err := divide(10, 0)
	fmt.Printf("%d %s", result, err)
}
