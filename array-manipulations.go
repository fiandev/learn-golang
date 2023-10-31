package main
import "fmt"

func print (anything any) {
    fmt.Println(anything)
}

func br () {
    fmt.Println("\n")
}

func main() {
  items := []int{1, 2, 3, 4, 5}
  items2 := items
  
  /**
   * change element of slice / array
   */
  
  items[0] = 11
  fmt.Printf("%v | items[0] = 0 => %v \n", items2, items)
  
  br()
  // Append Elements To a Slice
  fmt.Printf("%v | append(items2, 20) => %v \n", items2, append(items2, 20))
  
  slice1 := []string{"a", "b", "c"}
  slice2 := []string{"d", "e", "f"}
  slice3 := append(slice1, slice2...)
  // Append One Slice To Another Slice
  fmt.Printf("%v | append(slice1, slice2) => %v \n", slice1, slice3)
  
  // Change The Length of a Slice
  fmt.Printf("%v slice[0:3] => %v \n", slice3, slice3[0:3])
  fmt.Printf("length %v slice[0:3] => %v \n", len(slice3), len(slice3[0:3]))
  
  neededSlice := slice3[:len(slice3) - 3]
  sliceCopy := make([]string, len(neededSlice))
  
  fmt.Printf("sliceCopy = %v\n", sliceCopy)
  fmt.Printf("length = %d\n", len(sliceCopy))
  fmt.Printf("capacity = %d\n", cap(sliceCopy))
  
  print ("=== after ===")
  copy(sliceCopy, neededSlice)
  
  fmt.Printf("sliceCopy = %v\n", sliceCopy)
  fmt.Printf("length = %d\n", len(sliceCopy))
  fmt.Printf("capacity = %d\n", cap(sliceCopy))
  
}