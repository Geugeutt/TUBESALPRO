package main
import "fmt"

func introBeliSaham(dbSaham[]saham, pilihan, kunci int){
	var yakin string

	fmt.Println("Pilih saham yang akan Anda beli")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)
	fmt.Println()

	fmt.Println("=========================================")
	fmt.Println("Silahkan periksa kembali pilihan Anda")
	fmt.Printf("\nNama Produk: %s \nHarga/lembar: %d \nReturn: %.2f\n", dbSaham[pilihan-1].produksaham.namaProduk, dbSaham[pilihan-1].produksaham.hargaPerLembar, dbSaham[pilihan-1].produksaham.returnaset)
    fmt.Println()

	var stat bool
	for !stat {
		fmt.Print("Yakin untuk melanjutkan? (Ya/Tidak): ")
		fmt.Scan(&yakin)
		
		if yakin == "Ya" || yakin == "ya"{
			beliSaham(dbSaham, pilihan, kunci, yakin)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			introBeliSaham(dbSaham, pilihan, kunci)
			stat = true
		}else{
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func beliSaham(dbSaham[]saham, pilihan, kunci int, yakin string){
	var (
		unit, tambahanaset, tanda int
	)

	tanda = 1
	fmt.Print("Banyak lembar yang akan anda beli: ")
	fmt.Scan(&unit)

	tambahanaset = dbSaham[pilihan-1].produksaham.hargaPerLembar*unit
	fmt.Printf("Total nilai aset yang akan dibeli: %d", tambahanaset)

	var stat bool
	for !stat {
		fmt.Print("\nKonfirmasi pembayaran [Ya/Tidak]: ")
		fmt.Scan(&yakin)

		if yakin == "Ya" || yakin == "ya"{
			addPorto(&dbPorto, kunci, tambahanaset, pilihan, tanda)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			beliSaham(dbSaham, pilihan, kunci, yakin)
			stat = true
		}else{
			fmt.Println("Pilihan tidak valid")
		}
	}
}


func introBeliReksa(dbReksadana[]reksadana, j, i, kunci int){
	var yakin string

	fmt.Println("Pilih reksadana yang akan Anda beli")
	fmt.Print("Pilihan: ")
	fmt.Scan(&j)
	fmt.Println()

	fmt.Println("=========================================")
	fmt.Println("Silahkan periksa kembali pilihan Anda")
	fmt.Printf("\nNama Produk: %s \nBank Kustodian: %s \nBank Penampung: %s\nMinimal Pembelian: %d \n", dbReksadana[i].produkreksa[j-1].Produk, dbReksadana[i].produkreksa[j-1].kustodian, dbReksadana[i].produkreksa[j-1].penampung, dbReksadana[i].produkreksa[j-1].minimal)
	fmt.Printf("Return: %.2f\n", dbReksadana[i].produkreksa[j-1].returnreksa) 
	fmt.Println()

	var stat bool
	for !stat {
		fmt.Print("Yakin untuk melanjutkan? (Ya/Tidak): ")
		fmt.Scan(&yakin)

		if yakin == "Ya" || yakin == "ya"{
			beliReksa(dbReksadana, j, i, kunci, yakin)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			introBeliReksa(dbReksadana, j, i, kunci)
			stat = true
		}else{
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func beliReksa(dbReksadana[]reksadana, j, i, kunci int, yakin string){
	var (
		unit, tambahanaset, tanda int
	)

	tanda = 2

	fmt.Print("Banyak lembar yang akan anda beli: ")
	fmt.Scan(&unit)

	tambahanaset = dbReksadana[i].produkreksa[j-1].minimal*unit
	fmt.Print("Total nilai aset yang akan dibeli: ", tambahanaset)

	var stat bool
	for !stat {
		fmt.Print("\nKonfirmasi pembayaran [Ya/Tidak]: ")
		fmt.Scan(&yakin)

		if yakin == "Ya" || yakin == "ya"{
			addPorto(&dbPorto, kunci, tambahanaset, j, tanda)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			beliReksa(dbReksadana, j, i, kunci, yakin)
			stat = true
		}else{
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func introBeliObligasi(dbObligasi[]obligasi, kunci, j int){
	var yakin string

	fmt.Println("Pilih Obligasi yang akan Anda beli")
	fmt.Print("Pilihan: ")
	fmt.Scan(&j)
	fmt.Println()

	fmt.Println("=========================================")
	fmt.Println("Silahkan periksa kembali pilihan Anda")
	fmt.Printf("\nNama Produk: %s \nHarga/unit: %d \nReturn: %.2f\n", dbObligasi[j-1].produkobligasi.namaProduk, dbObligasi[j-1].produkobligasi.hargaPerLembar, dbObligasi[j-1].produkobligasi.returnaset)
	fmt.Printf("Jatuh Tempo: %s \nKupon Selanjutnya: %s \nPenerbit: %s \nPemberian Kupon: per %d bulan\n", dbObligasi[j-1].jatuhtempo, dbObligasi[j-1].nextkupon, dbObligasi[j-1].penerbit, dbObligasi[j-1].kupon)
	fmt.Println()

	var stat bool
	for !stat {
		fmt.Print("Yakin untuk melanjutkan? (Ya/Tidak): ")
		fmt.Scan(&yakin)

		if yakin == "Ya" || yakin == "ya"{
			beliObligasi(dbObligasi, j, kunci, yakin)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			introBeliObligasi(dbObligasi, kunci, j)
			stat = true
		}else{
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func beliObligasi(dbObligasi[]obligasi, pilihan, kunci int, yakin string){
	var (
		unit, tambahanaset, tanda int
	)

	tanda = 3

	fmt.Print("Banyak lembar yang akan anda beli: ")
	fmt.Scan(&unit)

	tambahanaset = dbObligasi[pilihan-1].produkobligasi.hargaPerLembar*unit
	fmt.Print("Total nilai aset yang akan dibeli: ", tambahanaset)

	var stat bool
	for !stat {
		fmt.Print("\nKonfirmasi pembayaran [Ya/Tidak]: ")
		fmt.Scan(&yakin)

		if yakin == "Ya" || yakin == "ya"{
			addPorto(&dbPorto, kunci, tambahanaset, pilihan, tanda)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			beliObligasi(dbObligasi, pilihan, kunci, yakin)
			stat = true
		}else{
			fmt.Println("Pilihan tidak valid")
		}
	}
}