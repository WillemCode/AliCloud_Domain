package utils

import (
	"fmt"
	"strings"

	"github.com/pterm/pterm"
)

const (
	STRING = "++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++"
)

func MARKER_STRING() {
	fmt.Println("")
	fmt.Println(STRING)
	fmt.Println("")
}

// 接收输入参数的函数
func TextInput(information string, defaultOption string) string {
	MARKER_STRING()
	// 创建单选菜单
	printer := pterm.DefaultInteractiveTextInput.
		WithDefaultText(information).
		WithDefaultValue(defaultOption) // 默认参数选项
	// 显示菜单并获取选择的选项
	TextResult, _ := printer.Show()
	return TextResult
}

// 提取子域名的函数
func ExtractDomainParts(fullDomain string) (string, string) {
	// 以 "." 分割域名
	parts := strings.Split(fullDomain, ".")
	if len(parts) < 2 {
		return "", fullDomain // 不足两个部分，返回空子域名和原始域名
	}

	// 提取主域名部分
	mainDomain := strings.Join(parts[len(parts)-2:], ".")
	// 提取子域名部分
	subDomain := strings.Join(parts[:len(parts)-2], ".")
	return subDomain, mainDomain
}
