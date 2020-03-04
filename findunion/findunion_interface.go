package findunion

//Interface - Interface defining the Find and Union Methods
type Interface interface {
	Connected(num1 int, num2 int) bool
	Connect(num1 int, num2 int)
}
