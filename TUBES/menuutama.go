package main

import "fmt"

const NMAX = 10

//nanti tambahin for biar kalau udah selesai ga langsung keluar aplikasi
func menuUtama(datauser user, i, pilihan int, kunci *int) {
	var (
		isKeluar bool
		tipereksadana int
	)

	for !isKeluar {
		fmt.Println("\n======= Menu Utama =======")
		fmt.Println("1. Katalog")
		fmt.Println("2. Portofolio") //bisa nambahin aset (beli) dan kurangin aset (jual)
		fmt.Println("3. Profil")
		fmt.Println("4. Kembali")
		fmt.Println("0. Keluar")
		fmt.Println()

		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		switch pilihan {
		case 1:
			katalog(&dbReksadana, &dbSaham, &dbObligasi, pilihan, kunci, &tipereksadana)
		case 2:
			tampilPortofolio(kunci, &tipereksadana)
		case 3:
			profil(datauser, &dbUser, kunci, pilihan)
		case 4:
			fiturLoginDaftar(datauser, i, pilihan, kunci)
			isKeluar = true
		case 0:
			fmt.Println("Terima kasih sudah menggunakan aplikasi ini")
			isKeluar = true
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func keluar() {
	fmt.Println("Terima kasih sudah menggunakan aplikasi ini")
}

//i merupakan jumlah user saat ini
//kunci dipakai buat jadi penghubung antara portofolio dengan user
//pilihan digunakan buat jadi variabel penampung aksi pilihan user
//datauser digunain buat nyimpen data pas orang daftar ataupun login (nampung data dari input sebelum dimasukkin ke arraynya)
func main() {
	seedPorto()
	seedTransaksi()
	seedreksa()
	seedobligasi()
	seedsaham()
	isiUser()

	var datauser user
	var i, pilihan, kunci int

	i = 2

	fiturLoginDaftar(datauser, i, pilihan, &kunci)
}
