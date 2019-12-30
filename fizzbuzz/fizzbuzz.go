// the classic FizzBuzz exercise

package main

import "fmt"

func main() {
	for i := 1; i < 101; i++ {
		if i%5 == 0 && i%3 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}
