package src

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

func init() {
	Bucket.AddCommand(createBucketCmd)
}

var createBucketCmd = &cobra.Command{
	Use:     "create_bucket",
	Short:   "create buckets",
	Long:    "This command will create buckets.",
	Example: "go run s3.go bucket create_bucket <bucket_name>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bucket := args[0]
		CreateBucket(bucket)
	},
}

func CreateBucket(bucket string) {
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
		if !errHasCode(err, "NotFound") {
			var awsErr awserr.Error
			if errors.As(err, &awsErr) {
				log.Printf("Failed to head bucket[%s], Error[%s]", bucket, awsErr.Message())
			}
			return
		}
	} else {
		log.Printf("Bucket [%s] already exist", bucket)
		return
	}

	_, err = s3Client.CreateBucketWithContext(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		log.Printf("Failed to create bucket %s, err: %v", bucket, err)
	} else {
		log.Printf("Bucket %s successfully created", bucket)
	}
}
