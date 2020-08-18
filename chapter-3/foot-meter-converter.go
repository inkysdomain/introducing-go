//6. Write another program that converts from feet into meters (1 ft = 0.3048 m).

package main
import "fmt"
func main() {
  fmt.Print("Enter a number in feet: ")
  var input float64
  fmt.Scanf("%f", &input)
  output := input * 0.3048

  fmt.Println(output)
}
