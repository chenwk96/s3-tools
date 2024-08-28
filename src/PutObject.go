package src

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/cobra"
)

// 以文件的形式上传到桶
var putObject = &cobra.Command{
	Use:     "put_object",
	Short:   "put file to object buckets",
	Long:    "This command will put file to object",
	Example: "go run s3.go bucket put2object file_path bucket_name object_name",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		bucket := args[1]
		object := args[2]
		PutObject(filePath, bucket, object)
	},
}

func PutObject(filePath, bucket, object string) {
	s3Client, err := initClient()
	if err != nil {
		log.Println("Failed to initial s3 client, err: ", err)
		return
	}

	ctx := context.Background()
	objInfo, err := s3Client.HeadObjectWithContext(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})

	if err == nil {
		log.Println("err: this object is exists, please change the object name.\nobject info: ", objInfo)
		return
	} else {
		if !errHasCode(err, "NotFound") {
			log.Println("Failed to head object, err: ", err)
			return
		}
	}

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("open file err : ", err)
		return
	}
	defer file.Close()

	uploader := s3manager.NewUploaderWithClient(s3Client)

	_, err = uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
		Body:   file,
	})

	if err != nil {
		log.Println("Failed to put object err : ", err)
		return
	}

	log.Println("Successfully uploaded")
}

func init() {
	Object.AddCommand(putObject)
}
