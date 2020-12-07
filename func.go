package main

import (
	"fmt"
	"runtime"
)

func main() {
	Foo()
}
func Foo() {
	fmt.Printf("my %s, %s call !\n", printMyName(), printCallerName())
	Bar()
}
func Bar() {
	fmt.Printf("my %s, %s and call!\n", printMyName(), printCallerName())
}
func printMyName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
func printCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}
