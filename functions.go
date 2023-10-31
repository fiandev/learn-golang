package main
import "fmt"

func print (anything any) {
  fmt.Println(anything)
}

/**
 * Basic Function Return
 */
func sqrt (value int) (result int) {
  for i := 1; i <= value; i++ {
    if i * i == value {
      result = i
      break
    }
  }
  
  return result
}

/**
 * Function Multiple Return Values
 */
func calcultate (x int, y int) (a int, b int) {
  a = x + y
  b = x - y
  
  return
}

func main() {
  fmt.Printf("√100 = %v \n", sqrt(100))
  fmt.Printf("√4 = %v \n", sqrt(4))
  fmt.Printf("√1024 = %v \n", sqrt(1024))
  
  a, b := calcultate(10, 20)
  
  fmt.Printf("a = %v \n", a)
  fmt.Printf("b = %v \n", b)
}