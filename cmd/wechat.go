/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/orangbus/cmd/console"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type WeChatArticle struct {
	Title    string `json:"title"`    // 文章标题
	Content  string `json:"content"`  // 文章内筒
	Platform string `json:"platform"` // 平台
	PushAt   string `json:"push_at"`  // 发布时间
	Url      string `json:"url"`      // 地址
	Area     string `json:"area"`     //地区
}

// wechatCmd represents the wechat command
var wechatCmd = &cobra.Command{
	Use:   "wechat",
	Short: "微信公众号文章采集",
	Run: func(cmd *cobra.Command, args []string) {
		url := "https://mp.weixin.qq.com/s/ew8YkLHnAWaBnIIUjkrPWw"
		if len(args) > 0 {
			url = args[0]
		}
		if url == "" {
			console.Error("请输入链接地址")
			return
		}
		log.Println("采集地址：", url)
		article, err := spider(url)
		if err != nil {
			console.Error(err.Error())
			return
		}
		log.Println(article.Title)
		console.Success("done")
	},
}

func init() {
	rootCmd.AddCommand(wechatCmd)
	// wechatCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func spider(url string) (WeChatArticle, error) {
	article := WeChatArticle{}
	response, err := http.Get(url)
	if err != nil {
		return article, err
	}
	if response.StatusCode != 200 {
		return article, errors.New("请求错误")
	}
	defer response.Body.Close()

	return parseHtml(response.Body, url)
}

func parseHtml(doc io.Reader, url string) (WeChatArticle, error) {
	article := WeChatArticle{
		Url: url,
	}

	reader, err := goquery.NewDocumentFromReader(doc)
	if err != nil {
		return WeChatArticle{}, err
	}
	article.Title = strings.TrimSpace(reader.Selection.Find(".rich_media_title").Text())
	article.Platform = strings.TrimSpace(reader.Selection.Find("#js_name").Text())
	article.Content = strings.TrimSpace(reader.Selection.Find(".rich_media_content").Text())
	return article, nil
}

func testParse() {
	file, err := os.Open("wechat.html")
	if err != nil {
		return
	}
	defer file.Close()
	article, err := parseHtml(file, "")
	if err != nil {
		return
	}
	log.Println(article.Title)
	log.Println(article.Platform)
	log.Println(article.Content)
}
