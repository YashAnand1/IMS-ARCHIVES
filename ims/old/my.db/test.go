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

	key := []byte("NAME")
	value := []byte("YashAnand")

	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("WHO_AM_I"))

		oldValue := b.Get(key)
		fmt.Printf("Previous value of %s: %s\n", key, oldValue)

		return b.Put(key, value)
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New value of %s: %s\n", key, value)
}
