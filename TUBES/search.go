package main
import "fmt"

func cariAsetSequential(target string) int{
	var (
		found bool
		i int
	)

	found = false
    for i = 0; i < NMAX; i++ {
		if dbSaham[i].produksaham.namaProduk == target {
			found = true
			return i
		}	
    }

    if !found {
        fmt.Println("Produk tidak tersedia")
    }

	return -i
}

func cariAsetBinary(kunci *int, target string) int{
	var i, high, low, mid int

	for i=0; i<NMAX; i++{
		if dbTempforSort[*kunci].riwayat[i].tanda == 0{
			high = i-1
			break
		}
	}


	for low <= high {
		mid = (low + high) / 2
		if dbTempforSort[*kunci].riwayat[mid].produk == target {
			return mid
		} else if dbTempforSort[*kunci].riwayat[mid].produk < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}