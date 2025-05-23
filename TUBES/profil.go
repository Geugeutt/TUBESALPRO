package main
import "fmt"

func profil(datauser user, kunci, asetDijual, jumlahDijual *int, pilihan int){
	var (
		anything string
		stat bool
		jenisTransaksi, count int
	)


	for !stat{
		fmt.Println("=========== PROFILE ===========")
		fmt.Printf("Username: %s\nPassword: %s\n", dbUser[*kunci].username, dbUser[*kunci].password)

		fmt.Println("[1] Ubah Password")
		fmt.Println("[2] Riwayat Transaksi")
		fmt.Println("[3] Switch Akun")
		fmt.Println("[0] Kembali")

		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			clearScreen()
			fmt.Print("Masukkan password baru: ")
			fmt.Scan(&anything)

			dbUser[*kunci].password = anything

			if dbUser[*kunci].password == anything{
				fmt.Println("Password berhasil diubah.")
				stat = true
			}else{
				clearScreen()
				fmt.Println("Gagal merubah password. Coba lagi")
			}
		case 2:
			fmt.Println("============================")
			fmt.Println("\n[1] Pembelian")
			fmt.Println("[2] Penjualan")
			fmt.Println("[0] Kembali")
			fmt.Print("Pilihan: ")
			fmt.Scan(&pilihan)

			if pilihan == 1{
				clearScreen()
				jenisTransaksi = tampilTransaksiPembelian(kunci, &count)
				filterTransaksi(kunci, pilihan, jenisTransaksi)
			}else if pilihan == 2{
				clearScreen()
				jenisTransaksi = tampilTransaksiPenjualan(kunci, *jumlahDijual, *asetDijual)
				filterTransaksi(kunci, pilihan, jenisTransaksi)
			}
		case 3:
			clearScreen()
			switchAkun(datauser, kunci)
		case 0:
			clearScreen()
			stat = true
		}
	}
}

func switchAkun (datauser user, kunci *int){
	fmt.Print("Masukkan username: ")
	fmt.Scan(&datauser.username)
	fmt.Print("Masukkan password: ")
	fmt.Scan(&datauser.password)

	for i:=0 ; i<NMAX ; i++{
		if dbUser[i].username == datauser.username && dbUser[i].password == datauser.password{
			*kunci = i
			break
		}
	}

	clearScreen()
	fmt.Printf("Anda sekarang login sebagai %s.\n", dbUser[*kunci].username)
}