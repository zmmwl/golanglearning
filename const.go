package golanglearning

import "fmt"

const(
	Sunday = iota
	Monday
	Tuesday
	Wednesday = iota
	Thursday
	Friday
	Saturday
)
func ConstTest(){
	fmt.Println("Saturday is: ", Saturday)
}
