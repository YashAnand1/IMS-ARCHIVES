package cmd

import (
	// ... (import statements)

	"context"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete mentioned key & value",
	Long:  " ",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		var key1 string
		var numberof_key int

		args[0] = strings.ToUpper(args[0])

		// ... (other code)

		if numberof_key == 4 {
			data1, err1 := fetchDataFromAPIWithKey(key1)
			if err1 != nil {
				fmt.Printf("Failed to fetch data from the etcd API: %v", err1)
			}
			fmt.Println(data1)
		}

		deleteSpecificKey() // Call the function here
	},
}

func deleteSpecificKey() {
	var etcdHost = "localhost:2379"
	ctx := context.TODO()
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Printf("Failed to connect to etcd: %v", err)
		return
	}
	defer etcdClient.Close()

	etcdKeyData := "your_etcd_key_here"

	_, err = etcdClient.Delete(ctx, etcdKeyData, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByCreateRevision, clientv3.SortAscend))
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("Key Deleted!")
}
