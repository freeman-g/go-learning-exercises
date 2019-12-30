// Given a positive integer num, return the sum of all odd Fibonacci numbers that are less than or equal to num.

package main

func sumFibs(num int) int {
	previousFibNumber := 0
	fibNumber := 1
	sum := 0
	for fibNumber <= num {
		if fibNumber%2 != 0 {
			// fmt.Printf("Fib number is %v, previous fib number is %v\r\n", fibNumber, previousFibNumber)
			sum += fibNumber
		}

		fibNumber += previousFibNumber
		previousFibNumber = fibNumber - previousFibNumber
	}

	return sum
}
