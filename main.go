package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/puneetakhouri/find-union/findunion"
	"github.com/puneetakhouri/find-union/quickfind"
	"github.com/puneetakhouri/find-union/quickunion"
	"github.com/puneetakhouri/find-union/weightedunion"
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

	switch choice {
	case 1:
		execute(arr, quickfind.Impl{arr})
	case 2:
		execute(arr, quickunion.Impl{arr})
	case 3:
		sizeArr := make([]int, count)
		for i := 0; i < len(sizeArr); i++ {
			sizeArr[i] = 1
		}
		execute(arr, weightedunion.Impl{arr, sizeArr})
	}

	fmt.Printf("The length of the array is " + strconv.Itoa(len(arr)) + "\n")

}

func execute(arr []int, findUnion findunion.Interface) {
	fmt.Println("Trying")

	//var findUnion findunion.Interface = quickfind.Impl{arr}

	fmt.Printf("Checking if 2 & 3 are connected\n")
	fmt.Println(findUnion.Connected(2, 3))

	fmt.Println("Connecting 1 & 2")
	findUnion.Connect(1, 2)

	fmt.Println("Connecting 4 and 5")
	findUnion.Connect(4, 5)

	fmt.Println("Connecting 7 & 8")
	findUnion.Connect(7, 8)

	fmt.Println("Connecting 4 & 7")
	findUnion.Connect(4, 7)

	fmt.Println("Connecting 1 & 3")
	findUnion.Connect(1, 3)

	fmt.Printf("Checking if 1 & 3 are connected\n")
	fmt.Println(findUnion.Connected(1, 3))

	fmt.Println("Connecting 1 & 8")
	findUnion.Connect(1, 8)

	fmt.Printf("Checking if 5 & 6 are connected\n")
	fmt.Println(findUnion.Connected(5, 6))

	fmt.Println("Connecting 3 & 6")
	findUnion.Connect(3, 6)

	fmt.Printf("Checking if 8 & 9 are connected\n")
	fmt.Println(findUnion.Connected(8, 9))
}
