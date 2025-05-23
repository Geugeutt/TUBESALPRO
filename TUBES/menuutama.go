package main

import "fmt"

const NMAX = 10

func menuUtama(datauser user, pilihan int, kunci *int, isNew *bool) {
	var (
		isKeluar bool
		tipereksadana, asetDijual, jumlahDijual int
	)

	for !isKeluar {
		dashboard(kunci, *isNew)
		
		clearScreen()
		
		fmt.Println("\n========== MENU UTAMA ==========")
		fmt.Println("[1] Katalog")
		fmt.Println("[2] Portofolio")
		fmt.Println("[3] Profil")
		fmt.Println("[4] Kembali")
		fmt.Println("[0] Keluar")
		fmt.Println()

		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		switch pilihan {
		case 1:
			clearScreen()
			katalog(pilihan, kunci, &tipereksadana)
		case 2:
			clearScreen()
			tampilPortofolio(kunci, &tipereksadana, &asetDijual, &jumlahDijual)
		case 3:
			clearScreen()
			profil(datauser, kunci, &asetDijual, &jumlahDijual, pilihan)
		case 4:
			isKeluar = true
		case 0:
			clearScreen()
			fmt.Println("Terima kasih sudah menggunakan aplikasi ini")
			isKeluar = true
		default:
			clearScreen()
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

	var (
		datauser user
		i, pilihan, kunci int
		isNew bool
	)

	i = 2

	clearScreen()
	fiturLoginDaftar(datauser, i, pilihan, &kunci, &isNew)
}
