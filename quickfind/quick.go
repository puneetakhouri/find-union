package quickfind

import (
	"fmt"
	"strconv"
)

//Impl - QuickFind Implementation for find union
type Impl struct {
	Arr []int
}

//Connected - Checks if they are connected
func (impl Impl) Connected(num1 int, num2 int) bool {
	return impl.Arr[num1] == impl.Arr[num2]
}

//Connect - Connects them using the quick approach
func (impl Impl) Connect(num1 int, num2 int) {
	originalNum1 := impl.Arr[num1]
	originalNum2 := impl.Arr[num2]

	for i := 0; i < len(impl.Arr); i++ {
		if impl.Arr[i] == originalNum1 {
			impl.Arr[i] = originalNum2
		}
	}

	fmt.Printf("Array after connection is \n")
	for i := 0; i < len(impl.Arr); i++ {
		fmt.Printf(strconv.Itoa(impl.Arr[i]) + "\t")
	}
	fmt.Println()
}
