package src

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
	"log"
	"tmp/config"
)

var removeObject = &cobra.Command{
	Use:     "remove_object",
	Short:   "remove object from bucket",
	Long:    "This command will remove a object from a bucket",
	Example: "go run s3.go object remove_object bucket_name object_name",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		BucketName = args[0]
		ObjectName := args[1]
		RemoveObject(BucketName, ObjectName)
	},
}

func RemoveObject(BucketName, ObjectName string) {
	Minioclient, err := initClient(config.Cfg.Version)
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()
	exist, _ := Minioclient.BucketExists(ctx, BucketName)
	if !exist {
		log.Fatalln("bucket not exist")
	}
	err = Minioclient.RemoveObject(ctx, BucketName, ObjectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Fatalln("remove object err : ", err)
	}
	log.Println("remove object success")
}

func init() {
	Object.AddCommand(removeObject)
}
