package main

import (
	"fmt"
	"log"
	"os"

	"github.com/olekukonko/tablewriter"
	"go.etcd.io/bbolt"
)

func main() {
	filePath := "/home/user/my.db/reader/default.etcd/member/snap/db"

	db, err := bbolt.Open(filePath, 0666, nil)
	if err != nil {
		log.Fatalf("Error opening the database: %v", err)
	}
	defer db.Close()

	err = db.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte("key"))
		if bucket == nil {
			return fmt.Errorf("Bucket not found")
		}

		// Prepare data for table
		var data [][]string
		bucket.ForEach(func(k, v []byte) error {
			data = append(data, []string{string(k), string(v)})
			return nil
		})

		// Create table
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Key", "Value"})
		table.SetBorder(false)
		table.AppendBulk(data)
		table.Render()

		return nil
	})

	if err != nil {
		log.Fatalf("Error reading the database: %v", err)
	}
}
