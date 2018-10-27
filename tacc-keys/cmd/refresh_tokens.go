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
	"github.com/spf13/viper"
)

// refreshTokenPair checks if the access token needs to b refreshed, if it does
// then it makes a request to create a new access and refresh token pair.
func refreshTokenPair() error {
	apiKey := viper.GetString("apikey")
	apiSecret := viper.GetString("apisecret")
	refreshToken := viper.GetString("refresh_token")
	baseURL := viper.GetString("baseurl")
	createdAt := viper.GetInt64("created_at")
	expiresIn := viper.GetInt64("expires_in")

	now := time.Now().Unix() - 100
	if (createdAt + expiresIn) < now {
		fmt.Fprintln(os.Stderr, "Refreshing token...")
		access, refresh, err := services.RefreshToken(
			baseURL, refreshToken, apiKey, apiSecret)
		if err != nil {
			return err
		}
		viper.Set("access_token", access)
		viper.Set("refresh_token", refresh)
		viper.Set("created_at", strconv.FormatInt(time.Now().Unix(), 10))
		viper.WriteConfig()
	}

	return nil
}
