package LevelDB

import (
	"fmt"
	"testing"
)

func TestNewIterator(t *testing.T) {
	testData:=map[string][]byte{
		"k1":[]byte("v1"),
		"k2":[]byte("v2"),
	}
	iterator := NewIterator(testData)
	if iterator.length != 2 {
		t.Fatal()
	}
	for iterator.Next() {
		fmt.Printf("%s:%s\n",string(iterator.Key()),string(iterator.Value()) )
	}
}
