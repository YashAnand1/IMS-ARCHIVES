package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var CREATE = &cobra.Command{
	Use:   "CREATE",
	Short: "Create and upload data to server",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CREATE called")
		//uploadToEtcd()

		var url string = addressServer + "/IMS/Anoop"
		etcdKey := key
		etcdValue := value

		line := "{" + "\"EtcdKey\"" + ":" + "\"" + etcdKey + "\","     //added user arg 1
		line = line + "\"EtcdValue\"" + ":" + "\"" + etcdValue + "\"}" //added user arg 2

		fmt.Println(line)
		fmt.Println(url)
		var jsonStr = []byte(line)
		responseBody := bytes.NewBuffer(jsonStr)
		
		resp, err := http.Post(url, "application/json", responseBody)
		if err != nil {
			fmt.Println(err)
		}
		if err == nil {
			defer resp.Body.Close()
		}

		if resp.StatusCode == 200 {
			// fmt.Printf("Key: %s is netered as %s succesfully", "servers/Physical/10.246.40.139/Hostname", "vahanapp00")
			fmt.Printf("Key: %s has been metered as %s succesfully\n", strings.ToUpper(etcdKey), strings.ToUpper(etcdValue))

		}

	},
}

var key string
var value string

func init() {
	itldims.AddCommand(CREATE)

	CREATE.PersistentFlags().StringVarP(&key, "key", "k", "", "A flag for the KEY")
	CREATE.PersistentFlags().StringVarP(&value, "value", "v", "", "A flag for the VALUE")

}
