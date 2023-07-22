// package main

// import "fmt"

// func main() {

// var angka int
// angka = 2

// var angka int = 3

// var angka = 3

// angka := 3

// var (
// 	angka  = 2
// 	angka2 = 3
// 	angka3 = 4
// )

// fmt.Println(angka)
// fmt.Println(angka2)
// fmt.Println(angka3)

// }
// ------------------------------------------------------------

// package main

// import "fmt"

// func main() {
// 	var benar bool
// 	var (
// 		angka  = 5
// 		angka2 = 4
// 	)
// 	benar = angka > angka2
// 	fmt.Println(benar)
// }
// ------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var jari float32 = 4
// 	var phi float32 = 3.14
// 	var hasil = phi * jari * jari

// 	fmt.Println(hasil)
// }
// ------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	const (
// 		nama        = "uki"
// 		jenis_rokok = "dji sum soe"
// 	)
// 	fmt.Println(nama)
// 	fmt.Println(jenis_rokok)
// }
// ------------------------------------------------------------
// package main

// import "fmt"

// func main() {

// 	var r float32
// 	var hasil float32

//		fmt.Print("masukksan nilai jari jari: ")
//		fmt.Scanln(&r)
//		hasil = 3.14 * r * r
//		fmt.Println("luas lingkarang adalah", hasil)
//	}
//
// ------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var nilai int

// 	fmt.Print("masukkan nilai anda:")
// 	fmt.Scan(&nilai)
// 	if nilai >= 70 {
// 		fmt.Println("selamat anda lulus")
// 	} else {
// 		fmt.Println("anda tidak lulus")
// 	}
// }
// ------------------------------------------------------------

// package main

// import "fmt"

//	func main() {
//		var nilai int
//		fmt.Print("masukkan nilai anda:")
//		fmt.Scan(&nilai)
//		if nilai >= 90 {
//			fmt.Println("selamat nilai anda A")
//		} else if nilai >= 80 {
//			fmt.Println("selamat nilai anda B")
//		} else if nilai >= 70 {
//			fmt.Println("selamat nilai anada C")
//		} else {
//			fmt.Println("Eror")
//		}
//	}
//
// ------------------------------------------------------------
// package main

// import "fmt"

//	func main() {
//		var nilai int
//		fmt.Print("masukkan nilai=")
//		fmt.Scan(&nilai)
//		if nilai%2 == 0 {
//			fmt.Println(nilai, "adalah nilai genap")
//		} else {
//			fmt.Println(nilai, " adalah nilai ganjil")
//		}
//	}
//
// ------------------------------------------------------------
// package main

// import "fmt"

// func main() {

// 	var (
// 		jenis       string
// 		jumlah      float32
// 		total_bayar float32
// 	)
// 	fmt.Println("1.sempurna 30000")
// 	fmt.Println("2.surya 32000")
// 	fmt.Println("3.rokok na uki 60000")

// 	fmt.Print("masukkan jenis=")
// 	fmt.Scan(&jenis)

// 	if jenis == "1" {
// 		fmt.Print("masukkan jumlah=")
// 		fmt.Scan(&jumlah)
// 		total_bayar = 30000 * jumlah
// 		fmt.Println("total bayarx adalah", total_bayar)
// 	} else if jenis == "2" {
// 		fmt.Print("masukkan jumlah=")
// 		fmt.Scan(&jumlah)
// 		total_bayar = 32000 * jumlah
// 		fmt.Println("total bayarx adalah", total_bayar)
// 	} else if jenis == "3" {
// 		fmt.Print("masukkan jumlah=")
// 		fmt.Scan(&jumlah)
// 		total_bayar = 60000 * jumlah
// 		fmt.Print("total bayarx adalah", total_bayar)
// 	} else {
// 		fmt.Println("jenis rokokx gak ada bestie")
// 	}

// }
// ------------------------------------------------------------

// package main

// import "fmt"

// func main() {
// 	var (
// 		jenis       string
// 		harga       int
// 		jumlah      int
// 		total_harga int
// 	)
// 	fmt.Print("jenis rokok \n 1. sempurna \n 2. surya \n 3.rokok na uki \n")
// 	fmt.Print("pilih jenis rokok")
// 	fmt.Scan(&jenis)

// 	switch jenis {
// 	case "1":
// 		harga = 30000
// 		fmt.Print("masukkan jumlah")
// 		fmt.Scan(&jumlah)
// 		total_harga = harga * jumlah
// 		fmt.Println("total harga adalah ", total_harga)
// 	case "2":
// 		harga = 32000
// 		fmt.Print("masukkan jumlah")
// 		fmt.Scan(&jumlah)
// 		total_harga = harga * jumlah
// 		fmt.Println("total harga adalah ", total_harga)
// 	case "3":
// 		harga = 60000
// 		fmt.Print("masukkan jumlah")
// 		fmt.Scan(&jumlah)
// 		total_harga = harga * jumlah
// 		fmt.Println("total harga adalah ", total_harga)
// 	default:
// 		fmt.Println("iput salah")
// 	}
// }

