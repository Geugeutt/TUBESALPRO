package main
import "fmt"

func sortPorto(kunci *int){
	var maxIdx int

	for i := 0; i < NMAX-1; i++ {
		maxIdx = i
		for j := i + 1; j < NMAX; j++ {
			if (dbTransaksi[*kunci].riwayat[j].nilaiAset + dbTransaksi[*kunci].riwayat[j].nilaiReturn) > (dbTransaksi[*kunci].riwayat[maxIdx].nilaiAset + dbTransaksi[*kunci].riwayat[maxIdx].nilaiReturn) {
				maxIdx = j
			}
		}

		dbTransaksi[*kunci].riwayat[i], dbTransaksi[*kunci].riwayat[maxIdx] = dbTransaksi[*kunci].riwayat[maxIdx], dbTransaksi[*kunci].riwayat[i]
	}
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
			dbSaham[j+1].produksaham.returnaset = dbSaham[j].produksaham.returnaset
			dbSaham[j+1].produksaham.namaProduk = dbSaham[j].produksaham.namaProduk
			dbSaham[j+1].produksaham.hargaPerLembar = dbSaham[j].produksaham.hargaPerLembar
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
			fmt.Printf("  %-7d %-40s %-15d %.2f\n", i+1, dbSaham[i].produksaham.namaProduk, dbSaham[i].produksaham.hargaPerLembar, dbSaham[i].produksaham.returnaset)
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
			dbReksadana[tipereksadana-1].produkreksa[j+1].returnreksa = dbReksadana[tipereksadana-1].produkreksa[j].returnreksa
			dbReksadana[tipereksadana-1].produkreksa[j+1].Produk = dbReksadana[tipereksadana-1].produkreksa[j].Produk
			dbReksadana[tipereksadana-1].produkreksa[j+1].minimal = dbReksadana[tipereksadana-1].produkreksa[j].minimal
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
			dbObligasi[j+1].produkobligasi.returnaset = dbObligasi[j].produkobligasi.returnaset
			dbObligasi[j+1].produkobligasi.namaProduk = dbObligasi[j].produkobligasi.namaProduk
			dbObligasi[j+1].produkobligasi.hargaPerLembar = dbObligasi[j].produkobligasi.hargaPerLembar
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