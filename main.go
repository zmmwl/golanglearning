package main

import "fmt"
import d1 "github.com/zmmwl/project1/d1"

func main() {
	fmt.Println("hello Tyler");
	fmt.Println(d1.Add(1,2));
	i:=1;
	j:=2;

	fmt.Println(i+j);

	b:=true;
	if b {
		fmt.Println("True!")
	}else{
		fmt.Println("False")
	}


	i, _ =M1(i,j)
	var b2 bool = false;

	fmt.Println("i is: ", i,b2)

	const(
		c0 = iota;
		c1 = iota;
		c2 = iota;
	)
	fmt.Println(c0,c1,c2)


	var value2 int32
	value1 := 23
	value1 = int(value2)
	fmt.Println(value1)



}

func M1( i int,  j int)(x int, y int){
   return j+1,i+2;
}