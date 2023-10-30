package main
import "fmt"

func print (anything any) {
    fmt.Println(anything)
}

func br () {
    fmt.Println("\n")
}

func main() {
    username := "fiandev"
    PI := 3.14159265358
    isSuccess := true
    percentage := 100
    value := 10
    
    /**
    * general formatting verbs
    * %v  - menampilkan nilai format secara default
    * %#v - menampilkan nilai format sesuai syntax go
    * %T  - menampilkan tipe data dari valuee
    * %%  - menambahkan simbol %
    */
    fmt.Printf("text %v (%%v) => %v \n", username, username)
    fmt.Printf("text %v (%%#v) => %#v \n", username, username)
    fmt.Printf("text %v (%%T) => %T \n", username, username)
    fmt.Printf("nilai %v (%%%%) => %v%% \n", percentage, percentage)
    
    br()
    /**
    * Integer Formatting verbs
    * %b   - base 2
    * %d   - base 10
    * %+d  - base 10 dan selalu menambahkan '+' di depannya
    * %o   - base 8
    * %O   - base 8 dan selalu menambahkan '0o' di depannya
    * %x   - base 16 lowercase
    * %X   - base 16 uppercase
    * %#x  - base 16 dan selau menambahkan '0x' di depannya
    * %4d  - menambahkan padding dengan width '4' pada sebelah kanan
    * %-4d - menambahkan padding dengan width '4' pada sebelah kiri
    * %04d - menambahkan Padding dengan width '4' dan menambahkan '0' pada bagian yang ksoong
    */
    
    fmt.Printf("nilai %v (%%b) => %b \n", value, value)
    fmt.Printf("nilai %v (%%d) => %d \n", value, value)
    fmt.Printf("nilai %v (%%+d) => %+d \n", value, value)
    fmt.Printf("nilai %v (%%o) => %o \n", value, value)
    fmt.Printf("nilai %v (%%O) => %O \n", value, value)
    fmt.Printf("nilai %v (%%X) => %x \n", value, value)
    fmt.Printf("nilai %v (%%X) => %X \n", value, value)
    fmt.Printf("nilai %v (%%#x) => %#x \n", value, value)
    fmt.Printf("nilai %v (%%4d) => %4d \n", value, value)
    fmt.Printf("nilai %v (%%-4d) => %-4d \n", value, value)
    fmt.Printf("nilai %v (%%04d) => %04d \n", value, value)

    br()
    /**
    * String Formatting verbs
    * %s   - menampilkan nilai sebagai string
    * %q   - menampilkan nilai dengan diapit dengan dua tanda petik '"..."'
    * %8s  - menambahkan padding dengan width '8' pada sebelah kiri
    * %-8s - menambahkan padding dengan width '8' pada sebelah kanan
    * %x   - mengubah string menjadi bit hexadecimal
    * % x  - mengubah string menjadi bit hexadecimal dengan dipisahkan spasi tiap dua digit nilai
    */
    
    fmt.Printf("text %v (%%s) => %s \n", username, username)
    fmt.Printf("text %v (%%q) => %q \n", username, username)
    fmt.Printf("text %v (%%8s) => %8s \n", username, username)
    fmt.Printf("text %v (%%-8s) => %-8s \n", username, username)
    fmt.Printf("text %v (%%x) => %x \n", username, username)
    fmt.Printf("text %v (%% x) => % x \n", username, username)
    
    br()
    /**
    * Boolean Formatting verbs
    * %t - menampilkan nilai boolean sama seperti menggunakan '%v'
    */
    fmt.Printf("nilai %v (%%t) => %t \n", isSuccess, isSuccess)
    
    br()
    /**
    * Float Formatting Verbs
    * %e    - mengubah nilai desimal kedalam bentuk eksponen
    * %f    - nilai desimal yang tidak di ubah ke bentuk eksponen
    * %.2f  - mengatur nilai presisi setelah koma sebnyak '2'
    * %6.2f - mengatur nilai presisi setelah koma sebnyak '2' dengan nilai width '6'
    * %g    - hanya menampilkan nilai digit eksponenyang dibutuhkan
    */
    fmt.Printf("nilai %v (%%e) => %e \n", PI, PI)
    fmt.Printf("nilai %v (%%f) => %f \n", PI, PI)
    fmt.Printf("nilai %v (%%.2f) => %.2f \n", PI, PI)
    fmt.Printf("nilai %v (%%6.2f) => %6.2f \n", PI, PI)
    fmt.Printf("nilai %v (%%g) => %g \n", PI, PI)
}