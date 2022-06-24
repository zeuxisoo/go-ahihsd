package info

import (
	"fmt"

	"github.com/zeuxisoo/go-ahihsd/internal/utils"
)

type CalibrationInfo struct {
	HeaderBlockNumber 										uint8
	BlockLength												uint16
	BandNumber												uint16
	CentralWaveLength										float64
	ValidNumberOfBitsPerPixel								uint16
	CountValueOfErrorPixels									uint16
	CountValueOfPixelsOutOfScanArea							uint16
	CountRadianceConversionEquationGain						float64
	CountRadianceConversionEquationConstant					float64
	/**
	 * (BandNumber >= 7 && &BasicInfo.SatelliteName == Himawari) ||
	 * (BandNumber >= 2 && &BasicInfo.SatelliteName == MTSAT-2)
	 *
	 CorrectionCoefficientRadianceToBrightnessTemperatureC00 float64
	 CorrectionCoefficientRadianceToBrightnessTemperatureC01 float64
	 CorrectionCoefficientRadianceToBrightnessTemperatureC02 float64
	 CorrectionCoefficientBrightnessTemperatureToRadianceC00 float64
	 CorrectionCoefficientBrightnessTemperatureToRadianceC01 float64
	 CorrectionCoefficientBrightnessTemperatureToRadianceC02 float64
	 LightSpeed												float64
	 PlanckConstant											float64
	 BoltzmannConstant										float64
	 Spare													[40]byte
	 */
	TransformationCoefficeintFromRadianceIToAlbedoA			float64
	ModifiedCalibrationCoefficientModifiedTime				float64
	ModifiedCalibrationCoefficientGain						float64
	ModifiedCalibrationCoefficientConstant					float64
	SpareV													[80]byte
}

func ShowCalibrationInfo(data CalibrationInfo) {
	fmt.Printf("\n# 5 Calibration information block -----\n")

	utils.ShowInfo("header block number", data.HeaderBlockNumber)
	utils.ShowInfo("block length", data.BlockLength)
	utils.ShowInfo("band number", data.BandNumber)
	utils.ShowInfo("central wave length", data.CentralWaveLength)
	utils.ShowInfo("valid number of bits per pixel", data.ValidNumberOfBitsPerPixel)
	utils.ShowInfo("count value of error pixels", data.CountValueOfErrorPixels)
	utils.ShowInfo("count value of pixels out of scan area", data.CountValueOfPixelsOutOfScanArea)

	utils.ShowTitle("count-radiance conversion equation")
	utils.ShowInfo("   gain", data.CountRadianceConversionEquationGain)
	utils.ShowInfo("   constant", data.CountRadianceConversionEquationConstant)

	utils.ShowTitle("transformation coefficeint(c') from radiance(I) to albedo(A)")
	utils.ShowInfo("transformation coefficeint(c')", data.TransformationCoefficeintFromRadianceIToAlbedoA)

	utils.ShowTitle("modified calibration coefficient")
	utils.ShowInfo("   modified time", data.ModifiedCalibrationCoefficientModifiedTime)
	utils.ShowInfo("   gain", data.ModifiedCalibrationCoefficientGain)
	utils.ShowInfo("   constant", data.ModifiedCalibrationCoefficientConstant)
}
