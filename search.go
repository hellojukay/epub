package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func search(name string) ([]Epub, error) {
	var books []Epub
	var raw = "select * from epubs where title like ?"
	err := store.Raw(raw, `%`+name+`%`).Scan(&books).Error
	return books, err
}
func Search(cmd *cobra.Command, args []string) {
	var name = cmd.Flags().String("name", "python", "book name")
	cmd.ParseFlags(args)
	books, err := search(`国`)
	if err != nil {
		log.Printf("检索: %s 失败 %s", *name, err)
		os.Exit(1)
	}
	for _, book := range books {
		fmt.Printf("%-30s%s\n", book.Title, book.Path)
	}
}
