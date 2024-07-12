/*
https://github.com/grafov/m3u8
*/
package cmd

import (
	"bufio"
	"fmt"
	"github.com/grafov/m3u8"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	url2 "net/url"
	"strings"
)

// var url = "https://svipsvip.ffzy-online5.com/20240711/29853_a88ef24d/index.m3u8"
var url = "https://svipsvip.ffzy-online5.com/20240711/29853_a88ef24d/2000k/hls/mixed.m3u8"

func init() {
	rootCmd.AddCommand(m3u8Cmd)
	//m3u8Cmd.Flags().StringVarP(&url, "url", "u", "", "m3u8视频地址")
}

// m3u8Cmd represents the m3u8 command
var m3u8Cmd = &cobra.Command{
	Use:   "m3u8",
	Short: "m3u8视频解析",
	Run: func(cmd *cobra.Command, args []string) {
		if url == "" {
			log.Println("请输入m3u8地址")
			return
		}
		// 得到全部的下载链接
		list, err := parseUrl(url)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(list)

		// 下载分片内容

		// 合并分片

	},
}

func parseUrl(url string) ([]string, error) {
	var list = []string{}
	response, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return list, err
	}
	defer response.Body.Close()

	p, listType, err := m3u8.DecodeFrom(bufio.NewReader(response.Body), true)
	if err != nil {
		log.Println(err)
		return list, err
	}
	switch listType {
	case m3u8.MEDIA: // 直接播放的列表
		mediapl := p.(*m3u8.MediaPlaylist)
		fmt.Printf("%+v\n", mediapl)
		log.Println(strings.Repeat("-", 10))
		for _, v := range mediapl.Segments {
			if v != nil {
				u, _ := getFullUrl(url, v.URI)
				list = append(list, u)
			}
		}
		return list, nil
	case m3u8.MASTER: // 直播列表
		realUrl := ""
		masterpl := p.(*m3u8.MasterPlaylist)
		variants := masterpl.Variants
		for k, item := range variants {
			if k == 0 {
				realUrl = item.URI
				break
			}
		}
		//fmt.Printf("%+v\n", masterpl)
		log.Println(realUrl)
		log.Println(strings.Repeat("=", 10))
		u, _ := getFullUrl(url, realUrl)
		return parseUrl(u)
	}
	return list, nil
}

/*
*
1、m3u8 地址解析
【http 开头】直接范围
【/】 url + path
*/
func getFullUrl(m3u8Url, path string) (string, error) {
	parse, err := url2.Parse(m3u8Url)
	if err != nil {
		return "", err
	}
	host := fmt.Sprintf("%s://%s", parse.Scheme, parse.Hostname()) //  https://svipsvip.ffzy-online5.com
	length := strings.LastIndex(m3u8Url, "/")
	preUrl := m3u8Url[:length] //  https://svipsvip.ffzy-online5.com/20240711/29853_a88ef24d
	//log.Println(host)
	//log.Println(preUrl)

	full_url := ""
	if strings.HasPrefix(path, "/") {
		full_url = fmt.Sprintf("%s%s", host, path)
	} else {
		full_url = fmt.Sprintf("%s/%s", preUrl, path)
	}
	return full_url, err
}
