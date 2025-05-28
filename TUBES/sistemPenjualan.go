package main
import "fmt"

func introJualAset(kunci, jumlah, asetDijual *int, pilihan, totalLokal int, saham, reksadana, obligasi float64, namaPorto string){
	var (
		yakin, produkDijual string
		count int
		stat bool
	)

	for !stat{
		sumber := "PORTOFOLIO"
		uiHeaderTable(sumber, 0, 0)
		uiHeaderPorto(dbUser[*kunci].username, namaPorto, totalLokal, saham, reksadana, obligasi)

		tampilArrayBantu(kunci, &count, namaPorto)

		uiFooterTablePanjang()
		if count != 0 {
			fmt.Println()
			fmt.Print("Pilih aset yang akan Anda jual: ")
			fmt.Scan(&pilihan)

			if dbTempforSort[*kunci].riwayat[pilihan-1].nilaiAset+dbTempforSort[*kunci].riwayat[pilihan-1].nilaiReturn <= 0{
				fmt.Println("Nilai aset Anda tidak valid (kerugian/loss)")
				fmt.Print("Pilih ulang aset yang akan Anda jual: ")
				fmt.Scan(&pilihan)
			}

			fmt.Println("=========================================")
			fmt.Println("Silahkan periksa kembali pilihan Anda")
			
			fmt.Printf("Nama Produk: %s\nNilai Aset: %d\nTersedia: %d\n", dbTempforSort[*kunci].riwayat[pilihan-1].produk, dbTempforSort[*kunci].riwayat[pilihan-1].nilaiAset+dbTempforSort[*kunci].riwayat[pilihan-1].nilaiReturn, dbTempforSort[*kunci].riwayat[pilihan-1].jumlah)
			produkDijual = dbTempforSort[*kunci].riwayat[pilihan-1].produk

			fmt.Print("Yakin untuk melanjutkan? (Ya/Tidak): ")
			fmt.Scan(&yakin)

			if yakin == "Ya" || yakin == "ya"{
				fmt.Print("Masukkan jumlah aset yang akan Anda jual (dalam satuan lembar): ")
				fmt.Scan(jumlah)
				clearScreen()
				hitungNilaiDijual(kunci, asetDijual, pilihan, *jumlah, produkDijual)
				jualAset(pilihan, *jumlah, *asetDijual, kunci)
				stat = true
			}else if yakin == "Tidak" || yakin == "tidak"{
				clearScreen()
				fmt.Println("Silahkan pilih ulang aset yang akan Anda jual")
				count = 0
			}else{
				clearScreen()
				fmt.Println("Pilihan tidak valid")
				count = 0
			}
		}else{
			clearScreen()
			stat = true
		}
	}
}

func hitungNilaiDijual(kunci, asetDijual *int, pilihan, jumlah int, produk string){

	switch dbTempforSort[*kunci].riwayat[pilihan-1].tipe {
	case "Saham":
		for i:=0; i<NMAX; i++{
			if dbSaham[i].produksaham.namaProduk == produk{
				*asetDijual = dbSaham[i].produksaham.hargaPerLembar*jumlah
				break
			}
		}
	case "Obligasi":
		for i:=0; i<NMAX; i++{
			if dbObligasi[i].produkobligasi.namaProduk == produk{
				*asetDijual = dbObligasi[i].produkobligasi.hargaPerLembar*jumlah
				break
			}
		}
	case "Reksadana Pasar Uang":
		for i:=0; i<NMAX; i++{
			if dbReksadana[0].produkreksa[i].Produk == produk{
				*asetDijual = dbReksadana[0].produkreksa[i].minimal*jumlah
				break
			}
		}
	case "Reksadana Saham":
		for i:=0; i<NMAX; i++{
			if dbReksadana[1].produkreksa[i].Produk == produk{
				*asetDijual = dbReksadana[1].produkreksa[i].minimal*jumlah
				break
			}
		}
	case "Reksadana Obligasi":
		for i:=0; i<NMAX; i++{
			if dbReksadana[2].produkreksa[i].Produk == produk{
				*asetDijual = dbReksadana[2].produkreksa[i].minimal*jumlah
				break
			}
		}
	}

	fmt.Printf("Total nilai yang akan dijual: %d", *asetDijual)
}

func untungRugi(db *[NMAX]transaksi, kunci *int){
	for i:=0 ; i<NMAX ; i++{
		if db[*kunci].riwayat[i].tanda == 1{
			db[*kunci].riwayat[i].nilaiReturn = int(db[*kunci].riwayat[i].untung * float64(db[*kunci].riwayat[i].nilaiAset))
		}
	}
}

func jualAset(pilihan, jumlahDijual, asetDijual int, kunci *int){
	var posisi int

	for i:=0; i<NMAX; i++{
		if dbPorto[*kunci].detailPorto[i].tipe == dbTempforSort[*kunci].riwayat[pilihan-1].portofolio{
			posisi = i
		}
	}

	switch dbTempforSort[*kunci].riwayat[pilihan-1].tipe {
	case "Saham":
		dbPorto[*kunci].detailPorto[posisi].saham -= asetDijual
	case "Obligasi":
		dbPorto[*kunci].detailPorto[posisi].obligasi -= asetDijual
	case "Reksadana Pasar Uang":
		dbPorto[*kunci].detailPorto[posisi].reksadana -= asetDijual
	case "Reksadana Saham":
		dbPorto[*kunci].detailPorto[posisi].reksadana -= asetDijual
	case "Reksadana Obligasi":
		dbPorto[*kunci].detailPorto[posisi].reksadana -= asetDijual
	}

	perubahanAfterPenjualan(kunci, pilihan, jumlahDijual)

	fmt.Println("\nAset berhasil dijual.")
}

func perubahanAfterPenjualan(kunci *int, pilihan, jumlahDijual int){
	dbTempforSort[*kunci].riwayat[pilihan-1].tanda = 2

	hargaSatuan := dbTempforSort[*kunci].riwayat[pilihan-1].nilaiAset/dbTempforSort[*kunci].riwayat[pilihan-1].jumlah

	dbTempforSort[*kunci].riwayat[pilihan-1].jumlah -= jumlahDijual
	dbTempforSort[*kunci].riwayat[pilihan-1].nilaiAset = hargaSatuan*dbTempforSort[*kunci].riwayat[pilihan-1].jumlah

	untungRugi(&dbTempforSort, kunci)

	for i:=0;i<NMAX;i++{
		if dbTransaksi[*kunci].riwayat[i].produk == dbTempforSort[*kunci].riwayat[pilihan-1].produk{
			dbTransaksi[*kunci].riwayat[i].tanda = 2
			break
		}
	}
}