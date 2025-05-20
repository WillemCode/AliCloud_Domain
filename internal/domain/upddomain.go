package domain

import (
	"bufio"
	"fmt"
	"os"

	"gitee.com/chuyio/aligo/internal/client"
	"gitee.com/chuyio/aligo/pkg/utils"
)

func UpdDomain(regionId string, accessKeyId string, accessKeySecret string, fullDomainInput string, mainDomain string, AccountName string) {

	DesDomain(regionId, accessKeyId, accessKeySecret, fullDomainInput, mainDomain, AccountName)

	fmt.Println("请根据提示输入修改后的 前缀/类型/值")
	domainRecordIdInput := utils.TextInput("请输入RecordId", "123456789987654321")
	subDomainInput := utils.TextInput("请输入修改后的前缀", "www")
	domainTypeInput := utils.TextInput("请输入修改后的类型", "A/CNAME/TXT")
	domainValueInput := utils.TextInput("请输入修改后记录的值", "IP/域名")

	response, err := client.NewUpdDomainClient(regionId, accessKeyId, accessKeySecret, domainRecordIdInput, subDomainInput, domainTypeInput, domainValueInput)
	if err != nil {
		// utils.MARKER_STRING()
		// fmt.Println(err)
		return
	}

	utils.MARKER_STRING()
	fmt.Println("云平台账户信息:", AccountName)
	fmt.Println("解析记录 修改成功!")
	fmt.Println("记录ID:", response.RecordId)
	fmt.Println("请求ID:", response.RequestId)
	utils.MARKER_STRING()
	fmt.Println("程序已执行完成，按任意键退出...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	os.Exit(0)
}
