package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/adilJamshad/harvestTime/internal/config"
)

func RunApp() {
	appConfig, _ := config.LoadConfig("config.json") // Adjust LoadConfig to work with binding.
	myApp := app.New()
	myWindow := myApp.NewWindow("Harvest Time")

	tabs := container.NewAppTabs(
		container.NewTabItem("Timer", TimerTab(appConfig)), // Convert timer_tab to *widget.TabItem
		container.NewTabItem("ToDo List", widget.NewLabel("ToDo List Content")),
		container.NewTabItem("Config", widget.NewLabel("Config Content")),
		container.NewTabItem("Settings", SettingsTab(appConfig)),
	)
	tabs.SelectIndex(0)
	myWindow.SetContent(tabs)

	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.ShowAndRun()
}
