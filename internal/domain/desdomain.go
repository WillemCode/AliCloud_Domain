package domain

import (
	"bufio"
	"fmt"
	"os"

	"gitee.com/chuyio/aligo/internal/client"
	"gitee.com/chuyio/aligo/pkg/utils"
	"github.com/pterm/pterm"
)

func DesDomain(regionId string, accessKeyId string, accessKeySecret string, fullDomainInput string, mainDomain string, AccountName string) {
	response, err := client.NewDesDomainClient(regionId, accessKeyId, accessKeySecret, fullDomainInput, mainDomain)
	if err != nil {
		// utils.MARKER_STRING()
		// fmt.Println(err)
		return
	}
	utils.MARKER_STRING()
	fmt.Println("云平台账户信息:", AccountName)
	fmt.Printf("\n查询到 %d 个结果\n", response.TotalCount)

	// 遍历 Record 数组并提取 DomainName, Value, RR, Status
	for _, record := range response.DomainRecords.Record {
		fmt.Println("")
		maxLineLength := 50
		status := fmt.Sprintf("%-*s", maxLineLength, "状态: "+record.Status)
		recordType := fmt.Sprintf("%-*s", maxLineLength, "类型: "+record.Type)
		recordValue := fmt.Sprintf("%-*s", maxLineLength, "解析值: "+record.Value)
		recordId := fmt.Sprintf("%-*s", maxLineLength, "RecordId: "+record.RecordId)
		pterm.DefaultBox.WithTitle(record.RR+"."+record.DomainName).WithTitleTopCenter().Printf("\n%s\n%s\n%s\n%s\n", status, recordType, recordValue, recordId)
		fmt.Println("")
	}
	utils.MARKER_STRING()
	fmt.Println("程序已执行完成，按任意键退出...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	os.Exit(0)
}
