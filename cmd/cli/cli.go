package main

import (
	_ "embed"

	"goalculator/cmd"

	"github.com/avamsi/climate"
)

//go:generate go run goalculator/cmd/cli --out=md.cli
//go:embed md.cli
var md []byte

func main() {
	app := cmd.NewApp()
	defer app.Close()
	climate.RunAndExit(climate.Func(app.RunCli), climate.WithMetadata(md))
}
