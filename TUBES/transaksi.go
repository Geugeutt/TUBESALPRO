package main

import (
	"fmt"
	"time"
)

func tampilTransaksiPembelian(kunci, count *int) int{
	var (
		jenisTransaksi int
	)

	sumber := "TRANSAKSI"
	uiHeaderTable(sumber, 1, 0)
	uiHeaderTransaksi()
	for i:=0; i<NMAX; i++{
		if dbTransaksi[*kunci].riwayat[i].nilaiAset > 0 {
			*count++
			fmt.Printf("  %-7d %-40s %-15d %-15d %-20s %-28s %-15s\n", *count, dbTransaksi[*kunci].riwayat[i].produk, dbTransaksi[*kunci].riwayat[i].jumlah, dbTransaksi[*kunci].riwayat[i].nilaiAset, dbTransaksi[*kunci].riwayat[i].portofolio, dbTransaksi[*kunci].riwayat[i].tipe, dbTransaksi[*kunci].riwayat[i].tanggal)
		}
	}
	uiFooterTableSuperPanjang()

	switch *count {
	case 0:
		fmt.Println("Belum ada riwayat transaksi pembelian aset.")
	}

	jenisTransaksi = 1
	return jenisTransaksi
}

func tampilTransaksiPenjualan(kunci *int, jumlahDijual, asetDijual int) int{
	var (
		count, jenisTransaksi int
	)

	sumber := "TRANSAKSI"
	uiHeaderTable(sumber, 2, 0)
	uiHeaderTransaksi()
	for i:=0; i<NMAX; i++{
		if dbTransaksi[*kunci].riwayat[i].nilaiAset > 0 && dbTransaksi[*kunci].riwayat[i].tanda == 2 {
			count++
			fmt.Printf("  %-7d %-40s %-15d %-15d %-20s %-28s %-15s\n", count, dbTransaksi[*kunci].riwayat[i].produk, jumlahDijual , asetDijual, dbTransaksi[*kunci].riwayat[i].portofolio, dbTransaksi[*kunci].riwayat[i].tipe, dbTransaksi[*kunci].riwayat[i].tanggal)
		}
	}
	uiFooterTableSuperPanjang()

	switch count {
	case 0:
		fmt.Println("Belum ada riwayat transaksi penjualan aset.")
	}

	jenisTransaksi = 2
	return jenisTransaksi
}

func filterTransaksi(kunci *int, pilihan, jenisTransaksi int){
	var (
		stat bool
		count int
	)

	if jenisTransaksi == 1{
		fmt.Println()
		fmt.Println("Filter: ")
		fmt.Println("[1] Riwayat Transaksi Pembelian Saham")
		fmt.Println("[2] Riwayat Transaksi Pembelian Reksadana")
		fmt.Println("[3] Riwayat Transaksi Pembelian Obligasi")
		fmt.Println("[0] Kembali")
	}else if jenisTransaksi == 2{
		fmt.Println()
		fmt.Println("Filter: ")
		fmt.Println("[1] Riwayat Transaksi Penjualan Saham")
		fmt.Println("[2] Riwayat Transaksi Penjualan Reksadana")
		fmt.Println("[3] Riwayat Transaksi Penjualan Obligasi")
		fmt.Println("[0] Kembali")
	}

	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	for !stat{
		switch pilihan {
		case 1:
			clearScreen()
			sumber := "TRANSAKSI"
			uiHeaderTable(sumber, 1, 0)
			uiHeaderTransaksi()
			for i:=0; i<NMAX; i++{
				if dbTransaksi[*kunci].riwayat[i].tipe == "Saham" && dbTransaksi[*kunci].riwayat[i].nilaiAset > 0{
					count++
					fmt.Printf("  %-7d %-40s %-15d %-15d %-20s %-28s %-15s\n", count, dbTransaksi[*kunci].riwayat[i].produk, dbTransaksi[*kunci].riwayat[i].jumlah, dbTransaksi[*kunci].riwayat[i].nilaiAset, dbTransaksi[*kunci].riwayat[i].portofolio, dbTransaksi[*kunci].riwayat[i].tipe, dbTransaksi[*kunci].riwayat[i].tanggal)
				}
			}
			uiFooterTableSuperPanjang()

			fmt.Println()
			fmt.Println("[0] Kembali")
			fmt.Print("Pilihan: ")
			fmt.Scan(&pilihan)

			if pilihan == 0{
				clearScreen()
				stat = true
			}
		case 2:
			clearScreen()
			sumber := "TRANSAKSI"
			uiHeaderTable(sumber, 1, 0)
			uiHeaderTransaksi()
			for i:=0; i<NMAX; i++{
				if (dbTransaksi[*kunci].riwayat[i].tipe == "Reksadana Pasar Uang" || dbTransaksi[*kunci].riwayat[i].tipe == "Reksadana Saham" || dbTransaksi[*kunci].riwayat[i].tipe == "Reksadana Obligasi") && dbTransaksi[*kunci].riwayat[i].tanda == jenisTransaksi && dbTransaksi[*kunci].riwayat[i].nilaiAset != 0{
					count++
					fmt.Printf("  %-7d %-40s %-15d %-15d %-20s %-28s %-15s\n", count, dbTransaksi[*kunci].riwayat[i].produk, dbTransaksi[*kunci].riwayat[i].jumlah, dbTransaksi[*kunci].riwayat[i].nilaiAset, dbTransaksi[*kunci].riwayat[i].portofolio, dbTransaksi[*kunci].riwayat[i].tipe, dbTransaksi[*kunci].riwayat[i].tanggal)
				}
			}
			uiFooterTableSuperPanjang()
			
			fmt.Println()
			fmt.Println("[0] Kembali")
			fmt.Print("Pilihan: ")
			fmt.Scan(&pilihan)

			if pilihan == 0{
				clearScreen()
				stat = true
			}
		case 3:
			clearScreen()
			sumber := "TRANSAKSI"
			uiHeaderTable(sumber, 1, 0)
			uiHeaderTransaksi()
			for i:=0; i<NMAX; i++{
				if dbTransaksi[*kunci].riwayat[i].tipe == "Obligasi" && dbTransaksi[*kunci].riwayat[i].nilaiAset > 0{
					count++
					fmt.Printf("  %-7d %-40s %-15d %-15d %-20s %-28s %-15s\n", count, dbTransaksi[*kunci].riwayat[i].produk, dbTransaksi[*kunci].riwayat[i].jumlah, dbTransaksi[*kunci].riwayat[i].nilaiAset, dbTransaksi[*kunci].riwayat[i].portofolio, dbTransaksi[*kunci].riwayat[i].tipe, dbTransaksi[*kunci].riwayat[i].tanggal)
				}
			}
			uiFooterTableSuperPanjang()
			
			fmt.Println()
			fmt.Println("[0] Kembali")
			fmt.Print("Pilihan: ")
			fmt.Scan(&pilihan)

			if pilihan == 0{
				clearScreen()
				stat = true
			}
		case 0:
			clearScreen()
			stat = true
		}
	}
}

