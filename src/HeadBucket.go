package src

import (
	"context"
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var headBucket = &cobra.Command{
	Use:     "head_bucket [bucket]",
	Short:   "check a bucket exist",
	Long:    "This commond will check the bucket whether in server",
	Example: "go run main.go head_bucket <bucket_name>",
	Run: func(cmd *cobra.Command, args []string) {
		bucket := args[0]
		HeadBucket(bucket)
	},
}

func HeadBucket(bucket string) {
	s3Client, err := initClient()
	if err != nil {
		log.Println("Failed to initial s3 client, err : ", err)
		return
	}

	ctx := context.Background()
	output, err := s3Client.HeadBucketWithContext(ctx, &s3.HeadBucketInput{
		Bucket: &bucket,
	})

	if err != nil {
		if errHasCode(err, "NoSuchBucket") {
			log.Println("This Bucket is not exist")
		} else {
			var awsErr awserr.Error
			if errors.As(err, &awsErr) {
				log.Printf("Failed to head bucket[%s], Error[%s]", bucket, awsErr.Message())
			}
		}
	} else {
		log.Printf("Bucket [%s]: \n%v", bucket, output)
	}
}

func init() {
	Bucket.AddCommand(headBucket)
}
