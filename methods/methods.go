//Go does not have classes. However, you can define methods on types.
//A method is a function with a special receiver argument.
//The receiver appears in its own argument list between the func keyword and the method name.
//You can only declare a method with a receiver whose type is defined in the same package as the method.
//methods with value receivers take either a value or a pointer as the receiver when they are called:
package main

import "fmt"

type Mystruct struct {
	x, y int
}

type Myfloat float64

//Methods are functions with special reciever types
func (s Mystruct) Add() int {
	return s.x + s.y
}

//Methods can be declared on non struct types too
func (f *Myfloat) Abs() Myfloat {
	if *f < Myfloat(0) {
		*f = -(*f)
	}
	return *f
}

//Methods with value recievers cannot change the value[call by value]
func (s Mystruct) ScaleByVal(sf int) {
	s.x = s.x * sf
	s.y = s.y * sf
}

//Methods with pointer reciever can change the input argument value[call by reference]
//Methods function with pointer reciever can be called by a value and a pointer of appropriate type
func (s *Mystruct) Scale(sf int) {
	s.x = s.x * sf
	s.y = s.y * sf
}

//function with pointer argument should be called with a pointer argument but this is not the case with methods as go allows implicit conversion here for method
func AddV(s Mystruct) int {
	return s.x + s.y
}

func AddP(s *Mystruct) int {
	return s.x + s.y
}

func main() {
	v := Mystruct{3, 4}
	var mf Myfloat = -43.43567

	//Call by value
	v.ScaleByVal(10)
	fmt.Println(v)

	//Call by reference
	v.Scale(10)
	fmt.Println(v)

	//Cann on non struct type and call by reference
	fmt.Println(mf.Abs())

	//Function call by pointer or reference
	fmt.Println(AddP(&v))

	//function call by value
	fmt.Println(AddV(v))

	//methods with value receivers take either a value or a pointer as the receiver when they are called:
	fmt.Println(v.Add())
	pv := &v
	fmt.Println(pv.Add())
}
