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
    * mendeklarasikan array
    * - ukuran array didefinisikan
    * var array_name = [<ukuran array>]<tipe data>{values}
    * - ukuran array tidak didefinisikan
    * var array_name = [...]<tipe data>{values} 
    */
    
    var arr1 = [3]int{1, 2, 3}
    var arr2 = [6]int{4, 5, 6}
    
    arr3 := [...]int{7, 8, 9}
    
    print (arr1)
    print (arr2)
    print (arr3)
    
    br()
    /**
    * Initialize Only Specific Elements
    * [<ukuran array>]<tipe data>{<index>:<value>, ...dst}
    */
    
    items := [5]string{0:"xiomi", 2:"samsung", 4:"iphone"}
    numbers := [5]int{0:10, 2:20, 4:30}
    
    print (items)
    print (numbers)
    
    br()
    /**
    * Find the Length of an Array
    * len(<variabel>)
    */
    fmt.Printf("array items %v len(items) => %v \n", items, len(items))
    fmt.Printf("array arr3 %v len(arr3) => %v \n", arr3, len(arr3))
}