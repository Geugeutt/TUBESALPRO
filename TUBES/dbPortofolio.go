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
			{nilaiAset: 100000,
			 nilaiReturn: 530000,
			 jumlah: 10, 
			 produk: "Majoris Pasar Uang Syariah Indonesia", 
			 tipe: "Reksadana Pasar Uang", 
			 portofolio: "Dana Darurat", 
			 untung: 5.30, 
			 tanggal: "10-04-2025", 
			 tanda: 1},

			{nilaiAset: 6400,
			 nilaiReturn: 10240,
			 jumlah: 100, 
			 produk: "HUMI", 
			 tipe: "Saham", 
			 portofolio: "Dana Darurat", 
			 untung: 1.60, 
			 tanggal: "15-04-2025", 
			 tanda: 1},

			{nilaiAset: 100000,
			 nilaiReturn: -627000,
			 jumlah: 1, 
			 produk: "BNP Paribas Infrastruktur Plus", 
			 tipe: "Reksadana Saham", 
			 portofolio: "Tabungan", 
			 untung: -6.27, 
			 tanggal: "17-04-2025", 
			 tanda: 1},

			{nilaiAset: 1000000,
			 nilaiReturn: 5990000,
			 jumlah: 10, 
			 produk: "BNI-AM ITB Harmoni", 
			 tipe: "Reksadana Obligasi", 
			 portofolio: "Tabungan", 
			 untung: 5.99, 
			 tanggal: "19-04-2025", 
			 tanda: 1},

			{nilaiAset: 1000000,
			 nilaiReturn: 5500000,
			 jumlah: 1, 
			 produk: "FR0086", 
			 tipe: "Obligasi", 
			 portofolio: "Tabungan", 
			 untung: 5.50, 
			 tanggal: "23-04-2025", 
			 tanda: 1},

			{nilaiAset: 6400,
			 nilaiReturn: 9984,
			 jumlah: 100, 
			 produk: "HUMI", 
			 tipe: "Saham", 
			 portofolio: "Dana Darurat", 
			 untung: 1.56, 
			 tanggal: "26-04-2025", 
			 tanda: 1},
		},
	}

	dbTransaksi[1] = transaksi{
		riwayat: [NMAX]detailPembelian{
			{nilaiAset: 642500,
			 nilaiReturn: 57825,
			 jumlah: 100, 
			 produk: "BREN", 
			 tipe: "Saham", 
			 portofolio: "Dana Pensiun", 
			 untung: 0.09, 
			 tanggal: "11-04-2025", 
			 tanda: 1},

			{nilaiAset: 1000000,
			 nilaiReturn: 5370000,
			 jumlah: 1, 
			 produk: "PBS036", 
			 tipe: "Obligasi", 
			 portofolio: "Dana Darurat", 
			 untung: 5.37, 
			 tanggal: "13-04-2025", 
			 tanda: 1},

			{nilaiAset: 1000000,
			 nilaiReturn: 1810000,
			 jumlah: 10, 
			 produk: "Grow Obligasi Optima Dinamis Kelas O", 
			 tipe: "Reksadana Obligasi", 
			 portofolio: "Dana Pensiun", 
			 untung: 1.81, 
			 tanggal: "14-04-2025", 
			 tanda: 1},

			{nilaiAset: 100000,
			 nilaiReturn: -46000,
			 jumlah: 1, 
			 produk: "Grow Saham Indonesia Plus Kelas O", 
			 tipe: "Reksadana Saham", 
			 portofolio: "Dana Darurat", 
			 untung: -0.46, 
			 tanggal: "20-04-2025", 
			 tanda: 1},

			{nilaiAset: 100000,
			 nilaiReturn: 486000, 
			 jumlah: 1,
			 produk: "Danamas Rupiah Plus", 
			 tipe: "Reksadana Pasar Uang", 
			 portofolio: "Dana Darurat", 
			 untung: 4.86, 
			 tanggal: "24-04-2025", 
			 tanda: 1},

			{nilaiAset: 897500,
			 nilaiReturn: 26925,
			 jumlah: 100, 
			 produk: "CUAN", 
			 tipe: "Saham", 
			 portofolio: "Dana Pensiun", 
			 untung: 0.03, 
			 tanggal: "26-04-2025", 
			 tanda: 1},
		},
	}

	dbTransaksi[2] = transaksi{
		riwayat: [NMAX]detailPembelian{
			{nilaiAset: 100000,
			 nilaiReturn: 530000,
			 jumlah: 10, 
			 produk: "Majoris Pasar Uang Syariah Indonesia", 
			 tipe: "Reksadana Pasar Uang", 
			 portofolio: "Liburan Keluarga", 
			 untung: 5.30, 
			 tanggal: "10-04-2025", 
			 tanda: 1},

			{nilaiAset: 6400,
			 nilaiReturn: 10240,
			 jumlah: 100, 
			 produk: "HUMI", 
			 tipe: "Saham", 
			 portofolio: "Dana Pendidikan", 
			 untung: 1.60, 
			 tanggal: "15-04-2025", 
			 tanda: 1},

			{nilaiAset: 100000,
			 nilaiReturn: -627000,
			 jumlah: 1, 
			 produk: "BNP Paribas Infrastruktur Plus", 
			 tipe: "Reksadana Saham", 
			 portofolio: "Dana Pendidikan", 
			 untung: -6.27, 
			 tanggal: "17-04-2025", 
			 tanda: 1},

			{nilaiAset: 1000000,
			 nilaiReturn: 5990000,
			 jumlah: 10, 
			 produk: "BNI-AM ITB Harmoni", 
			 tipe: "Reksadana Obligasi", 
			 portofolio: "Dana Pendidikan", 
			 untung: 5.99, 
			 tanggal: "19-04-2025", 
			 tanda: 1},	

			{nilaiAset: 1000000,
			 nilaiReturn: 5500000,
			 jumlah: 1, 
			 produk: "FR0086", 
			 tipe: "Obligasi", 
			 portofolio: "Dana Investasi", 
			 untung: 5.50, 
			 tanggal: "23-04-2025", 
			 tanda: 1},

			{nilaiAset: 6400,
			 nilaiReturn: 9984,
			 jumlah: 100, 
			 produk: "HUMI", 
			 tipe: "Saham", 
			 portofolio: "Dana Pendidikan", 
			 untung: 1.56, 
			 tanggal: "26-04-2025", 
			 tanda: 1},		
		},
	}
}

