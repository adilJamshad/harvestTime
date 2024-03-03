package config

import (
	"encoding/json"
	"os"
	"sync"

	"fyne.io/fyne/v2/data/binding"
)

// Config represents the application configuration with Fyne data bindings.
type Config struct {
	SessionTime       binding.Int
	BreakTime         binding.Int
	PushNotifications binding.Bool
}

// NewConfig creates a new instance of Config with initialized data bindings.
func NewConfig() *Config {
	return &Config{
		SessionTime:       binding.NewInt(),
		BreakTime:         binding.NewInt(),
		PushNotifications: binding.NewBool(),
	}
}

// configData is an intermediary struct for JSON operations.
type configData struct {
	SessionTime       int  `json:"sessionTime"`
	BreakTime         int  `json:"breakTime"`
	PushNotifications bool `json:"pushNotifications"`
}

var lock = &sync.Mutex{}
var conf *Config

// LoadConfig loads the configuration from the given file path into a Config instance.
func LoadConfig(filePath string) (*Config, error) {
	lock.Lock()
	defer lock.Unlock()

	if conf == nil {
		conf = NewConfig()
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data configData
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}

	// Set the loaded values into the Fyne data bindings.
	conf.SessionTime.Set(data.SessionTime)
	conf.BreakTime.Set(data.BreakTime)
	conf.PushNotifications.Set(data.PushNotifications)

	return conf, nil
}

// SaveConfig saves the configuration data from a Config instance to the given file path.
func SaveConfig(filePath string, data *Config) error {
	lock.Lock()
	defer lock.Unlock()

	// Retrieve values from data bindings.
	sessionTime, _ := data.SessionTime.Get()
	breakTime, _ := data.BreakTime.Get()
	pushNotifications, _ := data.PushNotifications.Get()

	// Prepare the intermediary struct for JSON serialization.
	dataToSave := configData{
		SessionTime:       sessionTime,
		BreakTime:         breakTime,
		PushNotifications: pushNotifications,
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(dataToSave)
}
