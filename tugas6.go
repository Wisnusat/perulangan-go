package main

import (
	"fmt"
	"strconv"
)

func genap() {
	var x, count int

	fmt.Println("Bilangan Genap")
	fmt.Scanln(&x)

	for x%2 == 0 {
		count++
		fmt.Scanln(&x)
	}
	fmt.Println("Count:", count)
}

func login() {
	fmt.Println("Login")

	var username_input, password_input, username_temp, password_temp string
	var loop int
	var isSuccess bool

	username_temp = "admin"
	password_temp = "admin"
	loop = 0
	isSuccess = false

	fmt.Scan(&username_input, &password_input)
	for !isSuccess {
		fmt.Scan(&username_input, &password_input)
		isSuccess = username_input != username_temp || password_input != password_temp
		loop++
	}

	fmt.Println(loop)
	fmt.Println("Login berhasil")
}

func digit() {
	fmt.Println("Digit")

	var i, total int
	var input_bilangan string
	total = 0

	fmt.Scan(&input_bilangan)

	// descending
	for i = len(input_bilangan) - 1; i >= 0; i-- {
		fmt.Print(string(input_bilangan[i]), " ")
		loop, _ := strconv.Atoi(string(input_bilangan[i]))
		total += loop
	}

	fmt.Print("\n", total)
}

func dompet() {
	fmt.Println("Dompet")

	var saldo, saldo_total int
	var isZero bool

	isZero = false
	saldo_total = 0

	for !isZero {
		fmt.Scan(&saldo)
		saldo_total += saldo
		isZero = saldo == 0
	}

	fmt.Println(saldo_total)
}

func konsekutif() {
	var input_bilangan string
	var temp_number, temp_number_before, i, before_index int
	var hasil bool

	fmt.Scanln(&input_bilangan)

	for i = 1; i <= len(input_bilangan)-1; i++ {
		before_index = i - 1
		temp_number, _ = strconv.Atoi(string(input_bilangan[i]))
		temp_number_before, _ = strconv.Atoi(string(input_bilangan[before_index]))
		hasil = temp_number-temp_number_before == 1 || temp_number-temp_number_before == -1
	}

	fmt.Println(hasil)
}

func tangki_air() {
	var tanki, ember, hasil int
	var condition bool

	fmt.Scanln(&tanki)

	hasil = 0

	for hasil <= tanki {
		fmt.Scan(&ember)
		hasil += ember
		condition = hasil >= tanki
		fmt.Println(condition)
	}
}

func cangkir_kopi() {

	var s_gula, s_kopi, gula, kopi, cangkir int

	fmt.Scan(&s_gula, &s_kopi, &gula, &kopi)

	for gula <= s_gula && kopi <= s_kopi {
		cangkir += 1
		s_gula -= gula
		s_kopi -= kopi
	}

	fmt.Println(cangkir)
}

func bil_prima() {
	var prima bool
	var i, bil int

	prima = true
	i = 2
	fmt.Scanln(&bil)
	if bil <= 1 {
		prima = false
	}
	for i <= bil-1 && prima {
		prima = bil%i != 0
		i++
	}
	fmt.Println(prima)
}

func main() {
	genap()
	login()
	digit()
	dompet()
	cangkir_kopi()
	konsekutif()
	tangki_air()
	bil_prima()
}
