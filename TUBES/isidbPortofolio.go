package main

func seedPorto(){
    dbPorto[0] = porto{
        total: 0,
        detailPorto: [NMAX]subtotal{
            {tanda: 1,
             tipe: "Tabungan",
             saham: 203646000,
             reksadana: 27750000,
             obligasi: 13870000,
            },
            {tanda: 1,
             tipe: "Dana Darurat",
             saham: 30923200,
             reksadana: 6620000,
             obligasi: 0,
            },
        },
    }

    dbPorto[1] = porto{
        total: 0,
        detailPorto: [NMAX]subtotal{
            {tanda: 1,
             tipe: "Dana Darurat",
             saham: 92442500,
             reksadana: 490000,
             obligasi: 0,
            },
            {tanda: 1,
             tipe: "Dana Pensiun",
             saham: 70032500,
             reksadana: 5400000,
             obligasi: 12370000,
            },
        },
    }

    dbPorto[2] = porto{
        total: 0,
        detailPorto: [NMAX]subtotal{
            {tanda: 1,
             tipe: "Dana Pendidikan",
             saham: 0,
             reksadana: 80870000,
             obligasi: 0,
            },
            {tanda: 1,
             tipe: "Liburan Keluarga",
             saham: 32228400,
             reksadana: 6470000,
             obligasi: 0,
            },
            {tanda: 1,
             tipe: "Data Investasi",
             saham: 0,
             reksadana: 0,
             obligasi: 6120000,
            },
        },
    }
}