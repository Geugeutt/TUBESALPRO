package investasi

func SelectionSortByDana(data []Investasi) []Investasi {
	res := make([]Investasi, len(data))
	copy(res, data)
	for i := 0; i < len(res)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(res); j++ {
			if res[j].Dana < res[minIdx].Dana {
				minIdx = j
			}
		}
		res[i], res[minIdx] = res[minIdx], res[i]
	}
	return res
}

func InsertionSortByKeuntungan(data []Investasi) []Investasi {
	res := make([]Investasi, len(data))
	copy(res, data)
	for i := 1; i < len(res); i++ {
		key := res[i]
		j := i - 1
		for j >= 0 && hitungKeuntungan(res[j]) > hitungKeuntungan(key) {
			res[j+1] = res[j]
			j--
		}
		res[j+1] = key
	}
	return res
}