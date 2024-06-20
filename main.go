package main

import (
	"nginx-config-parse/cmd"
	"nginx-config-parse/src/util"
)

func main() {

	ngntcmd := cmd.NewRootCmd()
	err := ngntcmd.Execute()
	if err != nil {
		util.Logger.Error("run cmd error", err.Error())
		return
	}
}
