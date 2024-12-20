/*
spesifikasi 
a.	Pengguna bisa melakukan penambahan, pengubahan, dan penghapusan data barang yang dijual.
b.	Pengguna bisa mencatat setiap transaksi.
c.	Pengguna bisa menampilkan daftar transaksi dan omzet harian.
*/

package main

import (
	"fmt"
	"time"
)

//struktur untuk menyimpan informasi barang
type Barang struct {
	nama      string
	harga     int
	kuantitas int
}

//struktur untuk menyimpan informasi transaksi
type Transaksi struct {
	namaBarang string
	kuantitas  int
	hargaTotal int
	waktu      time.Time
}

const NMAX int = 100

var daftarBarang [NMAX]Barang
var daftarTransaksi []Transaksi
var jumlahBarang int

func main() {
	var pilihan int
	for {
		fmt.Println("=========== M E N U ===========")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Ubah Barang")
		fmt.Println("3. Hapus Barang")
		fmt.Println("4. Catat Transaksi")
		fmt.Println("5. Tampilkan Daftar Transaksi")
		fmt.Println("6. Tampilkan Omzet Harian")
		fmt.Println("7. Keluar")
		fmt.Println("===============================")
		fmt.Print("Pilih menu: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahBarang()
		case 2:
			ubahBarang()
		case 3:
			hapusBarang()
		case 4:
			catatTransaksi()
		case 5:
			tampilkanDaftarTransaksi()
		case 6:
			tampilkanOmzetHarian()
		case 7:
			fmt.Println("Terima kasih")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

//fungsi untuk menambahkan barang yang dibeli
func tambahBarang() {
	if jumlahBarang >= NMAX {
		fmt.Println("Kapasitas daftar barang penuh.")
		return
	}
	var nama string
	var harga, kuantitas int
	fmt.Print("Masukkan nama barang: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan harga barang: ")
	fmt.Scanln(&harga)
	fmt.Print("Masukkan kuantitas barang: ")
	fmt.Scanln(&kuantitas)

	daftarBarang[jumlahBarang] = Barang{nama, harga, kuantitas}
	jumlahBarang++
	fmt.Println("Barang berhasil ditambahkan.")
}

//fungsi untuk mengubah data barang yang telah dimasukkan berdasarkan nama barang yang dirutkan menggunakan selection sort
func urutkanBarang() {
	var t Barang
	i := 1
	for i <= jumlahBarang - 1 {
		idx := i - 1
		j := i
		for j < jumlahBarang{
			if daftarBarang[idx].nama > daftarBarang[j].nama{
				idx = j
			}
			j++
		}
		t = daftarBarang[idx]
		daftarBarang[idx] = daftarBarang[i-1]
		daftarBarang[i-1] = t
		i++
	}
}
//fungsi untuk ubah barang berdasarkan nama menggunakan binary search
func binarySearch(nama string) int {
	left := 0
	right := jumlahBarang-1
	for left <= right {
		mid := left + (right-left)/2
		if daftarBarang[mid].nama == nama {
			return mid
		}
		if daftarBarang[mid].nama < nama {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func ubahBarang() {
	var nm string
	fmt.Print("Masukkan nama barang yang akan diubah: ")
	fmt.Scanln(&nm)

	// memastikan daftar barang diurutkan sebelum melakukan binary search
	urutkanBarang()

	index := binarySearch(nm)
	if index == -1 {
		fmt.Println("Barang tidak ditemukan.")
		return
	}

	var nama string
	var harga, kuantitas int
	fmt.Print("Masukkan nama baru: ")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan harga baru: ")
	fmt.Scanln(&harga)
	fmt.Print("Masukkan kuantitas baru: ")
	fmt.Scanln(&kuantitas)

	daftarBarang[index].nama = nama
	daftarBarang[index].harga = harga
	daftarBarang[index].kuantitas = kuantitas
	fmt.Println("Barang berhasil diubah.")
}

//fungsi untuk menghapus data barang yang dimasukkan menggunakan sequential search
func hapusBarang() {
	var nm string
	fmt.Print("Masukkan nama barang yang akan dihapus: ")
	fmt.Scanln(&nm)

	for i := 0; i < jumlahBarang; i++ {
		if daftarBarang[i].nama == nm {
			daftarBarang[i] = daftarBarang[jumlahBarang-1]
			jumlahBarang--
			fmt.Println("Barang berhasil dihapus.")
			return
		}
	}
	fmt.Println("Barang tidak ditemukan.")
}

//fungsi untuk mencatat transaksi
func catatTransaksi() {
	var nm string
	var kuantitas int
	fmt.Print("Masukkan nama barang: ")
	fmt.Scanln(&nm)
	fmt.Print("Masukkan kuantitas: ")
	fmt.Scanln(&kuantitas)

	for i := 0; i < jumlahBarang; i++ {
		if daftarBarang[i].nama == nm {
			if daftarBarang[i].kuantitas < kuantitas {
				fmt.Println("Kuantitas barang tidak mencukupi.")
				return
			}

			daftarBarang[i].kuantitas -= kuantitas
			totalHarga := daftarBarang[i].harga * kuantitas
			daftarTransaksi = append(daftarTransaksi, Transaksi{
				namaBarang: nm,
				kuantitas:  kuantitas,
				hargaTotal: totalHarga,
				waktu:      time.Now(),
			})
			fmt.Println("Transaksi berhasil dicatat.")
			return
		}
	}
	fmt.Println("Barang tidak ditemukan.")
}

//fungsi untuk mencetak transaksi
func tampilkanDaftarTransaksi() {
	waktuTransaksi()
	fmt.Println("Daftar Transaksi:")
	fmt.Printf("%15s %10s %12s %20s\n", "Barang", "Kuantitas", "Harga Total", "Tanggal")
	var total_Pembelian int
	for i := 0; i < len(daftarTransaksi); i++ {
		fmt.Printf("%15s %10d %12d %20s\n", daftarTransaksi[i].namaBarang, daftarTransaksi[i].kuantitas, daftarTransaksi[i].hargaTotal, daftarTransaksi[i].waktu.Format("02-01-2006 15:04:05"))
		total_Pembelian += daftarTransaksi[i].hargaTotal
	}
	fmt.Printf("Total: %d\n", total_Pembelian)
}

//fungsi mencetak omzet harian
func tampilkanOmzetHarian() {
	totalOmzet := 0
	for i := 0; i < len(daftarTransaksi); i++ {
		totalOmzet += daftarTransaksi[i].hargaTotal
	}
	fmt.Printf("Total Omzet: %d\n", totalOmzet)
}

//waktu transaksi menggunakan insertion search
func waktuTransaksi() {
	for i := 0; i < len(daftarTransaksi); i++ {
		x := daftarTransaksi[i]
		j := i - 1
		for j > 0 && daftarTransaksi[j].waktu.After(x.waktu) {
			daftarTransaksi[j+1] = daftarTransaksi[j]
			j--
		}
		daftarTransaksi[j+1] = x
	}
}
