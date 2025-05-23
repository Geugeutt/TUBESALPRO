package main

type detailPembelian struct{
	nilaiAset, tanda, jumlah, nilaiReturn int
	produk, portofolio, tanggal, tipe string
	untung float64
}

type transaksi struct{
	riwayat [NMAX]detailPembelian
}

type subtotal struct{
	tipe string
	tanda, saham, reksadana, obligasi int
}


type porto struct{
	total int
	detailPorto [NMAX]subtotal
}

var dbPorto [NMAX]porto
var dbTransaksi [NMAX]transaksi

func seedTransaksi(){
    dbTransaksi[0] = transaksi{
        riwayat: [NMAX]detailPembelian{
            {nilaiAset: 27400000,
             nilaiReturn: 1284860,
             jumlah: 10000,
             produk: "ANTM",
             tipe: "Saham",
             portofolio: "Tabungan",
             untung: 4.69,
             tanggal: "01-03-2025",
             tanda: 1},

            {nilaiAset: 30800000,
             nilaiReturn: 169400,
             jumlah: 10000,
             produk: "PTRO",
             tipe: "Saham",
             portofolio: "Tabungan",
             untung: 0.55,
             tanggal: "05-03-2025",
             tanda: 1},

            {nilaiAset: 1000000,
             nilaiReturn: 56200,
             jumlah: 100,
             produk: "BRI Seruni Pasar Uang III",
             tipe: "Reksadana PasarUang",
             portofolio: "Dana Darurat",
             untung: 5.62,
             tanggal: "10-03-2025",
             tanda: 1},

            {nilaiAset: 1000000,
             nilaiReturn: 65000,
             jumlah: 1,
             produk: "FR0081",
             tipe: "Obligasi",
             portofolio: "Tabungan",
             untung: 6.50,
             tanggal: "15-03-2025",
             tanda: 1},

            {nilaiAset: 1760000,
             nilaiReturn: 291632,
             jumlah: 10000,
             produk: "DMAS",
             tipe: "Saham",
             portofolio: "Dana Darurat",
             untung: 16.57,
             tanggal: "20-03-2025",
             tanda: 1},

            {nilaiAset: 5000000,
             nilaiReturn: 227500,
             jumlah: 100,
             produk: "Principal Cash Fund",
             tipe: "Reksadana PasarUang",
             portofolio: "Tabungan",
             untung: 4.55,
             tanggal: "25-03-2025",
             tanda: 1},

            {nilaiAset: 1000000,
             nilaiReturn: 53700,
             jumlah: 1,
             produk: "PBS036",
             tipe: "Obligasi",
             portofolio: "Tabungan",
             untung: 5.37,
             tanggal: "28-03-2025",
             tanda: 1},
        },
    }

    dbTransaksi[1] = transaksi{
        riwayat: [NMAX]detailPembelian{
            {nilaiAset: 64250000,
             nilaiReturn: 57825,
             jumlah: 10000,
             produk: "BREN",
             tipe: "Saham",
             portofolio: "Dana Pensiun",
             untung: 0.09,
             tanggal: "11-04-2025",
             tanda: 1},

            {nilaiAset: 10000000,
             nilaiReturn: -46000,
             jumlah: 100,
             produk: "Grow Saham Indonesia Plus Kelas O",
             tipe: "Reksadana Saham",
             portofolio: "Dana Pensiun",
             untung: -0.46,
             tanggal: "15-04-2025",
             tanda: 1},

            {nilaiAset: 1000000,
             nilaiReturn: 55000,
             jumlah: 1,
             produk: "FR0086",
             tipe: "Obligasi",
             portofolio: "Dana Pensiun",
             untung: 5.50,
             tanggal: "18-04-2025",
             tanda: 1},

            {nilaiAset: 89750000,
             nilaiReturn: 26925,
             jumlah: 10000,
             produk: "CUAN",
             tipe: "Saham",
             portofolio: "Dana Darurat",
             untung: 0.03,
             tanggal: "20-04-2025",
             tanda: 1},

            {nilaiAset: 1000000,
             nilaiReturn: -75100,
             jumlah: 100,
             produk: "Avrist IDX30",
             tipe: "Reksadana Saham",
             portofolio: "Dana Darurat",
             untung: -7.51,
             tanggal: "22-04-2025",
             tanda: 1},

            {nilaiAset: 1000000,
             nilaiReturn: 48700,
             jumlah: 1,
             produk: "PBS032",
             tipe: "Obligasi",
             portofolio: "Dana Pensiun",
             untung: 4.87,
             tanggal: "25-04-2025",
             tanda: 1},
        },
    }

    dbTransaksi[2] = transaksi{
        riwayat: [NMAX]detailPembelian{
            {nilaiAset: 640000,
             nilaiReturn: 9984,
             jumlah: 10000,
             produk: "HUMI",
             tipe: "Saham",
             portofolio: "Liburan Keluarga",
             untung: 1.56,
             tanggal: "01-05-2025",
             tanda: 1},

            {nilaiAset: 1000000,
             nilaiReturn: 53000,
             jumlah: 100,
             produk: "Majoris Pasar Uang Syariah Indonesia",
             tipe: "Reksadana PasarUang",
             portofolio: "Dana Pendidikan",
             untung: 5.30,
             tanggal: "05-05-2025",
             tanda: 1},

            {nilaiAset: 1000000,
             nilaiReturn: 60000,
             jumlah: 1,
             produk: "PBS003",
             tipe: "Obligasi",
             portofolio: "Data Investasi",
             untung: 6.00,
             tanggal: "08-05-2025",
             tanda: 1},

            {nilaiAset: 1750000,
             nilaiReturn: 288400,
             jumlah: 10000,
             produk: "TASA",
             tipe: "Saham",
             portofolio: "Liburan Keluarga",
             untung: 16.48,
             tanggal: "10-05-2025",
             tanda: 1},

            {nilaiAset: 10000000,
             nilaiReturn: 639000,
             jumlah: 100,
             produk: "ABF Indonesia Bond Index Fund",
             tipe: "Reksadana Obligasi",
             portofolio: "Dana Pendidikan",
             untung: 6.39,
             tanggal: "12-05-2025",
             tanda: 1},

            {nilaiAset: 1000000,
             nilaiReturn: 51200,
             jumlah: 1,
             produk: "FR0090",
             tipe: "Obligasi",
             portofolio: "Data Investasi",
             untung: 5.12,
             tanggal: "15-05-2025",
             tanda: 1},

            {nilaiAset: 1000000,
             nilaiReturn: 54700,
             jumlah: 100,
             produk: "Danamas Stabil",
             tipe: "Reksadana Obligasi",
             portofolio: "Liburan Keluarga",
             untung: 5.47,
             tanggal: "18-05-2025",
             tanda: 1},

            {nilaiAset: 1000000,
             nilaiReturn: -53300,
             jumlah: 10,
             produk: "Eastspring IDX ESG Leaders Plus Kelas A",
             tipe: "Reksadana Saham",
             portofolio: "Dana Pendidikan",
             untung: -5.33,
             tanggal: "20-05-2025",
             tanda: 1},
        },
    }
}

