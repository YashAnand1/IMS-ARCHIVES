package main

import (
	"fmt"
	"log"
	"time"
	"go.etcd.io/bbolt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("WHO_AM_I"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		fmt.Printf("Bucket called WHO_AM_I created\n\n")
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("WHO_AM_I"))
		err := b.Put([]byte("KEY_NAME"), []byte("YASH"))
		fmt.Printf("Key Called KEY_NAME created\n")
		return err
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("WHO_AM_I"))
		v := b.Get([]byte("KEY_NAME"))
		fmt.Printf("Value of 'KEY_NAME'= %s\n", v)
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
