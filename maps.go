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
   * map
   * digunakan untuk menyimpan banyak data dalam satu variabel
   * - seperti layaknya array, tetapi memiliki key yang dapat berupa string
   * - seperti array assosiatif pada bahasa PHP, tetapi dapat memiliki tipe data untuk yang lebih bervariatif
   * 
   * tipe sata yang diperbolehkan untuk key
   * -  Booleans
   * -  Numbers
   * -  Strings
   * -  Arrays
   * -  Pointers
   * -  Structs
   * -  Interfaces (selama tipe dinamis mendukung kesetaraan)
   * 
   * # format syntax
   * 1. map[<key type>]<value type>{<value>, ...dst}
   * 3. make(map[<key type>]<value type>)
   */
   
  a := map[string]string{"brand": "xiomi", "type": "smartphone", "model": "poco X5 pro"}
  b := make(map[int]string)
  
  b[0] = "redmi"
  b[1] = "poco"
  b[2] = "mi"
  
  /**
   * mengecek element pada map secara spesifik
   * val, ok := <map>[<key>]
   */
  
  key := "brand"
  value, ok := a[key]
  
  if ok {
    fmt.Printf("value key %v pada map a adalah = %v \n", key, value)
    
  } else {
    fmt.Printf("value key %v pada map a tidak ada \n", key)
    
  }
  
  br()
  /**
   * Iterate Over Maps
   */
  
  for key, value := range a {
    fmt.Printf("a['%v'] = %v \n", key, value)
  }
  
  for _, v := range b {
    print (v)
  }
  
  br()
  /**
   * delete item by key
   * 
   * delete(<map>, <key>)
   */
  delete(b, 0)
  delete(a, "type")
  
  print (a)
  print (b)
}