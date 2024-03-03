package main

import (
	"github.com/adilJamshad/harvestTime/internal/config"
	"github.com/adilJamshad/harvestTime/internal/eventManager"
	"github.com/adilJamshad/harvestTime/ui"
)

func main() {
	eventmanager := eventManager.NewEventManager()
	appConfig, _ := config.LoadConfig("config.json")
	ui.RunApp(eventmanager, appConfig)
}
