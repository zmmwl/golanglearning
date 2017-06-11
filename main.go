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

	var cmplx1 complex128
	cmplx1 = 32 + 12i
	cmplx2 := 23+ 23i
	cmplx3 := cmplx1 + cmplx2

	fmt.Println(cmplx3)

	var str string
	str = "Hello Tyler"
	fmt.Println(str)
	fmt.Println(str[0])
	fmt.Printf("str[0] is: %s,%t,%d,%c\n", str[0],str[0],str[0],str[0])

	fmt.Println("//////////////i is: ",i)
	for i:=0; i<len(str); i++ {
		fmt.Printf("%c\n", str[i])
		fmt.Println("-----------i is: ",i)
	}
	fmt.Println("//////////////i is: ",i)

	var i2 int
	i2 = 23
	//i2 := 43  /* can not := to an existing var */
	fmt.Println(i2)

	for i, ch := range str{
		fmt.Println(i,"is: ",ch)
	}

	var its [3]int
	its = [3]int{6,2,3}
	fmt.Println(its)

	for i, v := range its{
		fmt.Println(i,v)
	}

	ModifyArray(its)
	fmt.Println(its)


	var its2 []int
	its2 = its[:]
	ModifyArray2(its2)
	fmt.Println(its2)


	var mySlice = make ([]int,5,10)
	mySlice[0] = 1
	//mySlice[9] = 1
	mySlice = append(mySlice,6,7,8,9,10,11,12)
	fmt.Println(mySlice)

	type Person struct{
		Id string
		Name  string
	}

	var persons map[string] Person
	persons = make(map[string] Person)
	persons["1"] = Person{"1","Tyler"}
	persons["二"] = Person{"1","Ivy"}
	fmt.Println(persons)

	fmt.Println(TestIf());

	TestNolimitFunc(1,2,3)
	TestNolimitFunc2(Person{"1","Tyler"},2,3)


	var a = func(){
		fmt.Println(i*3)
	}

	a()
	i++
	a()

	/******************* The End! ********************/
	fmt.Println("\n")
	fmt.Println("\n")
	fmt.Println("\n")


}

func TestNolimitFunc2(args ...interface{}){
	fmt.Println(args)

	for _, arg := range args {
		fmt.Println(arg)
		switch arg.(type) {
		case int: fmt.Println("int type")
		}
	}
}
func TestNolimitFunc(args ...int){
	fmt.Println(args)
}

func TestIf()int{
	if 1<2 {
		return 0;
	}else{
		return 1;
	}
}

func ModifyArray2( its []int){
	its[0] = its[0]+1
}
func ModifyArray( its [3]int){
	its[0] = its[0]+1
}

func M1( i int,  j int)(x int, y int){
   return j+1,i+2;
}