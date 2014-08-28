package main

import (
    "os"
    //"runtime"
    "github.com/codegangsta/cli"
    "github.com/kureikain/go-ziplocate/cmd"
)

const APP_VER = "0.0.1"

func main() {
	app := cli.NewApp()
	app.Name = "gozip"
	app.Usage = "Geocoder for ZIP"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdWeb,
		cmd.CmdImport,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	app.Run(os.Args)
}
