package configuration

import (
	"encoding/json"
)

//BotConfig hold general configuration
type BotConfig struct {
	BotToken  string `json:"bot_token"`
	BotPrefix string `json:"bot_prefix"`
}

//LoadConfig loads the main config file for the application
func (config *BotConfig) LoadConfig(fileBytes []byte) error {
	err := json.Unmarshal(fileBytes, &config)
	if err != nil {
		return err
	}

	return nil
}
