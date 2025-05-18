package main
import "fmt"

func rekomendasiSaham(dbSaham *[NMAX]saham){
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

    fmt.Println("Rekomendasi Saham: ")
	for i:=0; i<NMAX; i++{
		fmt.Printf("%d. Nama Produk: %s || Harga Per Lembar: %d || Return: %.2f\n", i+1, dbSaham[i].produksaham.namaProduk, dbSaham[i].produksaham.hargaPerLembar, dbSaham[i].produksaham.returnaset)
	}
}

func rekomendasiReksa(dbReksadana *[3]reksadana, tipereksadana int){
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

    fmt.Println("Rekomendasi Reksadana: ")
	for i:=0; i<NMAX; i++{
		fmt.Printf("%d. Nama Produk: %s || Harga Per Lembar: %d || Return: %.2f\n", i+1, dbReksadana[tipereksadana-1].produkreksa[i].Produk, dbReksadana[tipereksadana-1].produkreksa[i].minimal, dbReksadana[tipereksadana-1].produkreksa[i].returnreksa)
	}
}

func rekomendasiObli(dbObligasi *[NMAX]obligasi){
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

    fmt.Println("Rekomendasi Obligasi: ")
	for i:=0; i<NMAX; i++{
		fmt.Printf("%d. Nama Produk: %s || Harga Per Lembar: %d || Return: %.2f\n", i+1, dbObligasi[i].produkobligasi.namaProduk, dbObligasi[i].produkobligasi.hargaPerLembar, dbObligasi[i].produkobligasi.returnaset)
	}
}

func katalog(dbReksadana *[3]reksadana, dbSaham *[NMAX]saham, dbObligasi *[NMAX]obligasi, pilihan int, kunci, tipereksadana *int){
	var (
        stat bool
    )

    for !stat {
        fmt.Println("======= Katalog ========")
        fmt.Println("Pilih jenis aset:")
        fmt.Println("1. Saham")
        fmt.Println("2. Reksadana")
        fmt.Println("3. Obligasi")
        fmt.Println("0. Kembali")
        fmt.Println()

        fmt.Print("Pilihan: ")
        fmt.Scan(&pilihan)

        if pilihan == 1{
            fmt.Println("\n============ Katalog Saham ============")
            fmt.Println()
            for j:=0;j<7;j++{
                fmt.Printf("%d. Nama Produk: %s \nHarga/lembar: %d \nReturn: %.2f\n", j+1, dbSaham[j].produksaham.namaProduk, dbSaham[j].produksaham.hargaPerLembar, dbSaham[j].produksaham.returnaset)
                fmt.Println()
            }
            fmt.Println("========================================")
            fmt.Println()

            maubeli(dbSaham, dbReksadana, dbObligasi, *tipereksadana, pilihan, kunci)
        }else if pilihan == 2{
            fmt.Println("Tipe Reksadana : ")
            fmt.Println("1. Pasar Uang")
            fmt.Println("2. Saham")
            fmt.Println("3. Obligasi")
            fmt.Print("Pilihan: ")
            fmt.Scan(tipereksadana)

            fmt.Println(*tipereksadana)

            fmt.Println("\n===== Katalog Reksadana",dbReksadana[*tipereksadana-1].jenis,"=====")
            fmt.Println()
            for j:=0;j<7;j++{ //batas j nya perlu diubah
                fmt.Printf("%d. Nama Produk: %s \nBank Kustodian: %s \nBank Penampung: %s\nMinimal Pembelian: %d \n", j+1, dbReksadana[*tipereksadana-1].produkreksa[j].Produk, dbReksadana[*tipereksadana-1].produkreksa[j].kustodian, dbReksadana[*tipereksadana-1].produkreksa[j].penampung, dbReksadana[*tipereksadana-1].produkreksa[j].minimal)
                fmt.Printf("Return: %.2f\n", dbReksadana[*tipereksadana-1].produkreksa[j].returnreksa)
                fmt.Println()
            }
            fmt.Println("========================================")
            fmt.Println()

            maubeli(dbSaham, dbReksadana, dbObligasi, *tipereksadana, pilihan, kunci)
        }else if pilihan == 3{
            fmt.Println("\n=========== Katalog Obligasi ===========")
            fmt.Println()
            for j:=0;j<7;j++{
                fmt.Printf("%d. Nama Produk: %s \nHarga/unit: %d \nReturn: %.2f\n", j+1, dbObligasi[j].produkobligasi.namaProduk, dbObligasi[j].produkobligasi.hargaPerLembar, dbObligasi[j].produkobligasi.returnaset)
                fmt.Printf("Jatuh Tempo: %s \nKupon Selanjutnya: %s \nPenerbit: %s \nPemberian Kupon: per %d bulan\n", dbObligasi[j].jatuhtempo, dbObligasi[j].nextkupon, dbObligasi[j].penerbit, dbObligasi[j].kupon)
                fmt.Println()
            }
            fmt.Println("=========================================")
            fmt.Println()

            maubeli(dbSaham, dbReksadana, dbObligasi, *tipereksadana, pilihan, kunci)
        }else if pilihan == 0{
            stat = true
        }else{
            fmt.Println("Pilihan tidak valid")
        } 
    }
}

func maubeli(dbSaham *[NMAX]saham, dbReksadana *[3]reksadana, dbObligasi *[NMAX]obligasi, tipereksadana, pilihan int, kunci *int){
    var (
        chose int
        stat bool
    )

    fmt.Println("1. Beli aset")
    fmt.Println("2. Lihat rekomendasi aset")
    fmt.Println("0. Kembali")

    for !stat{
        fmt.Print("Pilihan: ")
        fmt.Scan(&chose)
        fmt.Println()

        switch chose {
        case 1:
            if pilihan == 1{
                introBeliSaham(dbSaham, pilihan, kunci)
            }else if pilihan == 2{
                introBeliReksa(dbReksadana, pilihan, tipereksadana, kunci)
            }else if pilihan == 3{
                introBeliObligasi(dbObligasi, kunci, pilihan)
            }
            stat = true
        case 2:
            if pilihan == 1{
                rekomendasiSaham(dbSaham)
            }else if pilihan == 2{
                rekomendasiReksa(dbReksadana, tipereksadana)
            }else if pilihan == 3{
                rekomendasiObli(dbObligasi)
            }
            
            fmt.Println()
            fmt.Println("1. Beli aset")
            fmt.Println("0. Kembali")
        case 0:
            stat = true
        default:
            fmt.Println("Pilihan anda tidak valid")
        }
    }
}