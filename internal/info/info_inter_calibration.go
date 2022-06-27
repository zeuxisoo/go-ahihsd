package info

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/zeuxisoo/go-ahihsd/internal/utils"
)

type InterCalibrationInfo struct {
	HeaderBlockNumber 				uint8
	BlockLength						uint16
	GSICSCofficientIntercept		float64
	GSICSCofficientSlope			float64
	GSICSCofficientQuadraticTerm	float64
	RadianceBias 					float64
	RadianceUncertainty 			float64
	RadianceStandardScene 			float64
	StartTimeOfValidityPeriod 		float64
	EndTimeOfValidityPeriod 		float64
	RadianceValidRangeUpperLimit 	float32
	RadianceValidRangeLowerLimit 	float32
	GSICSFileName 					[128]byte
	Spare 							[56]byte
}

func NewInterCalibrationInfo() *InterCalibrationInfo {
	return &InterCalibrationInfo{}
}

func (i *InterCalibrationInfo) Read(reader io.Reader) *InterCalibrationInfo {
	binary.Read(reader, binary.LittleEndian, i)

	return i
}

func (i InterCalibrationInfo) Show() {
	fmt.Printf("\n# 6 Inter Calibration information block -----\n")

	utils.ShowInfo("header block number", i.HeaderBlockNumber)
	utils.ShowInfo("block length", i.BlockLength)

	utils.ShowTitle("GSICS calibration coefficient)")
	utils.ShowInfo("  coefficient (Intercept)", i.GSICSCofficientIntercept)
	utils.ShowInfo("  coefficient (Slope)", i.GSICSCofficientSlope)
	utils.ShowInfo("  coefficient (Quadratic Term)", i.GSICSCofficientQuadraticTerm)

	utils.ShowTitle("Radiance bias and its uncertainty")
	utils.ShowInfo("  Radiance bias", i.RadianceBias)
	utils.ShowInfo("  Uncertainty", i.RadianceUncertainty)
	utils.ShowInfo("  Radiance for standard scene", i.RadianceStandardScene)

	utils.ShowInfo("start time of validity period", i.StartTimeOfValidityPeriod)
	utils.ShowInfo("end time of validity period", i.EndTimeOfValidityPeriod)

	utils.ShowTitle("Radiance valid range of GSICS Calibration Coefficients")
	utils.ShowInfo("  upper limit", i.RadianceValidRangeUpperLimit)
	utils.ShowInfo("  lower limit", i.RadianceValidRangeLowerLimit)

	utils.ShowInfo("File name of GSICS Correction", i.GSICSFileName)
	utils.ShowInfo("spare", i.Spare)
}
