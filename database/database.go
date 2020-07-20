package database

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/bitcav/nitr/models"
	bolt "go.etcd.io/bbolt"
)

//SetupDB creates nitr database with default values
func SetupDB() (*bolt.DB, error) {
	db, err := bolt.Open("nitr.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("users"))
		if err != nil {
			return fmt.Errorf("could not create root bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	return db, nil
}

//SetUserData adds User data to nitr database with default values
func SetUserData(db *bolt.DB, id string, user models.User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("could not marshal entry json: %v", err)
	}
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.Bucket([]byte("users")).Put([]byte(id), []byte(userBytes))
		if err != nil {
			return fmt.Errorf("could not insert entry: %v", err)
		}

		return nil
	})
	return err
}

//GetUserByID returns User by ID
func GetUserByID(db *bolt.DB, id string) models.User {
	var userData models.User
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		user := b.Get([]byte(id))
		if err := json.Unmarshal(user, &userData); err != nil {
			panic(err)
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return userData
}

//GetApiKey returns current User Api Key
func GetApiKey() string {
	db, err := bolt.Open("nitr.db", 0600, nil)

	if err != nil {
		fmt.Println("could not open db")
	}
	nitrUser := GetUserByID(db, "1")
	db.Close()
	return nitrUser.Apikey

}