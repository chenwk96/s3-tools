package src

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
	"log"
	"tmp/config"
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
		BucketName = args[0]
		fmt.Printf("Endpoint: %s\n", config.Cfg.Endpoint)
		fmt.Printf("access_key_id: %s\n", config.Cfg.AccessKeyID)
		fmt.Printf("secret_access_key: %s\n", config.Cfg.SecretAccessKey)
		ctx := context.Background()

		minioClient, err := initClient(config.Cfg.Version)

		if err != nil {
			log.Fatalln(err)
		}

		exists, _ := minioClient.BucketExists(ctx, BucketName)
		if exists {
			log.Fatalln("err : This bucket is exist")
		} else {
			err = minioClient.MakeBucket(ctx, BucketName, minio.MakeBucketOptions{Region: config.Cfg.Region})
			if err != nil {
				log.Fatalln(err)
			} else {
				log.Println("Successfully created ", BucketName)
			}
		}
	},
}
