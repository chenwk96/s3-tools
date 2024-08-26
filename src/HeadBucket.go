package src

import (
	"context"
	"github.com/spf13/cobra"
	"log"
	"tmp/config"
)

var headBucket = &cobra.Command{
	Use:     "head_bucket [bucket]",
	Short:   "check a bucket exist",
	Long:    "This commond will check the bucket whether in server",
	Example: "go run main.go head_bucket <bucket_name>",
	Run: func(cmd *cobra.Command, args []string) {
		BucketName = args[0]
		HeadBucket()
	},
}

func HeadBucket() {
	minioClient, err := initClient(config.Cfg.Version)
	if err != nil {
		log.Fatalln("err : ", err)
	}

	ctx := context.Background()
	exists, _ := minioClient.BucketExists(ctx, BucketName)
	if exists {
		log.Println("This Bucket is exist")
	} else {
		log.Println("This Bucket is not exist")
	}
}

func init() {
	Bucket.AddCommand(headBucket)
}
