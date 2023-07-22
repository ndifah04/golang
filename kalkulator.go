package main

import "fmt"

func main() {

	var (
		a     int
		b     int
		hasil int
		op    string
	)

	fmt.Println("AYO COBA KALKULATOR SEDERHANA YUK")
	fmt.Println("pertama-tama pilih operator yang ingin kamu gunakan")
	fmt.Println("1.+ \n2.-\n3.:\n4.x")
	fmt.Scan(&op)

	if op == "1" {
		fmt.Print("masukkan nilai pertama: ")
		fmt.Scan(&a)

		fmt.Print("masukkan nilai kedua: ")
		fmt.Scan(&b)
		hasil = a + b
		fmt.Print(a, "+", b, "=", hasil)
	} else if op == "2" {
		fmt.Print("masukkan nilai pertama: ")
		fmt.Scan(&a)

		fmt.Print("masukkan nilai kedua: ")
		fmt.Scan(&b)
		hasil = a - b
		fmt.Print(a, "-", b, "=", hasil)
	} else if op == "3" {

		fmt.Print("masukkan nilai pertama: ")
		fmt.Scan(&a)

		fmt.Print("masukkan nilai kedua: ")
		fmt.Scan(&b)
		hasil = a / b
		fmt.Print(a, ":", b, "=", hasil)
	} else if op == "4" {
		fmt.Print("masukkan nilai pertama: ")
		fmt.Scan(&a)

		fmt.Print("masukkan nilai kedua: ")
		fmt.Scan(&b)
		hasil = a * b
		fmt.Print(a, "x", b, "=", hasil)
	} else {
		fmt.Print("input salah")
	}
}
