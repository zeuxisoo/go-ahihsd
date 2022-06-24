package info

import (
	"fmt"

	"github.com/zeuxisoo/go-ahihsd/internal/utils"
)

type NavigationInfo struct {
	HeaderBlockNumber 					uint8
	BlockLength							uint16
	NavigationInformationTime			float64
	SSPLongitude						float64
	SSPLatitude							float64
	DistanceFromEarthCenterToSatellite	float64
	NadirLongitude						float64
	NadirLatitude						float64
	SunPositionX						float64
	SunPositionY						float64
	SunPositionZ						float64
	MoonPositionX						float64
	MoonPositionY						float64
	MoonPositionZ						float64
	Spare								[40]byte
}

func ShowNavigationInfo(data NavigationInfo) {
	fmt.Printf("\n# 4 Navigation information block -----\n")

	utils.ShowInfo("header block number", data.HeaderBlockNumber)
	utils.ShowInfo("block length", data.BlockLength)
	utils.ShowInfo("navigation information time", data.NavigationInformationTime)
	utils.ShowInfo("SSP longitude", data.SSPLongitude)
	utils.ShowInfo("SSP latitude", data.SSPLatitude)
	utils.ShowInfo("distance from Earth's center to satellite", data.DistanceFromEarthCenterToSatellite)
	utils.ShowInfo("nadir longitude", data.NadirLongitude)
	utils.ShowInfo("nadir latitude", data.NadirLatitude)
	utils.ShowInfo("Sun's position  (x)", data.SunPositionX)
	utils.ShowInfo("                (y)", data.SunPositionY)
	utils.ShowInfo("                (z)", data.SunPositionZ)
	utils.ShowInfo("Moon's position (x)", data.MoonPositionX)
	utils.ShowInfo("                (y)", data.MoonPositionY)
	utils.ShowInfo("                (z)", data.MoonPositionZ)
	utils.ShowInfo("spare", data.Spare)
}
