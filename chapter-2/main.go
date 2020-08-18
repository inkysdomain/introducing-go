package main

import "fmt"

func main() {
/*  fmt.Println("1 + 1 =", 1 + 1)
  fmt.Println("1.0 + 1.0 =", 1.0 + 1.0)
*/
/*  fmt.Println(len("Hello, World"))
  fmt.Println("Hello, World"[1])
  fmt.Println("Hey"[1])
  fmt.Println("There"[1])
  fmt.Println("Hello, " + "World")
*/
  fmt.Println("true && true:", true && true)
  fmt.Println("true && false:", true && false)
  fmt.Println("true || true:", true || true)
  fmt.Println("true || false:", true || false)
  fmt.Println("!true:", !true)
}

/*
1. How are integers stored on a computer?
Integers are stored in a computer by counting up in binary, adding a digit for every time the number is filled with ones.

2. We know that (in base 10) the largest one-digit number is 9 and the largest two-
digit number is 99. Given that in binary the largest two-digit number is 11 (3),
the largest three-digit number is 111 (7) and the largest four-digit number is 1111
(15), what’s the largest eight-digit number? (Hint: 10^1−1 = 9 and 10^2−1 = 99)
2^8-1 = 255

3. Although overpowered for the task, you can use Go as a calculator. Write a pro‐
gram that computes 32,132 × 42,452 and prints it to the terminal (use the * oper‐
ator for multiplication).
4. What is a string? How do you find its length?
A string is a sequence of characters with a definite length
used to represent text. You can find its length by running
  fmt.Println(len("Desired String"))

5. What’s the value of the expression (true && false) || (false && true) || !
(false && false)?
False.
*/
