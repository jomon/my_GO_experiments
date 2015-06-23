package main
import "fmt"
 func main () {
   fmt.Println("Enter the temperature in Fahrenheit :")
   var x int
   fmt.Scanf("%d",&x)
   fmt.Println(" In celsius is:",((x - 32) * 5/9 ) )
   y:= .32
   result:= float64(x)*y
   fmt.Println("double temp is :",result )

 }