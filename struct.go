package main
import "fmt"

func print (anything any) {
    fmt.Println(anything)
}

func br () {
    fmt.Println("\n")
}

/**
 * struct
 * digunakan untuk mengelompokkan beberapa variabel dengan tipe data yang berbeda dapam satu grup
 * - layak nya type pada typescript
 * - untuk cetak suatu object (class)
 */
 
type Person struct {
  name string
  age int
  hobbies []string
  job string
}

func (person *Person) addHobby (hobby string) []string {
  person.hobbies = append(person.hobbies, hobby)
  
  return person.hobbies
}

func (person *Person) changeName (name string) *Person {
  person.name = name
  
  return person
}

func main() {
  var me Person
  hobbies := []string{"code", "anime", "music"}
  
  me.name = "fian"
  me.age = 18
  me.job = "software engineer"
  
  for _, value := range hobbies {
    me.addHobby(value)
  }
  
  o := me.changeName("fiandev")
  
  print (me)
  
  // memiliki simbol & ??
  print (o)
}