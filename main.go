package main

import "fmt"

func menu() {
	fmt.Println("=================================================")
	fmt.Println("        APLIKASI MANAJEMEN GYM MEMBERSHIP")
	fmt.Println("=================================================")
	fmt.Println("1. Input Data Member")
	fmt.Println("2. Input Data Pendatang")
	fmt.Println("3. Tampilkan Data Pengunjung")
	fmt.Println("4. Tampilkan Pengunjung Paling Aktif")
	fmt.Println("5. Keluar")
	fmt.Println("=================================================")
}

func main() {

	var pilihan int

	for pilihan != 5 {

		menu()

		fmt.Print("Pilih Menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			fmt.Println()
			fmt.Println("===== INPUT DATA MEMBER =====")
			fmt.Println()

		case 2:
			fmt.Println()
			fmt.Println("===== INPUT DATA PENDATANG =====")
			fmt.Println()

		case 3:
			fmt.Println()
			fmt.Println("===== TAMPILKAN DATA PENGUNJUNG =====")
			fmt.Println()

		case 4:
			fmt.Println()
			fmt.Println("===== PENGUNJUNG PALING AKTIF =====")
			fmt.Println()

		case 5:
			fmt.Println()
			fmt.Println("Terima kasih telah menggunakan program")
			fmt.Println()

		default:
			fmt.Println()
			fmt.Println("Menu tidak tersedia")
			fmt.Println()
		}
	}
}
