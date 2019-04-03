package util

import (
	"fmt"
	"log"

	orbit "github.com/keks/go-orbitdb"
)

//GlobalDB ...
var GlobalDB *orbit.KVStore

//ConnectOrbitDB ...
func ConnectOrbitDB(topic string) (*orbit.KVStore, error) {
	var kvDB *orbit.KVStore
	orbitDB, err := orbit.NewOrbitDB(topic) //connect to given database
	if err == nil {
		kvDB = orbit.NewKVStore(orbitDB)
		log.Println("connected to ipfs node")
	} else {
		log.Println("err in connecting to orbit db", err)
	}
	fmt.Println("-----------------------______NITP_______----------------------")
	GlobalDB = kvDB
	return kvDB, err
}
