package info

import (
	"encoding/binary"
	"fmt"
	"io"

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

func NewNavigationInfo() *NavigationInfo {
	return &NavigationInfo{}
}

func (n *NavigationInfo) Read(reader io.Reader) *NavigationInfo {
	binary.Read(reader, binary.LittleEndian, n)

	return n
}

func (n NavigationInfo) Show() {
	fmt.Printf("\n# 4 Navigation information block -----\n")

	utils.ShowInfo("header block number", n.HeaderBlockNumber)
	utils.ShowInfo("block length", n.BlockLength)
	utils.ShowInfo("navigation information time", n.NavigationInformationTime)
	utils.ShowInfo("SSP longitude", n.SSPLongitude)
	utils.ShowInfo("SSP latitude", n.SSPLatitude)
	utils.ShowInfo("distance from Earth's center to satellite", n.DistanceFromEarthCenterToSatellite)
	utils.ShowInfo("nadir longitude", n.NadirLongitude)
	utils.ShowInfo("nadir latitude", n.NadirLatitude)
	utils.ShowInfo("Sun's position  (x)", n.SunPositionX)
	utils.ShowInfo("                (y)", n.SunPositionY)
	utils.ShowInfo("                (z)", n.SunPositionZ)
	utils.ShowInfo("Moon's position (x)", n.MoonPositionX)
	utils.ShowInfo("                (y)", n.MoonPositionY)
	utils.ShowInfo("                (z)", n.MoonPositionZ)
	utils.ShowInfo("spare", n.Spare)
}
