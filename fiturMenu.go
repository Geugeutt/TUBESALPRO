package main
import "fmt"

func dashboard(kunci *int, isNew bool){
	var (
		imbal, nilaiAwal int
		gain float64
	)

	if !isNew{
		fmt.Println()
		sumber := "DASHBOARD"
		uiHeaderTable(sumber, 0, 0)
		uiHeaderDashboard()

		for i:=0 ; i<NMAX ; i++{
			if dbTransaksi[*kunci].riwayat[i].tanda == 1{
				nilaiAwal += dbTransaksi[*kunci].riwayat[i].nilaiAset
				imbal += dbTransaksi[*kunci].riwayat[i].nilaiReturn
			}
		}

		gain = float64(imbal*100/nilaiAwal)

		for j:=0;j<NMAX;j++{
			if dbPorto[*kunci].detailPorto[j].tanda == 1{
				dbPorto[*kunci].total += dbPorto[*kunci].detailPorto[j].saham + dbPorto[*kunci].detailPorto[j].reksadana + dbPorto[*kunci].detailPorto[j].obligasi
			}else{
				break
			}
		}
		fmt.Printf("%-4s %-15d %-15d %-21d %-.2f\n", " ", dbPorto[*kunci].total, nilaiAwal, imbal, gain)
		uiFooterTablePendek()
	}else if isNew {
		fmt.Println()
		sumber := "DASHBOARD"
		uiHeaderTable(sumber, 1, 0)
		uiHeaderDashboard()
		fmt.Printf("%-19s Anda belum memiliki portofolio %-19s\n", " ", " ")
		uiFooterTablePendek()
	}
}

func menuUtama(datauser user, pilihan int, kunci *int, isNew *bool) {
	var (
		isKeluar bool
		tipereksadana, asetDijual, jumlahDijual int
		namaPortoLama string
	)

	for !isKeluar {
		clearScreen()
		dashboard(kunci, *isNew)
		
		fmt.Println("\n========== MENU UTAMA ==========")
		fmt.Println("[1] Katalog")
		fmt.Println("[2] Portofolio")
		fmt.Println("[3] Profil")
		fmt.Println("[4] Kembali")
		fmt.Println("[0] Keluar")
		fmt.Println()

		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)
		fmt.Println()

		switch pilihan {
		case 1:
			clearScreen()
			katalog(pilihan, kunci, &tipereksadana)
		case 2:
			clearScreen()
			tampilPortofolio(kunci, &tipereksadana, &asetDijual, &jumlahDijual, &namaPortoLama)
		case 3:
			clearScreen()
			profil(datauser, kunci, &asetDijual, &jumlahDijual, pilihan)
		case 4:
			isKeluar = true
		case 0:
			clearScreen()
			fmt.Println("Terima kasih sudah menggunakan aplikasi ini")
			isKeluar = true
		default:
			clearScreen()
			fmt.Println("Pilihan tidak valid")
		}
	}
}
