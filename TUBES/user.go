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

func fiturLoginDaftar(datauser user,i, pilihan int, kunci *int, isNew *bool){
	var stat bool

	for !stat{
		fmt.Println("========= DAFTAR/LOGIN =========")
		fmt.Println("[1] Daftar")
		fmt.Println("[2] Login")
		fmt.Print("[0] Keluar\n\n")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		if pilihan == 1{
			fmt.Println()
			clearScreen()
			fmt.Println("============ DAFTAR ============")
			signUp(datauser, pilihan, i, kunci, isNew)
			stat = true
		}else if pilihan == 2{
			fmt.Println()
			clearScreen()
			fmt.Println("============ LOGIN ============")
			login(datauser, i, pilihan, kunci, isNew)
			stat = true
		}else if pilihan == 0{
			clearScreen()
			fmt.Println("Terima kasih sudah menggunakan aplikasi ini.")
			stat = true
		}else{
			fmt.Println()
			fmt.Println("Pilihan tidak valid")
			clearScreen()
		}
	}
}

func login(datauser user, i, pilihan int, kunci *int, isNew *bool){
	*isNew = false

	fmt.Print("Masukkan username anda: ")
	fmt.Scan(&datauser.username)
	fmt.Print("Masukkan password anda: ")
	fmt.Scan(&datauser.password)
	for k:=0;k<4;k++{
		if datauser.username == dbUser[k].username && datauser.password == dbUser[k].password{
			*kunci = k
			clearScreen()
			menuUtama(datauser, pilihan, kunci, isNew)
			break
		}else if k==3 && datauser.username != dbUser[k].username && datauser.password != dbUser[k].password{
			fmt.Println("Kombinasi username dan password tidak ditemukan")
			fmt.Println("[1] Coba lagi\n[0] Keluar")
			fmt.Print("Pilihan: ")
			fmt.Scan(&pilihan)
			if pilihan == 1{
				clearScreen()
				fiturLoginDaftar(datauser,i, pilihan, kunci, isNew)
			}else if pilihan == 0{
				clearScreen()
				keluar()
			}
		}
	}
}

func signUp(datauser user, pilihan, i int, kunci *int, isNew *bool){
	var stat bool

	*isNew = true

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
				clearScreen()
				menuUtama(datauser, pilihan, kunci, isNew)
				stat = true
				break
			}else if k>=4 && datauser.username != dbUser[k].username && datauser.password != dbUser[k].password{
				fmt.Println("Daftar Gagal, Coba Lagi")
				clearScreen()
			}
		}
	}
}