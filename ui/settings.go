package ui

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/adilJamshad/harvestTime/internal/config"
)

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func SettingsTab(appConfig *config.Config) fyne.CanvasObject {

	// Create UI elements with data binding.
	sessionEntry := widget.NewEntryWithData(binding.IntToString(appConfig.SessionTime))
	breakEntry := widget.NewEntryWithData(binding.IntToString(appConfig.BreakTime))
	pushNotificationCheck := widget.NewCheckWithData("Push Notifications", appConfig.PushNotifications)

	// Save button logic.
	saveButton := widget.NewButton("Save", func() {
		// Parse string values back to int for session and break times.
		sessionTimeValue, err := strconv.Atoi(sessionEntry.Text)
		handleErr(err)
		breakTimeValue, err := strconv.Atoi(breakEntry.Text)
		handleErr(err)

		// Update binding values.
		appConfig.SessionTime.Set(sessionTimeValue)
		appConfig.BreakTime.Set(breakTimeValue)
		// Push notification binding is directly linked and doesn't need parsing.

		// Persist the updated configuration.
		err = config.SaveConfig("config.json", appConfig) // Ensure SaveConfig accepts a Config type with bindings.
		handleErr(err)

		fmt.Println("Settings saved successfully.")
	})

	// Layout.
	return container.NewVBox(
		widget.NewForm(
			widget.NewFormItem("Session Time (minutes):", sessionEntry),
			widget.NewFormItem("Break Time (minutes):", breakEntry),
			widget.NewFormItem("", pushNotificationCheck),
		),
		saveButton,
	)
}
