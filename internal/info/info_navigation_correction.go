package info

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"

	"github.com/zeuxisoo/go-ahihsd/internal/utils"
)

type CorrectionColumnAndLineDirectionItem struct {
	LineNumberAfterRotation 		uint16
	ShiftAmountForColumnDirection 	float32
	ShiftAmountForLineDirection 	float32
}

type NavigationCorrectionInfo struct {
	HeaderBlockNumber 						uint8
	BlockLength 							uint16
	CenterColumnOfRotation 					float32
	CenterLineOfRotation 					float32
	AmountOfRotationalCorrection 			float64
	CorrectionColumnAndLineDirectionNumber 	uint16
	CorrectionColumnAndLineDirectionItems   []CorrectionColumnAndLineDirectionItem
	Spare 									[40]byte
}

func NewNavigationCorrectionInfo() *NavigationCorrectionInfo {
	return &NavigationCorrectionInfo{}
}

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
	n.CorrectionColumnAndLineDirectionNumber= binary.LittleEndian.Uint16(number);
	n.CorrectionColumnAndLineDirectionItems = make([]CorrectionColumnAndLineDirectionItem, binary.LittleEndian.Uint16(number));

	for i := 0; i < int(n.CorrectionColumnAndLineDirectionNumber); i++ {
		itemNumber := make([]byte, 2)
		itemColumn := make([]byte, 4)
		itemLine := make([]byte, 4)

		reader.Read(itemNumber)
		reader.Read(itemColumn)
		reader.Read(itemLine)

		n.CorrectionColumnAndLineDirectionItems[i] = CorrectionColumnAndLineDirectionItem{
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

func (n NavigationCorrectionInfo) Show() {
	fmt.Printf("\n# 8 Navigation Correction information block -----\n")

	utils.ShowInfo("header block number", n.HeaderBlockNumber)
	utils.ShowInfo("block length", n.BlockLength)
	utils.ShowInfo("center column of rotation", n.CenterColumnOfRotation)
	utils.ShowInfo("center line of rotation", n.CenterLineOfRotation)
	utils.ShowInfo("amount of rotational correction", n.AmountOfRotationalCorrection)

	utils.ShowTitle("correction information for column and line direction")
	utils.ShowInfo("number of correction info", n.CorrectionColumnAndLineDirectionNumber)

	for _, item := range n.CorrectionColumnAndLineDirectionItems {
		utils.ShowInfo("line number after the rotation", item.LineNumberAfterRotation)
		utils.ShowInfo("shift amount for column direction", item.ShiftAmountForColumnDirection)
		utils.ShowInfo("shift amount for line   direction", item.ShiftAmountForLineDirection)
	}

	utils.ShowInfo("spare", n.Spare)
}
