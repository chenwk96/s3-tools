package src

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
	"log"
	"tmp/config"
)

var removeObjects = &cobra.Command{
	Use:   "remove_objects",
	Short: "remove object from bucket",
	Long: "This command will remove objects from a bucket, you can use prefix_name " +
		"to filterate which you want to remove.",
	Example: "go run s3.go object remove_objects bucket_name <option: prefix_name>",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		BucketName = args[0]
		PrefexName := args[1]
		RemoveObjects(BucketName, PrefexName)
	},
}

func RemoveObjects(BucketName, PrefixName string) {
	Objectch := make(chan minio.ObjectInfo)
	client, err := initClient(config.Cfg.Version)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	go func() {
		defer close(Objectch)
		// List all objects from a bucket-name with a matching prefix.
		opts := minio.ListObjectsOptions{Prefix: PrefixName, Recursive: true}
		for object := range client.ListObjects(context.Background(), BucketName, opts) {
			if object.Err != nil {
				log.Fatalln(object.Err)
			}
			Objectch <- object
		}
	}()

	for rErr := range client.RemoveObjects(ctx, BucketName, Objectch, minio.RemoveObjectsOptions{}) {
		fmt.Println("Error detected during deletion: ", rErr)
	}

	log.Println("remove objects done")
}

func init() {
	Object.AddCommand(removeObjects)
}
