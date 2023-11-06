package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/cobra"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var etcdHost = "localhost:2379"

func deleteSpecificKeyCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a specific key from etcd",
		RunE: func deleteSpecificKey(cmd *cobra.Command, args []string) error {
	// Extract the etcd key from the command arguments
	if len(args) != 1 {
		return fmt.Errorf("usage: delete <etcd_key>")
	}
	etcdKeyData := args[0]

	ctx := context.TODO()
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: []string{etcdHost},
	})
	if err != nil {
		log.Printf("Failed to connect to etcd: %v", err)
		return err
	}
	defer etcdClient.Close()

	_, err = etcdClient.Delete(ctx, etcdKeyData, clientv3.WithPrefix(), clientv3.WithSort(clientv3.SortByCreateRevision, clientv3.SortAscend))
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}

	fmt.Println("Keys have been deleted")
	return nil
},
}

