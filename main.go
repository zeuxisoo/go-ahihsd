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
	dataInfo := info.DataInfo{}
	binary.Read(file, binary.LittleEndian, &dataInfo)

	info.ShowDataInfo(dataInfo)

	// Projection
	projectionInfo := info.ProjectionInfo{}
	binary.Read(file, binary.LittleEndian, &projectionInfo)

	info.ShowProjectionInfo(projectionInfo)

	// Navigation
	navigationInfo := info.NavigationInfo{}
	binary.Read(file, binary.LittleEndian, &navigationInfo)

	info.ShowNavigationInfo(navigationInfo)

	// Calibration
	calibrationInfo := info.CalibrationInfo{}
	binary.Read(file, binary.LittleEndian, &calibrationInfo)

	info.ShowCalibrationInfo(calibrationInfo)

	// Inter Calibration
	interCalibrationInfo := info.InterCalibrationInfo{}
	binary.Read(file, binary.LittleEndian, &interCalibrationInfo)

	info.ShowInterCalibrationInfo(interCalibrationInfo)

	// Segment
	segmentInfo := info.SegmentInfo{}
	binary.Read(file, binary.LittleEndian, &segmentInfo)

	info.ShowSegmentInfo(segmentInfo)

	// Navigation Correction
	navigationCorrectionInfo := info.NavigationCorrectionInfo{}

	info.ReadNavigationCorrectionInfo(file, &navigationCorrectionInfo)
	info.ShowNavigationCorrectionInfo(navigationCorrectionInfo, file)
}
