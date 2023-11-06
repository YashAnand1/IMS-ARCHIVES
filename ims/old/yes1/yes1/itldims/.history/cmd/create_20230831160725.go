package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var create = &cobra.Command{
	Use:   "create",
	Short: "Create and upload data to etcd",
	Run: func(cmd *cobra.Command, args []string) {
		uploadToEtcd()
		////  anoop new post function  ///////
		var key1 string
		var numberof_key int
		////////////  anoop ///////////
		if len(args) == 2 { // changed len from 1 to 2 to allow key-val in the form of 2 arguments
			key1 = args[0]
			numberof_key = len(strings.Split(key1, "/")) // number of key components
		}
		if numberof_key == 5 { // 5th is the value
			var url string = "http://localhost:8181/" + "servers/"
			var line string = args[0]
			etcdKey := key
			etcdValue := value
			line = "{" + "\"EtcdKey\"" + ":" + "\"" + etcdKey + "\","      //added user arg 1
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
				fmt.Printf("Key: %s has been metered as %s succesfully\n", etcdKey, string(etcdValue))
			}
			/////
		}

	},
}

func init() {
	itldims.AddCommand(create)

	create.PersistentFlags().StringVarP(&key, "key", "k", "", "A flag for the KEY")
	create.PersistentFlags().StringVarP(&value, "value", "v", "", "A flag for the VALUE")

	create.MarkPersistentFlagRequired("key")
	create.MarkPersistentFlagRequired("value")
}