package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Printf("Enter the number of elements in the array\n")
	scanner := bufio.NewScanner(os.Stdin)
	var count int = -1
	for scanner.Scan() {
		count, _ = strconv.Atoi(scanner.Text())
		fmt.Printf("You entered " + strconv.Itoa(count) + ", initializing an array with these many elements\n")
		break
	}

	var arr []int
	if count > -1 {
		arr = make([]int, count)
		for i := 0; i < count; i++ {
			arr[i] = i
		}
	}

	fmt.Printf("Please select 1 of the following\n")
	fmt.Printf("1. For Quick Find, Eager Union\n")
	fmt.Printf("2. For Quick Union\n")
	fmt.Printf("3. For Weighted Union\n")

	var choice int = -1
	for scanner.Scan() {
		choice, _ = strconv.Atoi(scanner.Text())
		if choice < 1 || choice > 3 {
			fmt.Printf("Wrong input, please try again\n")
			continue
		}
		break
	}

	fmt.Printf("The length of the array is " + strconv.Itoa(len(arr)) + "\n")

}
