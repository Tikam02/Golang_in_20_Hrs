
// Struct Literals

// A struct literal denotes a newly allocated struct value by listing the values of its fields.

// You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)

// The special prefix & returns a pointer to the struct value.
package main

import "fmt"

type void struct{
	x,y int
}

var (
	v1 = void{1,2}  //has type void
	v2 = void{x: 1} //Y:0 is implicit
	v3 = void{}     // x:0 and Y:0
	p = &void{1,2}  //has type *Void
)

func main(){
	fmt.Println(v1,p,v2,v3)
}