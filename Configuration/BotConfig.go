package Configuration

import (
	"encoding/json"
)

//BotConfig hold general configuration
type BotConfig struct {
	BotToken     string `json:"bot_token"`
	BotPrefix    string `json:"bot_prefix"`
	DatabaseIP   string `json:"database_ip"`
	DatabaseUser string `json:"database_user"`
	DatabasePass string `json:"database_password"`
	DatabasePort string `json:"database_port"`
	DatabaseName string `json:"database_name"`
	SSLMode      string `json:"ssl_mode"`
}

//LoadConfig loads the main config file for the application
func (config *BotConfig) LoadConfig(fileBytes []byte) error {
	err := json.Unmarshal(fileBytes, &config)
	if err != nil {
		return err
	}

	return nil
}
