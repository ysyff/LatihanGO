package main

import "fmt"

func main() {
	// Menampilkan tipe data dari variabel i
	var i int = 23
	fmt.Printf("Tipe data dari variabel i: %T\n", i)

	// Menampilkan nilai boolean j: true
	var j bool = true
	fmt.Printf("Menampilkan nilai boolean j: %t\n", j)
	
	// Menampilkan unicode Russia: Я (ya)
	fmt.Printf("Menampilkan unicode Russia: %c\n", '\u042F')

	// Menampilkan nilai base 10: 21
	var n int = 21
	fmt.Printf("Menampilkan nilai base 10: %d\n", n)

	// Menampilkan nilai base 8: 25
	fmt.Printf("Menampilkan nilai base 8: %o\n", n)

	// Menampilkan nilai base 16: f
	fmt.Printf("Menampilkan nilai base 16: %x\n", 15)

	// Menampilkan nilai base 16: F13
	fmt.Printf("Menampilkan nilai base 16: %X\n", 0xF13)

	// Menampilkan unicode karakter Я: U+042F
	fmt.Printf("Menampilkan unicode karakter Я: %U\n", 'Я')

	// Menampilkan float64 k: 123.456000
	var k float64 = 123.456
	fmt.Printf("Menampilkan float: %f\n", k)

	// Menampilkan float64 k dalam notasi scientific: 1.234560e+02
	fmt.Printf("Menampilkan float scientific: %e\n", k)

	
}