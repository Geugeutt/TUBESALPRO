package main

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

var dbReksadana [3]reksadana
var dbSaham [NMAX]saham
var dbObligasi [NMAX]obligasi

func seedsaham(){
    dbSaham[0] = saham{
        produksaham: detailProduk{
        namaProduk: "ANTM", 
        hargaPerLembar: 2740, 
        returnaset: 4.69,
        },
    }

    dbSaham[1] = saham{
        produksaham: detailProduk{
        namaProduk: "PTRO", 
        hargaPerLembar: 3080, 
        returnaset: 0.55,
        },
    }

    dbSaham[2] = saham{
        produksaham: detailProduk{
        namaProduk: "HUMI", 
        hargaPerLembar: 64, 
        returnaset: 1.56,
        },
    }

    dbSaham[3] = saham{
        produksaham: detailProduk{
        namaProduk: "BREN", 
        hargaPerLembar: 6425, 
        returnaset: 0.09,
        },
    }

    dbSaham[4] = saham{
        produksaham: detailProduk{
        namaProduk: "DMAS", 
        hargaPerLembar: 176, 
        returnaset: 16.57,
        },
    }

    dbSaham[5] = saham{
        produksaham: detailProduk{
        namaProduk: "CUAN", 
        hargaPerLembar: 8975, 
        returnaset: 0.03,
        },
    }

    dbSaham[6] = saham{
        produksaham: detailProduk{
        namaProduk: "TASA", 
        hargaPerLembar: 175, 
        returnaset: 16.48,
        },
    }
}

func seedreksa(){
    dbReksadana[0] = reksadana{
        jenis: "PasarUang",
        produkreksa: [NMAX]khususreksadana{
            {Produk: "BRI Seruni Pasar Uang III",
             kustodian: "CIMB NIAGA", 
             penampung: "BCA", 
             minimal: 10000, 
             returnreksa: 5.62},

            {Produk: "Danamas Rupiah Plus",
             kustodian: "CIMB NIAGA", 
             penampung: "BCA", 
             minimal: 100000, 
             returnreksa: 4.86},

            {Produk: "Bahana Dana Likuid Kelas G",
             kustodian: "STD CHARTERED", 
             penampung: "JAGO", 
             minimal: 10000, 
             returnreksa: 4.72},

            {Produk: "Mandiri Investa Pasar Uang Kelas A",
             kustodian: "CITIBANK N.A. - CUSTODY", 
             penampung: "BCA", 
             minimal: 10000, 
             returnreksa: 4.49},

            {Produk: "Principal Cash Fund",
             kustodian: "DEUTSCHE BANK", 
             penampung: "BCA", 
             minimal: 50000, 
             returnreksa: 4.55},

            {Produk: "TRIM Kas 2 Kelas A",
             kustodian: "DBS INDONESIA", 
             penampung: "JAGO", 
             minimal: 10000, 
             returnreksa: 5.48},

            {Produk: "Majoris Pasar Uang Syariah Indonesia",
             kustodian: "BNI", 
             penampung: "JAGO", 
             minimal: 10000, 
             returnreksa: 5.30},
        },
    }
    
    dbReksadana[1] = reksadana{
        jenis: "Saham",
        produkreksa: [NMAX]khususreksadana{
            {Produk: "Sucorinvest Maxi Fund",
             kustodian: "HSBC INDONESIA", 
             penampung: "BCA", 
             minimal: 100000, 
             returnreksa: 5.34},

            {Produk: "Grow Saham Indonesia Plus Kelas O",
             kustodian: "HSBC INDONESIA", 
             penampung: "BCA", 
             minimal: 100000, 
             returnreksa: -0.46},

            {Produk: "Grow SRI KEHATI Kelas O",
             kustodian: "HSBC INDONESIA", 
             penampung: "BCA", 
             minimal: 10000, 
             returnreksa: -4.60},

            {Produk: "BNP Paribas Infrastruktur Plus",
             kustodian: "CITIBANK", 
             penampung: "BCA", 
             minimal: 100000, 
             returnreksa: -6.27},

            {Produk: "TRAM Comsumption Plus Kelas A",
             kustodian: "HSBC INDONESIA", 
             penampung: "BCA", 
             minimal: 100000, 
             returnreksa: -7.00},

            {Produk: "Avrist IDX30",
             kustodian: "STD CHARTERED", 
             penampung: "BCA", 
             minimal: 10000, 
             returnreksa: -7.51},

            {Produk: "Eastspring IDX ESG Leaders Plus Kelas A",
             kustodian: "STD CHARTERED", 
             penampung: "STD CHARTERED", 
             minimal: 10000, 
             returnreksa: -5.33},
        },
    }

	dbReksadana[2] = reksadana{
		jenis: "Obligasi",
        produkreksa: [NMAX]khususreksadana{
            {Produk: "ABF Indonesia Bond Index Fund",
             kustodian: "HSBC INDONESIA", 
             penampung: "HSBC INDONESIA", 
             minimal: 100000, 
             returnreksa: 6.39},

            {Produk: "Grow Obligasi Optima Dinamis Kelas O",
             kustodian: "STD CHARTERED", 
             penampung: "BCA", 
             minimal: 100000, 
             returnreksa: 1.81},

            {Produk: "Trimegah Dana Tetap Syariah Kelas A",
             kustodian: "CIMB NIAGA", 
             penampung: "BCA", 
             minimal: 100000, 
             returnreksa: 8.87},

            {Produk: "TRIM Dana Tetap 2 Kelas A",
             kustodian: "DEUTSCHE BANK", 
             penampung: "BCA", 
             minimal: 100000, 
             returnreksa: 7.61},

            {Produk: "BNI-AM ITB Harmoni",
             kustodian: "CIMB NIAGA",
             penampung: "CIMB NIAGA",
             minimal: 100000,
             returnreksa: 5.99},

            {Produk: "Danamas Stabil",
             kustodian: "BANK CIMB NIAGA",
             penampung: "BCA", 
             minimal: 10000, 
             returnreksa: 5.47},

            {Produk: "Batavia Obligasi Negara Indonesia",
             kustodian: "HSBC INDONESIA", 
             penampung: "BCA", 
             minimal: 10000, 
             returnreksa: 5.03},
        },
	}
}

