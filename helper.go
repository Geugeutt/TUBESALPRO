package main
import "fmt"

func kosongkanTempforSort(kunci *int){
	for i:=0 ; i<NMAX ; i++{
		if dbTempforSort[*kunci].riwayat[i].nilaiAset != 0{
			dbTempforSort[*kunci].riwayat[i].jumlah = 0
			dbTempforSort[*kunci].riwayat[i].nilaiAset = 0
			dbTempforSort[*kunci].riwayat[i].nilaiReturn = 0
			dbTempforSort[*kunci].riwayat[i].portofolio = ""
			dbTempforSort[*kunci].riwayat[i].produk = ""
			dbTempforSort[*kunci].riwayat[i].tanda = 0
			dbTempforSort[*kunci].riwayat[i].tanggal = ""
			dbTempforSort[*kunci].riwayat[i].tipe = ""
			dbTempforSort[*kunci].riwayat[i].untung = 0.0
		}
	}
}

func pengisianArrayBantu(kunci *int, namaPorto string){
	var count int = -1

	for i:=0 ; i<NMAX ; i++{
		if dbTransaksi[*kunci].riwayat[i].nilaiAset != 0 && dbTransaksi[*kunci].riwayat[i].portofolio == namaPorto{
			count++
			switch count {
			case 0:
				dbTempforSort[*kunci].riwayat[count] = dbTransaksi[*kunci].riwayat[i]
			default:
				for k:=count-1; k>=0; k--{
					if dbTransaksi[*kunci].riwayat[i].produk == dbTempforSort[*kunci].riwayat[k].produk && dbTransaksi[*kunci].riwayat[i].tanda == dbTempforSort[*kunci].riwayat[k].tanda{
						dbTempforSort[*kunci].riwayat[k].jumlah += dbTransaksi[*kunci].riwayat[i].jumlah
						dbTempforSort[*kunci].riwayat[k].nilaiAset += dbTransaksi[*kunci].riwayat[i].nilaiAset
						break
					}else if k==0 && dbTransaksi[*kunci].riwayat[i].produk != dbTempforSort[*kunci].riwayat[k].produk || dbTransaksi[*kunci].riwayat[i].tanda != dbTempforSort[*kunci].riwayat[k].tanda{
						dbTempforSort[*kunci].riwayat[count] = dbTransaksi[*kunci].riwayat[i]
					}
				}	
			}
		}
	}

	untungRugi(&dbTempforSort, kunci)
}

func tampilArrayBantu(kunci, count *int, namaPorto string){
	for i:=0 ; i<NMAX ; i++{
		if dbTempforSort[*kunci].riwayat[i].nilaiAset != 0 && dbTempforSort[*kunci].riwayat[i].portofolio == namaPorto{
			*count++
			if dbTempforSort[*kunci].riwayat[i].nilaiReturn >= 0 {
				fmt.Printf("%-1s %-5d %-40s %-24s %-17d  %-.3f %-12s %-18d\n", " ", *count, dbTempforSort[*kunci].riwayat[i].produk, dbTempforSort[*kunci].riwayat[i].tipe, dbTempforSort[*kunci].riwayat[i].nilaiAset + dbTempforSort[*kunci].riwayat[i].nilaiReturn, dbTempforSort[*kunci].riwayat[i].untung, " ", dbTempforSort[*kunci].riwayat[i].nilaiReturn)
			}else{
				fmt.Printf("%-1s %-5d %-40s %-24s %-16d  %-.3f %-12s %-18d\n", " ", *count, dbTempforSort[*kunci].riwayat[i].produk, dbTempforSort[*kunci].riwayat[i].tipe, dbTempforSort[*kunci].riwayat[i].nilaiAset + dbTempforSort[*kunci].riwayat[i].nilaiReturn, dbTempforSort[*kunci].riwayat[i].untung, " ", dbTempforSort[*kunci].riwayat[i].nilaiReturn)
			}
		}
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
