package main

import "github.com/spf13/cobra"

func main() {

}

var (
	root = cobra.Command{
		Use:   "cafe_manager",
		Short: "Cafe manage program",
		Long:  "Cafe manager program",
	}
)
