/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/jdkato/prose/v2"
	"github.com/spf13/cobra"
	"log"
)

// proseCmd represents the prose command
var proseCmd = &cobra.Command{
	Use:   "prose",
	Short: "关键词提取",
	Run: func(cmd *cobra.Command, args []string) {
		article := "文件上传大小限制：当前 WebDAV 客户端和网页端上传大小的限制是一致的，默认为 500M（私有云可以通过相关设置调整）。\n\n访问频率限制：由于WebDAV协议比较占用系统资源，免费版用户限制访问频率为每30分钟不超过600次请求。付费用户限制访问频率为每30分钟不超过1500次请求。\n\n同步目录数限制：目前坚果云的WebDAV协议单次请求文件数（包含文件和文件夹）为750个，支持分多页多次加载。如果您使用WebDAV的三方工具未实现按分页多次加载，可能会出现文件同步不完整的情况，建议您使用坚果云客户端进行直接同步。\n\n该协议也可以用于定制企业内部的其他应用和服务访问存储在坚果云的文件，常见的应用包括邮件服务，OA服务或者CRM服务等。\n\n坚果云也可以为您的企业提供相关集成方案的有偿咨询顾问服务，请联系我们并告知您的需求。"
		document, err := prose.NewDocument(article)
		if err != nil {
			log.Println(err)
			return
		}
		tokens := document.Tokens()
		for _, item := range tokens {
			log.Println(item.Text)
		}

	},
}

func init() {
	rootCmd.AddCommand(proseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// proseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// proseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
