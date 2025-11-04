package main

import "fmt"

func main() {
    // fmt.Println("Hello, World!")
	// // what i learn is go dev dont write err for println because printing never uasuallly fails 
	// fmt.Println("hello")
	// fmt.Println("Nikunj")
	//  morre like any other language 
	 


	// var name string
    // fmt.Print("Enter your name: ")
    // fmt.Scan(&name) // note the & (address of variable)
    // fmt.Println("Hello,", name)


	result := Hello("Hello World ")       // call your own Hello() function
    fmt.Println(result)


} 
func Hello(name string ) string{
		return name 
}