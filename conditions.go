package main
import "fmt"

func print (anything any) {
    fmt.Println(anything)
}

func br () {
    fmt.Println("\n")
}

func main() {
  value := 11
  username := "hehe"
  /**
   * if, else if and else keywords
   */
  if value > 5 {
    if value > 10 {
      print ("value lebih dari 5 dan 10")
    } else {
      print ("value lebih dari 5")
    }
  } else if value > 0 {
    print ("value lebih dari 0")
  } else {
    print ("nilai value kurang dari 0")
  }
  
  /**
   * switch, case, default keywords
   */
  
  switch username {
    case "fiandev":
      print ("username adalah fiandev")
    case "fian":
      print ("username adalah fian")
    case "anu", "hehe":
      print ("username adalah anu / hehe")
    default:
      print ("tidak ada kondisi yang terpenuhi")
  }
}