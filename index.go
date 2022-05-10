package main

import (
	"io/fs"
	"log"
	"path/filepath"
	"strings"

	"github.com/sebojanko/epub/epub"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite" // Sqlite driver based on GGO

	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	"gorm.io/gorm"
)

var store *gorm.DB

func init() {

	// github.com/mattn/go-sqlite3
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	store = db
	store.AutoMigrate(Epub{})
}
func index(dir string) {
	//list directory
	filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if filepath.Ext(path) != ".epub" {
			return nil
		}

		bookInfo := epub.GetMetadata(path)
		descript := strings.Trim(bookInfo.Metadata.Description, " ")
		descript = strings.TrimSuffix(descript, "</p>")
		descript = strings.TrimPrefix(descript, `<p class="description">`)
		descript = strings.TrimPrefix(descript, `<div>`)
		descript = strings.TrimSuffix(descript, "</div>")

		book := Epub{
			Path:     path,
			Title:    bookInfo.Metadata.Title,
			Descript: descript,
		}
		if e := book.Insert(); e != nil {
			log.Printf("无法保存书籍 %s %s", book.Title, err)
		}
		return nil
	})
}

func Index(cmd *cobra.Command, args []string) {
	var dir = cmd.Flags().String("dir", "D:/data/", "book directory")
	cmd.ParseFlags(args)
	index(*dir)
}
