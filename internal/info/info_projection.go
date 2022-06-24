package info

import (
	"fmt"

	"github.com/zeuxisoo/go-ahihsd/internal/utils"
)

type ProjectionInfo struct {
	HeaderBlockNumber 		uint8
	BlockLength				uint16
	SubLon					float64
	CFAC					uint32
	LFAC					uint32
	COFF					float32
	LOFF					float32
	Rs						float64
	EarthEquatorialRadius	float64
	EarthPolarRadius		float64
	Req2Rpol2Req2			float64
	Rpol2Req2				float64
	Req2Rpol2               float64
	CoefficientForSd		float64
	ResamplingTypes			uint16
	ResamplingSize			uint16
	Spare					[40]byte
}

func ShowProjectionInfo(data ProjectionInfo) {
	fmt.Printf("\n# 3 Projection information block -----\n")

	utils.ShowInfo("header block number", data.HeaderBlockNumber)
	utils.ShowInfo("block length", data.BlockLength)
	utils.ShowInfo("sub_lon", data.SubLon)
	utils.ShowInfo("CFAC", data.CFAC)
	utils.ShowInfo("LFAC", data.LFAC)
	utils.ShowInfo("COFF", data.COFF)
	utils.ShowInfo("LOFF", data.LOFF)
	utils.ShowInfo("Rs", data.Rs)
	utils.ShowInfo("Earth's equatorial radius", data.EarthEquatorialRadius)
	utils.ShowInfo("Earth's polar radius", data.EarthPolarRadius)
	utils.ShowInfo("(req^2 - rpol^2)/req^2", data.Req2Rpol2Req2, "%13.11f")
	utils.ShowInfo("rpol^2 / req^2", data.Rpol2Req2, "%13.11f")
	utils.ShowInfo("req^2 / rpol^2", data.Req2Rpol2, "%13.11f")
	utils.ShowInfo("coefficient for sd", data.CoefficientForSd)
	utils.ShowInfo("resampling types", data.ResamplingTypes)
	utils.ShowInfo("resampling size", data.ResamplingSize)
	utils.ShowInfo("spare", data.Spare)
}
