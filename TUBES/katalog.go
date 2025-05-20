package main
import "fmt"

func katalog(pilihan int, kunci, tipereksadana *int){
	var (
        stat bool
    )

    for !stat {
        fmt.Println("============ KATALOG ============")
        fmt.Println("Pilih jenis aset:")
        fmt.Println("[1] Saham")
        fmt.Println("[2] Reksadana")
        fmt.Println("[3] Obligasi")
        fmt.Println("[0] Kembali")
        fmt.Println()

        fmt.Print("Pilihan: ")
        fmt.Scan(&pilihan)

        sumber := "KATALOG"
        if pilihan == 1{
            clearScreen()
            uiHeaderTable(sumber, pilihan, 0)
            uiHeaderKatalog(pilihan, 0)
            for j:=0;j<7;j++{
                if dbSaham[j].produksaham.returnaset >= 0.00 {
                    fmt.Printf("%-7s %-10d %-15s %-16d %-.2f %-7s\n", " ", j+1, dbSaham[j].produksaham.namaProduk, dbSaham[j].produksaham.hargaPerLembar, dbSaham[j].produksaham.returnaset, " ")
                }else {
                    fmt.Printf("%-7s %-10d %-15s %-14d %-.2f %-7s\n", " ", j+1, dbSaham[j].produksaham.namaProduk, dbSaham[j].produksaham.hargaPerLembar, dbSaham[j].produksaham.returnaset, " ")
                }
            }
            uiFooterTablePendek()
            fmt.Println()

            maubeli(*tipereksadana, pilihan, kunci)
        }else if pilihan == 2{
            clearScreen()
            fmt.Println("Tipe Reksadana : ")
            fmt.Println("[1] Pasar Uang")
            fmt.Println("[2] Saham")
            fmt.Println("[3] Obligasi")
            fmt.Print("Pilihan: ")
            fmt.Scan(tipereksadana)

            clearScreen()
            uiHeaderTable(sumber, pilihan, *tipereksadana)
            uiHeaderKatalog(pilihan, *tipereksadana)
            for j:=0;j<7;j++{ 
                fmt.Printf("  %-7d %-40s %-20d %-.2f %-10s %-20s %s\n", j+1, dbReksadana[*tipereksadana-1].produkreksa[j].Produk, dbReksadana[*tipereksadana-1].produkreksa[j].minimal, dbReksadana[*tipereksadana-1].produkreksa[j].returnreksa, " ", dbReksadana[*tipereksadana-1].produkreksa[j].penampung, dbReksadana[*tipereksadana-1].produkreksa[j].kustodian)
            }
            uiFooterTablePanjang()
            fmt.Println()

            maubeli(*tipereksadana, pilihan, kunci)
        }else if pilihan == 3{
            clearScreen()
            uiHeaderTable(sumber, pilihan, *tipereksadana)
            uiHeaderKatalog(pilihan, 0)
            for j:=0;j<7;j++{
                fmt.Printf(" %-7d %-13s %-15d %-.2f %-7s %-22s %d %s %-10s %-20s %-15s\n", j+1, dbObligasi[j].produkobligasi.namaProduk, dbObligasi[j].produkobligasi.hargaPerLembar, dbObligasi[j].produkobligasi.returnaset, " ", dbObligasi[j].jatuhtempo, dbObligasi[j].kupon, "bulan", " ", dbObligasi[j].nextkupon, dbObligasi[j].penerbit)
            }
            uiFooterTableSuperPanjang()
            fmt.Println()

            maubeli(*tipereksadana, pilihan, kunci)
        }else if pilihan == 0{
            clearScreen()
            stat = true
        }else{
            fmt.Println("Pilihan tidak valid")
            clearScreen()
        } 
    }
}

func maubeli(tipereksadana, pilihan int, kunci *int){
    var (
        chose int
        stat bool
    )

    fmt.Println("[1] Beli aset")
    fmt.Println("[2] Lihat rekomendasi aset")
    fmt.Println("[0] Kembali")

    for !stat{
        fmt.Print("Pilihan: ")
        fmt.Scan(&chose)
        fmt.Println()

        switch chose {
        case 1:
            if pilihan == 1{
                fmt.Println()
                introBeliSaham(kunci)
            }else if pilihan == 2{
                fmt.Println()
                introBeliReksa(tipereksadana, kunci)
            }else if pilihan == 3{
                fmt.Println()
                introBeliObligasi(kunci)
            }
            stat = true
        case 2:
            if pilihan == 1{
                clearScreen()
                rekomendasiSaham()
            }else if pilihan == 2{
                clearScreen()
                rekomendasiReksa(tipereksadana)
            }else if pilihan == 3{
                clearScreen()
                rekomendasiObli()
            }
            
            fmt.Println()
            fmt.Println("[1] Beli aset")
            fmt.Println("[0] Kembali")
        case 0:
            clearScreen()
            stat = true
        default:
            fmt.Println("Pilihan anda tidak valid")
            clearScreen()
        }
    }
}