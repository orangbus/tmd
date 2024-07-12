/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/orangbus/cmd/pkg/notice"
	"github.com/spf13/cobra"
	"log"
)

const notice_url = "https://open.feishu.cn/open-apis/bot/v2/hook/ea48b504-03d5-40bc-a8b5-5b999bb11fe9"

// noticeCmd represents the notice command
var noticeCmd = &cobra.Command{
	Use:   "notice",
	Short: "消息通知",
	Run: func(cmd *cobra.Command, args []string) {
		msg := notice.NewFeiShuNotice(notice_url)
		err := msg.SeedCard("消息通知", "在群组中创建，只能在群组中使用的特殊机器人。完成开发与配置后，你可以调用自定义机器人的 webhook 地址，以机器人的身份向群组中自动推送来自外部系统的消息。具体信息请参考自定义机器人使用指南。")
		if err != nil {
			log.Println(err.Error())
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(noticeCmd)
	// noticeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
