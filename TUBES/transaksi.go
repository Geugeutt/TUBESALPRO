package main
import "fmt"

func tampilTransaksiPembelian(kunci *int, pilihan int){
	for i:=0; i<NMAX; i++{
		if dbTransaksi[*kunci].riwayat[i].tanda == 1 && dbTransaksi[*kunci].riwayat[i].nilaiAset != 0{
			fmt.Printf("%d. %s || %s || %d || %s\n", i+1, dbTransaksi[*kunci].riwayat[i].produk, dbTransaksi[*kunci].riwayat[i].tipe, dbTransaksi[*kunci].riwayat[i].nilaiAset, dbTransaksi[*kunci].riwayat[i].tanggal)
		}
	}
}

func tampilTransaksiPenjualan(kunci *int, pilihan int){
	for i:=0; i<NMAX; i++{
		if dbTransaksi[*kunci].riwayat[i].tanda == 2 && dbTransaksi[*kunci].riwayat[i].nilaiAset != 0{
			fmt.Printf("%d. %s || %s || %d || %s\n", i+1, dbTransaksi[*kunci].riwayat[i].produk, dbTransaksi[*kunci].riwayat[i].tipe, dbTransaksi[*kunci].riwayat[i].nilaiAset, dbTransaksi[*kunci].riwayat[i].tanggal)
		}
	}
}

func filterTransaksi(kunci *int, pilihan int){
	var (
		stat bool
		count int
	)

	fmt.Println("\n=========================================")
	fmt.Println("Filter: ")
	fmt.Println("1. Riwayat Transaksi Pembelian Saham")
	fmt.Println("2. Riwayat Transaksi Pembelian Reksadana")
	fmt.Println("3. Riwayat Transaksi Pembelian Obligasi")
	fmt.Println("0. Kembali")
	fmt.Print("Pilihan: ")
	fmt.Scan(&pilihan)

	for !stat{
		switch pilihan {
		case 1:
			fmt.Println("\n============= Transaksi Pembelian Saham =============")
			for i:=0; i<NMAX; i++{
				if dbTransaksi[*kunci].riwayat[i].tipe == "Saham" && dbTransaksi[*kunci].riwayat[i].tanda == 1 && dbTransaksi[*kunci].riwayat[i].nilaiAset != 0{
					count++
					fmt.Printf("%d. %s || %s || %d || %s", count, dbTransaksi[*kunci].riwayat[i].produk, dbTransaksi[*kunci].riwayat[i].tipe, dbTransaksi[*kunci].riwayat[i].nilaiAset, dbTransaksi[*kunci].riwayat[i].tanggal)
				}
			}
			stat = true
		case 2:
			fmt.Println("\n============= Transaksi Pembelian Reksadana =============")
			for i:=0; i<NMAX; i++{
				if (dbTransaksi[*kunci].riwayat[i].tipe == "Reksadana Pasar Uang" || dbTransaksi[*kunci].riwayat[i].tipe == "Reksadana Saham" || dbTransaksi[*kunci].riwayat[i].tipe == "Reksadana Obligasi") && dbTransaksi[*kunci].riwayat[i].tanda == 1 && dbTransaksi[*kunci].riwayat[i].nilaiAset != 0{
					count++
					fmt.Printf("%d. %s || %s || %d || %s", count, dbTransaksi[*kunci].riwayat[i].produk, dbTransaksi[*kunci].riwayat[i].tipe, dbTransaksi[*kunci].riwayat[i].nilaiAset, dbTransaksi[*kunci].riwayat[i].tanggal)
				}
			}
			stat = true
		case 3:
			fmt.Println("\n============= Transaksi Pembelian Obligasi =============")
			for i:=0; i<NMAX; i++{
				if dbTransaksi[*kunci].riwayat[i].tipe == "Obligasi" && dbTransaksi[*kunci].riwayat[i].tanda == 1 && dbTransaksi[*kunci].riwayat[i].nilaiAset != 0{
					count++
					fmt.Printf("%d. %s || %s || %d || %s", count, dbTransaksi[*kunci].riwayat[i].produk, dbTransaksi[*kunci].riwayat[i].tipe, dbTransaksi[*kunci].riwayat[i].nilaiAset, dbTransaksi[*kunci].riwayat[i].tanggal)
				}
			}
			stat = true
		case 0:
			stat = true
		}
	}
}

func cekIndeksTransaksi(dbTransaksi *[NMAX]transaksi, kunci, i *int){
	for *i=0; *i<NMAX; *i++ {
		if dbTransaksi[*kunci].riwayat[*i].nilaiAset == 0{
			break
		}
	}
}

func addTransaksi(dbporto *[NMAX]porto, dbTransaksi *[NMAX]transaksi, kunci *int, tambahanaset, pilihan, tanda, unit int, produk string){
	var posisi int
	pilihPorto(kunci, &pilihan)
	cekIndeksTransaksi(dbTransaksi, kunci, &posisi)

	fmt.Println("Tambahan aset: ", tambahanaset)
	if tanda == 1{
		dbTransaksi[*kunci].riwayat[posisi].nilaiAset += tambahanaset
		dbPorto[*kunci].detailPorto[pilihan-1].saham += tambahanaset
		dbTransaksi[*kunci].riwayat[posisi].portofolio = dbPorto[*kunci].detailPorto[pilihan-1].tipe
	}else if tanda == 2{
		dbTransaksi[*kunci].riwayat[posisi].nilaiAset += tambahanaset
		dbPorto[*kunci].detailPorto[pilihan-1].reksadana += tambahanaset
		dbTransaksi[*kunci].riwayat[posisi].portofolio = dbPorto[*kunci].detailPorto[pilihan-1].tipe
	}else if tanda == 3{
		dbTransaksi[*kunci].riwayat[posisi].nilaiAset += tambahanaset
		dbPorto[*kunci].detailPorto[pilihan-1].obligasi += tambahanaset
		dbTransaksi[*kunci].riwayat[posisi].portofolio = dbPorto[*kunci].detailPorto[pilihan-1].tipe
	}

	dbTransaksi[*kunci].riwayat[posisi].jumlah = unit
	dbTransaksi[*kunci].riwayat[posisi].produk = produk
	dbTransaksi[*kunci].riwayat[posisi].tanda = 1
	fmt.Println("\nAset berhasil ditambahkan ke portofolio Anda.")
}

func hitungNilaiDijual(dbTransaksi *[NMAX]transaksi, kunci, asetDijual *int, pilihan, jumlah int, produk string){

	switch dbTransaksi[*kunci].riwayat[pilihan-1].tipe {
	case "Saham":
		for i:=0; i<NMAX; i++{
			fmt.Println("plis: ", dbSaham[i].produksaham.namaProduk, " produknya: ", produk)
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
	fmt.Println("\nAset berhasil dijual.")
}

//UI
func introJualAset(dbTransaksi *[NMAX]transaksi, kunci, tipereksadana *int, pilihan int){
	var (
		yakin, produkDijual string
		jumlah, asetDijual int
		stat bool
	)

	for !stat{
		tampilTransaksiPembelian(kunci, pilihan)
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
			fmt.Scan(&jumlah)
			hitungNilaiDijual(dbTransaksi, kunci, &asetDijual, pilihan, jumlah, produkDijual)
			jualAset(pilihan, asetDijual, kunci)
			stat = true
		}else if yakin == "Tidak" || yakin == "tidak"{
			fmt.Println(("Silahkan pilih ulang aset yang akan Anda jual"))
		}else{
			fmt.Println("Pilihan tidak valid")
		}
	}
}