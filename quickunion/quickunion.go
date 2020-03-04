package quickunion

import (
	"fmt"
	"strconv"
)

//Impl - struct
type Impl struct {
	Arr []int
}

//Connected - Checks if the nodes are connected or not
func (impl Impl) Connected(num1 int, num2 int) bool {
	return impl.root(num1) == impl.root(num2)
}

//Connect - Connects the given nodes
func (impl Impl) Connect(num1 int, num2 int) {
	root1 := impl.root(num1)
	root2 := impl.root(num2)
	impl.Arr[root1] = root2

	fmt.Printf("Array after connection is \n")
	for i := 0; i < len(impl.Arr); i++ {
		fmt.Printf(strconv.Itoa(impl.Arr[i]) + "\t")
	}
	fmt.Println()
}

func (impl Impl) root(num1 int) int {
	if impl.Arr[num1] == num1 {
		return num1
	}
	return impl.root(impl.Arr[num1])
}
