package src

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"tmp/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"gopkg.in/yaml.v3"
)

var (
	BucketName string
)

func initConfig() {
	//YAML 文件路径
	file_path := "./config/config.yaml"
	data, err := ioutil.ReadFile(file_path)
	if err != nil {
		fmt.Println("读取配置文件时出错: ", err)
		return
	}

	// 解析 YAML 数据
	err = yaml.Unmarshal(data, &config.Cfg)
	if err != nil {
		fmt.Println("解析配置文件时出错: ", err)
		return
	}
}

func initClient() (*s3.S3, error) {
	creds := credentials.NewStaticCredentials(config.Cfg.AccessKeyID, config.Cfg.SecretAccessKey, "")
	httpClient := &http.Client{
		Timeout: time.Second * time.Duration(30),
	}

	session, err := session.NewSession(&aws.Config{
		Credentials:      creds,
		Endpoint:         aws.String(config.Cfg.Endpoint),
		Region:           aws.String(config.Cfg.Region),
		S3ForcePathStyle: aws.Bool(true),
		HTTPClient:       httpClient,
	})

	if err != nil {
		fmt.Printf("Failed to create session: %v\n", err)
		return nil, errors.New("faild to create session")
	}

	s3Client := s3.New(session)

	if s3Client == nil {
		fmt.Printf("Failed to create s3 client: %v", err)
		return nil, errors.New("failed to create s3 client")
	}

	return s3Client, nil
}

func errHasCode(err error, code string) bool {
	if err == nil || code == "" {
		return false
	}

	var awsErr awserr.Error
	if errors.As(err, &awsErr) {
		if awsErr.Code() == code {
			return true
		}
	}

	var multiUploadErr s3manager.MultiUploadFailure
	if errors.As(err, &multiUploadErr) {
		return errHasCode(multiUploadErr.OrigErr(), code)
	}

	return false
}
