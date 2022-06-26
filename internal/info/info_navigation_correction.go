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

func ReadNavigationCorrectionInfo(file io.Reader, navigationCorrectionInfo *NavigationCorrectionInfo) {
	header := make([]byte, 1)
	block := make([]byte, 2)
	column := make([]byte, 4)
	line := make([]byte, 4)
	amount := make([]byte, 8)
	number := make([]byte, 2)

	file.Read(header)
	file.Read(block)
	file.Read(column)
	file.Read(line)
	file.Read(amount)
	file.Read(number)

	navigationCorrectionInfo.HeaderBlockNumber					   = header[0];
	navigationCorrectionInfo.BlockLength						   = binary.LittleEndian.Uint16(block);
	navigationCorrectionInfo.CenterColumnOfRotation				   = math.Float32frombits(binary.LittleEndian.Uint32(column));
	navigationCorrectionInfo.CenterLineOfRotation				   = math.Float32frombits(binary.LittleEndian.Uint32(line));
	navigationCorrectionInfo.AmountOfRotationalCorrection		   = math.Float64frombits(binary.LittleEndian.Uint64(amount));
	navigationCorrectionInfo.CorrectionColumnAndLineDirectionNumber= binary.LittleEndian.Uint16(number);
	navigationCorrectionInfo.CorrectionColumnAndLineDirectionItems = make([]CorrectionColumnAndLineDirectionItem, binary.LittleEndian.Uint16(number));

	for i := 0; i < int(navigationCorrectionInfo.CorrectionColumnAndLineDirectionNumber); i++ {
		itemNumber := make([]byte, 2)
		itemColumn := make([]byte, 4)
		itemLine := make([]byte, 4)

		file.Read(itemNumber)
		file.Read(itemColumn)
		file.Read(itemLine)

		navigationCorrectionInfo.CorrectionColumnAndLineDirectionItems[i] = CorrectionColumnAndLineDirectionItem{
			LineNumberAfterRotation 	 : binary.LittleEndian.Uint16(itemNumber),
			ShiftAmountForColumnDirection: math.Float32frombits(binary.LittleEndian.Uint32(itemColumn)),
			ShiftAmountForLineDirection	 : math.Float32frombits(binary.LittleEndian.Uint32(itemLine)),
		}
	}

	spare := make([]byte, 40)

	file.Read(spare)

	navigationCorrectionInfo.Spare = *(*[40]byte)(spare)
}

func ShowNavigationCorrectionInfo(data NavigationCorrectionInfo, file io.Reader) {
	fmt.Printf("\n# 8 Navigation Correction information block -----\n")

	utils.ShowInfo("header block number", data.HeaderBlockNumber)
	utils.ShowInfo("block length", data.BlockLength)
	utils.ShowInfo("center column of rotation", data.CenterColumnOfRotation)
	utils.ShowInfo("center line of rotation", data.CenterLineOfRotation)
	utils.ShowInfo("amount of rotational correction", data.AmountOfRotationalCorrection)

	utils.ShowTitle("correction information for column and line direction")
	utils.ShowInfo("number of correction info", data.CorrectionColumnAndLineDirectionNumber)

	for _, item := range data.CorrectionColumnAndLineDirectionItems {
		utils.ShowInfo("line number after the rotation", item.LineNumberAfterRotation)
		utils.ShowInfo("shift amount for column direction", item.ShiftAmountForColumnDirection)
		utils.ShowInfo("shift amount for line   direction", item.ShiftAmountForLineDirection)
	}

	utils.ShowInfo("spare", data.Spare)
}
