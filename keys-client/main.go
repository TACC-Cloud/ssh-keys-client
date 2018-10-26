package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	keyId        string
	username     string
	keysFilename string

	createKeys bool
	deleteKey  bool
	listKeys   bool
)

func init() {
	// Sources.
	pflag.StringVarP(&keyId, "id", "i", "", "key id")
	pflag.StringVarP(&username, "user", "u", "", "owner of public key")
	pflag.StringVarP(&keysFilename, "keys", "k", "id_rsa", "Name of file to save rsa keys to")

	// Operations.
	pflag.BoolVarP(&createKeys, "create", "c", false, "create ssh keys using rsa")
	pflag.BoolVarP(&deleteKey, "delete", "d", false, "delete a key given its id")
	pflag.BoolVarP(&listKeys, "list", "l", false, "list all keys for user")
}

func main() {
	// Parse command line arguments.
	pflag.Parse()

	// Read config file.
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME")
	var conf Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading configuration file: %s\n", err)
		os.Exit(1)
	}

	if err := viper.Unmarshal(&conf); err != nil {
		fmt.Printf("Error decoding into struct: %s\n", err)
		os.Exit(1)
	}
	conf.ConfigFile = viper.ConfigFileUsed()

	// Refresh Token.
	createdAt, err := strconv.ParseInt(conf.CreatedAt, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	ExpiresIn, err := strconv.ParseInt(conf.ExpiresIn, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	now := time.Now().Unix() - 100
	// Check if token needs to be refreshed.
	if (createdAt + ExpiresIn) < now {
		fmt.Fprintln(os.Stderr, "Refreshing token...")
		if err := conf.RefreshAPIToken(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Create a pair of public and provate rsa keys for ssh.
	if createKeys {
		// Check keys filename is set.
		if keysFilename == "" {
			os.Exit(1)
		}
		// If username is not set, set it to the username from the config file.
		if username == "" {
			username = conf.Username
		}
		// Create rsa key pair.
		if err := SaveRSAKeysToFile(keysFilename); err != nil {
			fmt.Printf("Error creating RSA keys: %s\n", err)
			os.Exit(1)
		}

		// Post public keys to key server.
		publicKey, err := ioutil.ReadFile(keysFilename + ".pub")
		if err != nil {
			fmt.Printf("Error opening pubkey file: %s\n", err)
			os.Exit(1)
		}
		if err := conf.PostUserPubKey(username, string(publicKey)); err != nil {
			fmt.Printf("Error posting public key to server: %s\n", err)
			os.Exit(1)
		}
	} else if deleteKey { // Delete public key from key server.
		// Check key ID is set.
		if keyId == "" {
			os.Exit(1)
		}
		// Delete public key from server.
		if err = conf.DeletePubKey(keyId); err != nil {
			fmt.Println("Error deleting keys: %s\n", err)
			os.Exit(1)
		}
	} else if listKeys { // List all public keys for a user.
		// Check username is set.
		if username == "" {
			os.Exit(1)
		}
		// List all public keys for a user.
		if err := conf.GetUserPubKeys(username); err != nil {
			fmt.Println("Error listing keys: %s\n", err)
			os.Exit(1)
		}
	}
}
