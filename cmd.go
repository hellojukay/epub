package main

import (
	"github.com/spf13/cobra"
)

var rootCmd = cobra.Command{}
var indexCmd = cobra.Command{
	Long: "index",
	Use:  "index",
	Run:  Index,
}
var searchCmd = cobra.Command{
	Long: "search",
	Use:  "search",
	Run:  Search,
}

func init() {
	rootCmd.AddCommand(&indexCmd, &searchCmd)
}

func Run() {
	rootCmd.Execute()
}
