package main

import "fmt"

const NMAX = 1000

var nextID int = 1

type member struct {
	id              string
	nama            string
	membership      string
	durasiMember    int
	kunjungan       int
	totalLatihan    int
	tanggalTerakhir int
	bulanTerakhir   int
	tahunTerakhir   int
}

type tabMember [NMAX]member

type kunjungan struct {
	idMember string
	tanggal  string
	latihan  int
}

type tabKunjungan [NMAX]kunjungan

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

	var data tabMember
	var n int

	for pilihan != 9 {

		tampilMenu()

		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {

		case 1:
			tambahMember(&data, &n)

		case 2:
			ubahMember(&data, n)

		case 3:
			hapusMember(&data, &n)

		case 4:
			catatKunjungan(&data, n)

		case 5:
			cariMember(data, n)

		case 6:
			tampilMember(data, n)

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

func tambahMember(A *tabMember, n *int) {

	var pilih int

	fmt.Println()
	fmt.Println("===== TAMBAH DATA MEMBER =====")

	A[*n].id = fmt.Sprintf("%03d", nextID)
	nextID++

	fmt.Println("ID Member :", A[*n].id)

	fmt.Print("Nama Member : ")
	fmt.Scan(&A[*n].nama)

	fmt.Println("Jenis Membership")
	fmt.Println("1. Basic")
	fmt.Println("2. Premium")
	fmt.Print("Pilih : ")
	fmt.Scan(&pilih)

	if pilih == 1 {
		A[*n].membership = "Basic"
	} else {
		A[*n].membership = "Premium"
	}

	fmt.Print("Durasi Membership (bulan) : ")
	fmt.Scan(&A[*n].durasiMember)

	A[*n].kunjungan = 0
	A[*n].totalLatihan = 0

	*n = *n + 1

	fmt.Println("Data berhasil ditambahkan")
}

func ubahMember(A *tabMember, n int) {

	var id string
	var i int
	var ketemu bool
	var pilih int

	fmt.Println()
	fmt.Println("===== UBAH DATA MEMBER =====")

	fmt.Print("Masukkan ID Member : ")
	fmt.Scan(&id)

	i = 0
	for i < n && !ketemu {

		if A[i].id == id {

			ketemu = true

			fmt.Print("Nama Baru : ")
			fmt.Scan(&A[i].nama)

			fmt.Println("Jenis Membership")
			fmt.Println("1. Basic")
			fmt.Println("2. Premium")
			fmt.Print("Pilih : ")
			fmt.Scan(&pilih)

			if pilih == 1 {
				A[i].membership = "Basic"
			} else {
				A[i].membership = "Premium"
			}

			fmt.Print("Durasi Baru : ")
			fmt.Scan(&A[i].durasiMember)

			fmt.Println("Data berhasil diubah")
		}

		i++
	}

	if !ketemu {
		fmt.Println("Member tidak ditemukan")
	}
}

func hapusMember(A *tabMember, n *int) {

	var id string
	var idx int = -1
	var i int

	fmt.Println()
	fmt.Println("===== HAPUS DATA MEMBER =====")

	fmt.Print("Masukkan ID Member : ")
	fmt.Scan(&id)

	for i = 0; i < *n; i++ {

		if A[i].id == id {
			idx = i
		}
	}

	if idx == -1 {

		fmt.Println("Member tidak ditemukan")

	} else {

		for i = idx; i < *n-1; i++ {
			A[i] = A[i+1]
		}

		*n = *n - 1

		fmt.Println("Data berhasil dihapus")
	}
}

func catatKunjungan(A *tabMember, n int) {

	var id string
	var durasi int
	var tanggal, bulan, tahun int

	var i int
	var ketemu bool

	fmt.Println()
	fmt.Println("===== CATAT KUNJUNGAN =====")

	fmt.Print("ID Member : ")
	fmt.Scan(&id)

	for i = 0; i < n && !ketemu; i++ {

		if A[i].id == id {

			ketemu = true

			fmt.Print("Tanggal (dd mm yyyy) : ")
			fmt.Scan(&tanggal, &bulan, &tahun)

			if A[i].tahunTerakhir != 0 {

				if tahun < A[i].tahunTerakhir {

					fmt.Println("Tanggal tidak valid")
					fmt.Println("Tahun tidak boleh lebih kecil dari kunjungan terakhir")
					return

				} else if tahun == A[i].tahunTerakhir &&
					bulan < A[i].bulanTerakhir {

					fmt.Println("Tanggal tidak valid")
					fmt.Println("Bulan tidak boleh lebih kecil dari kunjungan terakhir")
					return

				} else if tahun == A[i].tahunTerakhir &&
					bulan == A[i].bulanTerakhir &&
					tanggal < A[i].tanggalTerakhir {

					fmt.Println("Tanggal tidak valid")
					fmt.Println("Tanggal tidak boleh lebih kecil dari kunjungan terakhir")
					return
				}
			}

			fmt.Print("Durasi Latihan (menit) : ")
			fmt.Scan(&durasi)

			A[i].kunjungan++
			A[i].totalLatihan += durasi

			A[i].tanggalTerakhir = tanggal
			A[i].bulanTerakhir = bulan
			A[i].tahunTerakhir = tahun

			fmt.Println("Kunjungan berhasil dicatat")
		}
	}

	if !ketemu {
		fmt.Println("Member tidak ditemukan")
	}
}

func cariMember(A tabMember, n int) {

	var nama string
	var i int
	var ketemu bool

	fmt.Println()
	fmt.Println("===== CARI DATA MEMBER =====")

	fmt.Print("Masukkan Nama : ")
	fmt.Scan(&nama)

	for i = 0; i < n; i++ {

		if A[i].nama == nama {

			ketemu = true

			fmt.Println()
			fmt.Println("Data Ditemukan")
			fmt.Println("ID         :", A[i].id)
			fmt.Println("Nama       :", A[i].nama)
			fmt.Println("Membership :", A[i].membership)
			fmt.Println("Durasi     :", A[i].durasiMember)
			fmt.Println("Kunjungan  :", A[i].kunjungan)
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func tampilMember(A tabMember, n int) {

	var i int

	fmt.Println()
	fmt.Println("=========================================================================================================")
	fmt.Printf("%-6s %-15s %-12s %-8s %-10s %-15s %-15s\n",
		"ID", "Nama", "Member", "Durasi", "Kunjungan", "Total Latihan", "Tgl Terakhir")
	fmt.Println("=========================================================================================================")

	for i = 0; i < n; i++ {

		fmt.Printf("%-6s %-15s %-12s %-8d %-10d %-15s %02d/%02d/%04d\n",
			A[i].id,
			A[i].nama,
			A[i].membership,
			A[i].durasiMember,
			A[i].kunjungan,
			fmt.Sprintf("%d Menit", A[i].totalLatihan),
			A[i].tanggalTerakhir,
			A[i].bulanTerakhir,
			A[i].tahunTerakhir)
	}
}
