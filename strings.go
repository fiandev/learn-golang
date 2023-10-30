package main

import (
  "fmt"
  "strings"
)

func print (anything any) {
  fmt.Println(anything)
}

func main() {
  /**
   * built-in method unruk tipe data string
   * 
   */
  text := "hello world!"
  
  /**
   * .Compare() membandingkan 2 string sama atau tidak
   */
  print (strings.Compare("foo", "foo")) // sama return 0
  print (strings.Compare("hello", text)) // tidak sama return -1
  
  /**
   * .Contains mengecek apakah substring ada pada suatu string
   */
  print (strings.Contains(text, "hello")) // return true jika ada
  print (strings.Contains("halo dunia", "hello")) // return false jika tidak ada
}