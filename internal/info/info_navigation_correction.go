package info

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/zeuxisoo/go-ahihsd/internal/utils"
)

type NavigationCorrectionInfoDataItem struct {
	LineNumberAfterRotation 		uint16
	ShiftAmountForColumnDirection 	float32
	ShiftAmountForLineDirection 	float32
}

type NavigationCorrectionInfo struct {
	HeaderBlockNumber 					uint8
	BlockLength 						uint16
	CenterColumnOfRotation 				float32
	CenterLineOfRotation 				float32
	AmountOfRotationalCorrection 		float64
	NumberOfCorrectionInfoDataNumber	uint16
	CorrectionInfoDataItems  			[]NavigationCorrectionInfoDataItem
	Spare 								[40]byte
}

func NewNavigationCorrectionInfo() *NavigationCorrectionInfo {
	return &NavigationCorrectionInfo{}
}

func (n *NavigationCorrectionInfo) Read(reader io.Reader) *NavigationCorrectionInfo {
	binary.Read(reader, binary.LittleEndian, &n.HeaderBlockNumber)
	binary.Read(reader, binary.LittleEndian, &n.BlockLength)
	binary.Read(reader, binary.LittleEndian, &n.CenterColumnOfRotation)
	binary.Read(reader, binary.LittleEndian, &n.CenterLineOfRotation)
	binary.Read(reader, binary.LittleEndian, &n.AmountOfRotationalCorrection)
	binary.Read(reader, binary.LittleEndian, &n.NumberOfCorrectionInfoDataNumber)

	n.CorrectionInfoDataItems = make(
		[]NavigationCorrectionInfoDataItem,
		n.NumberOfCorrectionInfoDataNumber,
	)

	for i := 0; i < int(n.NumberOfCorrectionInfoDataNumber); i++ {
		item := NavigationCorrectionInfoDataItem{}

		binary.Read(reader, binary.LittleEndian, &item)

		n.CorrectionInfoDataItems[i] = item
	}

	binary.Read(reader, binary.LittleEndian, &n.Spare)

	return n
}

/*
func (n *NavigationCorrectionInfo) Read(reader io.Reader) *NavigationCorrectionInfo {
	header := make([]byte, 1)
	block := make([]byte, 2)
	column := make([]byte, 4)
	line := make([]byte, 4)
	amount := make([]byte, 8)
	number := make([]byte, 2)

	reader.Read(header)
	reader.Read(block)
	reader.Read(column)
	reader.Read(line)
	reader.Read(amount)
	reader.Read(number)

	n.HeaderBlockNumber					    = header[0];
	n.BlockLength						    = binary.LittleEndian.Uint16(block);
	n.CenterColumnOfRotation			    = math.Float32frombits(binary.LittleEndian.Uint32(column));
	n.CenterLineOfRotation				    = math.Float32frombits(binary.LittleEndian.Uint32(line));
	n.AmountOfRotationalCorrection		    = math.Float64frombits(binary.LittleEndian.Uint64(amount));
	n.NumberOfCorrectionInfoDataNumber= binary.LittleEndian.Uint16(number);
	n.CorrectionInfoDataItems = make([]NavigationCorrectionInfoDataItem, binary.LittleEndian.Uint16(number));

	for i := 0; i < int(n.NumberOfCorrectionInfoDataNumber); i++ {
		itemNumber := make([]byte, 2)
		itemColumn := make([]byte, 4)
		itemLine := make([]byte, 4)

		reader.Read(itemNumber)
		reader.Read(itemColumn)
		reader.Read(itemLine)

		n.CorrectionInfoDataItems[i] = NavigationCorrectionInfoDataItem{
			LineNumberAfterRotation 	 : binary.LittleEndian.Uint16(itemNumber),
			ShiftAmountForColumnDirection: math.Float32frombits(binary.LittleEndian.Uint32(itemColumn)),
			ShiftAmountForLineDirection	 : math.Float32frombits(binary.LittleEndian.Uint32(itemLine)),
		}
	}

	spare := make([]byte, 40)

	reader.Read(spare)

	n.Spare = *(*[40]byte)(spare)

	return n
}
*/

func (n NavigationCorrectionInfo) Show() {
	fmt.Printf("\n# 8 Navigation Correction information block -----\n")

	utils.ShowInfo("header block number", n.HeaderBlockNumber)
	utils.ShowInfo("block length", n.BlockLength)
	utils.ShowInfo("center column of rotation", n.CenterColumnOfRotation)
	utils.ShowInfo("center line of rotation", n.CenterLineOfRotation)
	utils.ShowInfo("amount of rotational correction", n.AmountOfRotationalCorrection)

	utils.ShowTitle("correction information for column and line direction")
	utils.ShowInfo("number of correction info", n.NumberOfCorrectionInfoDataNumber)

	for _, item := range n.CorrectionInfoDataItems {
		utils.ShowInfo("line number after the rotation", item.LineNumberAfterRotation)
		utils.ShowInfo("shift amount for column direction", item.ShiftAmountForColumnDirection)
		utils.ShowInfo("shift amount for line   direction", item.ShiftAmountForLineDirection)
	}

	utils.ShowInfo("spare", n.Spare)
}
