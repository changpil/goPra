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

	a := 10.0
	b := 3
	c := int(a)+b
	d := a + float64(b)
	fmt.Println("C is ", c, "d is ", d, " c type: ", reflect.TypeOf(c), " d type: ", reflect.TypeOf(d))

	
	num :=  strconv.Itoa(-42)
	num_string, err := strconv.Atoi("-42") 
	
}
