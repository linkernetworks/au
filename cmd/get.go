package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(getCmd)

}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Aurora display one or many resources",
	Long: `Aurora display one or many resources
	
	Valid resource type include:
	* ws: Worksapce
`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO
		fmt.Println("Usage: au get `resource` (ws)")
	},
}
