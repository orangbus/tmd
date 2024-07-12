/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"github.com/orangbus/cmd/app/models"
	"github.com/orangbus/cmd/pkg/database"
	"github.com/orangbus/cmd/pkg/debug"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// tmd import
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "导入csv文件",
	Run: func(cmd *cobra.Command, args []string) {
		//testFile("F:\\share\\Mongo\\zjzhangui.csv")
		//return

		openFile("F:\\share\\Mongo\\zjzhangui.csv", 6, 2, 3)

		err := openFile("F:\\share\\Mongo\\diyijuzi.csv", 6, 2, 3)
		if debug.HasError("导入失败", err) {

		}

		debug.Info("'done")
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
	// importCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func testFile(path string) {
	file, err := os.Open(path)
	if debug.HasError("文件打开失败", err) {
		return
	}
	defer file.Close()

	list, err := csv.NewReader(file).ReadAll()
	if debug.HasError("文件读取失败", err) {
		return
	}
	for k, v := range list {
		if k > 0 {
			article := models.Articles{
				Name:     v[6],
				CateName: v[2],
				Content:  v[3],
			}
			log.Println(article)
			break
		}
	}
}

func openFile(path string, name, cateName, content int) error {
	file, err := os.Open(path)
	if debug.HasError("文件打开失败", err) {
		return err
	}
	defer file.Close()

	list, err := csv.NewReader(file).ReadAll()
	if debug.HasError("文件读取失败", err) {
		return err
	}
	return SaveDataBase(list, name, cateName, content)
}
func SaveDataBase(list [][]string, name, cateName, content int) error {
	var data []models.Articles
	for k, v := range list {
		if k > 0 {
			article := models.Articles{
				Name:     v[name],
				CateName: v[cateName],
				Content:  v[content],
			}
			data = append(data, article)
		}
	}

	err := database.DB.Model(&models.Articles{}).Create(&data).Error
	if debug.HasError("数据导入失败", err) {
		return err
	}
	debug.Info("导入成功！")
	return nil
}
