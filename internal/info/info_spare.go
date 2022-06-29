package info

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/zeuxisoo/go-ahihsd/internal/utils"
)

type SpareInfo struct {
	HeaderBlockNumber	uint8
	BlockLength			uint16
	Spare 				[256]byte
}

func NewSpareInfo() *SpareInfo {
	return &SpareInfo{}
}

func (s *SpareInfo) Read(reader io.Reader) *SpareInfo {
	binary.Read(reader, binary.LittleEndian, &s.HeaderBlockNumber)
	binary.Read(reader, binary.LittleEndian, &s.BlockLength)
	binary.Read(reader, binary.LittleEndian, &s.Spare)

	return s
}

func (s SpareInfo) Show() {
	fmt.Printf("\n# 11 Spare information block -----\n")

	utils.ShowInfo("header block number", s.HeaderBlockNumber)
	utils.ShowInfo("block length", s.BlockLength)
	utils.ShowInfo("spare", s.Spare)
}
