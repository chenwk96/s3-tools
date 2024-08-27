package src

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var copyObject = &cobra.Command{
	Use:     "copy_object",
	Short:   "copy object from bucket_1 to bucket_2",
	Long:    "This command will copy object from bucket_1 to bucket_2",
	Example: "go run s3.go object copy_object src_bucket dst_bucket src_object dst_object",
	Args:    cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		srcBucket := args[0]
		dstBucket := args[1]
		srcObject := args[2]
		dstObject := args[3]
		CopyObject(srcBucket, dstBucket, srcObject, dstObject)
	},
}

func CopyObject(srcBucket, dstBucket, srcObject, dstObject string) {
	s3Client, err := initClient()
	if err != nil {
		log.Printf("Failed to initial S3 client: %v", err)
		return
	}
	ctx := context.Background()

	cpybucket := srcBucket
	cpyObject := strings.TrimPrefix(srcObject, "/")
	cpySource := fmt.Sprintf("%s/%s", cpybucket, cpyObject)
	_, err = s3Client.CopyObjectWithContext(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(dstBucket),
		Key:        aws.String(dstObject),
		CopySource: aws.String(cpySource),
	})

	if err != nil {
		log.Println("Failed to copy object, err: ", err)
	}

	log.Printf("Copied successfully.")
}

func init() {
	Object.AddCommand(copyObject)
}
