package src

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var listBucket = &cobra.Command{
	Use:     "list_bucket [bucket]",
	Short:   "list all bucket info",
	Long:    "This common will list all bucket info from server.",
	Example: "go run main.go bucket list_bucket",
	Run: func(cmd *cobra.Command, args []string) {
		ListBucket()
	},
}

func ListBucket() {
	s3Client, err := initClient()
	if err != nil {
		log.Println("Failed to initial s3 Client, err: ", err)
		return
	}

	ctx := context.Background()
	listOutput, err := s3Client.ListBucketsWithContext(ctx, &s3.ListBucketsInput{})

	if err != nil {
		log.Println("Failed to list buckets, err:", err)
		return
	}

	for i, bucket := range listOutput.Buckets {
		fmt.Printf("%d bucket info : %s \n", i+1, bucket)
	}
}

func init() {
	Bucket.AddCommand(listBucket)
}
