package main 
import (
	"fmt"
)
// type vertex struct{
// 	x int 
// 	y int
// }

// func (v vertex ) Abs()  int  {
// 	return v.x*v.y
// }

// (v vertex ) is a receiver of type so we add a method to this 
func main() {

	// i:=0
//  so basically a keynote in go is that if we  decalared a var in go we have to use it or we can't compile 

    //  var arr [5]int 
	//  arr[0] = 1
	//  fmt.Print(arr[0])
	//  b:=[]int{1 ,2 ,3 }
	//  fmt.Print(len(b))

	//  []int we have to use it as prefix if we directly assign values to a array 

//    temp := make([]int , 3 )
//    fmt.Print(temp)

// maps 

// m := make(map[string] string )
// m["name"] = "Nikunj"
// fmt.Print(m["name"])

// //  if key value does not exits it return empty string

// delete(m , "name")

// for delete
//   fmt.Print(split(10 ))
// //   so basically we can create a func in this format too 
// // the x ,y are initialised in the begining and this  func automatically detects the values needs to be return 
// 	//  we dont use this much 
	 
// }
//  func split(sum int )(x , y int ){
// 	x = sum * 2 
// 	y = sum + 10
// 	return 

//    i := 0
//    for i < 10 {
// 	 fmt.Println(i)
// 	 i++
//    }
// the while loop


// defer 
//  for example you open a file (function calls ) and now after  some times you forgot to close that functions so what you did is you use defer added the function call to defer after calling it , it has a stack memeory which takes account of everything

// type name struct{
// 	x int 
// 	y int 
// }
// fmt.Print(name{1,2 })
// v := name{1 ,3 }
// fmt.Print(v.x)

// so this is sturct in go


// arr := [5]int{10, 20, 30, 40, 50}
// 	s1 := arr[1:4] // 20,30,40
// 	s2 := arr[2:]  // 30,40,50

// 	fmt.Println("s1:", s1 , cap(s1))
// 	fmt.Println("s2:", s2 , cap(s2))
//  the zero value  slice is nil 

// v:= []int{10 , 20 ,30 }
// for i , m   := range v {
// 	fmt.Println(i , "  " , m )
// }


// maps

// var m map[int]int 
// m = make(map[int]int)
// m[10] = 90 
// fmt.Print(m[10])

// we can initialise a  map why literal symbol rightaway 

// v := vertex{3 , 4}

// fmt.Print(v.Abs())
// Intrrface in go

// type a interface {
//     abs() float64
// }
// Any type that wants to be considered an a must have a method named abs that takes no arguments and returns a float64 

// type MyNumber float64

// func (m MyNumber) abs() float64 {
//     if m < 0 {
//         return -float64(m)
//     }
//     return float64(m)
// } 

// a type alredy satisfy if it has a method and same return type so we dont write implements in go 


// var x a
// x = MyNumber(-5)

// fmt.Println(x.abs()) // 5

// type i interface{}
// we use this for using any type variable 

}


