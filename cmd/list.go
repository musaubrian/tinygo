package cmd

import (
	"fmt"

	"github.com/musaubrian/tinygo/pkg/model"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all records in the database",
	Long: `Lists all the records(sitename, username, password) available 
    in the database`,

	Run: func(cmd *cobra.Command, args []string) {
		sites := model.ListAll()
		for _, site := range sites {
			fmt.Println("\nSiteName: ", site.Name)
			fmt.Println("UserName: ", site.UserName)
			fmt.Println("Password: ", site.Password)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
