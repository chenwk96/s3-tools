package src

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Bucket = &cobra.Command{
	Use:   "bucket option",
	Short: "create buckets",
	Long:  "This command will create buckets.",
	Example: "go run s3.go Bucket create_bucket <Param>\t" +
		"useinfo: create_bucket commond can create new Bucket\n" +
		"go run s3.go Bucket remove_bucket <Param>\t" +
		"useinfo: remove_bucket commond can remove point Bucket but should check the bucket exist in server\n" +
		"go run s3.go Bucket list_bucket   <Param>\t" +
		"useinfo: list_bucket commond will list all bucket info\n" +
		"go run s3.go Bucket head_bucket   <Param>\t" +
		"useinfo: head_bucket commond can check the bucket whether exist in server\n" +
		"more commond info you can go on sub commond use -h to get\n",
	Run: func(cmd *cobra.Command, args []string) {
		BucketInfo()
	},
}

func BucketInfo() {
	fmt.Println("you can only opration below\n" +
		"bucket operation : \n" +
		"     create_bucket             create a bucket in your server\n" +
		"     remove_bucket             delete a bucket from you server\n" +
		"     list_bucket               list all bucket infomation\n" +
		"     head_bucket               check the bucket wheather exist in server\n")
}

func init() {
	rootCmd.AddCommand(Bucket)
}
