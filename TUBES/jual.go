package main
import "fmt"

func introJualAset(kunci, jumlah, asetDijual *int, pilihan int){
	var (
		yakin, produkDijual string
		count int
		stat bool
	)

	for !stat{
		tampilTransaksiPembelian(kunci, &count)
		if count != 0 {
			fmt.Println()
			fmt.Print("Pilih aset yang akan Anda jual: ")
			fmt.Scan(&pilihan)

			fmt.Println("=========================================")
			fmt.Println("Silahkan periksa kembali pilihan Anda")
			
			fmt.Printf("Nama Produk: %s\nNilai Aset: %d\nTersedia: %d\n", dbTransaksi[*kunci].riwayat[pilihan-1].produk, dbTransaksi[*kunci].riwayat[pilihan-1].nilaiAset, dbTransaksi[*kunci].riwayat[pilihan-1].jumlah)

			produkDijual = dbTransaksi[*kunci].riwayat[pilihan-1].produk

			fmt.Print("Yakin untuk melanjutkan? (Ya/Tidak): ")
			fmt.Scan(&yakin)

			if yakin == "Ya" || yakin == "ya"{
				fmt.Print("Masukkan jumlah aset yang akan Anda jual (dalam satuan lembar): ")
				fmt.Scan(jumlah)
				hitungNilaiDijual(kunci, asetDijual, pilihan, *jumlah, produkDijual)
				jualAset(pilihan, *asetDijual, kunci)
				stat = true
			}else if yakin == "Tidak" || yakin == "tidak"{
				clearScreen()
				fmt.Println("Silahkan pilih ulang aset yang akan Anda jual")
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

	switch dbTransaksi[*kunci].riwayat[pilihan-1].tipe {
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


func jualAset(pilihan, asetDijual int, kunci *int){
	var posisi int

	for i:=0; i<NMAX; i++{
		if dbPorto[*kunci].detailPorto[i].tipe == dbTransaksi[*kunci].riwayat[pilihan-1].portofolio{
			posisi = i
		}
	}

	switch dbTransaksi[*kunci].riwayat[pilihan-1].tipe {
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

	dbTransaksi[*kunci].riwayat[pilihan-1].tanda = 2
	clearScreen()
	fmt.Println("\nAset berhasil dijual.")
}