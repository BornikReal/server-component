package main

import (
	"fmt"

	"github.com/BornikReal/server-component/internal"
)

func main() {
	app := internal.NewApp()
	if err := app.Init(); err != nil {
		fmt.Println(err.Error())
	}

	app.Run()
}
