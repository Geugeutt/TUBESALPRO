package main
import "fmt"

func untungRugi(kunci *int){
	for i:=0 ; i<NMAX ; i++{
		if dbTransaksi[*kunci].riwayat[i].tanda == 1{
			dbTransaksi[*kunci].riwayat[i].nilaiReturn = int(dbTransaksi[*kunci].riwayat[i].untung * float64(dbTransaksi[*kunci].riwayat[i].nilaiAset))
		}
	}
}

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
			nilaiAwal += dbTransaksi[*kunci].riwayat[i].nilaiAset
		}
		for i:=0 ; i<NMAX ; i++{
			imbal += dbTransaksi[*kunci].riwayat[i].nilaiReturn
		}

		gain = float64(imbal*100/nilaiAwal)

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

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}