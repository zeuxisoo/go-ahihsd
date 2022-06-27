package info

import (
	"encoding/binary"
	"fmt"
	"io"

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

func NewCalibrationInfo() *CalibrationInfo {
	return &CalibrationInfo{}
}

func (c *CalibrationInfo) Read(reader io.Reader) *CalibrationInfo {
	binary.Read(reader, binary.LittleEndian, c)

	return c
}

func (c CalibrationInfo) Show() {
	fmt.Printf("\n# 5 Calibration information block -----\n")

	utils.ShowInfo("header block number", c.HeaderBlockNumber)
	utils.ShowInfo("block length", c.BlockLength)
	utils.ShowInfo("band number", c.BandNumber)
	utils.ShowInfo("central wave length", c.CentralWaveLength)
	utils.ShowInfo("valid number of bits per pixel", c.ValidNumberOfBitsPerPixel)
	utils.ShowInfo("count value of error pixels", c.CountValueOfErrorPixels)
	utils.ShowInfo("count value of pixels out of scan area", c.CountValueOfPixelsOutOfScanArea)

	utils.ShowTitle("count-radiance conversion equation")
	utils.ShowInfo("   gain", c.CountRadianceConversionEquationGain)
	utils.ShowInfo("   constant", c.CountRadianceConversionEquationConstant)

	utils.ShowTitle("transformation coefficeint(c') from radiance(I) to albedo(A)")
	utils.ShowInfo("transformation coefficeint(c')", c.TransformationCoefficeintFromRadianceIToAlbedoA)

	utils.ShowTitle("modified calibration coefficient")
	utils.ShowInfo("   modified time", c.ModifiedCalibrationCoefficientModifiedTime)
	utils.ShowInfo("   gain", c.ModifiedCalibrationCoefficientGain)
	utils.ShowInfo("   constant", c.ModifiedCalibrationCoefficientConstant)
}
