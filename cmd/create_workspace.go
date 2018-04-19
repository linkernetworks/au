package cmd

import (
	"fmt"
	"log"

	"github.com/linkernetworks/au/client"
	"github.com/spf13/cobra"
)

var numPage *int
var wsType *string

func init() {
	createCmd.AddCommand(crewsCmd)

	numPage = crewsCmd.Flags().IntP("numPreview", "n", 0, "Number items for preview")
	wsType = crewsCmd.Flags().StringP("workspaceType", "t", "general", "Specific workspace type")
}

var crewsCmd = &cobra.Command{
	Use:   "ws",
	Short: "Create workspace via datasetIDs",
	Long: `Create workspace
	
	To create workspace with specific dataset ID: 
	usege: "au create ws datasetID1,  datasetID2 ..."
	
	To create workspace with preview number of dataset:
	usege: "au create ws -n=100 datasetID1,  datasetID2 ..."
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		if _, err := client.Workspaces(*usingSSL).Create(*wsType, *numPage, args); err != nil {
			log.Println("Could not create:", err)
		}
	},
}
