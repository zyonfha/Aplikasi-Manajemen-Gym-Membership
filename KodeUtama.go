package main

import "fmt"

const NMAX = 1000

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

// Inisialisasi data dummy
func initDummyData(A *tabMember, n *int, nextID *int) {
	A[0] = member{id: "001", nama: "Budi", membership: "Basic", durasiMember: 6, kunjungan: 12, totalLatihan: 720, tanggalTerakhir: 1, bulanTerakhir: 6, tahunTerakhir: 2026}
	A[1] = member{id: "002", nama: "Siti", membership: "Premium", durasiMember: 12, kunjungan: 45, totalLatihan: 2700, tanggalTerakhir: 10, bulanTerakhir: 6, tahunTerakhir: 2026}
	A[2] = member{id: "003", nama: "Andi", membership: "Basic", durasiMember: 1, kunjungan: 2, totalLatihan: 120, tanggalTerakhir: 5, bulanTerakhir: 6, tahunTerakhir: 2026}
	A[3] = member{id: "004", nama: "Joko", membership: "Premium", durasiMember: 3, kunjungan: 0, totalLatihan: 0, tanggalTerakhir: 0, bulanTerakhir: 0, tahunTerakhir: 0}
	A[4] = member{id: "005", nama: "Rina", membership: "Basic", durasiMember: 24, kunjungan: 30, totalLatihan: 1800, tanggalTerakhir: 12, bulanTerakhir: 6, tahunTerakhir: 2026}

	*n = 5
	*nextID = 6 // Karena sudah ada 5 data
}

func tampilMenu() {
	fmt.Println()
	fmt.Println("=====================================================")
	fmt.Println("         APLIKASI MANAJEMEN GYM MEMBERSHIP")
	fmt.Println("=====================================================")
	fmt.Println("1. Tambah Data Member")
	fmt.Println("2. Ubah Data Member (Binary Search)")
	fmt.Println("3. Hapus Data Member (Binary Search)")
	fmt.Println("4. Catat Kunjungan Member")
	fmt.Println("5. Cari Data Member (Sequential Search by Nama)")
	fmt.Println("6. Urutkan Data Member by Kunjungan (Selection Sort)")
	fmt.Println("7. Urutkan Data Member by Durasi (Insertion Sort)")
	fmt.Println("8. Tampilkan Top 3 Member Terajin & Termalas")
	fmt.Println("9. Catat Kunjungan Harian (Non-Member)")
	fmt.Println("10. Laporan Pendapatan Gym")
	fmt.Println("11. Keluar")
	fmt.Println("=====================================================")
}

func main() {
	var pilihan int
	var data tabMember
	var n int
	var totalNonMember int
	var nextID int

	nextID = 1

	initDummyData(&data, &n, &nextID)

	for pilihan != 11 {
		tampilMenu()
		fmt.Print("Pilih menu : ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahMember(&data, &n, &nextID)
		case 2:
			ubahMember(&data, n)
		case 3:
			hapusMember(&data, &n)
		case 4:
			catatKunjungan(&data, n)
		case 5:
			cariMember(data, n)
		case 6:
			urutkanMember(data, n)
		case 7:
			urutkanMemberDurasi(data, n)
		case 8:
			topTigaMember(data, n)
		case 9:
			fmt.Println()
			fmt.Println("===== CATAT KUNJUNGAN NON-MEMBER =====")
			totalNonMember++
			fmt.Println("Berhasil mencatat 1 pengunjung harian (Rp 50.000)")
		case 10:
			laporanPendapatan(data, n, totalNonMember)
		case 11:
			fmt.Println()
			fmt.Println("Terima kasih telah menggunakan program")
		default:
			fmt.Println()
			fmt.Println("Menu tidak tersedia")
		}
	}
}

