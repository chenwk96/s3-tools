package src

import (
	"bytes"
	"context"
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

// 以文件的形式上传到桶
var putObject = &cobra.Command{
	Use:     "put2object",
	Short:   "put file to object buckets in bytes",
	Long:    "This command will put file to object in bytes",
	Example: "go run s3.go bucket put2object file_path bucket_name object_name",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		bucket := args[1]
		object := args[2]
		PutToObjectByte(filePath, bucket, object)
	},
}

func PutToObjectByte(filePath, bucket, object string) {
	s3Client, err := initClient()
	if err != nil {
		log.Println("Failed to initial s3 client, err: ", err)
		return
	}
	ctx := context.Background()
	_, err = s3Client.HeadBucketWithContext(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		if errHasCode(err, "NoSuchBucket") {
			log.Printf("bucket %s is not exist", bucket)
		} else {
			var awsErr awserr.Error
			if errors.As(err, &awsErr) {
				log.Printf("Failed to head bucket, err: %s", awsErr.Message())
			}
		}
		return
	}

	_, err = s3Client.HeadObjectWithContext(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})

	if err != nil {
		if !errHasCode(err, "NotFound") {
			var awsErr awserr.Error
			if errors.As(err, &awsErr) {
				log.Printf("Failed to head object, err: %s", awsErr.Message())
			}
		}
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("open file err : ", err)
		return
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Println("stat file err : ", err)
		return
	}

	buf := make([]byte, stat.Size())
	file.Read(buf)

	_, err = s3Client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
		Body:   bytes.NewReader(buf),
	})

	if err != nil {
		log.Println("Failed to put object err : ", err)
	}

	log.Println("Successfully uploaded")
}

func init() {
	Object.AddCommand(putObject)
}
