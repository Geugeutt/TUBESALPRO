package main

import (
	"fmt"
)

func introBeliSaham(dbSaham *[NMAX]saham, pilihan int, kunci *int){
	var (
		yakin, produk string
		stat bool
	)

	for !stat {
	fmt.Println("Pilih saham yang akan Anda beli")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)
	fmt.Println()

	fmt.Println("=========================================")
	fmt.Println("Silahkan periksa kembali pilihan Anda")
	fmt.Printf("\nNama Produk: %s \nHarga/lembar: %d \nReturn: %.2f\n", dbSaham[pilihan-1].produksaham.namaProduk, dbSaham[pilihan-1].produksaham.hargaPerLembar, dbSaham[pilihan-1].produksaham.returnaset)
    fmt.Println()

	produk = dbSaham[pilihan-1].produksaham.namaProduk

		fmt.Print("Yakin untuk melanjutkan? (Ya/Tidak): ")
		fmt.Scan(&yakin)
		
		if yakin == "Ya" || yakin == "ya"{
			beliSaham(dbSaham, pilihan, kunci, yakin, produk)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			fmt.Println(("Silahkan pilih ulang produk yang akan Anda beli"))
		}else{
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func beliSaham(dbSaham *[NMAX]saham, pilihan int, kunci *int, yakin, produk string){
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
				addTransaksi(&dbPorto, &dbTransaksi, kunci, tambahanaset, pilihan, tanda, unit, produk)
				cek = true
				stat = true
			}else if yakin == "Tidak" || yakin == "tidak"{
				fmt.Println("Silahkan ...")
				cek = true
			}else{
				fmt.Println("Pilihan tidak valid")
				fmt.Println()
			}
		}
	}
}


func introBeliReksa(dbReksadana *[3]reksadana, pilihan, tipereksadana int, kunci *int){
	var (
		yakin, produk string
		stat bool
	)

	for !stat {
		fmt.Println("Pilih reksadana yang akan Anda beli")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		fmt.Println("=========================================")
		fmt.Println("Silahkan periksa kembali pilihan Anda")
		fmt.Printf("\nNama Produk: %s \nBank Kustodian: %s \nBank Penampung: %s\nMinimal Pembelian: %d \n", dbReksadana[tipereksadana-1].produkreksa[pilihan-1].Produk, dbReksadana[tipereksadana-1].produkreksa[pilihan-1].kustodian, dbReksadana[tipereksadana-1].produkreksa[pilihan-1].penampung, dbReksadana[tipereksadana-1].produkreksa[pilihan-1].minimal)
		fmt.Printf("Return: %.2f\n", dbReksadana[tipereksadana-1].produkreksa[pilihan-1].returnreksa) 
		fmt.Println()

		produk = dbReksadana[tipereksadana-1].produkreksa[pilihan-1].Produk

		fmt.Print("Yakin untuk melanjutkan? (Ya/Tidak): ")
		fmt.Scan(&yakin)

		if yakin == "Ya" || yakin == "ya"{
			beliReksa(dbReksadana, pilihan, tipereksadana, kunci, yakin, produk)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			fmt.Println(("Silahkan pilih ulang produk yang akan Anda beli"))
		}else{
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func beliReksa(dbReksadana *[3]reksadana, pilihan, tipereksadana int, kunci *int, yakin, produk string){
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
				addTransaksi(&dbPorto, &dbTransaksi, kunci, tambahanaset, pilihan, tanda, unit, produk)
				cek = true
				stat = true
			}else if yakin == "Tidak" || yakin == "tidak"{
				fmt.Println("Silahkan ...")
				cek = true
			}else{
				fmt.Println("Pilihan tidak valid")
				fmt.Println()
			}
		}
	}
}

func introBeliObligasi(dbObligasi *[NMAX]obligasi, kunci *int, pilihan int){
	var (
		yakin, produk string
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

		fmt.Print("Yakin untuk melanjutkan? (Ya/Tidak): ")
		fmt.Scan(&yakin)

		if yakin == "Ya" || yakin == "ya"{
			beliObligasi(dbObligasi, pilihan, kunci, yakin, produk)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			fmt.Println(("Silahkan pilih ulang produk yang akan Anda beli"))
		}else{
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func beliObligasi(dbObligasi *[NMAX]obligasi, pilihan int, kunci *int, yakin, produk string){
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
				addTransaksi(&dbPorto, &dbTransaksi, kunci, tambahanaset, pilihan, tanda, unit, produk)
				cek = true
				stat = true
			}else if yakin == "Tidak" || yakin == "tidak"{
				fmt.Println("Silahkan ...")
				cek = true
			}else{
				fmt.Println("Pilihan tidak valid")
				fmt.Println()
			}
		}
	}
}