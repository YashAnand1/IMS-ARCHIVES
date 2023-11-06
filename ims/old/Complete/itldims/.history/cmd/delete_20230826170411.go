package cmd

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/cobra"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [key]",
	Short: "Delete mentioned key & value",
	Run:   delete,
}

func delete(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		fmt.Println("Usage: delete [key]")
		return
	}

	key1 := strings.ToUpper(args[0])
	numberof_key := len(strings.Split(key1, "/"))

	if numberof_key == 4 {
		data1, err1 := fetchDataFromAPIWithKey(key1)
		if err1 != nil {
			fmt.Printf("Failed to fetch data from the etcd API: %v", err1)
			return
		}

		fmt.Println(data1)

		// Perform deletion logic here
		ctx := context.TODO()
		etcdClient, err := clientv3.New(clientv3.Config{
			Endpoints: []string{etcdHost},
		})
		if err != nil {
			log.Printf("Failed to connect to etcd: %v", err)
			return
		}
		defer etcdClient.Close()

		etcdKeyData := key1

		_, err = etcdClient.Delete(ctx, etcdKeyData, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByCreateRevision, clientv3.SortAscend))
		if err != nil {
			fmt.Printf("%v", err)
			return
		}

		fmt.Println("Keys have been deleted")
	}
}

func fetchDataFromAPIWithKey(key string) (string, error) {
	// Implement your logic here to fetch data from the etcd API
	// and return the data along with any error that occurred.
	return "", nil
}

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(deleteCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
