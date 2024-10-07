package main

import (
	"fmt"
	"go-template/internal/config"
	"go-template/utils"
	"github.com/spf13/cobra"
)

func main() {
	cfg := config.LoadConfig()

	var name string

	var rootCmd = &cobra.Command{
		Use:   "helloapp",
		Short: "A simple app that greets people",
		Run: func(cmd *cobra.Command, args []string) {
			if name == "" {
				name = cfg.DefaultName
			}
			greeting := utils.Greet(name)
			fmt.Println(greeting)
		},
	}

	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the person to greet")
	rootCmd.Execute()
}
