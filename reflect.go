package golanglearning

import (
	"fmt"
	"reflect"
)

type MyInfo struct{
	Name string
	Age  int
}
func ReflectTest(){
	fmt.Println("reflect function ==========================")

	golang := MyInfo{"Golang", 7}
	//fmt.Println(golang)
	whoami(golang)
}

func whoami(itf interface{}){
	typ := reflect.ValueOf(itf).Interface()
	types := reflect.TypeOf(typ);
	fmt.Printf("member\tname\ttype\tvalue\n");
	for i:=0;i<types.NumField();i++ {
		f := reflect.ValueOf(itf).Field(i);
		name := types.Field(i).Name
		fmt.Printf("%d\t%s\t%s\t%v\n",i,name,f.Type(),f.Interface());
	}


	var x int = 1
	fmt.Println("ValueOf is: ", reflect.ValueOf(x).Kind())
	//fmt.Println("ValueOf is: ", itf.Kind())
	//fmt.Println("ValueOf is: ", reflect.ValueOf(x)+1)
	//fmt.Println("ValueOf is: ", reflect.TypeOf(reflect.ValueOf(x)))
}