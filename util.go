package investasi

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Reader = bufio.NewReader(os.Stdin)

func BacaInput(prompt string) string {
	fmt.Print(prompt)
	input, _ := Reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func ParseFloat(input string) float64 {
	nilai, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Input tidak valid. Dianggap 0.")
		return 0.0
	}
	return nilai
}

func hitungKeuntungan(i Investasi) float64 {
	return (i.HargaAkhir - i.HargaAwal) / i.HargaAwal * 100
}