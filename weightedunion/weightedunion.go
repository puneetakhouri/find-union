package weightedunion

//Impl - struct
type Impl struct {
	Arr  []int
	Size []int
}

//Connected - Checks if the nodes are connected or not
func (impl Impl) Connected(num1 int, num2 int) bool {
	return impl.root(num1) == impl.root(num2)
}

//Connect - Connects the given nodes
func (impl Impl) Connect(num1 int, num2 int) {

	rootNum1 := impl.root(num1)
	rootNum2 := impl.root(num2)

	if impl.Size[num1] < impl.Size[num2] {
		rootNum1 = impl.root(num2)
		rootNum2 = impl.root(num1)
	}

	impl.Arr[rootNum1] = rootNum2
	impl.refreshSize()
}

func (impl Impl) root(num1 int) int {
	if impl.Arr[num1] == num1 {
		return num1
	}
	return impl.root(impl.Arr[num1])
}

func (impl Impl) refreshSize() {
	//WIP
}
