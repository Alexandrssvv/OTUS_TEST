package main

import (
	"fmt"
)

type EmptuI interface {
}

type MyI interface {
	MyMethod(string, string) (string, error)
	ObjMethod()
}

type MyImpl struct {
	Name string
}

type MyImpl2 struct {
	Name string
}

func (m *MyImpl) MyMethod(s1, s2 string) (string, error) {
	return s1 + s2, nil
}

func (m *MyImpl) ObjMethod() {
	fmt.Println("ObjMethod called")
}

func (m *MyImpl2) MyMethod(s1, s2 string) (string, error) {
	return s1 + s2, nil
}

func (m *MyImpl2) ObjMethod() {
	fmt.Println("ObjMethod called")
}

func main() {

	var myInterface MyI
	var obj = &MyImpl{Name: "MyImpl Object"}
	obj.Name = "My Object"
	obj.ObjMethod()
	myInterface = obj
	myInterface.ObjMethod()
	//fmt.Printf("i is of type %T and value %v\n", myInterface, myInterface)

	processValue(myInterface)

	var a int = 5
	var b float64 = 3.14
	var c string = "hello, Go!"

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)

	//interface

	var i any

	i = a
	fmt.Printf("i is of type %T and value %v\n", i, i)

	i = b
	fmt.Printf("i is of type %T and value %v\n", i, i)

	i = c
	fmt.Printf("i is of type %T and value %v\n", i, i)
}

func processValue(i MyI) {

	switch v := i.(type) {
	case *MyImpl:
		fmt.Println("Processing MyImpl with Name:", v.Name)
	case *MyImpl2:
		fmt.Println("Processing MyImpl2")
	default:
		fmt.Println("Unkown type")
	}
}
