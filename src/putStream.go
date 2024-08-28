package src

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/spf13/cobra"
)

var putStream = &cobra.Command{
	Use:     "put_stream",
	Short:   "put file to object buckets in stream",
	Long:    "This command will put file to object in stream.",
	Example: "go run s3.go bucket pu_stream file_path bucket_name object_name",
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		bucket := args[1]
		object := args[2]
		PutFileStream(filename, bucket, object)
	},
}

func PutFileStream(filename, bucket, object string) {
	s3Client, err := initClient()
	if err != nil {
		log.Println("Failed to initial s3 client, err: ", err)
		return
	}

	ctx := context.Background()
	objInfo, err := s3Client.HeadObjectWithContext(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
	})

	if err == nil {
		log.Println("err: this object is exists, please change the object name.\nobject info: ", objInfo)
		return
	} else {
		if !errHasCode(err, "NotFound") {
			log.Println("Failed to head object, err: ", err)
			return
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Printf("Faile to open file %s, err: %v", filename, err)
		return
	}

	reader, writer := io.Pipe()

	go func() {
		_, err := io.Copy(writer, file)
		if err != nil {
			return
		}

		file.Close()
		writer.Close()
	}()

	uploader := s3manager.NewUploaderWithClient(s3Client)

	_, err = uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(object),
		Body:   reader,
	})

	if err != nil {
		log.Println("Failed to put file in stream, err: ", err)
	}

	fmt.Println("Successfully uploaded ")
}

func init() {
	Object.AddCommand(putStream)
}
