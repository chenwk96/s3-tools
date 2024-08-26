package src

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"tmp/config"
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
	var minioClient, err = initClient(config.Cfg.Version)
	ctx := context.Background()
	buckets, err := minioClient.ListBuckets(ctx)

	if err != nil {
		log.Fatalln("err : ", err)
	}
	for i, bucket := range buckets {
		fmt.Printf("%d bucket info : %s \n", i+1, bucket)
	}
}

func init() {
	Bucket.AddCommand(listBucket)
}
