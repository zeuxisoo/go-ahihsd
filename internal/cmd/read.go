package cmd

import (
	"errors"
	"os"

	"github.com/urfave/cli"
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

	return nil
}
