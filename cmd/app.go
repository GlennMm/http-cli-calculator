package cmd

import (
	"fmt"

	"goalculator/internal/application"
	"goalculator/internal/storage/db"

	"gorm.io/gorm"
)

type App struct {
	Calculator *application.Calculator
	Db         *gorm.DB
}

type AppCliOptions struct {
	Operation string  `short:"o" long:"operation" description:"Operation to perform" required:"true"`
	A         float64 `short:"a" long:"first" description:"First number" required:"true"`
	B         float64 `short:"b" long:"second" description:"Second number" required:"true"`
}

func NewApp() *App {
	dB := db.GetDb()
	calculator := application.NewCalculator(dB)
	return &App{calculator, dB}
}

func (a *App) Close() {
	a.Db = nil
	a.Calculator = nil
}

func (a *App) RunCli(opts *AppCliOptions) {
	result, err := a.Calculator.PerformOperation(opts.A, opts.B, opts.Operation)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}
