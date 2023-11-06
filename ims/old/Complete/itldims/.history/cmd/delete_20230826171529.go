package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var delete = &cobra.Command{ // Cobra command is 'get'
	Use:   "delete",
	Short: "Delete mentioned key & value",
	Long:  " ",
	Args:  cobra.ExactArgs(0), //number of allowed arguments=1&2
	Run: func(cmd *cobra.Command, args []string) {

		var key1 string
		var numberof_key int

		args[0] = strings.ToUpper(args[0])

		////////////  anoop ///////////
		if len(args) == 1 {
			key1 = args[0]
			numberof_key = len(strings.Split(key1, "/"))
		}

		if numberof_key == 4 {
			data1, err1 := fetchDataFromAPIWithKey(key1)
			if err1 != nil {
				fmt.Printf("Failed to fetch data from the etcd API: %v", err1) // Print an error message if error occured
			}
			fmt.Println(data1)
		}

		deleteSpecificKey()
	},
}

func deleteSpecificKey(w http.ResponseWriter, r *http.Request) {
	// Extract the etcd key from the URL path
	log.Printf("response %v", r.URL.Path)

	var etcdHost = "localhost:2379"
	ctx := context.TODO()
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Printf("Failed to connect to etcd: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer etcdClient.Close()

	etcdKeyData := fmt.Sprintf(r.URL.Path)

	_, err = etcdClient.Delete(ctx, etcdKeyData, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByCreateRevision, clientv3.SortAscend))
	if err != nil {
		fmt.Printf("%v", err)
	}
}
