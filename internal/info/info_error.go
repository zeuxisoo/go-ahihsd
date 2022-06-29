package info

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/zeuxisoo/go-ahihsd/internal/utils"
)

type ErrorItem struct {
	LineNumber						uint16
	NumberOfErrorPixelPerOneLine 	uint16
}

type ErrorInfo struct {
	HeaderBlockNumber 	uint8
	BlockLength 		uint16
	ErrorNumber			uint16
	ErrorItems			[]ErrorItem
	Spare				[40]byte
}

func NewErrorInfo() *ErrorInfo {
	return &ErrorInfo{}
}

func (e *ErrorInfo) Read(reader io.Reader) *ErrorInfo {
	binary.Read(reader, binary.LittleEndian, &e.HeaderBlockNumber)
	binary.Read(reader, binary.LittleEndian, &e.BlockLength)
	binary.Read(reader, binary.LittleEndian, &e.ErrorNumber)

	e.ErrorItems = make([]ErrorItem, e.ErrorNumber)

	for i := 0; i<int(e.ErrorNumber); i++ {
		item := ErrorItem{}

		binary.Read(reader, binary.LittleEndian, item.LineNumber)
		binary.Read(reader, binary.LittleEndian, item.NumberOfErrorPixelPerOneLine)

		e.ErrorItems[i] = item
	}

	binary.Read(reader, binary.LittleEndian, &e.Spare)

	return e
}

func (e ErrorInfo) Show() {
	fmt.Printf("\n# 10 Error information block -----\n")

	utils.ShowInfo("header block number", e.HeaderBlockNumber)
	utils.ShowInfo("block length", e.BlockLength)
	utils.ShowInfo("number of error information", e.ErrorNumber)

	for _, item := range e.ErrorItems {
		utils.ShowInfo("line number", item.LineNumber)
		utils.ShowInfo("number of error pixels per one line", item.NumberOfErrorPixelPerOneLine)
	}

	utils.ShowInfo("spare", e.Spare)
}
