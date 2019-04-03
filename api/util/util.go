package util

import (
	"encoding/json"
	"errors"
	"gintest/api/model"
	"io/ioutil"
	"reflect"

	"github.com/keks/go-ipfs-colog"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	orbit "github.com/keks/go-orbitdb"
)

//ValidateInsertInfo ...
func ValidateInsertInfo(c *gin.Context) (*orbit.KVStore, model.KeyValue, error) {
	var (
		insertBody model.KeyValue
		err        error
		kvDB       *orbit.KVStore
		reqPayload []byte
	)
	if reqPayload, err = ioutil.ReadAll(c.Request.Body); err == nil {
		if err = json.Unmarshal(reqPayload, &insertBody); err == nil {
			if _, err = govalidator.ValidateStruct(insertBody); err == nil {
				if kvDB, err = ConnectOrbitDB(insertBody.Topic); err != nil {
					err = errors.New("connectio error :" + err.Error())
				}
			} else {
				err = errors.New("invalid required param : " + err.Error())
			}
		} else {
			err = errors.New("unmarshal error : " + err.Error())
		}
	} else {
		err = errors.New("request body read err: " + err.Error())
	}

	return kvDB, insertBody, err
}

//InsertInfo ...
func InsertInfo(c *gin.Context, kvDB *orbit.KVStore, insrtData model.KeyValue) (interface{}, error) {

	var (
		resp    interface{}
		err     error
		newHash colog.Hash
	)
	newHash, err = kvDB.Put(insrtData.Key, insrtData.Value)
	if reflect.ValueOf(err).IsNil() {
		resp = newHash
	} else {
		err = errors.New("insertion error :" + err.Error())
		resp = `insertion failed`
	}
	return resp, err
}

//ValidateUpdateInfo ...
func ValidateUpdateInfo(c *gin.Context) (*orbit.KVStore, model.KeyValue, error) {
	var (
		updateBody model.KeyValue
		err        error
		kvDB       *orbit.KVStore
		reqPayload []byte
	)
	if reqPayload, err = ioutil.ReadAll(c.Request.Body); err == nil {
		if err = json.Unmarshal(reqPayload, &updateBody); err == nil {
			if _, err = govalidator.ValidateStruct(updateBody); err == nil {
				kvDB = GlobalDB
			} else {
				err = errors.New("invalid required param : " + err.Error())
			}
		} else {
			err = errors.New("unmarshal error : " + err.Error())
		}
	} else {
		err = errors.New("request body read err: " + err.Error())
	}
	return kvDB, updateBody, err
}

//UpdateInfo ...
func UpdateInfo(c *gin.Context, kvDB *orbit.KVStore, updateData model.KeyValue) (interface{}, error) {
	var (
		resp    interface{}
		err     error //Error
		newHash colog.Hash
	)
	_, err = kvDB.Get(updateData.Key)
	//fmt.Println(err, "time 123 pass", reflect.ValueOf(err), reflect.TypeOf(err))
	if err != nil {
		newHash, err = kvDB.Put(updateData.Key, updateData.Value)
		if reflect.ValueOf(err).IsNil() {
			resp = `key inserted successfully new hash :` + newHash
		} else {
			resp = `insertion in place of updation failed`
		}
	} else {
		newHash, err = kvDB.Put(updateData.Key, updateData.Value)
		if reflect.ValueOf(err).IsNil() {
			resp = `key updated successfully new hash: ` + newHash
		} else {
			resp = `updation failed`
		}
	}
	return resp, err
}

//ValidateDeleteInfo ...
func ValidateDeleteInfo(c *gin.Context) (*orbit.KVStore, string, error) {
	var (
		delBody    model.Delete
		err        error
		kvDB       *orbit.KVStore
		key        string
		reqPayload []byte
	)
	if reqPayload, err = ioutil.ReadAll(c.Request.Body); err == nil {
		if err = json.Unmarshal(reqPayload, &delBody); err == nil {
			if _, err = govalidator.ValidateStruct(delBody); err == nil {
				key = delBody.Key
				kvDB = GlobalDB
			} else {
				err = errors.New("invalid required param : " + err.Error())
			}
		} else {
			err = errors.New("unmarshal error : " + err.Error())
		}
	} else {
		err = errors.New("request body read err: " + err.Error())
	}
	return kvDB, key, err
}

//DeleteInfo ...
func DeleteInfo(c *gin.Context, kvDB *orbit.KVStore, key string) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)
	if _, err = kvDB.Get(key); err == nil {
		if err = kvDB.Delete(key); reflect.ValueOf(err).IsNil() {
			resp = `key deleted successfully`
		} else {
			err = errors.New("key deletion err : ")
		}
	} else {
		err = errors.New("key does not exist")
		resp = "key does not exist"
	}
	return resp, err
}

//ValidateGetKey ...
func ValidateGetKey(c *gin.Context) (*orbit.KVStore, string, error) {
	var (
		key, topic string
		kvDB       *orbit.KVStore
		err        error
	)
	key = c.Request.URL.Query().Get(`key`)
	topic = c.Request.URL.Query().Get(`topic`)
	if topic != `` {
		if key != `` {
			kvDB = GlobalDB
		} else {
			err = errors.New("key is empty")
		}
	} else {
		err = errors.New("topic is empty")
	}

	return kvDB, key, err
}

//GetKey value ...
func GetKey(c *gin.Context, kvDB *orbit.KVStore, key string) (interface{}, error) {
	var (
		resp  interface{}
		err   error
		value string
	)
	value, err = kvDB.Get(key)
	if err == nil {
		resp = value
	} else {
		err = errors.New("key does not exist " + err.Error())
		resp = "value does not exist"
	}
	return resp, err
}
