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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [username]",
	Short: "List all public ssh keys registered to a user",
	Long: `List all public ssh keys registered to a user. The keys will be printed
using the following format:
    
    key-id:int tag:string
    ssh-pubkey:string`,
	Args: cobra.ExactArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, agrs []string) {
		if err := refreshTokenPair(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		baseURL := viper.GetString("baseurl")
		accessToken := viper.GetString("access_token")

		err := services.GetUserPubKeys(baseURL, accessToken, username)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
