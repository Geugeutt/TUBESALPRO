package main
import "fmt"

func tampilPortofolio(kunci, tipereksadana, asetDijual, jumlahDijual *int){
	var (
		saham, reksadana, obligasi float64
		pilihan, pilihanporto, totalLokal, count int
		stat, sort bool
	)

	pilihPorto(kunci, &pilihanporto)
	for !sort {
		namaPorto := dbPorto[*kunci].detailPorto[pilihanporto-1].tipe
		if pilihanporto != 0{
			hitungDetail(&saham, &reksadana, &obligasi, kunci, &totalLokal, pilihanporto)

			clearScreen()
			sumber := "PORTOFOLIO"
			uiHeaderTable(sumber, 0, 0)
			uiHeaderPorto(dbUser[*kunci].username, namaPorto, totalLokal, saham, reksadana, obligasi)
			for i:=0 ; i<NMAX ; i++{
				if dbTransaksi[*kunci].riwayat[i].nilaiAset != 0 && dbTransaksi[*kunci].riwayat[i].portofolio == namaPorto{
					count++
					if dbTransaksi[*kunci].riwayat[i].nilaiReturn >= 0 {
						fmt.Printf("%-1s %-5d %-40s %-24s %-17d  %-.2f %-12s %-18d %-21s\n", " ", count, dbTransaksi[*kunci].riwayat[i].produk, dbTransaksi[*kunci].riwayat[i].tipe, dbTransaksi[*kunci].riwayat[i].nilaiAset + dbTransaksi[*kunci].riwayat[i].nilaiReturn, dbTransaksi[*kunci].riwayat[i].untung, " ", dbTransaksi[*kunci].riwayat[i].nilaiReturn, dbTransaksi[*kunci].riwayat[i].tanggal)
					}else{
						fmt.Printf("%-1s %-5d %-40s %-24s %-16d  %-.2f %-12s %-18d %-23s\n", " ", count, dbTransaksi[*kunci].riwayat[i].produk, dbTransaksi[*kunci].riwayat[i].tipe, dbTransaksi[*kunci].riwayat[i].nilaiAset + dbTransaksi[*kunci].riwayat[i].nilaiReturn, dbTransaksi[*kunci].riwayat[i].untung, " ", dbTransaksi[*kunci].riwayat[i].nilaiReturn, dbTransaksi[*kunci].riwayat[i].tanggal)
					}
				}
			}
			uiFooterTablePanjang()
		} 

		for !stat{
			fmt.Println()
			fmt.Println("[1] Cari aset")
			fmt.Println("[2] Urutkan aset berdasarkan nilai aset")
			fmt.Println("[3] Beli aset")
			fmt.Println("[4] Jual aset")
			fmt.Println("[5] Hapus portofolio")
			fmt.Println("[0] Kembali")
			fmt.Println()
			fmt.Print("Pilihan: ")
			fmt.Scan(&pilihan)

			switch pilihan{
			case 1:

			case 2:
				sortPorto(kunci)
				stat = true
			case 3:
				clearScreen()
				katalog(pilihan, kunci, tipereksadana)
				stat = true
				sort = true
			case 4:
				clearScreen()
				if pilihanporto != 0{
					introJualAset(kunci, jumlahDijual, asetDijual, pilihan)
					sort = true
				}else {
					fmt.Println("Anda belum memiliki aset apapun")
				}
			case 5:
				clearScreen()
				if pilihanporto != 0{
					hapusPorto(kunci, pilihanporto)
					sort = true
				}else{
					fmt.Print("Anda belum memiliki portofolio")
				}
			case 0:
				clearScreen()
				sort = true
				stat = true
			default:
				clearScreen()
				fmt.Println("Pilihan tidak valid")
			}
		}
	}
}

func hitungDetail(saham, reksadana, obligasi *float64, kunci, totalLokal *int, pilihan int){
	*totalLokal = dbPorto[*kunci].detailPorto[pilihan-1].saham + dbPorto[*kunci].detailPorto[pilihan-1].reksadana + dbPorto[*kunci].detailPorto[pilihan-1].obligasi

	if dbPorto[*kunci].detailPorto[pilihan-1].saham != 0{
		*saham = float64(dbPorto[*kunci].detailPorto[pilihan-1].saham)*100.0/float64(*totalLokal)
	}else{
		*saham = 0
	}

	if dbPorto[*kunci].detailPorto[pilihan-1].reksadana != 0{
		*reksadana = float64(dbPorto[*kunci].detailPorto[pilihan-1].reksadana)*100.0/float64(*totalLokal)
	}else{
		*reksadana = 0
	}

	if dbPorto[*kunci].detailPorto[pilihan-1].obligasi != 0{
		*obligasi = float64(dbPorto[*kunci].detailPorto[pilihan-1].obligasi)*100.0/float64(*totalLokal)
	}else{
		*obligasi = 0
	}
}

func pilihPorto(kunci, pilihan *int){
	var count, k int
	fmt.Println("\nPilih portofolio Anda: ")
	for k=0;k<NMAX;k++{
		if dbPorto[*kunci].detailPorto[k].tanda == 1 {
			fmt.Printf("[%d] %s\n", k+1, dbPorto[*kunci].detailPorto[k].tipe)
			count++
		}else {
			break
		}
	}

	switch count {
	case 0:
		clearScreen()
		fmt.Println("Anda belum memiliki portofolio. Harap buat satu.")
		fmt.Println()
		buatPortoBaru(kunci, k)
		*pilihan = k+1
	default:
		fmt.Print("[0] Buat portofolio baru\n")
		fmt.Print("Pilihan: ")
		fmt.Scan(pilihan)

		if *pilihan == 0{
			clearScreen()
			buatPortoBaru(kunci, k)
			*pilihan = k+1
		}
	}
}

func buatPortoBaru(kunci *int, pilihan int){
	var (
		portobaru string
		stat bool
	)

	for !stat{
		fmt.Print("Nama Portofolio: ")
		fmt.Scan(&portobaru)

		dbPorto[*kunci].detailPorto[pilihan].tipe = portobaru
		dbPorto[*kunci].detailPorto[pilihan].tanda = 1

		if dbPorto[*kunci].detailPorto[pilihan].tipe == portobaru{
			fmt.Print("Portofolio berhasil ditambahkan!\n\n")
			stat = true
		}else{
			clearScreen()
			fmt.Println("Gagal menambahkan portofolio, coba lagi.")
		}
	}
}

func hapusPorto(kunci *int, pilihPorto int){

	dbPorto[*kunci].total -= dbPorto[*kunci].detailPorto[pilihPorto-1].saham + dbPorto[*kunci].detailPorto[pilihPorto-1].obligasi + dbPorto[*kunci].detailPorto[pilihPorto-1].reksadana

	for k:=pilihPorto-1; k<NMAX; k++{
		if dbPorto[*kunci].detailPorto[k].tanda == 0 && dbPorto[*kunci].detailPorto[k+1].tanda == 0 {
			break
		}

		dbPorto[*kunci].detailPorto[k].obligasi = dbPorto[*kunci].detailPorto[k+1].obligasi
		dbPorto[*kunci].detailPorto[k].reksadana = dbPorto[*kunci].detailPorto[k+1].reksadana
		dbPorto[*kunci].detailPorto[k].saham = dbPorto[*kunci].detailPorto[k+1].saham
		dbPorto[*kunci].detailPorto[k].tanda = dbPorto[*kunci].detailPorto[k+1].tanda
		dbPorto[*kunci].detailPorto[k].tipe = dbPorto[*kunci].detailPorto[k+1].tipe			
	}

	fmt.Println("Portofolio berhasil dihapus")
}

