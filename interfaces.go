package main
import "fmt"

func print (anything any) {
    fmt.Println(anything)
}

func br () {
    fmt.Println("\n")
}

type People struct {
  name string
  age uint
  characters []string
}
type Human interface {
  getName() string
  getAge() uint
}

func (people *People) getName (name string) string {
  people.name = name
  
  return name
}

func (people *People) getAge (age uint8) uint {
  people.age = age
  
  return age
}


func main() {
  var people Human
  
  people.name = "fiandev"
  people.age = 18
  
  
  print (people.getName())
}