package models

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"tiktok_Demo/config"
)

var InitBucket *oss.Bucket

func init() {
	config := config.ReturnConfig()

	Endpoint := config.Endpoint
	AccessKeyID := config.AccessKeyID
	AccessKeySecret := config.AccessKeySecret
	Bucket := config.Bucket

	client, err := oss.New(Endpoint, AccessKeyID, AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 指定bucket
	InitBucket, err = client.Bucket(Bucket) // 根据自己的填写
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
