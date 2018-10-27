// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/TACC-Cloud/ssh-keys-client/tacc-keys/tacc-services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	name        string
	description string
)

// subscribeCmd represents the subscribe command
var subscribeCmd = &cobra.Command{
	Use:   "subscribe",
	Short: "Create an oauth client and subscribe to TACC's keys service",
	Long: `Create a TACC oauth client, obtain an access token, and subscribe the 
client to the keys service.
This will also generate a config file if one doesn't already exist.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set name to hostname if not defined.
		if name == "" {
			var err error
			name, err = os.Hostname()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		baseURL := "https://api.tacc.utexas.edu"

		// Ge user credentials.
		username, password, err := services.Credentials()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create an oauth client.
		key, secret, err := services.CreateClient(
			baseURL, name, description, username, password)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Subscribe client to TACC's keys service.
		err = services.SubscribeClient(baseURL, name, username, password)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		access, refresh, err := services.GetTokens(
			baseURL, key, secret, username, password)

		// Save configurations.
		viper.Set("apikey", key)
		viper.Set("apisecret", secret)
		viper.Set("baseurl", baseURL)
		viper.Set("username", username)
		viper.Set("access_token", access)
		viper.Set("refresh_token", refresh)
		viper.Set("created_at", strconv.FormatInt(time.Now().Unix(), 10))
		viper.WriteConfigAs(".tacc-keys.yaml")
	},
}

func init() {
	rootCmd.AddCommand(subscribeCmd)

	subscribeCmd.Flags().StringVarP(&name, "name", "n", "", "Name of aouth client")
	subscribeCmd.Flags().StringVarP(
		&description, "description", "d", "", "Oauth lient description")
}
