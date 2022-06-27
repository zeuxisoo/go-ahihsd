package info

import (
	"encoding/binary"
	"fmt"
	"io"

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

func NewProjectionInfo() *ProjectionInfo {
	return &ProjectionInfo{}
}

func (p *ProjectionInfo) Read(reader io.Reader) *ProjectionInfo {
	binary.Read(reader, binary.LittleEndian, p)

	return p
}

func (p ProjectionInfo) Show() {
	fmt.Printf("\n# 3 Projection information block -----\n")

	utils.ShowInfo("header block number", p.HeaderBlockNumber)
	utils.ShowInfo("block length", p.BlockLength)
	utils.ShowInfo("sub_lon", p.SubLon)
	utils.ShowInfo("CFAC", p.CFAC)
	utils.ShowInfo("LFAC", p.LFAC)
	utils.ShowInfo("COFF", p.COFF)
	utils.ShowInfo("LOFF", p.LOFF)
	utils.ShowInfo("Rs", p.Rs)
	utils.ShowInfo("Earth's equatorial radius", p.EarthEquatorialRadius)
	utils.ShowInfo("Earth's polar radius", p.EarthPolarRadius)
	utils.ShowInfo("(req^2 - rpol^2)/req^2", p.Req2Rpol2Req2, "%13.11f")
	utils.ShowInfo("rpol^2 / req^2", p.Rpol2Req2, "%13.11f")
	utils.ShowInfo("req^2 / rpol^2", p.Req2Rpol2, "%13.11f")
	utils.ShowInfo("coefficient for sd", p.CoefficientForSd)
	utils.ShowInfo("resampling types", p.ResamplingTypes)
	utils.ShowInfo("resampling size", p.ResamplingSize)
	utils.ShowInfo("spare", p.Spare)
}
