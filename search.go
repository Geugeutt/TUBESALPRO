package investasi

import "strings"

func SequentialSearch(keyword string) []Investasi {
	var hasil []Investasi
	for _, inv := range DataInvestasi {
		if strings.Contains(strings.ToLower(inv.Nama), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(inv.Jenis), strings.ToLower(keyword)) {
			hasil = append(hasil, inv)
		}
	}
	return hasil
}

func BinarySearch(data []Investasi, keyword string) *Investasi {
	low := 0
	high := len(data) - 1

	for low <= high {
		mid := (low + high) / 2
		if strings.ToLower(data[mid].Nama) == strings.ToLower(keyword) {
			return &data[mid]
		} else if strings.ToLower(data[mid].Nama) < strings.ToLower(keyword) {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}