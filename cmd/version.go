package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var BuildNumber string = "1.0"
var BuildRevision string = ""

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of au",
	Long:  `All software has versions. This version from git tags`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("au build number:%s  revision:%s version:%s build time:%s \n", BuildNumber, BuildRevision, versionNumber, time.Now().String())
	},
}

//FIXME we need to find a way to handle our major.minor.. version number
var versionNumber string = "1.0"

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
