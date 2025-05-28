package main

const NMAX = 10

func main() {
	var (
		datauser user
		i, pilihan, kunci int
		isNew bool
	)

	seedPorto()
	seedTransaksi()
	seedreksa()
	seedobligasi()
	seedsaham()
	isiUser()
	cekJumlahUser(&i)
	clearScreen()
	fiturLoginDaftar(datauser, i, pilihan, &kunci, &isNew)
}