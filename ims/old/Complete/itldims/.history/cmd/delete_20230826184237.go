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
	Long:  "deleteKe",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		args[0] = strings.ToUpper(args[0])

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

	var key1 string
	var numberof_key int
	if numberof_key == 4 {
		data1, err1 := fetchDataFromAPIWithKey(key1)
		if err1 != nil {
			fmt.Printf("Failed to fetch data from the etcd API: %v", err1)
		}
		fmt.Println(data1)
	}
	etcdKeyData := key1

	_, err = etcdClient.Delete(ctx, etcdKeyData, clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("Key Deleted!")
}
