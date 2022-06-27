package info

import (
	"encoding/binary"
	"fmt"
	"io"

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

func NewDataInfo() *DataInfo {
	return &DataInfo{}
}

func (d *DataInfo) Read(reader io.Reader) *DataInfo {
	binary.Read(reader, binary.LittleEndian, d)

	return d
}

func (d *DataInfo) Show() {
	fmt.Printf("\n# 2 Data information block -----\n")

	utils.ShowInfo("header block number", d.HeaderBlockNumber)
	utils.ShowInfo("block length", d.BlockLength)
	utils.ShowInfo("number of bits per pixel", d.NumberOfBitsPerPixel)
	utils.ShowInfo("number of columns", d.NumberOfColumns)
	utils.ShowInfo("number of lines", d.NumberOfLines)
	utils.ShowInfo("compression flag for d", d.CompressionFlagForData)
	utils.ShowInfo("spare", d.Spare)
}
