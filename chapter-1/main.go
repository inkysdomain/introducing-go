package main

import "fmt"

// Creates a program that references the fmt library in order to
// create the function main with no arguments.
// It calls upon Println from fmt to print two arguments,
// "Hello, World!" and "What's up?"

func main() {
  fmt.Println("Hello, World!","What's up?")
}

//syntax for searching document "go doc fmt.Println "

/*1. What is whitespace?
Newlines, spaces, and tabs (because you can't see them)

2. What is a comment? What are the two ways of writing a comment?
A comment is a statement or text that is ignored by the compiler. They can be
writtin in one line with // preceding the text, or in a block by enclosing the
text with /* and */

/*
3. Our program began with package main. What would the files in the fmt package
begin with?
Files in the fmt package being with "package fmt".

4. We used the Println function defined in the fmt package. If you wanted to use
the Exit function from the os package, what would you need to do?
import "os"
func main(){
  os.Exit()
}

5. Modify the program we wrote so that instead of printing Hello, World it prints
Hello, my name is followed by your name.

package main

import "fmt"
func main() {
  fmt.Println("Hello, my name is","Eva!")
}
*/"
