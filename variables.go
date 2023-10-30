package main
import "fmt"

func print (anything any) {
  fmt.Println(anything)
}

func main() {
  /**
   * mendefinisikan variabel berserta dsta tipe nya
   * data tipe dapat berupa 
   * - string
   * - int
   * - float
   * - bool
   * tipe data yang didefinisikan secara implisit/eksplisit sama² tidak dapat dirubah
   */
  var name string = "Aditia Akbar"
  var 
      username string = "fian"
  var isSuccess bool = true // default: false when the value is doesn't signed
  var PI float32 = 3.14 // default: 0 when the value is doesn't signed
  var a, b, c, d int = 1, 2, 3, 4
  var (
    e int
    f string = "hello"
    g bool = true
  )
  
  /**
   * mendefinisikan variabel sesuai dengan value yang di berikan secara implisit
   * - hanya dapat digunakan didalam function
   * - hanya bisa ditulis dalam satu bsris yang sama
   */
  age := 1
  print (username)
  print (age)
  print (name)
  print (isSuccess)
  print (PI)
  print (a)
  print (b)
  print (c)
  print (d)
  print (e)
  print (f)
  print (g)
}