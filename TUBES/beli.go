package main

import (
	"fmt"
)

func introBeliSaham(kunci *int){
	var (
		yakin, produk string
		pilihan int
		stat bool
	)

	sumber := "SAHAM"

	for !stat {
		fmt.Println("Pilih saham yang akan Anda beli")
		fmt.Print("Pilihan: ")
		fmt.Scan(&produk)
		fmt.Println()

		clearScreen()
		fmt.Println("==========================================")
		fmt.Println("Silahkan periksa kembali pilihan Anda")
		pilihan = cariAset(sumber) + 1
		fmt.Printf("\nNama Produk: %s \nHarga/lembar: %d \nReturn: %.2f\n", dbSaham[pilihan-1].produksaham.namaProduk, dbSaham[pilihan-1].produksaham.hargaPerLembar, dbSaham[pilihan-1].produksaham.returnaset)
		fmt.Println()

		produk = dbObligasi[pilihan-1].produkobligasi.namaProduk
		untung := dbSaham[pilihan-1].produksaham.returnaset

		fmt.Print("Yakin untuk melanjutkan? (Ya/Tidak): ")
		fmt.Scan(&yakin)
		
		if yakin == "Ya" || yakin == "ya"{
			beliSaham(untung, pilihan, kunci, yakin, produk)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			fmt.Println("Silahkan pilih ulang produk yang akan Anda beli")
			clearScreen()
		}else{
			fmt.Println("Pilihan tidak valid")
			clearScreen()
		}
	}
}

func beliSaham(untung float64, pilihan int, kunci *int, yakin, produk string){
	var (
		unit, tambahanaset, tanda int
		stat, cek bool
	)

	
	tanda = 1
	for !stat {
		fmt.Print("Banyak lembar yang akan anda beli: ")
		fmt.Scan(&unit)
		cek = false

		for !cek {
			tambahanaset = dbSaham[pilihan-1].produksaham.hargaPerLembar*unit
			fmt.Printf("Total nilai aset yang akan dibeli: %d", tambahanaset)
			fmt.Print("\nKonfirmasi pembayaran [Ya/Tidak]: ")
			fmt.Scan(&yakin)

			if yakin == "Ya" || yakin == "ya"{
				addTransaksi(kunci, tambahanaset, pilihan, tanda, 0, unit, untung, produk)
				cek = true
				stat = true
			}else if yakin == "Tidak" || yakin == "tidak"{
				clearScreen()
				fmt.Println("Silahkan lakukan transaksi ulang")
				cek = true
			}else{
				clearScreen()
				fmt.Println("Pilihan tidak valid")
				fmt.Println()
			}
		}
	}
}


