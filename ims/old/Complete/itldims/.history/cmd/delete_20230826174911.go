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

var deleteKey = &cobra.Command{
	Use:   "deleteKey",
	Short: "Delete mentioned key & value",
	Long:  " ",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var key1 string
		var numberof_key int

		args[0] = strings.ToUpper(args[0])

		if numberof_key == 4 {
			data1, err1 := fetchDataFromAPIWithKey(key1)
			if err1 != nil {
				fmt.Printf("Failed to fetch data from the etcd API: %v", err1)
			}
			fmt.Println(data1)
		}

		deleteSpecificKey()
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

	etcdKeyData := fmt.Sprintf(r.URL.Path)

	_, err = etcdClient.Delete(ctx, etcdKeyData, clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("Key Deleted!")
}
