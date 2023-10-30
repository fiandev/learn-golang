package main
import "fmt"

func print (anything any) {
  fmt.Println(anything)
}

func main() {
  const A int = 18
  const (
    B string = "hello"
    C string = "world"
    D bool = true
  )
  const E = "tipe data eksplisit terhadap nilai yang diberikan"
  
  print (A)
  print (B)
  print (C)
  print (D)
  print (E)
}