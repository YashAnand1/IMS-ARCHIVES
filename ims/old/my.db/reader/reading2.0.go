// package main

// import (
// 	"encoding/binary"
// 	"fmt"
// 	"log"

// 	"go.etcd.io/bbolt"
// )

// func main() {
// 	filePath := "/home/user/my.db/reader/default.etcd/member/snap/db"

// 	db, err := bbolt.Open(filePath, 0666, nil)
// 	if err != nil {
// 		log.Fatalf("Error opening the database: %v", err)
// 	}
// 	defer db.Close()

// 	err = db.View(func(tx *bbolt.Tx) error {

// 		bucketName := []byte("key")
// 		bucket := tx.Bucket(bucketName)

// 		fmt.Printf("Current Bucket: %s\n\nAll existing buckets:\n", bucketName)

// 		err = tx.ForEach(func(name []byte, _ *bbolt.Bucket) error {
// 			fmt.Println(string(name))
// 			return nil
// 		})

// 		fmt.Printf("\nKey-Value Data From The Bucket Is As Follows\n")

// 		err = bucket.ForEach(func(k, v []byte) error {
// 			versionBytes := v[:8]
// 			value := v[8:]
// 			version := binary.LittleEndian.Uint64(versionBytes) // Use LittleEndian for version

// 			fmt.Printf("Key: %s, Value: %s, Version: %d\n", k, value, version)
// 			return nil
// 		})

// 		return nil
// 	})

// 	if err != nil {
// 		log.Fatalf("Error reading the database: %v", err)
// 	}
// }
