package main

import (
	"encoding/json"
	"os"
)

// Configurations stores the credentials necessary to interact with the Agave
// API. This file corresponds to the agave cache dir (~/.agave/current).
type Configurations struct {
	TenantId     string `mapstructure:"tenantid" json:"tenantid"`
	BaseUrl      string `mapstructure:"baseurl" json:"baseurl"`
	ApiSecret    string `mapstructure:"apisecret" json:"apisecret"`
	ApiKey       string `mapstructure:"apikey" json:"apikey"`
	Username     string `mapstructure:"username" json:"username"`
	AccessToken  string `mapstructure:"access_token" json:"access_token"`
	RefreshToken string `mapstructure:"refresh_token" json:"refresh_token"`
	CreatedAt    string `mapstructure:"created_at" json:"created_at"`
	ExpiresIn    string `mapstructure:"expires_in" json:"expires_in"`
	ExpiresAt    string `mapstructure:"expires_at" json:"expires_at"`

	ConfigFile string
}

// SaveConfig updates the value of the configuration file based on the
// contents fo the Configurations struct.
func (c *Configurations) SaveConfig() error {
	// Open config file.
	configFile, err := os.OpenFile(c.ConfigFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}
	defer configFile.Close()

	// Write values to file.
	encoder := json.NewEncoder(configFile)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(c); err != nil {
		return err
	}

	return err
}
