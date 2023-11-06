package main

import (
	"log"

	"github.com/boltdb/bolt"
)

func main() {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Update(func(tx *bolt.Tx) error {
		// Access buckets and perform operations within the transaction
		bucket, err := tx.CreateBucketIfNotExists([]byte("mybucket"))
		if err != nil {
			return err
		}

		err = bucket.Put([]byte("key1"), []byte("value1"))
		if err != nil {
			return err
		}

		err = bucket.Put([]byte("key2"), []byte("value2"))
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
