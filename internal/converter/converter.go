package converter

import (
	"math"

	"github.com/zeuxisoo/go-ahihsd/internal/info"
)

const (
	DEGREE_TO_RADIAN float64 = math.Pi / 180.0
	RADIAN_TO_DEGREE float64 = 180.0 / math.Pi
	SCL_UNIT  float64 = 1.52587890625E-05
)

func PixelLineToLongitudeLatitude(info *info.Info, pixel, line float32, longitude, latitude *float64) {
	*longitude = -9999
	*latitude  = -9999

	x := float64(DEGREE_TO_RADIAN * float64(pixel - info.Projection.COFF) / (SCL_UNIT * float64(info.Projection.CFAC)))
	y := float64(DEGREE_TO_RADIAN * float64(line - info.Projection.LOFF) / (SCL_UNIT * float64(info.Projection.LFAC)))

	sd := (info.Projection.Rs * math.Cos(x) * math.Cos(y)) *
		  (info.Projection.Rs * math.Cos(x) * math.Cos(y)) -
		  (math.Cos(y) * math.Cos(y) + info.Projection.Req2Rpol2 * math.Sin(y) * math.Sin(y)) *
		  info.Projection.CoefficientForSd
	sd = math.Sqrt(sd)

	sn := (info.Projection.Rs * math.Cos(x) * math.Cos(y) - sd) /
		  (math.Cos(y) * math.Cos(y) + info.Projection.Req2Rpol2 * math.Sin(y) * math.Sin(y))

	s1 := info.Projection.Rs - (sn * math.Cos(x) * math.Cos(y))
	s2 := sn * math.Sin(x) * math.Cos(y)
	s3 := -sn * math.Sin(y)

	sxy := math.Sqrt(s1 * s1 + s2 * s2)

	*longitude = RADIAN_TO_DEGREE * math.Atan2(s2, s1) + info.Projection.SubLon
	*latitude  = RADIAN_TO_DEGREE * math.Atan(info.Projection.Req2Rpol2 * s3 / sxy)

	for *longitude > 180 {
		*longitude = *longitude - 360.0
	}

	for *longitude < -180 {
		*longitude = *longitude + 360.0
	}
}
