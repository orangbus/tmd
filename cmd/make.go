/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/orangbus/cmd/console"
	"strings"

	"github.com/spf13/cobra"
)

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "随机生成一个随机数，Flags:key,uuid",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			console.Error("请输入参数,Flags:key,uuid")
			return
		}
		switch args[0] {
		case "key":
			generateKey()
		case "uuid":
			generateUuid()
		default:
			console.Error("参数错误")
		}
		fmt.Println(args)
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func generateKey() {
	console.Success(strings.Repeat("---", 10))
	console.Success(fmt.Sprintf("key: %s", uuid.NewString()))
	console.Success(strings.Repeat("---", 10))
}

func generateUuid() {
	console.Success(strings.Repeat("---", 10))
	console.Success(fmt.Sprintf("UUID: %s", uuid.New().String()))
	console.Success(strings.Repeat("---", 10))
}
