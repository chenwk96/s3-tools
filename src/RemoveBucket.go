package src

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var removeBucketCmd = &cobra.Command{
	Use:     "remove_bucket [bucket]",
	Short:   "Remove a bucket",
	Long:    "Remove a bucket from server.",
	Example: "go run main.go bucket remove_bucket <bucket_name>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bucket := args[0]
		removeBucket(bucket)
	},
}

func removeBucket(bucket string) {
	s3Client, err := initClient()

	if err != nil {
		log.Println("Failed to initial s3 client, err: ", err)
		return
	}

	ctx := context.Background()
	_, err = s3Client.DeleteBucketWithContext(ctx, &s3.DeleteBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		log.Printf("Failed to delete bucket: %s\n, err: %v", bucket, err)
		return
	}

	log.Printf("Bucket %s is successfully deleted\n", bucket)
}

func init() {
	Bucket.AddCommand(removeBucketCmd)
}