func seedPorto(){
    dbPorto[0] = porto{
        total: 70108292,
        detailPorto: [NMAX]subtotal{
            {tanda: 1,
             tipe: "Tabungan",
             saham: 59654260,
             reksadana: 5227500,
             obligasi: 2118700,
            },
            {tanda: 1,
             tipe: "Dana Darurat",
             saham: 2051632,
             reksadana: 1056200,
             obligasi: 0,
            },
        },
    }

    dbPorto[1] = porto{
        total: 167067350,
        detailPorto: [NMAX]subtotal{
            {tanda: 1,
             tipe: "Dana Darurat",
             saham: 89776925,
             reksadana: 924900,
             obligasi: 0,
            },
            {tanda: 1,
             tipe: "Dana Pensiun",
             saham: 64307825,
             reksadana: 9954000,
             obligasi: 2103700,
            },
        },
    }

    dbPorto[2] = porto{
        total: 18492984,
        detailPorto: [NMAX]subtotal{
            {tanda: 1,
             tipe: "Dana Pendidikan",
             saham: 0,
             reksadana: 12638700,
             obligasi: 0,
            },
            {tanda: 1,
             tipe: "Liburan Keluarga",
             saham: 2688384,
             reksadana: 0,
             obligasi: 1054700,
            },
            {tanda: 1,
             tipe: "Data Investasi",
             saham: 0,
             reksadana: 0,
             obligasi: 2111200,
            },
        },
    }
}