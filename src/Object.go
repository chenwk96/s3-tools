package src

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Object = &cobra.Command{
	Use:   "object option",
	Short: "point object do some operation",
	Long:  "This command will offer some information about some object's operation",
	Example: "go run s3.go put_stream <Param>\t upload in the form of a file\n" +
		"go run s3.go put_object <Param>\t 	 upload in the form of bytes\n" +
		"go run s3.go download_object <Param>\t download file from remote server\n" +
		"go run s3.go copy_object <Param>\t copy a object to another bucket'object\n" +
		"go run s3.go remove_object <Param>\t remove object from remote server\n" +
		"go run s3.go remove_objects <Param>\t remove objects use prefix name\n" +
		"more commond info you can go on sub commond use -h to get\n",
	Run: func(cmd *cobra.Command, args []string) {
		ObjectInfo()
	},
}

func ObjectInfo() {
	fmt.Println("you can do opration below\n" +
		"object operation : \n" +
		"     put_object                upload file to server\n" +
		"     put_stream                upload file to server by bytes stream\n" +
		"     download_object           download file from server\n" +
		"     remove_object             delete object from bucket\n" +
		"     remove_objects            delete object from bucket use prefix_name\n" +
		"     copy_object               copy object from bucket_1 to bucket_2")
}

func init() {
	rootCmd.AddCommand(Object)
}
