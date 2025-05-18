package main
import "fmt"

func profil(datauser user, dbUser *[NMAX]user, kunci *int, pilihan int){
	var (
		anything string
		stat bool
	)
	fmt.Printf("Username: %s\nPassword: %s\n", dbUser[*kunci].username, dbUser[*kunci].password)

	fmt.Println("1. Ubah Password")
	fmt.Println("2. Riwayat Transaksi")
	fmt.Println("3. Switch Akun")
	fmt.Println("4. Hapus Akun")
	fmt.Println("0. Kembali")

	for !stat{
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			fmt.Print("Masukkan password baru: ")
			fmt.Scan(&anything)

			dbUser[*kunci].password = anything

			if dbUser[*kunci].password == anything{
				fmt.Println("Password berhasil diubah.")
				stat = true
			}else{
				fmt.Println("Gagal merubah password. Coba lagi")
			}
		case 2:
			fmt.Println("\n1. Pembelian")
			fmt.Println("2. Penjualan")
			fmt.Println("0. Kembali")
			fmt.Print("Pilihan: ")
			fmt.Scan(&pilihan)

			if pilihan == 1{
				tampilTransaksiPembelian(kunci, pilihan)
				stat = true
			}else if pilihan == 2{
				tampilTransaksiPenjualan(kunci, pilihan)
				stat = true
			}
		case 3:
			switchAkun(datauser, dbUser, kunci)
			stat = true
		case 4:
			cek := false
			for !cek{
				fmt.Println("Apakah Anda yakin akan menghapus akun ini?")
				fmt.Scan(&anything)

				if anything == "Ya" || anything == "ya"{
					// hapusAkun(kunci)
					cek = true
					stat = true
				}else if anything == "Tidak" || anything == "tidak"{
					cek = true
					stat = true
				}else{
					fmt.Println("Pilihan tidak valid")
				}
			}
		case 0:
			stat = true
		}
	}
}

func switchAkun (datauser user, dbUser *[NMAX]user, kunci *int){
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

	fmt.Printf("Anda sekarang login sebagai %s.\n", dbUser[*kunci].username)
}