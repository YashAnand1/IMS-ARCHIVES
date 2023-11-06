package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var get = &cobra.Command{
	Use:   "get",
	Short: "Search Attributes & Values from etcd API",
	Long: `Data retrieval can be done using the following combinations:
 - 'itldims get <KeyComponent> <KeyComponent/Value>'
 - 'itldims get <KeyComponent/Value>'
`,
	Args: cobra.RangeArgs(1, 2),
	Run: func(cmd *cobra.Command, args []string) {

		data, err := fetchDataFromAPI()
		if err != nil {
			fmt.Printf("Failed to fetch data from the etcd API: %v", err)
		}

		if args[0] == "hardware" {
			displayAllHardwareAttributes(ddata)
			return
		}

		for key, value := range parseKeyValuePairs(data) {

			if strings.Contains(key, "{") || strings.Contains(key, "}") || strings.Contains(key, "data") ||
				strings.Contains(value, "{") || strings.Contains(value, "}") {
				continue
			}

			ATs := make(map[string]string)
			splitKey := strings.Split(key, "/")
			serverAtr := splitKey[4]
			serverIP := splitKey[3]
			ATs[serverAtr] = serverAtr

			if strings.Contains(args[0], ".") && (strings.Contains(key, args[0])) {
				if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
					continue
				}
				fmt.Printf("%s:\n%s\n\n", serverAtr, value)
			} else {
				if strings.Contains(key, args[0]) || strings.Contains(value, args[0]) {
					if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
						continue
					}
					fmt.Printf("Server IP:\n%s\n%s:%s\n\n", serverIP, serverAtr, value)
				}
			}
		}
	},
}

func displayAllHardwareAttributes(data map[string]string) {
	attributesToDisplay := []string{"RAM", "CPU", "External_Disk"}
	for key, value := range parseKeyValuePairs(data) {
		if strings.Contains(key, "Physical") {
			for _, attr := range attributesToDisplay {
				if strings.Contains(key, attr) {
					splitKey := strings.Split(key, "/")
					serverIP := splitKey[3]
					fmt.Printf("Server IP:\n%s\n%s:%s\n\n", serverIP, attr, value)
				}
			}
		}
	}
}

func init() {
	get.AddCommand(attributes)
	get.AddCommand(Types)
	get.AddCommand(servers)
}
