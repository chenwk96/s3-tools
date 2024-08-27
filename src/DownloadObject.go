package src

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var downloadObject = &cobra.Command{
	Use:     "download_object",
	Short:   "point object do some operation",
	Long:    "This command will offer some information about some object's operation",
	Example: "go run s3.go download_object <bucket_name> <object_name> <local_file_path>",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		bucket := args[0]
		object := args[1]
		localFileName := args[2]
		DownloadObject(localFileName, bucket, object)
	},
}

func DownloadObject(localFileName, bucketName, objectName string) {
	s3Client, err := initClient()
	if err != nil {
		log.Println("Failed to initial s3 client, err: ", err)
		return
	}

	ctx := context.Background()
	output, err := s3Client.GetObjectWithContext(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	})

	if err != nil {
		log.Println("Failed to get object, err: ", err)
		return
	}
	defer output.Body.Close()

	localFile, err := os.Create(localFileName)
	if err != nil {
		log.Println("create local file err : ", err)
		return
	}
	defer localFile.Close()

	if _, err = io.Copy(localFile, output.Body); err != nil {
		log.Println("copy object to local file err : ", err)
	}

	log.Println("download object from remote server success")
}

func init() {
	Object.AddCommand(downloadObject)
}
