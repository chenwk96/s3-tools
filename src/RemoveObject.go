package src

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var removeObject = &cobra.Command{
	Use:     "remove_object",
	Short:   "remove object from bucket",
	Long:    "This command will remove a object from a bucket",
	Example: "go run s3.go object remove_object bucket_name object_name",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		bucket := args[0]
		object := args[1]
		RemoveObject(bucket, object)
	},
}

func RemoveObject(bucket, object string) {
	s3Client, err := initClient()
	if err != nil {
		log.Println("Failed to initial s3 Client, err: ", err)
		return
	}

	ctx := context.Background()
	_, err = s3Client.HeadObjectWithContext(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})

	if err != nil {
		log.Println("Faile to head object, err: ", err)
		return
	}

	_, err = s3Client.DeleteObjectWithContext(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})

	if err != nil {
		log.Println("Failed to remove object, err : ", err)
		return
	}

	log.Println("remove object success")
}

func init() {
	Object.AddCommand(removeObject)
}
