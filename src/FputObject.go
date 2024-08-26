package src

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
	"log"
	"tmp/config"
)

// 以文件的形式上传到桶
var putFObject = &cobra.Command{
	Use:     "putfile2object",
	Short:   "put file to object buckets",
	Long:    "This command will put file to object use file path.",
	Example: "go run s3.go bucket putfile2object file_path bucket_name object_name",
	Run: func(cmd *cobra.Command, args []string) {
		FilePath := args[0]
		BucketName = args[1]
		ObjectName := args[2]
		PutFileToObject(FilePath, ObjectName)
	},
}

func PutFileToObject(FilePath, ObjectName string) {
	clientMinIo, err := initClient(config.Cfg.Version)
	if err != nil {
		log.Fatalln(err)
	}
	ctx := context.Background()

	exists, _ := clientMinIo.BucketExists(ctx, BucketName)
	if !exists {
		log.Fatalln("err : This bucket is not exist")
	}
	object, err := clientMinIo.StatObject(ctx, BucketName, ObjectName, minio.StatObjectOptions{})
	if err == nil {
		log.Fatalln("err : This object is exist, please change the object name \n"+
			"object info : ", object)
	}

	info, err := clientMinIo.FPutObject(ctx, BucketName, ObjectName, FilePath, minio.PutObjectOptions{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully uploaded ", info)
}

func init() {
	Object.AddCommand(putFObject)
}