// Algoritma Binary Search berdasarkan ID
func binarySearchID(A tabMember, n int, id string) int {
	var kiri, kanan, tengah int
	kiri = 0
	kanan = n - 1

	for kiri <= kanan {
		tengah = (kiri + kanan) / 2
		if A[tengah].id == id {
			return tengah
		} else if A[tengah].id < id {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}
	return -1
}

func tambahMember(A *tabMember, n *int, nextID *int) {
	var pilih int
	fmt.Println()
	fmt.Println("===== TAMBAH DATA MEMBER =====")

	A[*n].id = fmt.Sprintf("%03d", *nextID)
	*nextID++
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
	var pilih int
	var idx int

	fmt.Println()
	fmt.Println("===== UBAH DATA MEMBER =====")
	fmt.Print("Masukkan ID Member : ")
	fmt.Scan(&id)

	idx = binarySearchID(*A, n, id)

	if idx != -1 {
		fmt.Print("Nama Baru : ")
		fmt.Scan(&A[idx].nama)

		fmt.Println("Jenis Membership")
		fmt.Println("1. Basic")
		fmt.Println("2. Premium")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)

		if pilih == 1 {
			A[idx].membership = "Basic"
		} else {
			A[idx].membership = "Premium"
		}

		fmt.Print("Durasi Baru (bulan) : ")
		fmt.Scan(&A[idx].durasiMember)

		fmt.Println("Data berhasil diubah")
	} else {
		fmt.Println("Member tidak ditemukan")
	}
}

func hapusMember(A *tabMember, n *int) {
	var id string
	var idx int
	var i int

	fmt.Println()
	fmt.Println("===== HAPUS DATA MEMBER =====")
	fmt.Print("Masukkan ID Member : ")
	fmt.Scan(&id)

	idx = binarySearchID(*A, *n, id)

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
	var idx int

	fmt.Println()
	fmt.Println("===== CATAT KUNJUNGAN =====")
	fmt.Print("ID Member : ")
	fmt.Scan(&id)

	idx = binarySearchID(*A, n, id)

	if idx != -1 {
		fmt.Print("Tanggal (dd mm yyyy) : ")
		fmt.Scan(&tanggal, &bulan, &tahun)

		if A[idx].tahunTerakhir != 0 {
			if tahun < A[idx].tahunTerakhir {
				fmt.Println("Tanggal tidak valid: Tahun lebih kecil")
				return
			} else if tahun == A[idx].tahunTerakhir && bulan < A[idx].bulanTerakhir {
				fmt.Println("Tanggal tidak valid: Bulan lebih kecil")
				return
			} else if tahun == A[idx].tahunTerakhir && bulan == A[idx].bulanTerakhir && tanggal < A[idx].tanggalTerakhir {
				fmt.Println("Tanggal tidak valid: Tanggal lebih kecil")
				return
			}
		}

		fmt.Print("Durasi Latihan (menit) : ")
		fmt.Scan(&durasi)

		A[idx].kunjungan++
		A[idx].totalLatihan += durasi
		A[idx].tanggalTerakhir = tanggal
		A[idx].bulanTerakhir = bulan
		A[idx].tahunTerakhir = tahun

		fmt.Println("Kunjungan berhasil dicatat")
	} else {
		fmt.Println("Member tidak ditemukan")
	}
}

func cariMember(A tabMember, n int) {
	var nama string
	var ketemu bool
	var i int

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
			fmt.Println("Durasi     :", A[i].durasiMember, "Bulan")
			fmt.Println("Kunjungan  :", A[i].kunjungan, "Kali")
		}
	}

	if !ketemu {
		fmt.Println("Data tidak ditemukan")
	}
}

