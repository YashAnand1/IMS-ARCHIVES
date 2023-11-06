package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var delete = &cobra.Command{ // Cobra command is 'get'
	Use:   "delete",
	Short: "Delete mentioned key & value",
	Long:  " ",

	Args: cobra.ExactArgs(0), //number of allowed arguments=1&2
	Run: func(cmd *cobra.Command, args []string) {
		var key1 string
		var numberof_key int

		args[0] = strings.ToUpper(args[0])

		////////////  anoop ///////////
		if len(args) == 1 {
			key1 = args[0]
			numberof_key = len(strings.Split(key1, "/"))
		}

		//fmt.Printf("key1:%s", key1)

		if numberof_key == 4 {
			data1, err1 := fetchDataFromAPIWithKey(key1)
			if err1 != nil {
				fmt.Printf("Failed to fetch data from the etcd API: %v", err1) // Print an error message if error occured
			}

			fmt.Println(data1)
		}

		if numberof_key != 4 {
			//////////////////////////////
			data, err := fetchDataFromAPI()
			if err != nil {
				fmt.Printf("Failed to fetch data from the etcd API: %v", err) // Print an error message if error occured
			}

			for key, value := range parseKeyValuePairs(data) { // data is gone through and key value variables entered

				// For skipping/continuing data key completely
				if strings.Contains(key, "{") || strings.Contains(key, "}") || strings.Contains(key, "data") ||
					strings.Contains(value, "{") || strings.Contains(value, "}") {
					continue
				}

				ATs := make(map[string]string) // a MAP for attributes created for storing server attributes
				splitKey := strings.Split(key, "/")
				serverAtr := splitKey[4]
				serverIP := splitKey[3]
				ATs[serverAtr] = serverAtr // Store the server attribute in the map

				// IF SERVERIP IS MENTIONED
				if strings.Contains(args[0], ".") && (strings.Contains(key, args[0])) {
					if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
						continue
					}
					fmt.Printf("%s:\n%s\n\n", serverAtr, value)
				} else { // IF SERVERIP IS NOT MENTIONED
					if strings.Contains(key, args[0]) || strings.Contains(value, args[0]) {
						if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
							continue
						}
						fmt.Printf("Server IP:\n%s\n%s:%s\n\n", serverIP, serverAtr, value) // Print server IP, attribute, and value
					}
				}
			}
		}
	},
}
