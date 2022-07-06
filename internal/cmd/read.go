package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/zeuxisoo/go-ahihsd/internal/converter"
	"github.com/zeuxisoo/go-ahihsd/internal/info"
)

var Read = cli.Command{
	Name: "read",
	Usage: "Read the .dat file",
	Description: "Read the AHI HSD dat file and show it",
	Action: runRead,
	Flags: []cli.Flag{
		stringFlag("file, f", "", "which file should be read"),
	},
}

func runRead(c *cli.Context) error {
	filePath := c.String("file")

	if len(filePath) <= 0 {
		return errors.New("please enter dat path")
	}

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		return errors.New("dat file not exists")
	}

	file, _ := os.Open(filePath)
	defer file.Close()

	info := info.NewInfo()
	info.Read(file).Show()

	showPixelLineAndLongitudeLatitude(info)

	return nil
}

func showPixelLineAndLongitudeLatitude(info *info.Info) {
	longitude := [5]float64{}
	latitude  := [5]float64{}
	pixel     := [5]float32{}
	line      := [5]float32{}

	pixel[0] = float32(info.Data.NumberOfColumns / 9 * 3)
	pixel[1] = float32(info.Data.NumberOfColumns / 9 * 4)
	pixel[2] = float32(info.Data.NumberOfColumns / 9 * 5)
	pixel[3] = float32(info.Data.NumberOfColumns / 9 * 6)
	pixel[4] = float32(info.Data.NumberOfColumns / 9 * 7)

	line[0] = float32(info.Segment.FirstLineNumberOfTheImageSegment + info.Data.NumberOfLines / 6 * 1)
	line[1] = float32(info.Segment.FirstLineNumberOfTheImageSegment + info.Data.NumberOfLines / 6 * 2)
	line[2] = float32(info.Segment.FirstLineNumberOfTheImageSegment + info.Data.NumberOfLines / 6 * 3)
	line[3] = float32(info.Segment.FirstLineNumberOfTheImageSegment + info.Data.NumberOfLines / 6 * 4)
	line[4] = float32(info.Segment.FirstLineNumberOfTheImageSegment + info.Data.NumberOfLines / 6 * 5)

	fmt.Println("\n# convert from (pixel,line) to (longitude,latitude) ---")

	for i := 0; i<5; i++ {
		converter.PixelLineToLongitudeLatitude(info, pixel[i], line[i], &longitude[i], &latitude[i])

		fmt.Printf(
			"(Pix,Lin)(%8.1f,%8.1f) ==> (Lon,Lat)(%9.3f,%9.3f)\n",
			pixel[i], line[i], longitude[i], latitude[i],
		)
	}

	fmt.Printf("\n# convert from (longitude,latitude) to (pixel,line) ---\n")

	for i := 0; i<5; i++ {
		converter.LongitudeLatitudeToPixelLine(info, longitude[i], latitude[i], &pixel[i], &line[i])

		fmt.Printf(
			"(Lon,Lat)(%9.3f,%9.3f) ==> (Pix,Lin)(%6.1f,%6.1f)\n",
			longitude[i], latitude[i], pixel[i], line[i],
		)
	}
}