func seedobligasi(){
    dbObligasi[0] = obligasi{
        produkobligasi: detailProduk{
            namaProduk: "FR0081", 
            hargaPerLembar: 1000000, 
            returnaset: 6.50,
        },
            jatuhtempo: "15 Juni 2025",
            nextkupon: "15 Juni 2025",
            penerbit: "Pemerintah Republik Indonesia",
            kupon: 6,
    }

    dbObligasi[1] = obligasi{
        produkobligasi: detailProduk{
            namaProduk: "PBS036", 
            hargaPerLembar: 1000000, 
            returnaset: 5.37,
        },
            jatuhtempo: "15 Agustus 2025",
            nextkupon: "15 Agustus 2025",
            penerbit: "Pemerintah Republik Indonesia",
            kupon: 6,
    }

    dbObligasi[2] = obligasi{
        produkobligasi: detailProduk{
            namaProduk: "FR0086", 
            hargaPerLembar: 1000000, 
            returnaset: 5.50,
        },
            jatuhtempo: "15 April 2026",
            nextkupon: "15 Oktober 2025",
            penerbit: "Pemerintah Republik Indonesia",
            kupon: 6,
    }

    dbObligasi[3] = obligasi{
        produkobligasi: detailProduk{
            namaProduk: "PBS032", 
            hargaPerLembar: 1000000, 
            returnaset: 4.87,
        },
            jatuhtempo: "15 Juli 2026",
            nextkupon: "15 Juli 2025",
            penerbit: "Pemerintah Republik Indonesia",
            kupon: 6,
    }

    dbObligasi[4] = obligasi{
        produkobligasi: detailProduk{
            namaProduk: "PBS003", 
            hargaPerLembar: 1000000, 
            returnaset: 6.00,
        },
            jatuhtempo: "15 Januari 2027", 
            nextkupon: "15 Juni 2025",
            penerbit: "Pemerintah Republik Indonesia",
            kupon: 6,
    }

    dbObligasi[5] = obligasi{
        produkobligasi: detailProduk{
            namaProduk: "FR0090", 
            hargaPerLembar: 1000000, 
            returnaset: 5.12,
        },
            jatuhtempo: "15 April 2027",
            nextkupon: "15 Oktober 2025",
            penerbit: "Pemerintah Republik Indonesia",
            kupon: 6,
    }
    
    dbObligasi[6] = obligasi{
        produkobligasi: detailProduk{
            namaProduk: "PBS030", 
            hargaPerLembar: 1000000, 
            returnaset: 5.87,
        },
            jatuhtempo: "15 Juli 2028",
            nextkupon: "15 Juli 2025",
            penerbit: "Pemerintah Republik Indonesia",
            kupon: 6,
    }
}