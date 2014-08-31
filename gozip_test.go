package main

import (
	//"runtime"
	"github.com/codegangsta/cli"
	"github.com/kureikain/go-ziplocate/cmd"
	"testing"
)

func Test(t *testing.T) {
	app := cli.NewApp()
	app.Name = "gozip"
	app.Usage = "Geocoder for code ZIP code"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdWeb,
		cmd.CmdImport,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	t.Error("Missing ", 1)
}
