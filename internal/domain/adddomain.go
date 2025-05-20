package domain

import (
	"bufio"
	"fmt"
	"os"

	"gitee.com/chuyio/aligo/internal/client"
	"gitee.com/chuyio/aligo/pkg/utils"
)

func AddDomain(regionId string, accessKeyId string, accessKeySecret string, subDomain string, mainDomain string, domainType string, domainValue string, AccountName string) {

	response, err := client.NewAddDomainClient(regionId, accessKeyId, accessKeySecret, subDomain, mainDomain, domainType, domainValue)
	if err != nil {
		// utils.MARKER_STRING()
		// fmt.Println(err)
		return
	}

	utils.MARKER_STRING()
	fmt.Println("云平台账户信息:", AccountName)
	fmt.Println("域名添加成功:", subDomain+"."+mainDomain)
	fmt.Println("记录ID:", response.RecordId)
	fmt.Println("请求ID:", response.RequestId)
	utils.MARKER_STRING()
	fmt.Println("程序已执行完成，按任意键退出...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	os.Exit(0)
}
