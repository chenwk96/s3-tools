package src

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "command",
	Short: "A brief description of your command",
	Long:  `A longer description of your command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("you can choose a level : \n" +
			"1. bucket \n" +
			"2. object \n" +
			"For example: \n" +
			"go run s3.go bucket -h\n" +
			"go run s3.go object -h")
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&BucketName, "bucket", "", "Bucket name")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
