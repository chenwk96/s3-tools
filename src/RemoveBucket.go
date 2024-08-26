package src

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"tmp/config"
)

var removeBucketCmd = &cobra.Command{
	Use:     "remove_bucket [bucket]",
	Short:   "Remove a bucket",
	Long:    "Remove a bucket from server.",
	Example: "go run main.go bucket remove_bucket <bucket_name>",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		BucketName = args[0]
		removeBucket(BucketName)
	},
}

func removeBucket(BucketName string) {

	minioClient, err := initClient(config.Cfg.Version)

	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	exists, err := minioClient.BucketExists(ctx, BucketName)

	if !exists {
		log.Fatalln("Bucket is not exists")
	}

	err = minioClient.RemoveBucket(ctx, BucketName)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Bucket %s is successfully deleted\n", BucketName)
}

func init() {
	Bucket.AddCommand(removeBucketCmd)
}
