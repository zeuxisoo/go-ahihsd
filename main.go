package main

import (
	"encoding/binary"
	"os"

	"github.com/zeuxisoo/go-ahihsd/internal/info"
)

func main() {
	file, _ := os.Open("./data/HS_H08_20170623_0250_B01_R301_R10_S0101.DAT")

	defer file.Close()

	// Basic
	basicInfo := info.NewBasicInfo()
	basicInfo.Read(file).Show()

	// Data
	dataInfo := info.NewDataInfo()
	dataInfo.Read(file).Show()

	// Projection
	projectionInfo := info.NewProjectionInfo()
	projectionInfo.Read(file).Show()

	// Navigation
	navigationInfo := info.NewNavigationInfo()
	navigationInfo.Read(file).Show()

	// Calibration
	calibrationInfo := info.NewCalibrationInfo()
	calibrationInfo.Read(file).Show()

	// Inter Calibration
	interCalibrationInfo := info.NewInterCalibrationInfo()
	interCalibrationInfo.Read(file).Show()

	// Segment
	segmentInfo := info.SegmentInfo{}
	binary.Read(file, binary.LittleEndian, &segmentInfo)

	info.ShowSegmentInfo(segmentInfo)

	// Navigation Correction
	navigationCorrectionInfo := info.NavigationCorrectionInfo{}

	info.ReadNavigationCorrectionInfo(file, &navigationCorrectionInfo)
	info.ShowNavigationCorrectionInfo(navigationCorrectionInfo, file)
}
