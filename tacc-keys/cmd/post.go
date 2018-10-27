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
	"io/ioutil"
	"os"

	"github.com/TACC-Cloud/ssh-keys-client/tacc-keys/tacc-services"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	createKeys bool
	keyName    string
)

// postCmd represents the post command
var postCmd = &cobra.Command{
	Use:   "post [username]",
	Short: "Post a public ssh key to TACC's keys service",
	Long: `Post a public ssh key to TACC's keys service.
You can also create a public and provate key pair and post the public key to 
the service.`,
	Args: cobra.MaximumNArgs(1),
	PersistentPreRun: func(cmd *cobra.Command, agrs []string) {
		if err := refreshTokenPair(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		var username string
		if len(args) == 0 {
			username = viper.GetString("username")
		} else {
			username = args[0]
		}
		baseURL := viper.GetString("baseurl")
		accessToken := viper.GetString("access_token")

		// Create rsa key pair.
		if createKeys {
			if err := services.SaveRSAKeysToFile(keyName); err != nil {
				fmt.Printf("Error creating RSA keys: %s\n", err)
				os.Exit(1)
			}
		}

		// Read public key.
		pubKey, err := ioutil.ReadFile(keyName + ".pub")
		if err != nil {
			fmt.Printf("Error opening pubkey file: %s\n", err)
			os.Exit(1)
		}
		publicKey := string(pubKey)

		// Post public key to service.
		err = services.PostUserPubKey(baseURL, accessToken, username, publicKey)
		if err != nil {
			fmt.Printf("Error posting key: %s\n", err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(postCmd)

	postCmd.Flags().BoolVarP(
		&createKeys, "create", "c", false, "Create a pair of public and private ssh keys")
	postCmd.Flags().StringVarP(
		&keyName, "key-name", "k", "id_rsa", "Name of key (no extension)")
}
