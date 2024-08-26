package src

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"tmp/config"
)

var downloadObject = &cobra.Command{
	Use:     "download_object",
	Short:   "point object do some operation",
	Long:    "This command will offer some information about some object's operation",
	Example: "go run s3.go download_object <bucket_name> <object_name> <local_file_path>",
	Args:    cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		BucketName = args[0]
		ObjectName := args[1]
		LocalFileName := args[2]
		DownloadObject(LocalFileName, BucketName, ObjectName)
	},
}

func DownloadObject(LocalFileName, BucketName, ObjectName string) {
	client, err := initClient(config.Cfg.Version)
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()
	object, err := client.GetObject(ctx, BucketName, ObjectName, minio.GetObjectOptions{})

	if err != nil {
		log.Fatalln("getObject err : ", err)
	}
	defer object.Close()

	localFile, err := os.Create(LocalFileName)
	if err != nil {
		log.Fatalln("create local file err : ", err)
	}
	defer localFile.Close()

	if _, err = io.Copy(localFile, object); err != nil {
		log.Fatalln("copy object to local file err : ", err)
	}
	log.Println("download object from remote server success")
}

func init() {
	Object.AddCommand(downloadObject)
}
