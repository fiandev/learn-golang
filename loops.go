package main

import "fmt"

func print (anything any) {
  fmt.Println(anything)
}

func main() {
  each := 3000
  items := [5]string{"apple", "melon", "banana", "mango", "orange"} 
  
  /**
   * The For Keyword
   */
  for i := 0; i < each; i++ {
    print ("i love you :)")
  }
  
  for i := 0; i < 100; i++ {
    if i % 2 == 0 {
      fmt.Printf("%v adalah bilangan genap \n", i)
    } else {
      continue
    }
  }
  
  /**
   * The Range Keyword
   */
  for i, v := range items {
    fmt.Printf("index ke %v adalah %v \n", i, v)
  }
}