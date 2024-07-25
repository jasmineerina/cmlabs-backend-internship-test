package main

import "fmt"

func fizzBuzz(num int) {
	if num < 0 {
		fmt.Println("Input tidak boleh kurang dari 0")
	}

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if num%3 == 0 {
		fmt.Println("Fizz")
	} else if num%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(num)
	}
}

func main() {
	for i := 1; i <= 100; i++ {
		fizzBuzz(i)
	}
}
