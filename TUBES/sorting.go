package main
import "fmt"

var dbTempforSort[NMAX]transaksi

func sortPortobyAset(kunci *int){
	var maxIdx int

	for i := 0; i < NMAX-1; i++ {
		maxIdx = i
		for j := i + 1; j < NMAX; j++ {
			if (dbTempforSort[*kunci].riwayat[j].nilaiAset + dbTempforSort[*kunci].riwayat[j].nilaiReturn) > (dbTempforSort[*kunci].riwayat[maxIdx].nilaiAset + dbTempforSort[*kunci].riwayat[maxIdx].nilaiReturn) {
				maxIdx = j
			}else if dbTempforSort[*kunci].riwayat[j].nilaiAset == 0 {
				break
			}
		}

		dbTempforSort[*kunci].riwayat[i], dbTempforSort[*kunci].riwayat[maxIdx] = dbTempforSort[*kunci].riwayat[maxIdx], dbTempforSort[*kunci].riwayat[i]
	}
}

func sortforBinary(kunci *int){
		var maxIdx int

	for i := 0; i < NMAX-1; i++ {
		maxIdx = i
		for j := i + 1; j < NMAX; j++ {
			if dbTempforSort[*kunci].riwayat[j].produk < dbTempforSort[*kunci].riwayat[maxIdx].produk && dbTempforSort[*kunci].riwayat[j].produk != ""{
				maxIdx = j
			}
		}

		dbTempforSort[*kunci].riwayat[i], dbTempforSort[*kunci].riwayat[maxIdx] = dbTempforSort[*kunci].riwayat[maxIdx], dbTempforSort[*kunci].riwayat[i]
	}
}

func afterSort(sumber, namaPorto string, kunci *int, totalLokal int, saham, reksadana, obligasi float64){
	var count int

	clearScreen()
		uiHeaderTable(sumber, 0, 0)
		uiHeaderPorto(dbUser[*kunci].username, namaPorto, totalLokal, saham, reksadana, obligasi)
		for i:=0 ; i<NMAX ; i++{
			if dbTempforSort[*kunci].riwayat[i].nilaiAset != 0{
				count++
				if dbTransaksi[*kunci].riwayat[i].nilaiReturn >= 0 {
					fmt.Printf("%-1s %-5d %-40s %-24s %-17d  %-.2f %-12s %-18d %-21s\n", " ", count, dbTempforSort[*kunci].riwayat[i].produk, dbTempforSort[*kunci].riwayat[i].tipe, dbTempforSort[*kunci].riwayat[i].nilaiAset + dbTempforSort[*kunci].riwayat[i].nilaiReturn, dbTempforSort[*kunci].riwayat[i].untung, " ", dbTempforSort[*kunci].riwayat[i].nilaiReturn, dbTempforSort[*kunci].riwayat[i].tanggal)
				}else{
					fmt.Printf("%-1s %-5d %-40s %-24s %-16d  %-.2f %-12s %-18d %-23s\n", " ", count, dbTempforSort[*kunci].riwayat[i].produk, dbTempforSort[*kunci].riwayat[i].tipe, dbTempforSort[*kunci].riwayat[i].nilaiAset + dbTempforSort[*kunci].riwayat[i].nilaiReturn, dbTempforSort[*kunci].riwayat[i].untung, " ", dbTempforSort[*kunci].riwayat[i].nilaiReturn, dbTempforSort[*kunci].riwayat[i].tanggal)
				}
			}
		}
		uiFooterTablePanjang()
}

func rekomendasiSaham(){
	var (
	    key1 float64
	    key2 string
	    key3 int
	)

	for i:=1; i<NMAX; i++{
		key1 = dbSaham[i].produksaham.returnaset
		key2 = dbSaham[i].produksaham.namaProduk
		key3 = dbSaham[i].produksaham.hargaPerLembar
		j := i-1
		for j>=0 && dbSaham[j].produksaham.returnaset < key1{
			dbSaham[j+1].produksaham = dbSaham[j].produksaham
			j--
		}
		dbSaham[j+1].produksaham.returnaset = key1
		dbSaham[j+1].produksaham.namaProduk = key2
		dbSaham[i].produksaham.hargaPerLembar = key3
	}

	sumber := "REKOMENDASI"

    uiHeaderTable(sumber, 1, 0)
	uiHeaderRekomendasi()
	for i:=0; i<NMAX; i++{
		if dbSaham[i].produksaham.hargaPerLembar > 0{
			fmt.Printf("  %-7d %-40s %-20d %.2f\n", i+1, dbSaham[i].produksaham.namaProduk, dbSaham[i].produksaham.hargaPerLembar, dbSaham[i].produksaham.returnaset)
		}
	}
	uiFooterTablePanjang()
}

func rekomendasiReksa(tipereksadana int){
    var (
        key1 float64
        key2 string
        key3 int
    )

    	for i:=1; i<NMAX; i++{
		key1 = dbReksadana[tipereksadana-1].produkreksa[i].returnreksa
		key2 = dbReksadana[tipereksadana-1].produkreksa[i].Produk
		key3 = dbReksadana[tipereksadana-1].produkreksa[i].minimal
		j := i-1
		for j>=0 && dbSaham[j].produksaham.returnaset < key1{
			dbReksadana[tipereksadana-1].produkreksa[j+1] = dbReksadana[tipereksadana-1].produkreksa[j]
			j--
		}
		dbReksadana[tipereksadana-1].produkreksa[j+1].returnreksa = key1
		dbReksadana[tipereksadana-1].produkreksa[j+1].Produk = key2
		dbReksadana[tipereksadana-1].produkreksa[j+1].minimal = key3
	}

	sumber := "REKOMENDASI"

    uiHeaderTable(sumber, 2, tipereksadana)
	uiHeaderRekomendasi()
	for i:=0; i<NMAX; i++{
		if dbReksadana[tipereksadana-1].produkreksa[i].minimal > 0{
			fmt.Printf("  %-7d %-40s %-15d %.2f\n", i+1, dbReksadana[tipereksadana-1].produkreksa[i].Produk, dbReksadana[tipereksadana-1].produkreksa[i].minimal, dbReksadana[tipereksadana-1].produkreksa[i].returnreksa)
		}
	}
	uiFooterTablePanjang()
}

func rekomendasiObli(){
	var (
	    key1 float64
	    key2 string
	    key3 int
	)

	for i:=1; i<NMAX; i++{
		key1 = dbObligasi[i].produkobligasi.returnaset
		key2 = dbObligasi[i].produkobligasi.namaProduk
		key3 = dbObligasi[i].produkobligasi.hargaPerLembar
		j := i-1
		for j>=0 && dbSaham[j].produksaham.returnaset < key1{
			dbObligasi[j+1].produkobligasi = dbObligasi[j].produkobligasi
			j--
		}
		dbObligasi[j+1].produkobligasi.returnaset = key1
		dbObligasi[j+1].produkobligasi.namaProduk = key2
		dbObligasi[j+1].produkobligasi.hargaPerLembar = key3
	}

	sumber := "REKOMENDASI"

    uiHeaderTable(sumber, 3, 0)
	uiHeaderRekomendasi()
	for i:=0; i<NMAX; i++{
		if dbObligasi[i].produkobligasi.hargaPerLembar > 0{
			fmt.Printf("  %-7d %-40s %-15d %.2f\n", i+1, dbObligasi[i].produkobligasi.namaProduk, dbObligasi[i].produkobligasi.hargaPerLembar, dbObligasi[i].produkobligasi.returnaset)
		}
	}
	uiFooterTablePanjang()
}