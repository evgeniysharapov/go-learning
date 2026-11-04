package main

import (
	"fmt"
	"errors"
)

func divide( a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("b cannot be 0")
	}
	return a / b, a % b, nil 
}

func sum(nums ...int) int {
	s := 0
	for i := 0; i < len(nums); i++ {
		s += nums[i]
	}
	return s	
}


func main(){
	arr := []int{1,2,3,4,5,6}

	q,r, err := divide(2,3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result: ", q, " ", r)
	}

	q,r, err = divide(2,0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result: ", q, " ", r)
	}
	
	fmt.Printf("Sum %d\n", sum(1,2,3))
	fmt.Printf("Sum %d\n", sum(arr...))

}
