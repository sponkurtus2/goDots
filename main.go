package main

import (
	"dots/dotsis"
	"github.com/pterm/pterm"
	"os"
	"time"
)

func main() {
	var options []string
	centerArea, _ := pterm.DefaultArea.WithCenter().Start()

	options = append(options, "Cats")
	options = append(options, "Surprise")

	selectedDotStyle, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show("Which category do you want?")

	switch selectedDotStyle {
	case "Cats":
		centerArea.Update(pterm.Sprint(dotsis.GetFirstCat()))
		time.Sleep(time.Second * 3000)
		os.Exit(0) // Manually exit with Ctrl + c
		break
	case "Surprise":
		centerArea.Update(pterm.Sprint(dotsis.GetSecondCat()))
		time.Sleep(time.Second * 3000)
		os.Exit(0) // Manually exit with Ctrl + c
		break
	}
}
