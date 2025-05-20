package domain

import (
	"fmt"
	"io/ioutil"

	"gitee.com/chuyio/aligo/pkg/utils"
	"github.com/pterm/pterm"
	"gopkg.in/yaml.v2"
)

type Account struct {
	Name         string `yaml:"name"`
	RegionId     string `yaml:"regionId"`
	AccessKey    string `yaml:"access_key"`
	AccessSecret string `yaml:"access_secret"`
}

type Config struct {
	AliyunAccounts []Account `yaml:"aliyun_accounts"`
}

func loadConfig(filename string) (*Config, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func Domain() {
	config, err := loadConfig("./config/config.yaml")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	// totalAccounts := len(config.AliyunAccounts)

	// 菜单选项
	options := []string{
		"1. 添加域名解析记录",
		"2. 删除域名解析记录",
		"3. 修改域名解析记录",
		"4. 查询域名详细信息",
		"5. 模糊搜索域名信息",
	}

	selectedOption, _ := pterm.DefaultInteractiveSelect.WithDefaultText("请选择你要操作的内容...").WithOptions(options).Show()

	switch selectedOption {
	case "1. 添加域名解析记录":
		fmt.Println("请根据提示输入相关信息...")
		fullDomainInput := utils.TextInput("请输入要添加记录的完整域名", "www.example.com")
		domainValueInput := utils.TextInput("请输入要添加记录的值", "IP/域名")
		fmt.Println(domainValueInput)
		domainTypeInput := utils.TextInput("请输入要解析类型", "A/CNAME/TXT...")
		fmt.Println(domainTypeInput)
		subDomain, mainDomain := utils.ExtractDomainParts(fullDomainInput)
		// 遍历多个阿里云账户
		for _, account := range config.AliyunAccounts {
			// fmt.Printf("Account: %s\n", account.Name)
			AccountName := account.Name
			regionId := account.RegionId
			accessKeyId := account.AccessKey
			accessKeySecret := account.AccessSecret
			AddDomain(regionId, accessKeyId, accessKeySecret, subDomain, mainDomain, domainTypeInput, domainValueInput, AccountName)
		}
	case "2. 删除域名解析记录":
		fmt.Println("请根据提示输入相关信息...")
		fullDomainInput := utils.TextInput("请输入要删除记录的完整域名", "www.example.com")
		subDomain, mainDomain := utils.ExtractDomainParts(fullDomainInput)
		for _, account := range config.AliyunAccounts {
			AccountName := account.Name
			regionId := account.RegionId
			accessKeyId := account.AccessKey
			accessKeySecret := account.AccessSecret
			DelDomain(regionId, accessKeyId, accessKeySecret, subDomain, mainDomain, AccountName)
		}
	case "3. 修改域名解析记录":
		fmt.Println("请根据提示输入相关信息...")
		fullDomainInput := utils.TextInput("请输入要修改记录的完整域名", "www.example.com")
		_, mainDomain := utils.ExtractDomainParts(fullDomainInput)
		for _, account := range config.AliyunAccounts {
			AccountName := account.Name
			regionId := account.RegionId
			accessKeyId := account.AccessKey
			accessKeySecret := account.AccessSecret
			UpdDomain(regionId, accessKeyId, accessKeySecret, fullDomainInput, mainDomain, AccountName)
		}
	case "4. 查询域名详细信息":
		fmt.Println("正在执行 查询域名详细信息 的相关程序...")
		fullDomainInput := utils.TextInput("请输入要查询记录的完整域名", "www.example.com")
		_, mainDomain := utils.ExtractDomainParts(fullDomainInput)
		for _, account := range config.AliyunAccounts {
			AccountName := account.Name
			regionId := account.RegionId
			accessKeyId := account.AccessKey
			accessKeySecret := account.AccessSecret
			DesDomain(regionId, accessKeyId, accessKeySecret, fullDomainInput, mainDomain, AccountName)
		}
	case "5. 模糊搜索域名信息":
		fmt.Println("正在执行 查询域名详细信息 的相关程序...")
		fullDomainInput := utils.TextInput("请输入要搜索记录的完整域名", "example.com")
		subDomain, mainDomain := utils.ExtractDomainParts(fullDomainInput)
		for _, account := range config.AliyunAccounts {
			AccountName := account.Name
			regionId := account.RegionId
			accessKeyId := account.AccessKey
			accessKeySecret := account.AccessSecret
			SeaDomain(regionId, accessKeyId, accessKeySecret, subDomain, mainDomain, AccountName)
		}
	}
}
