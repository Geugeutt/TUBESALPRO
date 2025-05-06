package main

import (
	"fmt"
	"sort"
	"investasiapp/investasi"
)

func main() {
	for {
		fmt.Println("<== Menu Aplikasi Investasi ==>")
		fmt.Println("1. Tambah Investasi")
		fmt.Println("2. Ubah Investasi")
		fmt.Println("3. Hapus Investasi")
		fmt.Println("4. Tampilkan Semua Investasi")
		fmt.Println("5. Cari Investasi (Sequential)")
		fmt.Println("6. Cari Investasi (Binary)")
		fmt.Println("7. Urutkan berdasarkan Dana (Selection Sort)")
		fmt.Println("8. Urutkan berdasarkan Keuntungan (Insertion Sort)")
		fmt.Println("9. Keluar")

		pilih := investasi.BacaInput("Pilih menu: ")

		switch pilih {
		case "1":
			inv := inputDataInvestasi()
			investasi.TambahInvestasi(inv)
		case "2":
			index := investasi.ParseFloat(investasi.BacaInput("Masukkan index yang ingin diubah: "))
			inv := inputDataInvestasi()
			investasi.UbahInvestasi(int(index), inv)
		case "3":
			index := investasi.ParseFloat(investasi.BacaInput("Masukkan index yang ingin dihapus: "))
			investasi.HapusInvestasi(int(index))
		case "4":
			investasi.TampilkanInvestasi()
		case "5":
			keyword := investasi.BacaInput("Masukkan keyword pencarian: ")
			hasil := investasi.SequentialSearch(keyword)
			fmt.Println("\nHasil pencarian:")
			for _, inv := range hasil {
				fmt.Printf("%s - %s | Dana: %.2f\n", inv.Nama, inv.Jenis, inv.Dana)
			}
		case "6":
			sort.Slice(investasi.DataInvestasi, func(i, j int) bool {
				return investasi.DataInvestasi[i].Nama < investasi.DataInvestasi[j].Nama
			})
			keyword := investasi.BacaInput("Masukkan nama investasi: ")
			hasil := investasi.BinarySearch(investasi.DataInvestasi, keyword)
			if hasil != nil {
				fmt.Printf("Ditemukan: %s - %s | Dana: %.2f\n", hasil.Nama, hasil.Jenis, hasil.Dana)
			} else {
				fmt.Println("Data tidak ditemukan.")
			}
		case "7":
			hurut := investasi.SelectionSortByDana(investasi.DataInvestasi)
			fmt.Println("\nData setelah diurutkan berdasarkan Dana:")
			for _, inv := range hurut {
				fmt.Printf("%s - Dana: %.2f\n", inv.Nama, inv.Dana)
			}
		case "8":
			hurut := investasi.InsertionSortByKeuntungan(investasi.DataInvestasi)
			fmt.Println("\nData setelah diurutkan berdasarkan Keuntungan:")
			for _, inv := range hurut {
				fmt.Printf("%s - Keuntungan: %.2f%%\n", inv.Nama, (inv.HargaAkhir - inv.HargaAwal)/inv.HargaAwal*100)
			}
		case "9":
			fmt.Println("Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func inputDataInvestasi() investasi.Investasi {
	nama := investasi.BacaInput("Nama investasi: ")
	jenis := investasi.BacaInput("Jenis (saham/obligasi/reksadana): ")
	dana := investasi.ParseFloat(investasi.BacaInput("Jumlah dana: "))
	hargaAwal := investasi.ParseFloat(investasi.BacaInput("Harga awal aset: "))
	hargaAkhir := investasi.ParseFloat(investasi.BacaInput("Harga akhir aset: "))

	return investasi.Investasi{
		Nama:       nama,
		Jenis:      jenis,
		Dana:       dana,
		HargaAwal:  hargaAwal,
		HargaAkhir: hargaAkhir,
	}
}