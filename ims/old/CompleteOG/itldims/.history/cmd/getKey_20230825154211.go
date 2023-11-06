package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	getKey = &cobra.Command{
		Use:   "getKey",
		Short: "Displays values of an attribute from a server IP",
		Long:  "Find the value of a specific attribute from a Server IP",

		Run: func(cmd *cobra.Command, args []string) {

			attribute := args[0]
			serverType := "VM"
			etcdKey := fmt.Sprintf("/servers/%s/%s/%s", serverType, server, attribute)

			response, err := http.Get("http://localhost:8181" + etcdKey)
			if err != nil {
				log.Fatalf("Failed to connect to the etcd API.")
			}
			defer response.Body.Close()

			if response.StatusCode == http.StatusOK {
				body, err := io.ReadAll(response.Body)
				if err != nil {
					log.Fatalf("Failed to read response body: %v", err)
				}
				fmt.Printf("Attribute value for server IP %s and attribute %s: %s\n", server, attribute, string(body))
			}
		},
	}
)

func init() {
	getKey.AddCommand(attributes)
	getKey.AddCommand(Types)
	getKey.AddCommand(servers)
}
