package info

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/zeuxisoo/go-ahihsd/internal/utils"
)

type BasicInfo struct {
	HeaderBlockNumber 			uint8
	BlockLength 				uint16
	TotalNumberOfHeaderBlocks 	uint16
	ByteOrder 					uint8
	SatelliteName 				[16]byte
	ProcessingCenterName 		[16]byte
	ObservationArea 			[4]byte
	OtherObservationInformation [2]byte
	ObservationTimeline			uint16
	ObservationStartTime		float64
	ObservationEndTime			float64
	FileCreationTime 			float64
	TotalHeaderLength			uint32
	TotalDataLength				uint32
	QualityFlag1				uint8
	QualityFlag2				uint8
	QualityFlag3				uint8
	QualityFlag4 				uint8
	FileFormatVersion			[32]byte
	FileName 					[128]byte
	Spare  						[40]byte
}

func NewBasicInfo() *BasicInfo {
	return &BasicInfo{}
}

func (b *BasicInfo) Read(reader io.Reader) *BasicInfo {
	binary.Read(reader, binary.LittleEndian, b)

	return b
}

func (b BasicInfo) Show() {
	fmt.Printf("\n# 1 Basic information block -----\n")

	utils.ShowInfo("header block number", b.HeaderBlockNumber)
	utils.ShowInfo("block length", b.BlockLength)
	utils.ShowInfo("total number of header blocks", b.TotalNumberOfHeaderBlocks)
	utils.ShowInfo("byte order", b.ByteOrder)
	utils.ShowInfo("satellite name", b.SatelliteName)
	utils.ShowInfo("processing center name", b.ProcessingCenterName)
	utils.ShowInfo("observation area", b.ObservationArea)
	utils.ShowInfo("other observation information", b.OtherObservationInformation)
	utils.ShowInfo("observation timeline", b.ObservationTimeline)
	utils.ShowInfo("observation start time", b.ObservationStartTime)
	utils.ShowInfo("observation end time", b.ObservationEndTime)
	utils.ShowInfo("file creation time", b.FileCreationTime)
	utils.ShowInfo("total header length", b.TotalHeaderLength)
	utils.ShowInfo("total data length", b.TotalDataLength)
	utils.ShowInfo("quality flag 1", b.QualityFlag1)
	utils.ShowInfo("quality flag 2 (spare)", b.QualityFlag2)
	utils.ShowInfo("quality flag 3", b.QualityFlag3, "%c")
	utils.ShowInfo("quality flag 4", b.QualityFlag4)
	utils.ShowInfo("file format version", b.FileFormatVersion)
	utils.ShowInfo("file name", b.FileName)
	utils.ShowInfo("spare", b.Spare)
}
