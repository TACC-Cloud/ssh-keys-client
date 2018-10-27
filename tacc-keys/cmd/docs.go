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

	"github.com/spf13/cobra"
    "github.com/spf13/cobra/doc"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "docs",
	Short: "Generate bash completion scripts and documentation",
	Long: `docs will generate bash completion scripts, write them to standard out
and generate markdown documentation in your current directory.

To load completion run
    
    . <(tacc-keys completion)

To configure your bash shell to load completions for each session add to your bashrc

    # ~/.bashrc or ~/.profile
    . <(tacc-keys completion)
    `,
	Run: func(cmd *cobra.Command, args []string) {
        rootCmd.GenBashCompletion(os.Stdout);
        if err := doc.GenMarkdownTree(rootCmd, "./"); err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
    },
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