func seedPorto(){
	dbPorto[0] = porto{
		total : 13626024,
		detailPorto: [NMAX]subtotal{
			{tanda: 1,
			 tipe: "Tabungan",
			 saham: 0,
			 reksadana: 6463000,
			 obligasi: 6500000,},

			{tanda: 1,
			 tipe: "Dana Darurat",
			 saham: 33024,
			 reksadana: 630000,
			 obligasi: 0,},
		},
	}

	dbPorto[1] = porto{
		total: 11444750,
		detailPorto: [NMAX]subtotal{
			{tanda : 1,
			 tipe: "Dana Darurat",
			 saham: 0,
			 reksadana: 640000,
			 obligasi: 6370000,},

			{tanda: 1,
			 tipe: "Dana Pensiun",
			 saham: 1624750,
			 reksadana: 2810000,
			 obligasi: 0,},
		},
	}

	dbPorto[2] = porto{
		total: 13626024,
		detailPorto: [NMAX]subtotal{
			{tanda: 1,
			 tipe: "Dana Pendidikan",
			 saham: 33024,
			 reksadana: 6463000,
			 obligasi: 0,},

			{tanda: 1,
			 tipe: "Liburan Keluarga",
			 saham: 0,
			 reksadana: 630000,
			 obligasi: 0,},

			{tanda: 1,
			 tipe: "Dana Investasi",
			 saham: 0,
			 reksadana: 0,
			 obligasi: 6500000,},
		},
	}
}