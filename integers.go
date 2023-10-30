package main
import "fmt"

func print (anything any) {
  fmt.Println(anything)
}

func main() {
  /**
   * tipe integer memiliki dua tipe
   * - unsigned (tidak dapat menyimpan nilai negatif)
   * # dituliskan int<besar bit>
   * int   -2147483648 - 2147483647 pada sistem 32 bit
   *       -9223372036854775808 - 9223372036854775807 pada sistem 64 bit
   * int8  -128 - 127
   * int16 -32768 - 32767
   * int32 -2147483648 - 2147483647
   * int64 -9223372036854775808 to 9223372036854775807
   * 
   * - signed (dapat menyimpan nilai negatif maupun positif)
   * # dituliskan uint<besar bit>
   * uint   0 to 4294967295 - 32 bit pada sistem 32 bit
            0 to 18446744073709551615 - 64 pada sistem 64 bit
   * uint8  0 - 255
   * uint16 0 - 65535
   * uint32 0 - 4294967295
   * uint64 0 - 18446744073709551615
   */
   
   var intMaxNegativeValue int = -2147483648
   var intMaxPostiveValue int = 2147483647
   
   var uintMaxValue uint = 4294967295
   var uintMinValue uint = 0
   
   print (intMaxNegativeValue)
   print (intMaxPostiveValue)
   print (uintMaxValue)
   print (uintMinValue)
}