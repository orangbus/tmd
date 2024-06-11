/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/golang-module/carbon/v2"
	"github.com/orangbus/cmd/console"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	t   string
	dir string
	exd string
	ext string

	space    = 2 // 空格
	dateType = 1 // 1:指定时间。2：时间范围
	stime    time.Time
	etime    time.Time
	exdList  = []string{} // 排除目录
	extList  = []string{} // 指定文件后缀
	total    = 0          // 累计扫描的文件
	number   = 0          // 满足要求的文件
)

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "查询一个时间范围的修改文件",
	Run: func(cmd *cobra.Command, args []string) {
		timestart := carbon.Now().Timestamp()
		abs, err := filepath.Abs(dir)
		if err != nil {
			console.Error(err.Error())
			return
		}
		dir = abs

		stat, err2 := os.Stat(dir)
		if err2 != nil {
			console.Error(err2.Error())
			return
		}

		if !stat.IsDir() {
			console.Error("请输入正确的目录")
			return
		}

		// 时间处理
		if strings.Contains(t, ",") {
			dateType = 2
			dateBt := strings.Split(t, ",")
			stime = carbon.Parse(dateBt[0]).StdTime()
			etime = carbon.Parse(dateBt[1]).StdTime()
		} else {
			stime = carbon.Parse(t).StdTime()
		}

		if exd != "" {
			if strings.Contains(exd, ",") {
				extList = strings.Split(exd, ",")
			} else {
				extList = append(extList, dir)
			}
		}

		if ext != "" {
			if strings.Contains(ext, ",") {
				extList = strings.Split(ext, ",")
			} else {
				extList = append(extList, ext)
			}
		}

		fmt.Println(strings.Repeat("=", 50))
		fmt.Printf("文件目录：%s\n", dir)
		fmt.Printf("时间：%s\n", stime)
		fmt.Printf("排除目录：%s\n", exd)
		fmt.Printf("指定文件类型：%s\n", ext)
		fmt.Println(strings.Repeat("=", 50))
		scannerDir(dir, 0)
		duration := carbon.Now().Timestamp() - timestart
		fmt.Printf("累计扫描%d个文件，满足条件的文件有%d个,耗时：%d秒\n", total, number, duration)
	},
}

func init() {
	rootCmd.AddCommand(findCmd)

	t = carbon.Now().ToDateString()
	dirPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 接收参数
	findCmd.Flags().StringVarP(&t, "date", "t", t, "时间: 2024-01-01|2024-01-01,2024-01-31")
	findCmd.Flags().StringVarP(&dir, "dir", "d", dirPath, "文件目录")
	findCmd.Flags().StringVarP(&exd, "ignore", "i", "", "排除目录：.git,node_modules,...")
	findCmd.Flags().StringVarP(&ext, "ext", "e", "", "指定文件类型：png,md,...")
}

func scannerDir(path string, level int) {
	files, err := filepath.Glob(filepath.Join(path, "*"))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		fileInfo, err2 := os.Stat(file)
		if err2 != nil {
			return
		}
		// 排除指定目录
		if len(extList) > 0 && !DirIsInclude(fileInfo.Name()) {
			continue
		}
		// 排除指定文件类型
		if !fileInfo.IsDir() && len(extList) > 0 && !ExtIsInclude(fileInfo.Name()) {
			continue
		}
		// 是否在指定时间范围内
		if !fileInfo.IsDir() && !CheckFileDate(fileInfo.ModTime()) {
			continue
		}

		fmt.Printf("%s%s %s\n", strings.Repeat("--", level*space), fileInfo.Name(), carbon.Parse(fileInfo.ModTime().String()).ToDateString())
		if fileInfo.IsDir() {
			scannerDir(file, level+1)
		}

		// 排除指定文件类型

		// 是否在指定时间范围内

		fmt.Printf("%s%s %s\n", strings.Repeat("--", level*space), fileInfo.Name(), carbon.Parse(fileInfo.ModTime().String()).ToDateString())
		if fileInfo.IsDir() {
			scannerDir(file, level+1)
		}
	}
}

// 判断某个时间是否在某个时间段内
func CheckFileDate(modTime time.Time) bool {
	if dateType == 0 {
		if modTime.Equal(stime) {
			return true
		}
	} else {
		if modTime.After(etime) && modTime.Before(stime) {
			return true
		}
	}
	return false
}

// 包含目录
func DirIsInclude(name string) bool {
	result := false
	for _, v := range extList {
		if v == name {
			result = true
			break
		}
	}
	return result
}

// 扩展名是否包含
func ExtIsInclude(name string) bool {
	if !strings.Contains(name, ".") {
		return false
	}
	split := strings.Split(name, ".")
	e := split[len(split)-1] // png
	result := false
	for _, v := range extList {
		if e == v {
			result = true
			break
		}
	}
	return result
}
