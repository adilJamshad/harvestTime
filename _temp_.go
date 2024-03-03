package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Pomodoro Timer")

	initialDuration := 25 * time.Minute // Using a shorter duration for demonstration.

	timerLabel := widget.NewLabel(fmtDuration(initialDuration))

	var startTime time.Time
	var elapsedTime time.Duration
	var ticker *time.Ticker
	var timerRunning bool

	updateTimerDisplay := func(d time.Duration) {
		timerLabel.SetText(fmtDuration(d))
	}

	startTimer := func() {
		if !timerRunning {
			timerRunning = true
			if ticker == nil {
				startTime = time.Now().Add(-elapsedTime)
				ticker = time.NewTicker(1 * time.Second)
				go func() {
					for range ticker.C {
						elapsedTime = time.Since(startTime)
						remainingTime := initialDuration - elapsedTime
						if remainingTime <= 0 {
							ticker.Stop()
							timerRunning = false
							updateTimerDisplay(0)
							return
						}
						updateTimerDisplay(remainingTime)
					}
				}()
			}
		}
	}

	stopTimer := func() {
		if timerRunning && ticker != nil {
			ticker.Stop()
			ticker = nil
			timerRunning = false
		}
	}

	resetTimer := func() {
		stopTimer()
		elapsedTime = 0
		updateTimerDisplay(initialDuration)
	}

	startButton := widget.NewButton("Start/Resume", func() {
		startTimer()
	})

	stopButton := widget.NewButton("Stop/Pause", func() {
		stopTimer()
	})

	resetButton := widget.NewButton("Reset", func() {
		resetTimer()
	})

	content := container.NewVBox(
		timerLabel,
		startButton,
		stopButton,
		resetButton,
	)
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func fmtDuration(d time.Duration) string {
	minutes := int(d / time.Minute)
	seconds := int(d % time.Minute / time.Second)
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}
