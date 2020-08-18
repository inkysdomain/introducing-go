package main
import "fmt"
func main() {
//  var chi string = "Hello, World"
//  fmt.Println(chi)
  var chi string
  chi = "an integer, I swear"
  fmt.Println("Chi is", chi)
//  var chi int6
//  chi = 42
  chi ="a float now"
  fmt.Println("Chi is", chi)
  chi += "this is bad grammar?"
  fmt.Println(chi)
  fmt.Println(chi=="seven")
  }

/*
1. What are two ways to create a new variable?
There are two ways to create new variables. Declaring one by using "var $varfiable_name
type = $data" is one way. Another way is to write "$VARIABLE := $data"

2. What is the value of x after running x := 5; x += 1?
6.
3. What is scope? How do you determine the scope of a variable in Go?
The scope of a variable is the range of places where it can be used.

4. What is the difference between var and const?
A var can be changed but a const remains the same.
#5 is in fahrenheit-celsius-converter.go
6 is in foot-meter-converter.go

*/
