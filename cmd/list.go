// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yarbelk/grabtrello/trello"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the boards",
	Long:  `List the boards you have access too`,
	Run: func(cmd *cobra.Command, args []string) {
		key, token := viper.GetString("key"), viper.GetString("token")
		var userName string
		if len(args) != 1 {
			userName = viper.GetString("user")
		} else {
			userName = args[0]
		}

		user, err := trello.Member(userName, key, &token)
		if err != nil {
			log.Fatal(err)
		}

		boards, err := user.Boards()
		if err != nil {
			log.Fatal(err)
		}
		for _, board := range boards {
			fmt.Printf("* %v (%v)\n", board.Name, board.ShortUrl)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
