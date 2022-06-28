package info

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/zeuxisoo/go-ahihsd/internal/utils"
)

type ObservationItem struct {
	LineNumber 			uint16
	ObservationTimeMJD 	float64
}

type ObservationInfo struct {
	HeaderBlockNumber 	uint8
	BlockLength 		uint16
	ObservationNumber   uint16
	ObservationItems    []ObservationItem
	Spare				[40]byte
}

func NewObservationInfo() *ObservationInfo {
	return &ObservationInfo{}
}

func (o *ObservationInfo) Read(reader io.Reader) *ObservationInfo {
	binary.Read(reader, binary.LittleEndian, &o.HeaderBlockNumber)
	binary.Read(reader, binary.LittleEndian, &o.BlockLength)
	binary.Read(reader, binary.LittleEndian, &o.ObservationNumber)

	o.ObservationItems = make(
		[]ObservationItem,
		o.ObservationNumber,
	)

	for i := 0; i < int(o.ObservationNumber); i++ {
		item := ObservationItem{}

		binary.Read(reader, binary.LittleEndian, &item)

		o.ObservationItems[i] = item
	}

	binary.Read(reader, binary.LittleEndian, &o.Spare)

	return o
}

func (o ObservationInfo) Show() {
	fmt.Printf("\n# 9 Observation information block -----\n")

	utils.ShowInfo("header block number", o.HeaderBlockNumber)
	utils.ShowInfo("block length", o.BlockLength)
	utils.ShowInfo("number of observation time information", o.ObservationNumber)

	for _, item := range o.ObservationItems {
		utils.ShowInfo("line number", item.LineNumber)
		utils.ShowInfo("observation time (MJD)", item.ObservationTimeMJD)
	}

	utils.ShowInfo("spare", o.Spare)
}
