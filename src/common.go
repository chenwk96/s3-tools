package src

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"tmp/config"
)

var (
	BucketName string
)

func initConfig() {
	//YAML 文件路径
	file_path := "/Users/chh/GolandProjects/tmp/config/config.yaml"
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println("读取配置文件时出错: ", err)
	}

	// 解析 YAML 数据
	err = yaml.Unmarshal(data, &config.Cfg)
	if err != nil {
		fmt.Println("解析配置文件时出错: ", err)
	}
}

func initClient(version string) (*minio.Client, error) {

	var minioClient *minio.Client
	var err error
	if config.Cfg.Version == "V2" {
		minioClient, err = minio.New(config.Cfg.Endpoint, &minio.Options{
			//Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Creds:  credentials.NewStaticV2(config.Cfg.AccessKeyID, config.Cfg.SecretAccessKey, ""),
			Secure: config.Cfg.UseSSL,
		})
	} else if config.Cfg.Version == "V4" {
		minioClient, err = minio.New(config.Cfg.Endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(config.Cfg.AccessKeyID, config.Cfg.SecretAccessKey, ""),
			Secure: config.Cfg.UseSSL,
		})
	}
	return minioClient, err
}

func checkBucketExist(ctx context.Context, BucketName string, Minioclient *minio.Client) bool {
	exists, _ := Minioclient.BucketExists(ctx, BucketName)
	return exists
}
