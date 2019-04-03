package controller

import (
	"gintest/api/model"
	"gintest/api/util"
	"reflect"

	"github.com/gin-gonic/gin"
	orbit "github.com/keks/go-orbitdb"
)

//InsertInfo ...
func InsertInfo(c *gin.Context) {
	var (
		kvDB      *orbit.KVStore
		insrtData model.KeyValue
		err       error
		response  interface{}
	)
	if kvDB, insrtData, err = util.ValidateInsertInfo(c); err == nil {
		response, err = util.InsertInfo(c, kvDB, insrtData)
	}
	if err != nil {
		if reflect.ValueOf(err).IsNil() {
			c.JSON(200, gin.H{"data": response})
		} else {
			c.JSON(200, gin.H{"err": err.Error(), "data": response})
		}
	} else {
		c.JSON(200, gin.H{"data": response})
	}

}

//UpdateInfo ...
func UpdateInfo(c *gin.Context) {
	var (
		kvDB       *orbit.KVStore
		updateData model.KeyValue
		response   interface{}
		err        error
	)
	if kvDB, updateData, err = util.ValidateUpdateInfo(c); err == nil {
		response, err = util.UpdateInfo(c, kvDB, updateData)
	}
	if err != nil {
		if reflect.ValueOf(err).IsNil() {
			c.JSON(200, gin.H{"data": response})
		} else {
			c.JSON(200, gin.H{"err": err.Error(), "data": response})
		}
	} else {
		c.JSON(200, gin.H{"data": response})
	}
	//c.JSON(200, gin.H{"data": response})
}

//DeleteInfo ...
func DeleteInfo(c *gin.Context) {
	var (
		kvDB     *orbit.KVStore
		key      string
		response interface{}
		err      error
	)
	if kvDB, key, err = util.ValidateDeleteInfo(c); err == nil {
		response, err = util.DeleteInfo(c, kvDB, key)
	}
	if err != nil {
		if reflect.ValueOf(err).IsNil() {
			c.JSON(200, gin.H{"data": response})
		} else {
			c.JSON(200, gin.H{"err": err.Error(), "data": response})
		}
	} else {
		c.JSON(200, gin.H{"data": response})
	}

}

//GetKey ...
func GetKey(c *gin.Context) {
	var (
		kvDB     *orbit.KVStore
		key      string
		response interface{}
		err      error
	)
	if kvDB, key, err = util.ValidateGetKey(c); err == nil {
		response, err = util.GetKey(c, kvDB, key)
	}

	if err != nil {
		if reflect.ValueOf(err).IsNil() {
			c.JSON(200, gin.H{"data": response})
		} else {
			c.JSON(200, gin.H{"err": err.Error(), "data": response})
		}
	} else {
		c.JSON(200, gin.H{"data": response})
	}
}
