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
	basicInfo := info.BasicInfo{}
	binary.Read(file, binary.LittleEndian, &basicInfo)

	info.ShowBasicInfo(basicInfo)

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
}
