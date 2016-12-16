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
	"os"
	"path"

	mytrello "github.com/VojtechVitek/go-trello"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yarbelk/grabtrello/trello"
)

func writeBoards(outputDir string, boards []mytrello.Board) error {
	if _, err := os.Stat(outputDir); err != nil {
		return err
	}

	os.Chdir(outputDir)
	fd, err := os.Create("index.md")
	defer fd.Close()
	if err != nil {
		return err
	}

	for _, board := range boards {
		os.Mkdir(board.Name, os.ModeDir|os.ModePerm)
		fmt.Fprintf(fd, "- [%s](%s/index.md)\n", board.Name, board.Name)
		if err := writeLists(path.Join(outputDir, board.Name), board); err != nil {
			return err
		}
	}
	return nil
}

func writeLists(outputDir string, board mytrello.Board) error {
	if _, err := os.Stat(outputDir); err != nil {
		return err
	}

	fd, err := os.Create(path.Join(outputDir, "index.md"))
	defer fd.Close()
	if err != nil {
		return err
	}

	lists, err := board.Lists()
	if err != nil {
		return err
	}

	fmt.Fprintf(fd, "# %s\n", board.Name)

	for _, list := range lists {
		fmt.Fprintf(fd, "\n# %s\n", list.Name)
		cards, err := list.Cards()
		if err != nil {
			return err
		}
		for _, card := range cards {
			fmt.Fprintf(fd, "- [%s](%s.md)\n", card.Name, path.Join(list.Name, card.Name))
			writeCard(outputDir, list, card)
		}
	}
	return nil
}

func writeCard(outputDir string, list mytrello.List, card mytrello.Card) error {
	if _, err := os.Stat(outputDir); err != nil {
		return err
	}
	outputDir = path.Join(outputDir, list.Name)

	os.Mkdir(outputDir, os.ModeDir|os.ModePerm)
	fd, err := os.Create(path.Join(outputDir, fmt.Sprintf("%s.md", card.Name)))
	defer fd.Close()
	if err != nil {
		return err
	}

	fmt.Fprintf(fd, "# %s\n\n", card.Name)
	fmt.Fprintf(fd, "link: %s\n\n", card.ShortUrl)
	fmt.Fprintf(fd, "## Desc\n\n%s\n", card.Desc)

	return nil
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, token := viper.GetString("key"), viper.GetString("token")
		var userName string
		if len(args) != 2 {
			log.Fatal("need 2 args")
		}
		userName = args[0]
		outputDir := args[1]

		user, err := trello.Member(userName, key, &token)
		if err != nil {
			log.Fatal(err)
		}

		boards, err := user.Boards()
		if err != nil {
			log.Fatal(err)
		}
		writeBoards(outputDir, boards)
	},
}

func init() {
	RootCmd.AddCommand(getCmd)
}
