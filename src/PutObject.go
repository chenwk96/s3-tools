package src

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
	"log"
	"os"
	"tmp/config"
)

// 以文件的形式上传到桶
var putObject = &cobra.Command{
	Use:     "put2object",
	Short:   "put file to object buckets in bytes",
	Long:    "This command will put file to object in bytes",
	Example: "go run s3.go bucket put2object file_path bucket_name object_name",
	Run: func(cmd *cobra.Command, args []string) {
		FilePath := args[0]
		BucketName = args[1]
		ObjectName := args[2]
		PutToObjectByte(FilePath, ObjectName)
	},
}

func PutToObjectByte(FilePath, ObjectName string) {
	file, err := os.Open(FilePath)
	if err != nil {
		log.Fatalln("open file err : ", err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		log.Fatalln("stat file err : ", err)
	}

	MinioClient, err := initClient(config.Cfg.Version)
	if err != nil {
		log.Fatalln("init client err : ", err)
	}
	ctx := context.Background()

	exists, err := MinioClient.BucketExists(ctx, BucketName)
	if !exists {
		log.Fatalln("err : This bucket is not exist")
	}

	object, err := MinioClient.StatObject(ctx, BucketName, ObjectName, minio.StatObjectOptions{})
	if err == nil {
		log.Fatalln("err : This object is exist, please change the object name \n"+
			"object info : ", object)
	}
	info, err := MinioClient.PutObject(ctx, BucketName, ObjectName, file, stat.Size(), minio.PutObjectOptions{})
	if err != nil {
		log.Fatalln("put object err : ", err)
	}
	log.Println("Successfully uploaded ", info)
}

func init() {
	Object.AddCommand(putObject)
}
