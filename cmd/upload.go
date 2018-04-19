package cmd

import (
	"fmt"
	"os"

	"github.com/linkernetworks/aurora/src/cmd/au/client"
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload file entry to specific workspace",
	Long: `
	Upload file entry to specific workspace
	Usage: "au upload WORKSPACE_ID FILE_PATH"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("You need specific your WORKSPACE_ID and file path")
			return
		}

		//validate input path
		var pathList []string
		for k, v := range args {
			if k >= 1 {
				_, err := os.Stat(v)
				if os.IsNotExist(err) {
					fmt.Println("File not eixst:", v)
					continue
				} else {
					pathList = append(pathList, v)
				}
			}
		}
		if len(pathList) > 0 {
			//client.UploadPathsToWorkspace(*usingSSL, args[0], pathList)
			client.Workspaces(*usingSSL).Find(args[0]).UploadFile(pathList)
		}
	},
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
