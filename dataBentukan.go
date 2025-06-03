package main

//tipe data bentukan pengguna
type user struct{
	username string
	password string
}

//tipe bentukan produk-produk aset
type detailProduk struct{
    namaProduk string
    hargaPerLembar int
	returnaset float64
}
type khususreksadana struct{
    Produk, kustodian, penampung string
    minimal int
    returnreksa float64
}
type reksadana struct {
    jenis string
    produkreksa [NMAX]khususreksadana
}
type saham struct {
    produksaham detailProduk 
}
type obligasi struct {
    produkobligasi detailProduk
    jatuhtempo, penerbit, nextkupon string
    kupon int
}

//tipe data bentukan transaksi
type detailPembelian struct{
	nilaiAset, tanda, jumlah, nilaiReturn int
	produk, portofolio, tanggal, tipe string
	untung float64
}
type transaksi struct{
	riwayat [NMAX]detailPembelian
}

//tipe data bentukan portofolio
type subtotal struct{
	tipe string
	tanda, saham, reksadana, obligasi int
}
type porto struct{
	total int
	detailPorto [NMAX]subtotal
}