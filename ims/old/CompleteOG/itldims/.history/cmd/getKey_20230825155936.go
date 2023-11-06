package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	getKey = &cobra.Command{
		Use:   "getKey",
		Short: "Displays values of an attribute from a server IP",
		Long:  "Find the value of a specific attribute from a Server IP",

		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 3 {
				fmt.Println("Usage: itldims getKey <server_Type <server_IP> <attribute_name>")
			}

			serverType := args[0]
			server := args[1]
			attribute := args[2]
			etcdKey := fmt.Sprintf("/servers/%s/%s/%s", serverType, server, attribute)

			response, err := http.Get("http://localhost:8181" + etcdKey)
			if err != nil {
				fmt.Printf("Failed to connect to the etcd API.")
			}
			defer response.Body.Close()

			if response.StatusCode == http.StatusOK {
				body, err := io.ReadAll(response.Body)
				if err != nil {
					fmt.Printf("Failed to read response body: %v", err)
				}
				fmt.Printf("Attribute: %s\n%s\n", attribute, string(body))
			}
		},
	}
)

func init() {
	getKey.AddCommand(attributes)
	getKey.AddCommand(Types)
	getKey.AddCommand(servers)
}
