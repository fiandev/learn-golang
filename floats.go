package main
import "fmt"

func print (anything any) {
  fmt.Println(anything)
}

func main() {
  /**
   * tipe data float menyimpan nilai number yang berupa desimal
   * - tetap dapat menyimpan nilai integer (bilangan bulat)
   *   karna akan di terjemahkan ke desimal (5) => 5.0
   * - dapat menyimpan nilai desimal positif/negatif 
   * 
   * # float memiliki dua tipe
   * float32 -3.4e+38 - 3.4e+38.
   * float64 1.7e+308 - +1.7e+308.
   */
   
   var float32MaxPositiveValue float32 = 3.4e+38
   var float32MaxNegativeValue float32 = -3.4e+38
   var test = 5
   
   print (float32MaxPositiveValue)
   print (float32MaxNegativeValue)
   print (test)
   
}