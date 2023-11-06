package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var delete = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

		if numberof_key == 5 
			var url string = addressServer + "/delete"
			//var url string = "http://192.168.122.128:8181/" + "servers/delete"

			////////////////////////

			var line string = "{" + "\"EtcdKey\"" + ":" + "\"/Physical/10.246.40.139/Hostname\"" + "}"
			fmt.Println(line)
			fmt.Println(url)
			var jsonStr = []byte(line)
			responseBody := bytes.NewBuffer(jsonStr)

			////////////////////////////////////////
			// create a new HTTP client
			client := &http.Client{}

			// create a new DELETE request
			req, err := http.NewRequest("DELETE", url, responseBody)
			if err != nil {
				panic(err)
			}

			// send the request
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode == 200 {
				fmt.Printf("Key: %s deleted succesfully", line)
			}
	},
}
