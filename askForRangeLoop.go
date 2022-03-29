package main

import (
	"fmt"
)

func main() {
	//Insert code for only even here

	var value1, value2 int
	var dummy string
	fmt.Println("Enter First Value")
	fmt.Scanf("%d", &value1)
	fmt.Scanf("%s", &dummy) // to take care of CR

	fmt.Println("Enter Second Value")
	fmt.Scanf("%d", &value2)

	// make sure value 1 is less than value 2
	if value2 < value1 {
		value1, value2 = value2, value1
	}

	// make sure value 1 is less than value 2
	fmt.Println("Odd Numbers")
	for x := value1; x <= value2; x++ {
		if x%2 == 1 {
			fmt.Printf("%d\n", x)
		}
	}

	fmt.Println("Even Numbers")
	for x := value1; x <= value2; x++ {
		if x%2 == 0 {
			fmt.Printf("%d\n", x)
		}
	}
}
