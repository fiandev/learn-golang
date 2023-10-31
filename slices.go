package main
import "fmt"

func print (anything any) {
    fmt.Println(anything)
}

func br () {
    fmt.Println("\n")
}

func main() {
  /**
   * slice
   * slice hampir sama seperti array yang membedakan adalah
   * - ukuran slice dapat ditambah maupun dikurangi
   * 
   * # fungsi utilitas
   * len() - mendapatkan nilai ukuran dari sebuah array/slice
   * cap() - mendapatkan nilai kapasitas dari sebuah array/slice
   */
  
  /**
   * membuat slice
   */
  items := []int{1, 2, 3}
  
  print (items)
  
  br()
  /**
   * membuat slice dari sebuah array
   * # contoh
   * arr[start:end]
   */
  arr := [...]int{1,2,3,4,5}
  items2 := arr[0:2]
  
  print (items2)
  
  /**
   * membuat slice menggunakan fungsi make()
   * # contoh
   * make([]<tipe data>, <ukuran>, <kapasitas>)
   */
  
  items3 := make([]string, 3, 10)
  
  print (items3)
  
  br()
  
  fmt.Printf("array %v cap(arr) => %v \n", arr, cap(arr))
  fmt.Printf("array %v len(arr) => %v \n", arr, len(arr))
  fmt.Printf("slice %v len(items3) => %v \n", items3, len(items3))
  fmt.Printf("slice %v cap(items3) => %v \n", items3, cap(items3))
}