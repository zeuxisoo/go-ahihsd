package info

import (
	"encoding/binary"
	"fmt"
	"io"

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

func NewSegmentInfo() *SegmentInfo {
	return &SegmentInfo{}
}

func (s *SegmentInfo) Read(reader io.Reader) *SegmentInfo {
	binary.Read(reader, binary.LittleEndian, s)

	return s
}

func (s SegmentInfo) Show() {
	fmt.Printf("\n# 7 Segment information block -----\n")

	utils.ShowInfo("header block number", s.HeaderBlockNumber)
	utils.ShowInfo("block length", s.BlockLength)
	utils.ShowInfo("total number of segments", s.TotalNumberOfSegment)
	utils.ShowInfo("segment sequence number", s.SegmentSequenceNumber)
	utils.ShowInfo("first line number of the image segment", s.FirstLineNumberOfTheImageSegment)
	utils.ShowInfo("spare", s.Spare)
}
