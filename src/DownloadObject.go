package src

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
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
		localpath := args[2]
		DownloadObject(localpath, bucket, object)
	},
}

func DownloadObject(localFileName, bucketName, objectName string) {
	s3Client, err := initClient()
	if err != nil {
		log.Println("Failed to initial s3 client, err: ", err)
		return
	}

	localFile, err := os.Create(localFileName)
	if err != nil {
		log.Println("create local file err : ", err)
		return
	}
	defer localFile.Close()

	downloader := s3manager.NewDownloaderWithClient(s3Client)

	numBytes, err := downloader.Download(localFile, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectName),
	})

	if err != nil {
		log.Println("Failed to download object\nerr: ", err)
		return
	}

	log.Printf("download %s, %d bytes\n", localFileName, numBytes)
}

func init() {
	Object.AddCommand(downloadObject)
}
