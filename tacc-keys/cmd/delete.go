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

	"github.com/TACC-Cloud/ssh-keys-client/tacc-keys/tacc-services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [key ID]",
	Short: "Delete a public ssh key from TACC's keys service",
	Long: `Delete a public ssh key from TACC's keys service. You can use the 
command "list" if you need to figure out the ID for the key you wish to delete.`,
	Args: cobra.ExactArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, agrs []string) {
		if err := refreshTokenPair(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		keyID := args[0]
		baseURL := viper.GetString("baseurl")
		accessToken := viper.GetString("access_token")

		// Delete public key from server.
		if err := services.DeletePubKey(baseURL, accessToken, keyID); err != nil {
			fmt.Printf("Error deleting key: %s\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
