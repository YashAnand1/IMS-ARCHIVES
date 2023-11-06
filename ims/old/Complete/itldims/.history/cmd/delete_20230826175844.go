package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(deleteKey)
	rootCmd.Execute()
}

var deleteKey = &cobra.Command{
	Use:   "delete",
	Short: "Delete an etcd key from the API server",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		etcdKey := args[0]
		var url string = "http://localhost:8181/" + etcdKey

		resp, err := http.Delete(url)
		if err != nil {
			fmt.Printf("Failed to delete etcd key: %v\n", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			fmt.Println("Key deleted successfully")
		} else {
			fmt.Printf("Failed to delete etcd key. Status code: %d\n", resp.StatusCode)
		}
	},
}