func tampilMember(A tabMember, n int) {
	var i int
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

func urutkanMember(A tabMember, n int) {
	var i, j, maxIdx int
	var temp member
	var pilihan int
	var isAscending bool

	fmt.Println()
	fmt.Println("===== URUTKAN DATA MEMBER (SELECTION SORT - KUNJUNGAN) =====")

	if n == 0 {
		fmt.Println("Belum ada data member.")
		return
	}
	fmt.Println("1. Ascending (Terkecil ke Terbesar)")
	fmt.Println("2. Descending (Terbesar ke Terkecil)")
	fmt.Print("Pilih Urutan: ")
	fmt.Scan(&pilihan)

	isAscending = false
	if pilihan == 1 {
		isAscending = true
	}

	for i = 0; i < n-1; i++ {
		maxIdx = i
		for j = i + 1; j < n; j++ {
			if isAscending {
				if A[j].kunjungan < A[maxIdx].kunjungan {
					maxIdx = j
				}
			} else {
				if A[j].kunjungan > A[maxIdx].kunjungan {
					maxIdx = j
				}
			}

		}
		temp = A[i]
		A[i] = A[maxIdx]
		A[maxIdx] = temp
	}

	fmt.Println("Data berhasil diurutkan berdasarkan kunjungan terbanyak!")
	tampilMember(A, n)
}

func urutkanMemberDurasi(A tabMember, n int) {
	var pilihan int
	var isAscending bool
	var i, j int
	var key member

	fmt.Println()
	fmt.Println("===== URUTKAN DATA MEMBER (INSERTION SORT - DURASI) =====")
	if n == 0 {
		fmt.Println("Belum ada data member.")
		return
	}

	fmt.Println("1. Ascending (Terkecil ke Terbesar)")
	fmt.Println("2. Descending (Terbesar ke Terkecil)")
	fmt.Print("Pilih Urutan: ")
	fmt.Scan(&pilihan)

	isAscending = false
	if pilihan == 1 {
		isAscending = true
	}

	for i = 1; i < n; i++ {
		key = A[i]
		j = i - 1

		if isAscending {
			for j >= 0 && A[j].durasiMember > key.durasiMember {
				A[j+1] = A[j]
				j--
			}
		} else {
			for j >= 0 && A[j].durasiMember < key.durasiMember {
				A[j+1] = A[j]
				j--
			}
		}
		A[j+1] = key
	}

	fmt.Println("Data berhasil diurutkan berdasarkan durasi!")
	tampilMember(A, n)
}

func topTigaMember(A tabMember, n int) {
	var i, j, batas int
	var key member
	var terajin tabMember

	fmt.Println()

	if n == 0 {
		fmt.Println("Belum ada data member.")
		return
	}

	for i = 1; i < n; i++ {
		key = A[i]
		j = i - 1
		for j >= 0 && A[j].kunjungan > key.kunjungan {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}

	batas = 3
	if n < 3 {
		batas = n
	}

	fmt.Println("===== TOP 3 MEMBER TERMALAS =====")
	tampilMember(A, batas)

	for i = 0; i < n; i++ {
		terajin[i] = A[n-1-i]
	}

	fmt.Println("\n===== TOP 3 MEMBER TERAJIN =====")
	tampilMember(terajin, batas)
}

func laporanPendapatan(A tabMember, n int, kunjunganNonMember int) {
	var totalBasic, totalPremium int
	var pendapatanNonMember, pendapatanMember, totalKeseluruhan int
	var i int

	fmt.Println()
	fmt.Println("=====================================================")
	fmt.Println("               LAPORAN PENDAPATAN GYM")
	fmt.Println("=====================================================")

	for i = 0; i < n; i++ {
		if A[i].membership == "Basic" {
			totalBasic += A[i].durasiMember * 250000
		} else if A[i].membership == "Premium" {
			totalPremium += A[i].durasiMember * 350000
		}
	}

	pendapatanNonMember = kunjunganNonMember * 50000
	pendapatanMember = totalBasic + totalPremium
	totalKeseluruhan = pendapatanNonMember + pendapatanMember

	fmt.Printf("1. Kunjungan Harian (%d kali) : Rp %d\n", kunjunganNonMember, pendapatanNonMember)
	fmt.Printf("2. Total Pendaftaran Basic   : Rp %d\n", totalBasic)
	fmt.Printf("3. Total Pendaftaran Premium : Rp %d\n", totalPremium)
	fmt.Println("-----------------------------------------------------")
	fmt.Printf("TOTAL PENDAPATAN KESELURUHAN : Rp %d\n", totalKeseluruhan)
	fmt.Println("=====================================================")
}
