package main
import "fmt"

type user struct{
	username string
	password string
}

var dbUser[NMAX]user

func isiUser(){
	dbUser[0]=user{
		username:"Geugeut",
		password: "12345",
	}
	dbUser[1]=user{
		username: "Nabil",
		password: "34567",
	}
	dbUser[2]=user{
		username: "Humaira",
		password: "56789",
	}
}

func fiturLoginDaftar(datauser user,i, pilihan int, kunci *int){
	var stat bool

	for !stat{
		fmt.Println("======= Daftar/Login =======")
		fmt.Println("1. Daftar")
		fmt.Println("2. Login")
		fmt.Print("0. Keluar\n\n")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1{
			fmt.Println()
			fmt.Println("========= Daftar =========")
			signUp(datauser, pilihan, i, kunci)
			stat = true
		}else if pilihan == 2{
			fmt.Println()
			fmt.Println("========= Login =========")
			login(datauser, i, pilihan, kunci)
			stat = true
		}else if pilihan == 0{
			fmt.Println("Terima kasih sudah menggunakan aplikasi ini.")
			stat = true
		}else{
			fmt.Println()
			fmt.Println("Pilihan tidak valid")
		}
	}
}

func login(datauser user, i, pilihan int, kunci *int){
	fmt.Print("Masukkan username anda: ")
	fmt.Scan(&datauser.username)
	fmt.Print("Masukkan password anda: ")
	fmt.Scan(&datauser.password)
	for k:=0;k<3;k++{
		if datauser.username == dbUser[k].username && datauser.password == dbUser[k].password{
			*kunci = k
			menuUtama(datauser, pilihan, i, kunci)
			break
		}else if k==2 && datauser.username != dbUser[k].username && datauser.password != dbUser[k].password{
			fmt.Println("Kombinasi username dan password tidak ditemukan")
			fmt.Println("1. Coba lagi\n0. Kembali")
			fmt.Scan(&pilihan)
			if pilihan == 1{
				fiturLoginDaftar(datauser,i, pilihan, kunci)
			}else if pilihan == 0{
				keluar()
			}
		}
	}
}

func signUp(datauser user, pilihan, i int, kunci *int){
	var stat bool
	for !stat {
		fmt.Print("Buat username anda: ")
		fmt.Scan(&datauser.username)
		fmt.Print("[Password harus 5 karakter]\nBuat password anda: ")
		fmt.Scan(&datauser.password)

		if len(datauser.password)!=5{
			fmt.Print("\nPassword tidak memenuhi kriteria, masukkan password baru!\nPassword baru: ")
			fmt.Scan(&datauser.password)
		}

		i++
		dbUser[i].username = datauser.username
		dbUser[i].password = datauser.password

		for k:=0;k<4;k++{
			if datauser.username == dbUser[k].username && datauser.password == dbUser[k].password{
				*kunci = k
				fmt.Println()
				fmt.Println("Daftar Berhasil!")
				menuUtama(datauser, pilihan, i, kunci)
				stat = true
				break
			}else if k>=4 && datauser.username != dbUser[k].username && datauser.password != dbUser[k].password{
				fmt.Println("Daftar Gagal, Coba Lagi")
			}
		}
	}
}