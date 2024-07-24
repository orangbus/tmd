/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/withlin/canal-go/client"
)

// canalCmd represents the canal command
var canalCmd = &cobra.Command{
	Use:   "canal",
	Short: "mysql binlog 解析",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("canal called")
	},
}

func init() {
	rootCmd.AddCommand(canalCmd)
	// canalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func parse() {
	conn := client.NewSimpleCanalConnector("127.0.0.1", 3306, "root", "admin666", "test", 60000, 60*60*1000)
	err := conn.Connect()
	if err != nil {
		panic(err)
	}

	// 订阅
	err = conn.Subscribe("test.*")
	if err != nil {
		panic(err)
	}

	for {
		message, err := conn.Get(100, nil, nil)
		if err != nil {
			panic(err)
		}
		fmt.Println(message)
	}
}
