package investasi

import "fmt"

type Investasi struct {
	Nama        string
	Jenis       string
	Dana        float64
	HargaAwal   float64
	HargaAkhir  float64
}

var DataInvestasi []Investasi

func TambahInvestasi(i Investasi) {
	DataInvestasi = append(DataInvestasi, i)
	fmt.Println("Data berhasil ditambahkan.")
}

func UbahInvestasi(index int, i Investasi) {
	if index >= 0 && index < len(DataInvestasi) {
		DataInvestasi[index] = i
		fmt.Println("Data berhasil diubah.")
	} else {
		fmt.Println("Index tidak valid.")
	}
}

func HapusInvestasi(index int) {
	if index >= 0 && index < len(DataInvestasi) {
		DataInvestasi = append(DataInvestasi[:index], DataInvestasi[index+1:]...)
		fmt.Println("Data berhasil dihapus.")
	} else {
		fmt.Println("Index tidak valid.")
	}
}

func TampilkanInvestasi() {
	fmt.Println("\n=== Daftar Investasi ===")
	for i, inv := range DataInvestasi {
		fmt.Printf("[%d] %s - %s | Dana: %.2f | Keuntungan: %.2f%%\n",
			i, inv.Nama, inv.Jenis, inv.Dana,
			((inv.HargaAkhir - inv.HargaAwal) / inv.HargaAwal * 100))
	}
}