package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/threatgrid/jqpipe-go"
	"io"
	"log"
	"os"
)

type Cache struct {
	db *bolt.DB
}

func Open() (*Cache, error) {
	db, err := bolt.Open(".tinyelk.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte("Logs"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	return &Cache{db: db}, nil
}

func (cache *Cache) Close() {
	cache.db.Close()
}

func (cache *Cache) Put(timestamp string, line []byte) {
	cache.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Logs"))
		err := b.Put([]byte(timestamp), line)
		return err
	})
}

func (cache *Cache) SomeTest() {
	stats := cache.db.Stats()
	// Encode stats to JSON and print to STDERR.
	json.NewEncoder(os.Stderr).Encode(stats)
}

// e.g all entrys for an hour
func (cache *Cache) QueryPrefix(result io.Writer, prefix string, query string) (string, error) {
	reader, writer := io.Pipe()
	jqpipe, err := jq.New(reader, query)
	if err != nil {
		return "TODO", err
	}
	defer jqpipe.Close()

	cache.db.View(func(tx *bolt.Tx) error {
		defer writer.Close()
		//errChan := make(chan error)
		//go func() {
		//	errChan <- myFTP.Stor(path, reader)
		//}()

		// Assume bucket exists and has keys
		c := tx.Bucket([]byte("Logs")).Cursor()

		log.Println(tx.Bucket([]byte("Logs")).Stats())

		prefix := []byte(prefix)
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			writer.Write(v)
			// handle err
			//err = <-errChan
			// handle err

			//fmt.Printf("key=%s, value=%s\n", k, v)
		}
		return nil
	})

	firstEntry := true
	result.Write([]byte("["))
	for {
		if rawmessage, err := jqpipe.Next(); err != nil {
			if err == io.EOF {
				result.Write([]byte("]"))
				break // end of stream
			}
			fmt.Println(err)
			return "TODO", err // jq error
		} else {

			if j, err := json.Marshal(&rawmessage); err != nil {
				fmt.Println(err)
				return "TODO", err // marshal error
			} else {
				//fmt.Println(string(j)) // successfully marshalled

				if !firstEntry {
					result.Write([]byte(","))
				}
				firstEntry = false

				result.Write(j)
			}
		}
	}

	return "TODO", nil
}
