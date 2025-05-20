package domain

import (
	"bufio"
	"fmt"
	"os"

	"gitee.com/chuyio/aligo/internal/client"
	"gitee.com/chuyio/aligo/pkg/utils"
)

func DelDomain(regionId string, accessKeyId string, accessKeySecret string, subDomain string, mainDomain string, AccountName string) {
	response, err := client.NewDelDomainClient(regionId, accessKeyId, accessKeySecret, subDomain, mainDomain)
	if err != nil {
		// utils.MARKER_STRING()
		// fmt.Println(err)
		return
	}

	utils.MARKER_STRING()
	fmt.Println("云平台账户信息:", AccountName)
	fmt.Println("域名删除成功:", subDomain+"."+mainDomain)
	fmt.Println("请求ID:", response.RequestId)
	fmt.Println("删除的解析记录总数:", response.TotalCount)
	utils.MARKER_STRING()
	fmt.Println("程序已执行完成，按任意键退出...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	os.Exit(0)
}
