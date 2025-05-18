package main
import "fmt"

// bisa tampilin jumlah nilai aset semuanya, nilai aset tiap jenis invest, presentase persebaran aset ke tiap jenis invest
func tampilPortofolio(kunci, tipereksadana *int){
	var (
		saham, reksadana, obligasi float64
		pilihan int
		stat bool
	)

	pilihPorto(kunci, &pilihan)
	hitungDetail(&saham, &reksadana, &obligasi, kunci, pilihan)

	fmt.Printf("\n======= Portofolio %s ======\n", dbUser[*kunci].username)
	fmt.Printf("Jenis Portofolio: %s", dbPorto[*kunci].detailPorto[pilihan-1].tipe)
	fmt.Printf("\nTotal Aset: %d", dbPorto[*kunci].total)
	fmt.Println("\n======= Detail Portofolio ======")
	fmt.Printf("Saham: %d\nReksadana: %d\nObligasi: %d\n", dbPorto[*kunci].detailPorto[pilihan-1].saham, dbPorto[*kunci].detailPorto[pilihan-1].reksadana, dbPorto[*kunci].detailPorto[pilihan-1].obligasi)
	fmt.Println("======= Persentase Portofolio ======")
	fmt.Printf("Saham: %.2f persen\nReksadana: %.2f persen\nObligasi: %.2f persen\n\n", saham, reksadana, obligasi)

	fmt.Println("1. Beli aset")
	fmt.Println("2. Jual aset")
	fmt.Println("3. Analisis aset")
	fmt.Println("4. Hapus portofolio")
	fmt.Println("0. Kembali")
	fmt.Println()
	for !stat{
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan{
		case 1:
			katalog(&dbReksadana, &dbSaham, &dbObligasi, pilihan, kunci, tipereksadana)
			stat = true
		case 2:
			introJualAset(&dbTransaksi, kunci, tipereksadana, pilihan)
			stat = true
		case 3:
			//panggil prosedur analisis porto
		case 4:
			hapusPorto(kunci, pilihan)
			stat = true
		case 0:
			stat = true
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func hitungDetail(saham, reksadana, obligasi *float64, kunci *int, pilihan int){
	var totalLokal int
	for i:=0 ; i<NMAX ; i++{
		if dbPorto[*kunci].detailPorto[i].tanda == 0{
			break
		}

		dbPorto[*kunci].total += dbPorto[*kunci].detailPorto[i].saham + dbPorto[*kunci].detailPorto[i].reksadana + dbPorto[*kunci].detailPorto[i].obligasi
	}

	totalLokal = dbPorto[*kunci].detailPorto[pilihan-1].saham + dbPorto[*kunci].detailPorto[pilihan-1].reksadana + dbPorto[*kunci].detailPorto[pilihan-1].obligasi

	if dbPorto[*kunci].detailPorto[pilihan-1].saham != 0{
		*saham = float64(dbPorto[*kunci].detailPorto[pilihan-1].saham)*100.0/float64(totalLokal)
	}else{
		*saham = 0
	}

	if dbPorto[*kunci].detailPorto[pilihan-1].reksadana != 0{
		*reksadana = float64(dbPorto[*kunci].detailPorto[pilihan-1].reksadana)*100.0/float64(totalLokal)
	}else{
		*reksadana = 0
	}

	if dbPorto[*kunci].detailPorto[pilihan-1].obligasi != 0{
		*obligasi = float64(dbPorto[*kunci].detailPorto[pilihan-1].obligasi)*100.0/float64(totalLokal)
	}else{
		*obligasi = 0
	}
}

func pilihPorto(kunci, pilihan *int){
	var count, k int
	fmt.Println("\nPilih portofolio Anda: ")
	for k=0;k<NMAX;k++{
		if dbPorto[*kunci].detailPorto[k].tanda == 1 {
			fmt.Printf("%d. %s\n", k+1, dbPorto[*kunci].detailPorto[k].tipe)
			count++
		}else {
			break
		}
	}

	switch count {
	case 0:
		fmt.Println("\nAnda belum memiliki portofolio. Harap buat satu.")
		buatPortoBaru(kunci, k)
	default:
		fmt.Print("0. Buat portofolio baru\n")
		fmt.Print("Pilihan: ")
		fmt.Scan(pilihan)

		if *pilihan == 0{
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
			fmt.Println("Gagal menambahkan portofolio, coba lagi.")
		}
	}
}

func hapusPorto(kunci *int, pilihan int){
	var (
		l int
		saham, obligasi, reksadana float64
	)

	for i:=0; i<NMAX; i++{
		if dbPorto[*kunci].detailPorto[i].tanda !=0 {
			fmt.Printf("%d. total = %d\nsubtotal obligasi: %d\nsubtotal reksadana: %d\nsubtotal saham: %d\ntanda: %d\ntipe: %s\n", i+1, dbPorto[*kunci].total, dbPorto[*kunci].detailPorto[i].obligasi, dbPorto[*kunci].detailPorto[i].reksadana, dbPorto[*kunci].detailPorto[i].saham, dbPorto[*kunci].detailPorto[i].tanda, dbPorto[*kunci].detailPorto[i].tipe)
		}
	}

	if dbPorto[*kunci].detailPorto[pilihan-1].obligasi != 0 || dbPorto[*kunci].detailPorto[pilihan-1].reksadana != 0 || dbPorto[*kunci].detailPorto[pilihan-1].saham != 0{
		dbPorto[*kunci].detailPorto[pilihan-1].obligasi = 0
		dbPorto[*kunci].detailPorto[pilihan-1].reksadana = 0
		dbPorto[*kunci].detailPorto[pilihan-1].saham = 0
	}

	hitungDetail(&saham, &reksadana, &obligasi, kunci, pilihan)

	l=1
	for k:=0; k<NMAX; k++{
		l++
		if l == NMAX || dbPorto[*kunci].detailPorto[l].tanda == 0 {
			break
		}

		dbPorto[*kunci].detailPorto[k].obligasi = dbPorto[*kunci].detailPorto[l].obligasi
		dbPorto[*kunci].detailPorto[k].reksadana = dbPorto[*kunci].detailPorto[l].reksadana
		dbPorto[*kunci].detailPorto[k].saham = dbPorto[*kunci].detailPorto[l].saham
		dbPorto[*kunci].detailPorto[k].tanda = dbPorto[*kunci].detailPorto[l].tanda
		dbPorto[*kunci].detailPorto[k].tipe = dbPorto[*kunci].detailPorto[l].tipe			
	}


	fmt.Println()
	fmt.Println()
	fmt.Println()

	for i:=0; i<NMAX; i++{
		fmt.Println(dbPorto[*kunci].detailPorto[i].tanda !=0)
		if dbPorto[*kunci].detailPorto[i].tanda !=0 {
			fmt.Printf("%d. total = %d\nsubtotal obligasi: %d\nsubtotal reksadana: %d\nsubtotal saham: %d\ntanda: %d\ntipe: %s\n", i+1, dbPorto[*kunci].total, dbPorto[*kunci].detailPorto[i].obligasi, dbPorto[*kunci].detailPorto[i].reksadana, dbPorto[*kunci].detailPorto[i].saham, dbPorto[*kunci].detailPorto[i].tanda, dbPorto[*kunci].detailPorto[i].tipe)
		}
	}

	fmt.Println("Portofolio berhasil dihapus")
}

