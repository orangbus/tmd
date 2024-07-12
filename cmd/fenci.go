/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/go-ego/gse"
	"github.com/spf13/cobra"
	"log"
)

var (
	seg gse.Segmenter
	//posSeg pos.Segmenter

	//new, _ = gse.New("zh,testdata/test_en_dict3.txt", "alpha")

	text = "你好世界, Hello world, Helloworld."
)

// fenciCmd represents the fenci command
var fenciCmd = &cobra.Command{
	Use:   "fenci",
	Short: "中文分词",
	Run: func(cmd *cobra.Command, args []string) {
		list := seg.Cut(text, true)
		log.Println(list)
	},
}

func init() {
	rootCmd.AddCommand(fenciCmd)
	// fenciCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
