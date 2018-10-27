package services

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// Credentials ask for username and password at runtime.
func Credentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}

	fmt.Print("Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", "", err
	}
	password := string(bytePassword)

	return strings.TrimSpace(username), strings.TrimSpace(password), nil
}
