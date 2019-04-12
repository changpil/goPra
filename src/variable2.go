package main

import ( "fmt"
        "reflect")
import "strconv"
/*
var (
	version, name string
	module float64
)
*/
var (
	version, name, module = "3.4", "go proactice", 3.4
)

func main() {

    fmt.Println("name is ", name, "and is type is ", reflect.TypeOf(name))
    fmt.Println("name is ",module, "and is type is ", reflect.TypeOf(module))
    fmt.Println("name is ", name, " version = ", version, " module= ", module)

	a := "a"
	fmt.Println(a)
	changeVal(&a)
	fmt.Println(a)
}

func changeVal(course *string) string {
   *course = "b"
   fmt.Println(b)
   return course
}