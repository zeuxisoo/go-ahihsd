package info

import (
	"fmt"

	"github.com/zeuxisoo/go-ahihsd/internal/utils"
)

type SegmentInfo struct {
	HeaderBlockNumber 					uint8
	BlockLength 						uint16
	TotalNumberOfSegment 				byte
	SegmentSequenceNumber   			byte
	FirstLineNumberOfTheImageSegment	uint16
	Spare 								[40]byte
}

func ShowSegmentInfo(data SegmentInfo) {
	fmt.Printf("\n# 7 Segment information block -----\n")

	utils.ShowInfo("header block number", data.HeaderBlockNumber)
	utils.ShowInfo("block length", data.BlockLength)
	utils.ShowInfo("total number of segments", data.TotalNumberOfSegment)
	utils.ShowInfo("segment sequence number", data.SegmentSequenceNumber)
	utils.ShowInfo("first line number of the image segment", data.FirstLineNumberOfTheImageSegment)
	utils.ShowInfo("spare", data.Spare)
}
