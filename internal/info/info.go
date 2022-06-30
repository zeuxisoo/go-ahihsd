package info

import "io"

type Info struct {
	Basic 					*BasicInfo
	Data 					*DataInfo
	Projection 				*ProjectionInfo
	Navigation 				*NavigationInfo
	Calibration 			*CalibrationInfo
	InterCalibration 		*InterCalibrationInfo
	Segment 				*SegmentInfo
	NavigationCorrection 	*NavigationCorrectionInfo
	Observation				*ObservationInfo
	Error					*ErrorInfo
	Spare					*SpareInfo
}

func NewInfo() *Info {
	return &Info{}
}

func (i *Info) Read(reader io.Reader) *Info {
	i.Basic                = NewBasicInfo().Read(reader)
	i.Data                 = NewDataInfo().Read(reader)
	i.Projection           = NewProjectionInfo().Read(reader)
	i.Navigation           = NewNavigationInfo().Read(reader)
	i.Calibration          = NewCalibrationInfo().Read(reader)
	i.InterCalibration     = NewInterCalibrationInfo().Read(reader)
	i.Segment              = NewSegmentInfo().Read(reader)
	i.NavigationCorrection = NewNavigationCorrectionInfo().Read(reader)
	i.Observation          = NewObservationInfo().Read(reader)
	i.Error                = NewErrorInfo().Read(reader)
	i.Spare                = NewSpareInfo().Read(reader)

	return i
}

func (i Info) Show() {
	i.Basic.Show()
	i.Data.Show()
	i.Projection.Show()
	i.Navigation.Show()
	i.Calibration.Show()
	i.InterCalibration.Show()
	i.Segment.Show()
	i.NavigationCorrection.Show()
	i.Observation.Show()
	i.Error.Show()
	i.Spare.Show()
}
