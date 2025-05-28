package main

//database pengguna aplikasi
var dbUser[NMAX]user
//database produk-produk aset
var dbReksadana [3]reksadana
var dbSaham [NMAX]saham
var dbObligasi [NMAX]obligasi
//database portofolio
var dbPorto [NMAX]porto
//database transaksi
var dbTransaksi [NMAX]transaksi
//database sementara untuk pengolahan
var dbTempforSort[NMAX]transaksi