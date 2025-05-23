package main

import (
	"fmt"
	"strings"
)

func uiHeaderPorto(user, namaPorto string, totalAset int, saham, reksadana, obligasi float64){
	fmt.Println(strings.Repeat("-", 142))
	fmt.Printf("%-19s: %s\n%-19s: %s\n%-19s: %d\n%-19s: %.2f\n%-19s: %.2f\n%-19s: %.2f\n","User", user, "Jenis Portofolio", namaPorto, "Total Aset", totalAset, "Saham", saham, "Reksadana", reksadana, "Obligasi", obligasi)

	total := 142
	title := "=== DETAIL PORTO ==="
	left := (total - len(title)) / 2
	right := total - len(title) - left
	fmt.Println(strings.Repeat("=", left) + title + strings.Repeat("=", right))

	fmt.Printf("%-1s %-5s %-40s %-24s %-17s %-15s %-20s\n", " ", "No", "Produk", "Jenis Aset", "Nilai Aset", "Return", "Profit/Loss (Rp)")
	fmt.Println(strings.Repeat("-", 142))
}

func uiHeaderDashboard(){
	fmt.Printf("%-4s %-15s %-15s %-21s %-15s\n"," ", "Total aset", "Nilai awal", "Profit/Loss (Rp)", "Return")
	fmt.Println(strings.Repeat("-", 70))
}

func uiHeaderKatalog(jenisI, jenisII int){
	switch jenisI {
	case 1:
		fmt.Println(strings.Repeat("-", 70))
		fmt.Printf("%-7s %-10s %-15s %-15s %-25s %-7s\n"," ", "No", "Nama Produk", "Harga/lembar", "Return", " ")
		fmt.Println(strings.Repeat("-", 70))
	case 2:
		if jenisII == 1{
			fmt.Println(strings.Repeat("-", 142))
			fmt.Printf(" %-7s %-38s %-20s %-11s %-25s %s\n", "No", "Nama Produk", "Minimal Pembelian", "Return", "Bank Penampung", "Bank Kustodian")
			fmt.Println(strings.Repeat("-", 142))
		}else if jenisII == 2{
			fmt.Println(strings.Repeat("-", 142))
			fmt.Printf(" %-7s %-40s %-20s %-12s %-25s %s\n", "No", "Nama Produk", "Minimal Pembelian", "Return", "Bank Penampung", "Bank Kustodian")
			fmt.Println(strings.Repeat("-", 142))
		}else if jenisII == 3{
			fmt.Println(strings.Repeat("-", 142))
			fmt.Printf(" %-7s %-40s %-20s %-15s %-20s %s\n", "No", "Nama Produk", "Minimal Pembelian", "Return", "Bank Penampung", "Bank Kustodian")
			fmt.Println(strings.Repeat("-", 142))
		}
	case 3:
		fmt.Println(strings.Repeat("-", 145))
		fmt.Printf(" %-4s %-15s %-15s %-13s %-19s %-20s %-30s %-15s\n", "No", "Nama Produk", "Harga/unit", "Return", "Jatuh Tempo", "Pemberian Kupon", "Kupon Selanjutnya", "Penerbit")
		fmt.Println(strings.Repeat("-", 145))
	}
}

func uiHeaderRekomendasi() {
	fmt.Println(strings.Repeat("-", 142))
	fmt.Printf("  %-7s %-38s %-23s %-25s\n", "No", "Nama Produk", "Harga/Lembar", "Return")
	fmt.Println(strings.Repeat("-", 142))
}

func uiHeaderTransaksi(){
	fmt.Println(strings.Repeat("-", 145))
	fmt.Printf("  %-7s %-35s %-19s %-16s %-24s %-25s %-20s\n", "No", "Nama Produk", "Jumlah Lembar", "Nilai Aset", "Portofolio", "Jenis Aset", "Tanggal")
	fmt.Println(strings.Repeat("-", 145))
}

func uiHeaderTable(sumber string, jenisI, jenisII int ){
	var title string

	total := 142
	switch sumber {
	case "PORTOFOLIO":
		title = "=== PORTOFOLIO ==="
	case "DASHBOARD":
		total = 70
		title = "=== DASHBOARD ==="
	case "KATALOG":
		if jenisI == 1{
			total = 70
			title = "=== LIST SAHAM ==="
		}else if jenisI == 2{
			total = 142
			if jenisII == 1 {
				title = "=== LIST REKSADANA PASAR UANG ==="
			}else if jenisII == 2 {
				title = "=== LIST REKSADANA SAHAM ==="
			}else if jenisII == 3 {
				title = "=== LIST REKSADANA OBLIGASI ==="
			}
		}else if jenisI == 3{
			total = 145
			title = "=== LIST OBLIGASI ==="
		}
	case "REKOMENDASI":
		if jenisI == 1 {
			title = "=== REKOMENDASI SAHAM ==="
		}else if jenisI == 2 {
			if jenisII == 1 {
				title = "=== REKOMENDASI REKSADANA PASAR UANG ==="
			}else if jenisII == 2 {
				title = "=== REKOMENDASI REKSADANA SAHAM ==="
			}else if jenisII == 3 {
				title = "=== REKOMENDASI REKSADANA OBLIGASI ==="
			}
		}else if jenisI == 3 {
			title = "=== REKOMENDASI OBLIGASI ==="
		}
	case "TRANSAKSI":
		total = 145
		if jenisI == 1 {
			title = "=== TRANSAKSI PEMBELIAN ==="
		}else if jenisI == 2{
			title = "=== TRANSAKSI PENJUALAN ==="
		}
	}
	
	left := (total - len(title)) / 2
	right := total - len(title) - left

	fmt.Println(strings.Repeat("=", left) + title + strings.Repeat("=", right))

}

func uiFooterTablePanjang(){
	fmt.Println(strings.Repeat("=", 142))
}

func uiFooterTablePendek(){
	fmt.Println(strings.Repeat("=", 70))
}

func uiFooterTableSuperPanjang(){
	fmt.Println(strings.Repeat("=", 145))
}