package LevelDB

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	db, err := New("")
	if err != nil {
		t.Error("connect failed")
	}
	db.Put([]byte("k1"),[]byte("v1"))
	db.Put([]byte("k2"),[]byte("v2"))
	db.Put([]byte("k3"),[]byte("v3"))
	db.Put([]byte("k4"),[]byte("v4"))
	db.Put([]byte("k5"),[]byte("v5"))

	value, err := db.Get([]byte("k1"))
	if !bytes.Equal(value,[]byte("v1")) {
		t.Error("test failed")
	}
	err = db.Delete([]byte("k2"))
	if err != nil {
		t.Error(err)
	}
	i := db.Iterator()
	for i.Next(){
		fmt.Println(string(i.Key()),string(i.Value()))
	}
}