func cekIndeksTransaksi(kunci, i *int){
	for *i=0; *i<NMAX; *i++ {
		if dbTransaksi[*kunci].riwayat[*i].nilaiAset == 0{
			break
		}
	}
}

func addTransaksi(kunci *int, tambahanaset, pilihan, tanda, tipereksadana, unit int, untung float64, produk string){
	var (
		posisi int
		jenis string
	)

	switch tipereksadana {
	case 1:
		jenis = "Reksadana Pasar Uang"
	case 2:
		jenis = "Reksadana Saham"
	case 3:
		jenis = "Reksadana Obligasi"
	}

	pilihPorto(kunci, &pilihan)
	cekIndeksTransaksi(kunci, &posisi)

	dbTransaksi[*kunci].riwayat[posisi].untung = untung
	dbTransaksi[*kunci].riwayat[posisi].tanda = 1

	fmt.Println("Tambahan aset: ", tambahanaset)
	if tanda == 1{
		jenis = "Saham"
		dbTransaksi[*kunci].riwayat[posisi].nilaiAset += tambahanaset
		untungRugi(&dbTransaksi, kunci)
		dbPorto[*kunci].detailPorto[pilihan-1].saham += tambahanaset + dbTransaksi[*kunci].riwayat[pilihan-1].nilaiReturn
		dbTransaksi[*kunci].riwayat[posisi].portofolio = dbPorto[*kunci].detailPorto[pilihan-1].tipe
	}else if tanda == 2{
		dbTransaksi[*kunci].riwayat[posisi].nilaiAset += tambahanaset
		untungRugi(&dbTransaksi, kunci)
		dbPorto[*kunci].detailPorto[pilihan-1].reksadana += tambahanaset + dbTransaksi[*kunci].riwayat[pilihan-1].nilaiReturn
		dbTransaksi[*kunci].riwayat[posisi].portofolio = dbPorto[*kunci].detailPorto[pilihan-1].tipe
	}else if tanda == 3{
		jenis = "Obligasi"
		dbTransaksi[*kunci].riwayat[posisi].nilaiAset += tambahanaset
		untungRugi(&dbTransaksi, kunci)
		dbPorto[*kunci].detailPorto[pilihan-1].obligasi += tambahanaset + dbTransaksi[*kunci].riwayat[pilihan-1].nilaiReturn
		dbTransaksi[*kunci].riwayat[posisi].portofolio = dbPorto[*kunci].detailPorto[pilihan-1].tipe
	}

	dbPorto[*kunci].total += tambahanaset
	dbTransaksi[*kunci].riwayat[posisi].jumlah = unit
	dbTransaksi[*kunci].riwayat[posisi].produk = produk
	dbTransaksi[*kunci].riwayat[posisi].tipe = jenis
	dbTransaksi[*kunci].riwayat[posisi].tanggal = time.Now().Format("2006-01-02")

	clearScreen()
	fmt.Println("\nAset berhasil ditambahkan ke portofolio Anda.")
}
