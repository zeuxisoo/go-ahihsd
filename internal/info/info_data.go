package info

import (
	"fmt"

	"github.com/zeuxisoo/go-ahihsd/internal/utils"
)

type DataInfo struct {
	HeaderBlockNumber 		uint8
	BlockLength				uint16
	NumberOfBitsPerPixel	uint16
	NumberOfColumns			uint16
	NumberOfLines			uint16
	CompressionFlagForData	uint8
	Spare					[40]byte
}

func ShowDataInfo(data DataInfo) {
	fmt.Printf("\n# 2 Data information block -----\n")

	utils.ShowInfo("header block number", data.HeaderBlockNumber)
	utils.ShowInfo("block length", data.BlockLength)
	utils.ShowInfo("number of bits per pixel", data.NumberOfBitsPerPixel)
	utils.ShowInfo("number of columns", data.NumberOfColumns)
	utils.ShowInfo("number of lines", data.NumberOfLines)
	utils.ShowInfo("compression flag for data", data.CompressionFlagForData)
	utils.ShowInfo("spare", data.Spare)
}