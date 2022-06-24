package info

import (
	"fmt"

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

func ShowBasicInfo(data BasicInfo) {
	fmt.Printf("\n# 1 Basic information block -----\n")

	utils.ShowInfo("header block number", data.HeaderBlockNumber)
	utils.ShowInfo("block length", data.BlockLength)
	utils.ShowInfo("total number of header blocks", data.TotalNumberOfHeaderBlocks)
	utils.ShowInfo("byte order", data.ByteOrder)
	utils.ShowInfo("satellite name", data.SatelliteName)
	utils.ShowInfo("processing center name", data.ProcessingCenterName)
	utils.ShowInfo("observation area", data.ObservationArea)
	utils.ShowInfo("other observation information", data.OtherObservationInformation)
	utils.ShowInfo("observation timeline", data.ObservationTimeline)
	utils.ShowInfo("observation start time", data.ObservationStartTime)
	utils.ShowInfo("observation end time", data.ObservationEndTime)
	utils.ShowInfo("file creation time", data.FileCreationTime)
	utils.ShowInfo("total header length", data.TotalHeaderLength)
	utils.ShowInfo("total data length", data.TotalDataLength)
	utils.ShowInfo("quality flag 1", data.QualityFlag1)
	utils.ShowInfo("quality flag 2 (spare)", data.QualityFlag2)
	utils.ShowInfo("quality flag 3", data.QualityFlag3)
	utils.ShowInfo("quality flag 4", data.QualityFlag4)
	utils.ShowInfo("file format version", data.FileFormatVersion)
	utils.ShowInfo("file name", data.FileName)
	utils.ShowInfo("spare", data.Spare)
}
