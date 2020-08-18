/* 5. Using the example program as a starting point, write a program that converts
 from Fahrenheit into Celsius (C = (F − 32) * 5/9). */

 package main
 import "fmt"
 func main() {
   fmt.Print("Enter a number in Fahrenheit: ")
   var input float64
   fmt.Scanf("%f", &input)
   output := ((input - 32) * 5 /9)

   fmt.Println(output)
 }