// ------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var nilai1 int
// 	var nilai2 int
// 	var operator string
// 	var hasil int
// 	fmt.Print("pilih operator \n 1.+ \n 2.* \n 3. / \n 4.- \n")
// 	fmt.Print("masukkan perator :")
// 	fmt.Scan(&operator)

// 	fmt.Print("masukkan nilai 1=")
// 	fmt.Scan(&nilai1)
// 	fmt.Print("masukkan nilai 2=")
// 	fmt.Scan(&nilai2)

// 	if operator == "1" {
// 		hasil = nilai1 + nilai2
// 		fmt.Println(hasil)

// 	} else if operator == "2" {
// 		hasil = nilai1 * nilai2
// 		fmt.Println(hasil)

// 	} else if operator == "3" {
// 		hasil = nilai1 / nilai2
// 		fmt.Println(hasil)

// 	} else if operator == "4" {
// 		hasil = nilai1 - nilai2
// 		fmt.Println(hasil)

// 	} else {
// 		fmt.Println("input salah")
// 	}

// }
// ------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var nilai1 int
// 	var nilai2 int
// 	var operator string
// 	var hasil int

// 	fmt.Print("pilih operator \n 1.+ \n 2.* \n 3. / \n 4.- \n")
// 	fmt.Print("masukkan perator :")
// 	fmt.Scan(&operator)

// 	fmt.Print("masukkan nilai 1= ")
// 	fmt.Scan(&nilai1)
// 	fmt.Print("masukkan nilai 2=  ")
// 	fmt.Scan(&nilai2)

//		switch operator {
//		case "1":
//			hasil = nilai1 + nilai2
//			fmt.Println(hasil)
//		case "2":
//			hasil = nilai1 * nilai2
//			fmt.Println(hasil)
//		case "3":
//			hasil = nilai1 / nilai2
//			fmt.Println(hasil)
//		case "4":
//			hasil = nilai1 - nilai2
//			fmt.Println(hasil)
//		default:
//			fmt.Println("input salah")
//		}
//	}
//
// ------------------------------------------------------------
// package main

// import "fmt"

// func main() {

// 	var a int
// 	var b int
// 	var c int
// 	var d int
// 	var e int
// 	var f int
// 	var g int
// 	var h int

// 	fmt.Println("masukkan nilai pertama")
// 	fmt.Scan(&a)

// 	fmt.Println("masukkan nilai kedua")
// 	fmt.Scan(&b)

// 	fmt.Println("masukkan nilai ketiga")
// 	fmt.Scan(&c)

// 	fmt.Println("masukkan nilai keempat")
// 	fmt.Scan(&d)

// 	fmt.Println("masukkan nilai kelima")
// 	fmt.Scan(&e)

// 	fmt.Println("masukkan nilai keenam")
// 	fmt.Scan(&f)

// 	fmt.Println("masukkan nilai ketujuh")
// 	fmt.Scan(&g)

// 	fmt.Println("masukkan nilai kedelapan")
// 	fmt.Scan(&h)

// 	switch {
// 	case (a > b) && (a > c) && (a > d) && (a > e) && (a > f) && (a > g) && (a > h):
// 		fmt.Print(a, " adalah angkah terbesar")
// 	case (b > c) && (b > d) && (b > e) && (b > f) && (b > g) && (b > h):
// 		fmt.Print(b, " adalah angkah terbesar")
// 	case (c > d) && (c > e) && (c > f) && (c > g) && (c > h):
// 		fmt.Print(c, " adalah angkah terbesar")
// 	case (d > e) && (d > f) && (d > g) && (d > h):
// 		fmt.Print(d, " adalah angkah terbesar")
// 	case (e > f) && (e > g) && (e > h):
// 		fmt.Print(e, " adalah angkah terbesar")
// 	case (f > g) && (f > h):
// 		fmt.Print(f, " adalah angkah terbesar")
// 	case (g > h):
// 		fmt.Print(g, " adalah angkah terbesar")
// 	default:
// 		fmt.Print(h, " adalah angkah terbesar")
// 	}
// }

// ------------------------------------------------------------
// package main

// import "fmt"

//	func main() {
//		var i int
//		fmt.Println("masukkan nilai awal")
//		fmt.Scan(&i)
//		for i < 5 {
//			fmt.Println("angkah", i)
//			i++
//		}
//	}
//
// ------------------------------------------------------------
// package main

// import "fmt"

// func main() {
// 	var (
// 		a int
// 		b int
// 		c int
// 	)
// 	for a = 1; a < 5; a++ {
// 		for b = 0; b < a; b++ {
// 			c = c + 1
// 			fmt.Print(c)
// 		}
// 		fmt.Println()
// 	}
// }