func introBeliReksa(tipereksadana int, kunci *int){
	var (
		yakin, produk string
		pilihan int
		stat bool
	)

	for !stat {
		fmt.Println("Pilih reksadana yang akan Anda beli")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		clearScreen()
		fmt.Println("=========================================")
		fmt.Println("Silahkan periksa kembali pilihan Anda")
		fmt.Printf("\nNama Produk: %s \nBank Kustodian: %s \nBank Penampung: %s\nMinimal Pembelian: %d \n", dbReksadana[tipereksadana-1].produkreksa[pilihan-1].Produk, dbReksadana[tipereksadana-1].produkreksa[pilihan-1].kustodian, dbReksadana[tipereksadana-1].produkreksa[pilihan-1].penampung, dbReksadana[tipereksadana-1].produkreksa[pilihan-1].minimal)
		fmt.Printf("Return: %.2f\n", dbReksadana[tipereksadana-1].produkreksa[pilihan-1].returnreksa) 
		fmt.Println()

		produk = dbReksadana[tipereksadana-1].produkreksa[pilihan-1].Produk
		untung := dbReksadana[tipereksadana-1].produkreksa[pilihan-1].returnreksa

		fmt.Print("Yakin untuk melanjutkan? (Ya/Tidak): ")
		fmt.Scan(&yakin)

		if yakin == "Ya" || yakin == "ya"{
			beliReksa(untung, pilihan, tipereksadana, kunci, yakin, produk)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			clearScreen()
			fmt.Println("Silahkan pilih ulang produk yang akan Anda beli")
		}else{
			clearScreen()
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func beliReksa(untung float64, pilihan, tipereksadana int, kunci *int, yakin, produk string){
	var (
		unit, tambahanaset, tanda int
		stat, cek bool
	)

	tanda = 2

	for !stat {
		fmt.Print("Banyak lembar yang akan anda beli: ")
		fmt.Scan(&unit)
		cek = false

		for !cek{
			tambahanaset = dbReksadana[tipereksadana-1].produkreksa[pilihan-1].minimal*unit
			fmt.Print("Total nilai aset yang akan dibeli: ", tambahanaset)

			fmt.Print("\nKonfirmasi pembayaran [Ya/Tidak]: ")
			fmt.Scan(&yakin)

			if yakin == "Ya" || yakin == "ya"{
				addTransaksi(kunci, tambahanaset, pilihan, tanda, tipereksadana, unit, untung, produk)
				cek = true
				stat = true
			}else if yakin == "Tidak" || yakin == "tidak"{
				clearScreen()
				fmt.Println("Silahkan lakukan transaksi ulang")
				cek = true
			}else{
				clearScreen()
				fmt.Println("Pilihan tidak valid")
				fmt.Println()
			}
		}
	}
}

func introBeliObligasi(kunci *int){
	var (
		yakin, produk string
		pilihan int
		stat bool
	)

	for !stat {
		fmt.Println("Pilih Obligasi yang akan Anda beli")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		fmt.Println("=========================================")
		fmt.Println("Silahkan periksa kembali pilihan Anda")
		fmt.Printf("\nNama Produk: %s \nHarga/unit: %d \nReturn: %.2f\n", dbObligasi[pilihan-1].produkobligasi.namaProduk, dbObligasi[pilihan-1].produkobligasi.hargaPerLembar, dbObligasi[pilihan-1].produkobligasi.returnaset)
		fmt.Printf("Jatuh Tempo: %s \nKupon Selanjutnya: %s \nPenerbit: %s \nPemberian Kupon: per %d bulan\n", dbObligasi[pilihan-1].jatuhtempo, dbObligasi[pilihan-1].nextkupon, dbObligasi[pilihan-1].penerbit, dbObligasi[pilihan-1].kupon)
		fmt.Println()

		produk = dbObligasi[pilihan-1].produkobligasi.namaProduk
		untung := dbObligasi[pilihan-1].produkobligasi.returnaset

		fmt.Print("Yakin untuk melanjutkan? (Ya/Tidak): ")
		fmt.Scan(&yakin)

		if yakin == "Ya" || yakin == "ya"{
			beliObligasi(untung, pilihan, kunci, yakin, produk)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			clearScreen()
			fmt.Println("Silahkan pilih ulang produk yang akan Anda beli")
		}else{
			clearScreen()
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func beliObligasi(untung float64, pilihan int, kunci *int, yakin, produk string){
	var (
		unit, tambahanaset, tanda int
		stat, cek bool
	)

	tanda = 3

	for !stat {
		fmt.Print("Banyak lembar yang akan anda beli: ")
		fmt.Scan(&unit)
		cek = false

		for !cek{
			tambahanaset = dbObligasi[pilihan-1].produkobligasi.hargaPerLembar*unit
			fmt.Print("Total nilai aset yang akan dibeli: ", tambahanaset)

			fmt.Print("\nKonfirmasi pembayaran [Ya/Tidak]: ")
			fmt.Scan(&yakin)

			if yakin == "Ya" || yakin == "ya"{
				addTransaksi(kunci, tambahanaset, pilihan, tanda, 0, unit, untung, produk)
				cek = true
				stat = true
			}else if yakin == "Tidak" || yakin == "tidak"{
				clearScreen()
				fmt.Println("Silahkan lakukan transaksi ulang")
				cek = true
			}else{
				clearScreen()
				fmt.Println("Pilihan tidak valid")
				fmt.Println()
			}
		}
	}
}