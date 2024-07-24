/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/orangbus/cmd/console"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var (
	spider_url string // 爬取的地址
	imgRegex   = regexp.MustCompile(`https?://[^\s]+?\.(jpg|jpeg|png|gif|bmp)`)
	videoRegex = regexp.MustCompile(`https?://[^\s]+?\.(mp4|avi|mov|wmv|flv|mkv)`)
	docRegex   = regexp.MustCompile(`https?://[^\s]+?\.(doc|docx|xls|xlsx|pdf|zip)`)
)

// regexCmd represents the regex command
var regexCmd = &cobra.Command{
	Use:   "regex",
	Short: "提取文章中的图片，下载文件，图片，视频，doc,excel,pdf,zip",
	Run: func(cmd *cobra.Command, args []string) {
		list := parseDoc()
		total := len(list)
		if total == 0 {
			console.Success("未匹配到下载资源")
			return
		}
		log.Printf("累计下载条数：%d", total)
		for i, download_url := range list {
			go download(download_url)
			log.Printf("已经下载：%d/%d  %.2f%s", i+1, total, float64(i+1)/float64(total)*100, "%")
		}
		console.Success("download all done")
	},
}

func init() {
	rootCmd.AddCommand(regexCmd)
	regexCmd.Flags().StringVarP(&spider_url, "url", "u", spider_url, "采集地址")
}

func parseDoc() []string {
	list := []string{}
	content := ""
	if spider_url == "" {
		file, err := os.Open("demo.html")
		if err != nil {
			console.Error(err.Error())
			return list
		}
		defer file.Close()
		all, err := io.ReadAll(file)
		if err != nil {
			console.Error(err.Error())
			return list
		}
		content = string(all)
	} else {
		response, err := http.Get(spider_url)
		if err != nil {
			console.Error(err.Error())
			return nil
		}
		defer response.Body.Close()

		all, err := io.ReadAll(response.Body)
		if err != nil {
			console.Error(err.Error())
			return list
		}
		content = string(all)
	}

	images := extractLinks(content, imgRegex)
	videos := extractLinks(content, videoRegex)
	docs := extractLinks(content, docRegex)

	fmt.Println("Image Links:")
	for _, link := range images {
		fmt.Println(link)
		list = append(list, link)
	}

	fmt.Println("\nVideo Links:")
	for _, link := range videos {
		fmt.Println(link)
		list = append(list, link)
	}

	fmt.Println("\nDocument Links:")
	for _, link := range docs {
		fmt.Println(link)
		list = append(list, link)
	}
	return list
}

func extractLinks(content string, regex *regexp.Regexp) []string {
	matches := regex.FindAllString(content, -1)
	return matches
}

func download(download_url string) {
	index := strings.LastIndex(download_url, "/")
	fileName := download_url[index+1:]

	response, err := http.Get(download_url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	// 下载链接地址文件
	outFile, err := os.Create(fmt.Sprintf("images/%s", fileName))
	if err != nil {
		console.Error(err.Error())
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, response.Body)
	if err != nil {
		return
	}
}

func uploadMinio() {

}
