package main

import (
	"os"

	"github.com/zeuxisoo/go-ahihsd/internal/info"
)

func main() {
	file, _ := os.Open("./data/HS_H08_20170623_0250_B01_R301_R10_S0101.DAT")

	defer file.Close()

	info := info.NewInfo()
	info.Read(file).Show()
}
