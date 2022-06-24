package main

import (
	"encoding/binary"
	"fmt"
	"os"
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

func ShowInfo(name string, value interface{}) {
	fmt.Printf("%-30s: ", name)

	switch v := value.(type) {
	case uint8, uint16, uint32:
		if name == "quality flag 3" {
			fmt.Printf("%c", value)
		}else{
			fmt.Printf("%d", v)
		}
	case float32, float64:
		fmt.Printf("%f", v)
	default:
		fmt.Printf("%s", v)
	}

	fmt.Printf("\n")
}

func ShowBasicInfo(data BasicInfo) {
	fmt.Printf("\n# 1 Basic information block -----\n")

	ShowInfo("header block number", data.HeaderBlockNumber)
	ShowInfo("block length", data.BlockLength)
	ShowInfo("total number of header blocks", data.TotalNumberOfHeaderBlocks)
	ShowInfo("byte order", data.ByteOrder)
	ShowInfo("satellite name", data.SatelliteName)
	ShowInfo("processing center name", data.ProcessingCenterName)
	ShowInfo("observation area", data.ObservationArea)
	ShowInfo("other observation information", data.OtherObservationInformation)
	ShowInfo("observation timeline", data.ObservationTimeline)
	ShowInfo("observation start time", data.ObservationStartTime)
	ShowInfo("observation end time", data.ObservationEndTime)
	ShowInfo("file creation time", data.FileCreationTime)
	ShowInfo("total header length", data.TotalHeaderLength)
	ShowInfo("total data length", data.TotalDataLength)
	ShowInfo("quality flag 1", data.QualityFlag1)
	ShowInfo("quality flag 2 (spare)", data.QualityFlag2)
	ShowInfo("quality flag 3", data.QualityFlag3)
	ShowInfo("quality flag 4", data.QualityFlag4)
	ShowInfo("file format version", data.FileFormatVersion)
	ShowInfo("file name", data.FileName)
	ShowInfo("spare", data.Spare)
}

func main() {
	file, _ := os.Open("./data/HS_H08_20170623_0250_B01_R301_R10_S0101.DAT")

	defer file.Close()

	basicInfo := BasicInfo{}
	binary.Read(file, binary.LittleEndian, &basicInfo)

	ShowBasicInfo(basicInfo)
}
