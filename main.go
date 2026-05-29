package main

import "fmt"

func tampilMenu() {
	fmt.Println()
	fmt.Println("=====================================================")
	fmt.Println("         APLIKASI MANAJEMEN GYM MEMBERSHIP")
	fmt.Println("=====================================================")
	fmt.Println("1. Tambah Data Member")
	fmt.Println("2. Ubah Data Member")
	fmt.Println("3. Hapus Data Member")
	fmt.Println("4. Catat Kunjungan Pengunjung")
	fmt.Println("5. Cari Data Member")
	fmt.Println("6. Tampilkan Data Member")
	fmt.Println("7. Urutkan Data Member")
	fmt.Println("8. Tampilkan Top 3 Member Paling Aktif")
	fmt.Println("9. Keluar")
	fmt.Println("=====================================================")
}

func main() {
	var pilihan int

	for pilihan != 9 {

		tampilMenu()

		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			fmt.Println()
			fmt.Println("===== TAMBAH DATA MEMBER =====")

		case 2:
			fmt.Println()
			fmt.Println("===== UBAH DATA MEMBER =====")

		case 3:
			fmt.Println()
			fmt.Println("===== HAPUS DATA MEMBER =====")

		case 4:
			fmt.Println()
			fmt.Println("===== CATAT KUNJUNGAN PENGUNJUNG =====")

		case 5:
			fmt.Println()
			fmt.Println("===== CARI DATA MEMBER =====")

		case 6:
			fmt.Println()
			fmt.Println("===== TAMPILKAN DATA MEMBER =====")

		case 7:
			fmt.Println()
			fmt.Println("===== URUTKAN DATA MEMBER =====")

		case 8:
			fmt.Println()
			fmt.Println("===== TOP 3 MEMBER PALING AKTIF =====")

		case 9:
			fmt.Println()
			fmt.Println("Terima kasih telah menggunakan program")

		default:
			fmt.Println()
			fmt.Println("Menu tidak tersedia")
		}
	}
}
