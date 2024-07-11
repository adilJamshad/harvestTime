package ui

import (
	"fmt"
	"image/color"
	"time"

	"github.com/adilJamshad/harvestTime/internal/config"
	"github.com/adilJamshad/harvestTime/internal/eventManager"

	// Adjust this import path to match your project structure
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func TimerTab(appConfig *config.Config, event_manager *eventManager.EventManager) fyne.CanvasObject {
	// Create a circular timer with a total duration based on the appConfig.
	duration, _ := appConfig.SessionTime.Get() // Get session time in minutes
	circularTimer := NewCircularTimer(time.Duration(duration)*time.Minute, 200, 5, color.Gray{Y: 0xcc}, color.RGBA{R: 255, G: 165, B: 0, A: 255})

	// Create control buttons and a label to display the time left.
	startButton := widget.NewButton("Start", func() {
		circularTimer.Start()
	})
	stopButton := widget.NewButton("Stop", func() {
		circularTimer.Stop()
	})
	resetButton := widget.NewButton("Reset", func() {
		circularTimer.Reset() // Ensure you have implemented Reset in your CircularTimer
	})
	timeLabel := widget.NewLabel("")
	timeLabel.Bind(binding.IntToStringWithFormat(appConfig.SessionTime, "%d Minutes"))

	// Update the label and circular timer when the session time is changed in settings.
	appConfig.SessionTime.AddListener(binding.NewDataListener(func() {
		newDuration, _ := appConfig.SessionTime.Get()
		timeLabel.SetText(fmt.Sprintf("%d Minutes", newDuration))
		circularTimer.UpdateDuration(time.Duration(newDuration) * time.Minute) // Ensure you have a method to update the timer's duration
	}))

	return container.NewVBox(
		circularTimer,
		container.NewHBox(startButton, stopButton, resetButton),
		timeLabel,
	)
}

func fmtDuration(d time.Duration) string {
	minutes := int(d / time.Minute)
	seconds := int(d % time.Minute / time.Second)
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
