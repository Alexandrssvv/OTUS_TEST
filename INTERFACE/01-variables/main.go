package main

import "fmt"

type MyI interface {
	MyMethod(string, string) (string, error)
	ObjMethod()
}

type MyImpl struct {
}

func (m MyImpl) MyMethod(s1, s2 string) (string, error) {
	return s1 + s2, nil
}

func (m MyImpl) ObjMethod() {
	fmt.Println("ObjMethod called")
}
func main() {
	var a int = 5
	var b float64 = 3.14
	var c string = "hello, Go!"

	fmt.Println(a)

	fmt.Println(b)

	fmt.Println(c)

}
