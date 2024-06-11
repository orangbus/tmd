/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/levigross/grequests"
	"github.com/orangbus/cmd/console"
	"net/url"
	"time"

	"github.com/spf13/cobra"
)

// baiduCmd represents the baidu command
var baiduCmd = &cobra.Command{
	Use:   "baidu",
	Short: "百度收录查询",
	Run: func(cmd *cobra.Command, args []string) {
		urlStr := cmd.Flag("url").Value.String()
		parse, err := url.Parse(urlStr)
		if err != nil {
			console.Error(err.Error())
			return
		}
		query_url := fmt.Sprintf("https://www.baidu.com/s?wd=site:%s&tn=78040160_26_pg&ch=8", parse.Hostname())
		console.Info(query_url)
		ro := &grequests.RequestOptions{DialTimeout: 10 * time.Second}
		response, err := grequests.Get(query_url, ro)
		if err != nil {
			console.Error(err.Error())
			return
		}
		fmt.Println(response.String())
	},
}

func init() {
	rootCmd.AddCommand(baiduCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// baiduCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	baiduCmd.Flags().StringP("url", "u", "", "目标网站")
}
