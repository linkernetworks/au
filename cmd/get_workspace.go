package cmd

import (
	"log"

	"bitbucket.org/linkernetworks/aurora/src/cmd/au/client"
	"github.com/spf13/cobra"
)

var pageNum *int
var filterString *string

func init() {
	getCmd.AddCommand(wsCmd)

	pageNum = wsCmd.Flags().IntP("page", "p", 1, "pagination your result")
	filterString = wsCmd.Flags().StringP("filter", "f", "", "Qeury filter string")
}

var wsCmd = &cobra.Command{
	Use:   "ws",
	Short: "Query workspace status",
	Long:  `Get detail workspace detail command, include page and filter string.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := client.Workspaces(*usingSSL).Browse(*pageNum, *filterString); err != nil {
			log.Println("err=", err)
		}
	},
}
