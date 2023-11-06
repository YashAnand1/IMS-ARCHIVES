package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var (
	itldims = &cobra.Command{ //Command will check if connection with API URL is set
		Use:   "itldims",
		Short: "Interact with the etcd API",
		Long:  "A command-line tool to interact with the etcd API and check connection",
		Run: func(cmd *cobra.Command, args []string) { //here it is defined that when the itldims command is run, it will take a single argument
			response, err := http.Get("http://localhost:8181/servers/") //The response from API is stored in response & error variable
			if err != nil {                                             //if error variable is filled, then the following will be logged
				log.Fatalf("Failed to connect to the etcd API.")
			}
			defer response.Body.Close() //The itldims command & localhost:8181 will be then closed

			if response.StatusCode == http.StatusOK {
				fmt.Println("Connected to API. Interaction with etcd can be done.") //However if the statuscode is ok and connection is set then, message displayed
			}
		},
	}

	keyOnly bool
)

var get = &cobra.Command{ //the command will be used for retrieving data by filtering content of the API Server
	Use:   "get",
	Short: "Search Attributes & Values from etcd API",
	Long: `Data retrieval can be done using 'itldims get <input1> <input2>' or 'itldims get <input1>'.

Command combinations that can be utilised:
- itldims get <Servers>          		| Displays all the running Servers with their Server IPs | Working
- itldims get <Server IP>        		| Displays all Attribute values of a specific Server IP | Working
- itldims get <Server Type>        		| Displays all Attribute values of a specific Server Type | Working
- itldims get <Attribute>        		| Display Servers with a specific Attribute
- itldims get <Value>            		| Display Servers with a specific Attribute value
- itldims get <Server Type> <Attribute> | Display specific Attribute values of a Server Type |Working
- itldims get <Server Type> <Value>     | Display Server Types containing a specific value
- itldims get <Value> <Server Type>     | Display Server Types containing a specific value
- itldims get <Attribute> <Server IP>   | Display specific Attribute values of a Server IP
- itldims get <Server IP> <Attribute>   | Display specific Attribute values of a Server IP
- itldims get <Server IP> <Value>       | Display Server IPs containing a specific value
- itldims get <Value> <Server IP>       | Display Server IPs containing a specific value
- itldims get <Server IP> <Server Type> | Display Attribute values of a specific Server
	`,
	Args: cobra.RangeArgs(1, 2), //Here it is defined that 1 or 2 arguments are only allowed
	Run: func(cmd *cobra.Command, args []string) { //This is the function which gets executed on running get get <argumetns> comand
		data, err := fetchDataFromAPI() //The data from the API is fetched and stored in the data variable
		if err != nil {                 // if there is an erorr, its logged in err variable
			log.Fatalf("Failed to fetch data from the etcd API: %v", err)
		}

		if len(args) == 1 { //Here it is defined that if only one argument is entered, the second argument of 'servers' will be added to the
			args = append(args, "servers")
		}

		for key, value := range parseKeyValuePairs(data) {
			if strings.Contains(key, "{") || strings.Contains(key, "}") ||
				strings.Contains(value, "{") || strings.Contains(value, "}") {
				continue
			}

			if !strings.Contains(key, "data") && strings.Contains(key, args[0]) || strings.Contains(value, args[0]) {
				if len(args) > 1 && !strings.Contains(key, args[1]) && !strings.Contains(value, args[1]) {
					continue
				}

				fmt.Println(key) // at the end
				if !keyOnly {
					fmt.Println(value)
				}

				// fmt.println()
			}
		}
	},
}

func fetchDataFromAPI() (string, error) {
	response, err := http.Get("http://localhost:8181/servers/")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch data. Status code: %d", response.StatusCode)
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func parseKeyValuePairs(data string) map[string]string {
	result := make(map[string]string)

	keyValuePairs := strings.Split(data, "Key:")

	for _, kv := range keyValuePairs {
		kv = strings.TrimSpace(kv)
		if len(kv) == 0 {
			continue
		}

		lines := strings.Split(kv, "Value:")
		if len(lines) == 2 {
			key := strings.TrimSpace(lines[0])
			value := strings.TrimSpace(lines[1])
			result[key] = value
		}
	}
	return result
}

func init() {
	itldims.AddCommand(get)
	get.Flags().BoolVar(&keyOnly, "listall", false, "Display only keys without values") //this is not part of cobra but works
}

func main() {
	if err := itldims.Execute(); err != nil {
		log.Fatal(err)
	}
}
