package main
import "fmt"

func main(){
//  i := 1
//  for i <= 10 {
for i:=1; i<=10; i++{
  if i % 2 == 0 {
    fmt.Println(i, "Even")
  } else {
    fmt.Println(i, "Odd")
  }

}}
/*
1. What does the following program print?
i := 10
if i > 10 {
fmt.Println("Big")
} else {
fmt.Println("Small")
}
It prints Small.

2 Refer to exercise-2.go

3 Refer to exercise-3.go

*/
