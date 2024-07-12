/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var file = "site.txt"
var site = ""

// siteokCmd represents the siteok command
var siteokCmd = &cobra.Command{
	Use:   "siteok",
	Short: "检查一个网站返回是否正常",
	Run: func(cmd *cobra.Command, args []string) {
		if file != "" {

		}

		if site != "" {

		}
		fmt.Println("执行结束")
	},
}

func init() {
	rootCmd.AddCommand(siteokCmd)

	siteokCmd.Flags().StringVarP(&file, "file", "f", file, "指定一个 txt文件")
	siteokCmd.Flags().StringVarP(&site, "site", "s", file, "指定一个url地址，多个地址使用 , 分隔")
}
