package client

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

func NewAddDomainClient(regionId string, accessKeyId string, accessKeySecret string, subDomain string, mainDomain string, domainType string, domainValue string) (*alidns.AddDomainRecordResponse, error) {
	// 创建客户端
	client, err := alidns.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Print(err.Error())
	}
	request := alidns.CreateAddDomainRecordRequest()
	request.DomainName = mainDomain
	request.RR = subDomain
	request.Type = domainType
	request.Value = domainValue
	// 发送请求
	response, err := client.AddDomainRecord(request)
	if err != nil {
		// utils.MARKER_STRING()
		// fmt.Println()
		return nil, err
	}
	return response, nil
}

func NewDelDomainClient(regionId string, accessKeyId string, accessKeySecret string, subDomain string, mainDomain string) (*alidns.DeleteSubDomainRecordsResponse, error) {

	client, err := alidns.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Print(err.Error())
	}
	request := alidns.CreateDeleteSubDomainRecordsRequest()
	request.DomainName = mainDomain
	request.RR = subDomain

	response, err := client.DeleteSubDomainRecords(request)
	if err != nil {
		// utils.MARKER_STRING()
		// fmt.Println(err)
		return nil, err
	}
	return response, nil
}

func NewUpdDomainClient(regionId string, accessKeyId string, accessKeySecret string, domainRecordIdInput string, subDomainInput string, domainTypeInput string, domainValueInput string) (*alidns.UpdateDomainRecordResponse, error) {
	client, err := alidns.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Print(err.Error())
	}
	request := alidns.CreateUpdateDomainRecordRequest()
	request.RecordId = domainRecordIdInput
	request.RR = subDomainInput
	request.Type = domainTypeInput
	request.Value = domainValueInput
	response, err := client.UpdateDomainRecord(request)
	if err != nil {
		// utils.MARKER_STRING()
		// fmt.Println(err)
		return nil, err
	}
	return response, nil
}

func NewDesDomainClient(regionId string, accessKeyId string, accessKeySecret string, fullDomainInput string, mainDomain string) (*alidns.DescribeSubDomainRecordsResponse, error) {
	client, err := alidns.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Print(err.Error())
	}
	request := alidns.CreateDescribeSubDomainRecordsRequest()
	request.SubDomain = fullDomainInput
	request.DomainName = mainDomain

	// 发送请求
	response, err := client.DescribeSubDomainRecords(request)
	if err != nil {
		// utils.MARKER_STRING()
		// fmt.Println(err)
		return nil, err
	}
	return response, nil
}

func NewSeaDomainClient(regionId string, accessKeyId string, accessKeySecret string, subDomain string, mainDomain string) (*alidns.DescribeDomainRecordsResponse, error) {
	client, err := alidns.NewClientWithAccessKey(regionId, accessKeyId, accessKeySecret)
	if err != nil {
		fmt.Print(err.Error())
	}
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = mainDomain
	request.RRKeyWord = subDomain
	// 发送请求
	response, err := client.DescribeDomainRecords(request)
	if err != nil {
		// utils.MARKER_STRING()
		// fmt.Println(err)
		return nil, err
	}
	return response, nil
}
