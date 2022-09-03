package main

import "fmt"

type MyInterface interface {
	String() string
}

type MyType struct {
	value string
}

// MyType has a method "String()", which has a pointer to MyType as a receiver.
func (ptrString *MyType) String() string {
	return ptrString.value
}

// It seems that pointer to an interface can't be a parameter, because
// a pointer to particular concrete type itself is not an implementation of pointer to the interface.
// (Because "pointer" to concrete type doesn't have a method called String())
func ReceiveInterface(itf *MyInterface) {
	// fmt.Println(itf.String()) // An error occurs : a pointer parameter cannot use . operator?
	// As we can see in main function below, a pointer to MyType(secondConcrete) can use . operator to call String() fucntion.
	fmt.Println((*itf).String())
}

func main() {
	concrete := MyType{"concrete value"}
	fmt.Println(concrete.String())
	secondConcrete := &MyType{"secondConcrete value"}
	fmt.Println(secondConcrete.String())

	ReceiveInterface(secondConcrete)
}
