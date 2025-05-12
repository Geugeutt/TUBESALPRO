package main
import "fmt"

type subtotal struct{
	tipe string
	saham, reksadana, obligasi int
	tanda int
}

type porto struct{
	total int
	detailporto [NMAX]subtotal
}

var dbPorto [NMAX]porto

func seedPorto(){
	dbPorto[0] = porto{
		total : 0,
		detailporto: [NMAX]subtotal{
			{tanda: 1,
			 tipe: "Tabungan",
			 saham: 300000,
			 reksadana: 400000,
			 obligasi: 1000000,
			},
			{tanda: 1,
			 tipe: "Dana Darurat",
			 saham: 2000000,
			 reksadana: 3500000,
			 obligasi: 1500000,
			},
		},
	}

	dbPorto[1] = porto{
		total: 0,
		detailporto: [NMAX]subtotal{
			{tanda : 1,
			 tipe: "Dana Darurat",
			 saham: 2000000,
			 reksadana: 6000000,
			 obligasi: 7000000,
			},
			{tanda: 1,
			 tipe: "Dana Pensiun",
			 saham: 800000,
			 reksadana: 900000,
			 obligasi: 1500000,
			},
		},
	}
}

// bisa tampilin jumlah nilai aset semuanya, nilai aset tiap jenis invest, presentase persebaran aset ke tiap jenis invest
func tampilPortofolio(datauser user, dbUser[]user, dbPorto*[NMAX]porto, i, kunci int){
	var (
		saham, reksadana, obligasi float64
		pilihan int
		stat bool
	)

	pilihporto(dbPorto, kunci, &pilihan)
	hitungDetail(&saham, &reksadana, &obligasi, kunci, pilihan)

	fmt.Printf("======= Portofolio %s ======\n", dbUser[kunci].username)
	fmt.Printf("Jenis Portofolio: %s", dbPorto[kunci].detailporto[pilihan-1].tipe)
	fmt.Printf("\nTotal Aset: %d", dbPorto[kunci].total)
	fmt.Println("\n======= Detail Portofolio ======")
	fmt.Printf("Saham: %d\nReksadana: %d\nObligasi: %d\n", dbPorto[kunci].detailporto[pilihan-1].saham, dbPorto[kunci].detailporto[pilihan-1].reksadana, dbPorto[kunci].detailporto[pilihan-1].obligasi)
	fmt.Println("======= Persentase Portofolio ======")
	fmt.Printf("Saham: %.2f persen\nReksadana: %.2f persen\nObligasi: %.2f persen\n", saham, reksadana, obligasi)

	fmt.Println("1. Beli aset")
	fmt.Println("2. Jual aset")
	fmt.Println("0. Kembali")
	fmt.Println()

	for !stat{
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan{
		case 1:
			katalog(datauser, dbReksadana[:], dbSaham[:], dbObligasi[:], pilihan, i, kunci)
			stat = true
		case 2:
			//panggil prosedur delporto
			stat = true
		case 0:
			stat = true
		default:
			fmt.Println("Pilihan tidak valid")
		}
	}

	//masih salah pas milih katalog, dia langsung nampilin saham (salah di nilai variabel pilihan)
}

func hitungDetail(saham, reksadana, obligasi *float64, kunci, pilihan int){
	dbPorto[kunci].total = dbPorto[kunci].detailporto[pilihan-1].saham + dbPorto[kunci].detailporto[pilihan-1].reksadana + dbPorto[kunci].detailporto[pilihan-1].obligasi

	if dbPorto[kunci].detailporto[pilihan-1].saham != 0{
		*saham = float64(dbPorto[kunci].detailporto[pilihan-1].saham)*100.0/float64(dbPorto[kunci].total)
	}else{
		*saham = 0
	}

	if dbPorto[kunci].detailporto[pilihan-1].reksadana != 0{
		*reksadana = float64(dbPorto[kunci].detailporto[pilihan-1].reksadana)*100.0/float64(dbPorto[kunci].total)
	}else{
		*reksadana = 0
	}

	if dbPorto[kunci].detailporto[pilihan-1].reksadana != 0{
		*obligasi = float64(dbPorto[kunci].detailporto[pilihan-1].obligasi)*100.0/float64(dbPorto[kunci].total)
	}else{
		*obligasi = 0
	}
}

func pilihporto(dbporto*[NMAX]porto, kunci int, pilihan *int){

	fmt.Println("\nPilih portofolio Anda: ")
	for k:=0;k<NMAX;k++{
		if dbPorto[kunci].detailporto[k].tanda == 1{
			fmt.Printf("%d. %s\n", k+1, dbporto[kunci].detailporto[k].tipe)
		}else{
			break
		}
	}

	fmt.Print("Pilihan: ")
	fmt.Scan(pilihan)
}

func addPorto(dbporto *[NMAX]porto, kunci, tambahanaset, pilihan, tanda int){
	pilihporto(dbporto, kunci, &pilihan)

	fmt.Println("Tambahan aset: ", tambahanaset)
	if tanda == 1{
		dbPorto[kunci].detailporto[pilihan-1].saham += tambahanaset
	}else if tanda == 2{
		dbPorto[kunci].detailporto[pilihan-1].reksadana += tambahanaset
	}else if tanda == 3{
		dbPorto[kunci].detailporto[pilihan-1].obligasi += tambahanaset
	}

	fmt.Println("\nAset berhasil ditambahkan ke portofolio Anda.")
}

// func delPorto(){
// 	//hilangin nilai di porto
// }

