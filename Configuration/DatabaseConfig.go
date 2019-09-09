package configuration

import "encoding/json"

//DatabaseConfig holds our database struct
type DatabaseConfig struct {
	DatabaseIP   string `json:"database_ip"`
	DatabaseUser string `json:"database_user"`
	DatabasePass string `json:"database_password"`
	DatabasePort string `json:"database_port"`
	DatabaseName string `json:"database_name"`
	SSLMode      string `json:"ssl_mode"`
}

//LoadConfig Unmarshals the json passed to it into a DatabaseConfig struct
func (config *DatabaseConfig) LoadConfig(fileBytes []byte) error {
	err := json.Unmarshal(fileBytes, &config)
	if err != nil {
		return err
	}

	return nil
}

//CreateDatabaseString is a helper method for creating the initial connection string.
func (config *DatabaseConfig) CreateDatabaseString() string {
	return "postgres://" +
		config.DatabaseUser + ":" +
		config.DatabasePass + "@" +
		config.DatabaseIP + ":" +
		config.DatabasePort + "/" +
		config.DatabaseName + "?sslmode=" +
		config.SSLMode
}
