/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/orangbus/cmd/app/service/spider_serve"
	"github.com/spf13/cobra"
)

// spiderCmd represents the spider command
var spiderCmd = &cobra.Command{
	Use:   "spider",
	Short: "采集文章",
	Run: func(cmd *cobra.Command, args []string) {
		spider_serve.Start()
	},
}

func init() {
	rootCmd.AddCommand(spiderCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// spiderCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// spiderCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
