package src

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/spf13/cobra"
	"log"
	"time"
	"tmp/config"
)

var copyObject = &cobra.Command{
	Use:     "copy_object",
	Short:   "copy object from bucket_1 to bucket_2",
	Long:    "This command will copy object from bucket_1 to bucket_2",
	Example: "go run s3.go object copy_object src_bucket dst_bucket src_object dst_object",
	Args:    cobra.ExactArgs(4),
	Run: func(cmd *cobra.Command, args []string) {
		srcBucket := args[0]
		dstBucket := args[1]
		srcObject := args[2]
		dstObject := args[3]
		CopyObject(srcBucket, dstBucket, srcObject, dstObject)
	},
}

func CopyObject(srcBucket, dstBucket, srcObject, dstObject string) {
	client, err := initClient(config.Cfg.Version)
	if err != nil {
		return
	}
	// Source object
	src := minio.CopySrcOptions{
		Bucket: srcBucket,
		Object: srcObject,
		// All following conditions are allowed and can be combined together.
		// Set modified condition, copy object modified since 2014 April.
		MatchModifiedSince: time.Date(2014, time.April, 0, 0, 0, 0, 0, time.UTC),
		// Set unmodified condition, copy object unmodified since 2014 April.
		// MatchUnmodifiedSince: time.Date(2014, time.April, 0, 0, 0, 0, 0, time.UTC),
		// Set matching ETag condition, copy object which matches the following ETag.
		// MatchETag: "31624deb84149d2f8ef9c385918b653a",
		// Set matching ETag copy object which does not match the following ETag.
		// NoMatchETag: "31624deb84149d2f8ef9c385918b653a",
	}

	// Destination object
	dst := minio.CopyDestOptions{
		Bucket: dstBucket,
		Object: dstObject,
	}
	ctx := context.Background()
	ui, err := client.CopyObject(ctx, dst, src)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Copied %s, successfully to %s - UploadInfo %v\n", dst, src, ui)

}

func init() {
	Object.AddCommand(copyObject)
}
