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

func (basic *BasicInfo) Read(file io.Reader) *BasicInfo {
	binary.Read(file, binary.LittleEndian, basic)

	return basic
}

func (basic BasicInfo) Show() {
	fmt.Printf("\n# 1 Basic information block -----\n")

	utils.ShowInfo("header block number", basic.HeaderBlockNumber)
	utils.ShowInfo("block length", basic.BlockLength)
	utils.ShowInfo("total number of header blocks", basic.TotalNumberOfHeaderBlocks)
	utils.ShowInfo("byte order", basic.ByteOrder)
	utils.ShowInfo("satellite name", basic.SatelliteName)
	utils.ShowInfo("processing center name", basic.ProcessingCenterName)
	utils.ShowInfo("observation area", basic.ObservationArea)
	utils.ShowInfo("other observation information", basic.OtherObservationInformation)
	utils.ShowInfo("observation timeline", basic.ObservationTimeline)
	utils.ShowInfo("observation start time", basic.ObservationStartTime)
	utils.ShowInfo("observation end time", basic.ObservationEndTime)
	utils.ShowInfo("file creation time", basic.FileCreationTime)
	utils.ShowInfo("total header length", basic.TotalHeaderLength)
	utils.ShowInfo("total data length", basic.TotalDataLength)
	utils.ShowInfo("quality flag 1", basic.QualityFlag1)
	utils.ShowInfo("quality flag 2 (spare)", basic.QualityFlag2)
	utils.ShowInfo("quality flag 3", basic.QualityFlag3, "%c")
	utils.ShowInfo("quality flag 4", basic.QualityFlag4)
	utils.ShowInfo("file format version", basic.FileFormatVersion)
	utils.ShowInfo("file name", basic.FileName)
	utils.ShowInfo("spare", basic.Spare)
}
