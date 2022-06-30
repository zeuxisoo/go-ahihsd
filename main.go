package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"github.com/zeuxisoo/go-ahihsd/internal/cmd"
)

const (
	APP_VERSION = "0.1.0"
)

func main() {
	app := cli.NewApp()
	app.Name = "AHI HSD"
	app.Usage = "A AHI HSD Reader"
	app.Version = APP_VERSION
	app.Commands = []cli.Command{
		cmd.Read,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Printf("Cannot start the application: %v", err)
	}
}
