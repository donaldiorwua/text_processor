package TEXT_PROCESSOR

import (
	"fmt"
)

func BasicArithmetis() {
	fmt.Println("============================================")
	fmt.Println(",,,, WELCOME TO OUR CALCULATOR PROGRAM!,,,,")
	fmt.Println("=============================================")
	for {
		var calculation, input string
		var num1, num2 int

		defer fmt.Scan(&input)
		if input == "exit" {
			fmt.Println("Thanks for Using our System")
			return
		}

		fmt.Println("Please, choose an operation: \n add \n subtract \n multiply \n divide \n Enter 'exit' to terminate the program")

		fmt.Scan(&calculation)

		switch calculation {
		case "add":
			fmt.Println("Enter Two numbers to add!")
			fmt.Println("Enter a number: ")
			fmt.Scan(&num1)
			fmt.Println("Enter another number: ")
			fmt.Scan(&num2)
			fmt.Println("Result: ", add(num1, num2))
			return

		case "multiply":
			fmt.Println("Enter Two numbers to multiply!")
			fmt.Println("Enter a number: ")
			fmt.Scan(&num1)
			fmt.Println("Enter another number: ")
			fmt.Scan(&num2)
			fmt.Println("Result: ", multiply(num1, num2))
			return

		case "subtract":
			fmt.Println("Enter Two numbers to subtract!")
			fmt.Println("Enter a number: ")
			fmt.Scan(&num1)
			fmt.Println("Enter another number: ")
			fmt.Scan(&num2)
			fmt.Println("Result: ", subtract(num1, num2))
			return

		case "divide":
			var num1, num2 float64
			fmt.Println("Enter Two numbers to divide!")
			fmt.Println("Enter a number: ")
			fmt.Scan(&num1)
			fmt.Println("Enter another number: ")
			fmt.Scan(&num2)
			fmt.Println("Result: ", divide(num1, num2))

		default:
			fmt.Println("Invalid Operation, Try again!")
			return
		}

	}
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b float64) float64 {
	return a / b
}
