package main

import "fmt"

func main() {
	Waktu()
	Voucher()
	Jumlah_bilangan()
	Konsekutif()
	fibonacci()
}

func Waktu() {
	var masukkan, menit, detik, jam int

	fmt.Print("Input Waktu:")
	fmt.Scan(&masukkan)

	detik = masukkan
	for detik >= 60 {
		detik = detik - 60
		menit += 1
	}

	for menit >= 60 {
		jam += 1
		menit -= 60
	}

	fmt.Println(jam, " Jam", menit, " Menit", detik, " Detik")
}

func Voucher() {
	var bil, d1, d2, d3, d4 int
	var diskon, cb bool

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bil)
	d1 = bil / 1000 % 10
	d2 = bil / 100 % 10
	d3 = bil / 10 % 10
	d4 = bil / 1 % 10
	fmt.Println(d1, d2, d3, d4)

	//diskon true jika d3 genap
	diskon = d3%2 == 0
	fmt.Println(diskon)

	//cb true jika d1+d3 habis dibagi d4
	cb = (d4 != 0) && (d1+d3)%d4 == 0
	fmt.Println(cb)
}

func Jumlah_bilangan() {
	var x, y, digit int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	for x != 0 {
		digit = x % 10
		y += digit
		x = x / 10
	}
	fmt.Println(y)
}

func Konsekutif() {
	var bil, d1, d2 int
	var konsekutif bool

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&bil)
	konsekutif = true
	for bil != 0 && konsekutif {
		d1 = bil / 1 % 10
		d2 = bil / 10 % 10
		konsekutif = d1-d2 == 1 || d2-d1 == 1
		bil = bil / 10
	}
	fmt.Println(konsekutif)
}

func fibonacci() {
	var N, f0, f1, i, f_selanjutnya int
	f0 = 0
	f1 = 1

	fmt.Scan(&N)
	fmt.Print(f0, " ")
	for i = 1; i <= N; i++ {
		fmt.Print(f1, " ")
		f_selanjutnya = f0 + f1
		f0 = f1
		f1 = f_selanjutnya
	}
}
