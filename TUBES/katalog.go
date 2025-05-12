package main
import "fmt"

func introKatalog(datauser user, kunci, i, pilihan int){
    fmt.Println("======= Katalog ========")
    fmt.Println("Pilih jenis aset:")
    fmt.Println("1. Saham")
    fmt.Println("2. Reksadana")
    fmt.Println("3. Obligasi")
    fmt.Println("0. Kembali")
    fmt.Println()

    fmt.Print("Pilihan: ")
    fmt.Scan(&pilihan)
    katalog(datauser, dbReksadana[:], dbSaham[:], dbObligasi[:], pilihan, i, kunci)
}

func katalog(datauser user, dbReksadana[]reksadana, dbSaham[]saham, dbObligasi[]obligasi, pilihan, i, kunci int){
	var tipe int
	var jenis, count, chose int

    if pilihan == 1{
        seedsaham()
        fmt.Println("\n============ Katalog Saham ============")
        fmt.Println()
        for j:=0;j<7;j++{
            fmt.Printf("%d. Nama Produk: %s \nHarga/lembar: %d \nReturn: %.2f\n", j+1, dbSaham[j].produksaham.namaProduk, dbSaham[j].produksaham.hargaPerLembar, dbSaham[j].produksaham.returnaset)
            fmt.Println()
        }
        fmt.Println("========================================")
        fmt.Println()
    }else if pilihan == 2{
        seedreksa()
        fmt.Println("Tipe Reksadana : ")
        fmt.Println("1. Pasar Uang")
        fmt.Println("2. Saham")
        fmt.Println("3. Obligasi")
        fmt.Print("Pilihan: ")
        fmt.Scan(&tipe)

        fmt.Println("\n===== Katalog Reksadana",dbReksadana[tipe-1].jenis,"=====")
        fmt.Println()
        for j:=0;j<7;j++{ //batas j nya perlu diubah
            fmt.Printf("%d. Nama Produk: %s \nBank Kustodian: %s \nBank Penampung: %s\nMinimal Pembelian: %d \n", j+1, dbReksadana[tipe-1].produkreksa[j].Produk, dbReksadana[tipe-1].produkreksa[j].kustodian, dbReksadana[tipe-1].produkreksa[j].penampung, dbReksadana[tipe-1].produkreksa[j].minimal)
            fmt.Printf("Return: %.2f\n", dbReksadana[tipe-1].produkreksa[j].returnreksa)
            fmt.Println()
        }
        fmt.Println("========================================")
        fmt.Println()
    }else if pilihan == 3{
        seedobligasi()
        fmt.Println("\n=========== Katalog Obligasi ===========")
        fmt.Println()
        for j:=0;j<7;j++{
            fmt.Printf("%d. Nama Produk: %s \nHarga/unit: %d \nReturn: %.2f\n", j+1, dbObligasi[j].produkobligasi.namaProduk, dbObligasi[j].produkobligasi.hargaPerLembar, dbObligasi[j].produkobligasi.returnaset)
            fmt.Printf("Jatuh Tempo: %s \nKupon Selanjutnya: %s \nPenerbit: %s \nPemberian Kupon: per %d bulan\n", dbObligasi[j].jatuhtempo, dbObligasi[j].nextkupon, dbObligasi[j].penerbit, dbObligasi[j].kupon)
            fmt.Println()
        }
        fmt.Println("=========================================")
        fmt.Println()
    }else if pilihan == 0{
        menuUtama(datauser, i, pilihan, kunci)
    }else{
        fmt.Println("Pilihan tidak valid")
        introKatalog(datauser, kunci, i, pilihan)
    } 

    fmt.Println("1. Beli aset")
    fmt.Println("0. Kembali")

    for count == 0{
        fmt.Print("Pilihan: ")
        fmt.Scan(&chose)
        fmt.Println()

        if chose == 1{
            if pilihan == 1{
                introBeliSaham(dbSaham, pilihan, kunci)
                count++
            }else if pilihan == 2{
                introBeliReksa(dbReksadana, pilihan, jenis, kunci)
                count++
            }else if pilihan == 3{
                introBeliObligasi(dbObligasi, kunci, pilihan)
                count++
            }
        }else if chose == 0{
            introKatalog(datauser, kunci, i, pilihan)
            count++
        }else{
            fmt.Println("Pilihan anda tidak valid")
        }
    }
}