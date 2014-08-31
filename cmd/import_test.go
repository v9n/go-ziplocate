package cmd

import (
	//"log"
	//"github.com/codegangsta/cli"
	"testing"
)

func TestDef(t *testing.T) {
	if CmdImport.Name != "import" {
		t.Fatal("Wrong Command Name")
	}
}

