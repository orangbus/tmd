/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"crypto/rand"
	"fmt"
	"github.com/google/uuid"
	"github.com/orangbus/cmd/console"
	"strings"

	"github.com/spf13/cobra"
)

const (
	letters        = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits         = "0123456789"
	specialChars   = "!@#$%^&*()-_+=<>?/"
	passwordLength = 12 // 你可以根据需要调整密码长度
)

// makeCmd represents the make command
var makeCmd = &cobra.Command{
	Use:   "make",
	Short: "随机生成一个随机数，flags:key,uuid",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			console.Error("请输入参数,Flags:key,uuid")
			return
		}
		switch args[0] {
		case "key":
			generateKey()
		case "uuid":
			generateUuid()
		default:
			console.Error("参数错误")
		}
	},
}

func init() {
	rootCmd.AddCommand(makeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// makeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// makeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func generateRandomRune(pool string) (rune, error) {
	b, err := generateRandomBytes(1)
	if err != nil {
		return 0, err
	}
	index := int(b[0]) % len(pool)
	return rune(pool[index]), nil
}
func generatePassword() (string, error) {
	var password []rune

	// 确保密码包含至少一个大写字母、一个小写字母、一个数字和一个特殊字符
	for _, charPool := range []string{letters, letters, digits, specialChars} {
		r, err := generateRandomRune(charPool)
		if err != nil {
			return "", err
		}
		password = append(password, r)
	}

	// 随机填充剩余的密码字符
	for i := 0; i < passwordLength-4; i++ {
		r, err := generateRandomRune(letters + digits + specialChars)
		if err != nil {
			return "", err
		}
		password = append(password, r)
	}

	// 打乱密码字符的顺序以增加复杂度
	passwordStr := string(password)
	return passwordStr, nil
}

func generateKey() {
	key, err := generatePassword()
	if err != nil {
		console.Error(err.Error())
		return
	}
	console.Success(strings.Repeat("---", 10))
	console.Success(fmt.Sprintf("key: %s", key))
	console.Success(strings.Repeat("---", 10))
}

func generateUuid() {
	console.Success(strings.Repeat("---", 10))
	console.Success(fmt.Sprintf("UUID: %s", uuid.New().String()))
	console.Success(strings.Repeat("---", 10))
}
