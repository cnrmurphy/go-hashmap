package gohashmap

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
)

type Entry[T any] struct {
	ID    int
	Value T
}

func TestPutString(t *testing.T) {
	hashmap := NewHashmap[string, string]()
	err := hashmap.Put("hello", "world")
	if err != nil {
		t.Error(err)
	}
	spew.Dump(hashmap)
}

func TestPutStruct(t *testing.T) {
	hashmap := NewHashmap[int, *Entry[string]]()
	entry := &Entry[string]{
		ID:    0,
		Value: "foo",
	}

	err := hashmap.Put(entry.ID, entry)
	if err != nil {
		t.Error(err)
	}
	spew.Dump(hashmap)
}

func TestGet(t *testing.T) {
	hashmap := NewHashmap[string, string]()
	hashmap.Put("hello", "world")
	v, ok := hashmap.Get("hello")
	if !ok {
		t.Errorf("expected to find entry with key \"%s\", but it was not found", "hello")
	}
	spew.Dump(v)
}
