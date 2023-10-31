package main
import "fmt"

func print (anything any) {
    fmt.Println(anything)
}

func br () {
    fmt.Println("\n")
}

func main() {
  const VALUE = 100
  var (
    value int = VALUE
    score float32 = 87.5
  )
  
  /**
   * artimetics operator
   * +  | addition (penjumlahan)
   * -  | subtraction (pengurangan)
   * /  | division (pembagian)
   * *  | multiplication (perkalian)
   * %  | modulus (persen bagi)
   * ++ | increment (penambahan nilai 1)
   * -- | decrement (pengurangan nilai 1)
   */
  
  fmt.Printf("1 + 1 = %v \n", 1 + 1)
  fmt.Printf("1 - 1 = %v \n", 1 - 1)
  fmt.Printf("10 / 2 = %v \n", 10 / 2)
  fmt.Printf("2 * 5 = %v \n", 2 * 5)
  fmt.Printf("100 %% 7 = %v \n", 100 % 7)
  
  score++
  fmt.Printf("score++ = %v \n", score)
  score--
  fmt.Printf("score-- = %v \n", score)
  
  br()
  /**
   * assignment operators
   * =	  x = 5	    x = 5	
   * +=	  x += 3  	x = x + 3	
   * -=	  x -= 3  	x = x - 3	
   * *=	  x *= 3  	x = x * 3	
   * /=	  x /= 3  	x = x / 3	
   * %=	  x %= 3	  x = x % 3	
   * &=	  x &= 3	  x = x & 3	
   * |=	  x |= 3	  x = x | 3	
   * ^=	  x ^= 3	  x = x ^ 3	
   * >>=	x >>= 3	  x = x >> 3	
   * <<=	x <<= 3 	x = x << 3
   */
  
  value >>= 2
  fmt.Printf("%v | value >>= 2 %v \n", VALUE, value)
  
  value = VALUE
  value <<= 2
  fmt.Printf("%v | value <<= 2 %v \n", VALUE, value)
  
  br()
  /**
   * Comparison operators
   * >  lebih  dari
   * <  kurang dari
   * >= lebih dari sama dengan
   * <= kurang dari sama dengan
   * == sama dengan
   * != tidak sama dengan
   */
  
  print (10 > 5) // true
  print (10 < 5) // false
  print (10 <= 10) // true
  print (1 >= 0) // true
  print (10 == 5) // false
  print (10 != 5) // true
  
  br()
  /**
   * Logical Operators
   * && logical and
   * || logical or
   * !  logical not
   */
  
  print (1 != 0 && 0 < 1) // false
  print (1 > 0 || 10 < 5) // true
  print (!(1 == 1)) // false
  
  br()
  /**
   * Bitwise Operators
   * &  AND
   * |  OR
   * ^  XOR
   * << Zero fill left shift
   * >> Signed right shift
   */
   
   var x, y, z int = 1, 10, 100
   
   fmt.Printf("x = %v \ny = %v \nz = %v \n", x, y, z)
   
   fmt.Printf("x & y = %v \n", x & y)
   fmt.Printf("x | y = %v \n", x | y)
   fmt.Printf("x ^ y = %v \n", x ^ y)
   fmt.Printf("y << z = %v \n", y << z)
   fmt.Printf("x >> z = %v \n", x >> z)
}