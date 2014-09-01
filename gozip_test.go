package main

import (
	//"runtime"
	"github.com/codegangsta/cli"
	"github.com/kureikain/go-ziplocate/cmd"
	"testing"
)

func TestCmd(t *testing.T) {
	app := cli.NewApp()
	app.Name = "gozip"
	app.Usage = "Geocoder for code ZIP code"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdWeb,
		cmd.CmdImport,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	if app.Commands[1].Name != "import" {
		t.Fatal("Missing web command. got ", app.Commands[0].Name)
	}

	if app.Commands[0].Name != "web" {
		t.Fatal("Missing import command")
	}

}
