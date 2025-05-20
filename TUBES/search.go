package main
import "fmt"

func cariAset(target string) int{
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

