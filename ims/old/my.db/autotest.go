package main

import (
	"fmt"
	"log"
	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

 	db.View(func(tx *bolt.Tx) error {
	b := tx.Bucket([]byte("WHO_AM_I"))
 		err := b.ForEach(func(k, v []byte) error {
			if string(k) == "KEY_NAME_1" {
				fmt.Printf("Key: %s, Value: %s\n", k, v)
			}
			return nil
		})
		return err
	})
 	if err != nil {
		log.Fatal(err)

}}
