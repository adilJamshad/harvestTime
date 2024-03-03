package ui

import (
	"fmt"
	"sync"
	"time"

	"github.com/adilJamshad/harvestTime/internal/config"
	"github.com/adilJamshad/harvestTime/internal/timer" // Adjust this import path to match your project structure

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func TimerTab(appConfig *config.Config) fyne.CanvasObject {
	sessiontime, _ := appConfig.SessionTime.Get()
	pomodoroTimer := timer.NewTimer(time.Duration(sessiontime) * time.Minute) // Use a shorter duration for demonstration.
	timerLabel := widget.NewLabel(fmtDuration(pomodoroTimer.Duration))

	var wg sync.WaitGroup

	updateLabel := func() {
		// Directly update the timer label to show the remaining time.
		timerLabel.SetText(fmtDuration(pomodoroTimer.Remaining()))
	}

	startButton := widget.NewButton("Start", func() {
		if !pomodoroTimer.IsRunning {
			pomodoroTimer.Start()
			wg.Add(1)
			go func() {
				defer wg.Done()
				ticker := time.NewTicker(1 * time.Second)
				defer ticker.Stop()
				for range ticker.C {
					if !pomodoroTimer.IsRunning {
						break
					}
					updateLabel()
				}
			}()
		}
	})

	stopButton := widget.NewButton("Stop", func() {
		pomodoroTimer.Stop()
		updateLabel() // Ensure the label is updated immediately when the timer is stopped.
	})

	resetButton := widget.NewButton("Reset", func() {
		pomodoroTimer.Reset()
		updateLabel() // Update the label to reflect the reset timer duration.
	})

	// Initialize the label with the full duration at startup.
	updateLabel()

	content := container.NewVBox(
		timerLabel,
		startButton,
		stopButton,
		resetButton,
	)
	return content
}

func fmtDuration(d time.Duration) string {
	minutes := int(d / time.Minute)
	seconds := int(d % time.Minute / time.Second)
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
