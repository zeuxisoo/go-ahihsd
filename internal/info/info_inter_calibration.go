package info

import (
	"fmt"

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

func ShowInterCalibrationInfo(data InterCalibrationInfo) {
	fmt.Printf("\n# 6 Inter Calibration information block -----\n")

	utils.ShowInfo("header block number", data.HeaderBlockNumber)
	utils.ShowInfo("block length", data.BlockLength)

	utils.ShowTitle("GSICS calibration coefficient)")
	utils.ShowInfo("  coefficient (Intercept)", data.GSICSCofficientIntercept)
	utils.ShowInfo("  coefficient (Slope)", data.GSICSCofficientSlope)
	utils.ShowInfo("  coefficient (Quadratic Term)", data.GSICSCofficientQuadraticTerm)

	utils.ShowTitle("Radiance bias and its uncertainty")
	utils.ShowInfo("  Radiance bias", data.RadianceBias)
	utils.ShowInfo("  Uncertainty", data.RadianceUncertainty)
	utils.ShowInfo("  Radiance for standard scene", data.RadianceStandardScene)

	utils.ShowInfo("start time of validity period", data.StartTimeOfValidityPeriod)
	utils.ShowInfo("end time of validity period", data.EndTimeOfValidityPeriod)

	utils.ShowTitle("Radiance valid range of GSICS Calibration Coefficients")
	utils.ShowInfo("  upper limit", data.RadianceValidRangeUpperLimit)
	utils.ShowInfo("  lower limit", data.RadianceValidRangeLowerLimit)

	utils.ShowInfo("File name of GSICS Correction", data.GSICSFileName)
	utils.ShowInfo("spare", data.Spare)
}
