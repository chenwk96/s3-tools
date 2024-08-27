package src

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/spf13/cobra"
)

var removeObjects = &cobra.Command{
	Use:   "remove_objects",
	Short: "remove object from bucket",
	Long: "This command will remove objects from a bucket, you can use prefix_name " +
		"to filterate which you want to remove.",
	Example: "go run s3.go object remove_objects bucket_name <option: prefix_name>",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bucket := args[0]
		prefix := args[1]
		RemoveObjects(bucket, prefix)
	},
}

func RemoveObjects(bucket, prefix string) {
	// Objectch := make(chan minio.ObjectInfo)
	s3Client, err := initClient()
	if err != nil {
		log.Println("Failed to init s3 client, err: ", err)
		return
	}

	ctx := context.Background()
	err = s3Client.ListObjectsPagesWithContext(ctx, &s3.ListObjectsInput{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	}, func(p *s3.ListObjectsOutput, last bool) (shouldContinue bool) {
		for _, obj := range p.Contents {
			fmt.Println("Object:", *obj.Key)
		}
		return true
	})

	if err != nil {
		log.Println("Failed to list objects, err: ", err)
		return
	}
}

func init() {
	Object.AddCommand(removeObjects)
}
