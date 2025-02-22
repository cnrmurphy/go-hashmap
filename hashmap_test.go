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
	hashmap := NewHashMap[string, string]()
	err := hashmap.Put("hello", "world")
	if err != nil {
		t.Error(err)
	}
}

func TestPutStruct(t *testing.T) {
	hashmap := NewHashMap[int, *Entry[string]]()
	entry := &Entry[string]{
		ID:    0,
		Value: "foo",
	}

	err := hashmap.Put(entry.ID, entry)
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	hashmap := NewHashMap[string, string]()
	hashmap.Put("hello", "world")
	_, ok := hashmap.Get("hello")
	if !ok {
		t.Errorf("expected to find entry with key \"%s\", but it was not found", "hello")
	}
}

func TestDelete(t *testing.T) {
	hashmap := NewHashMap[string, string]()
	hashmap.Put("hello", "world")
	hashmap.Put("goodbye", "world")
	spew.Dump(hashmap)
	ok := hashmap.Delete("goodbye")
	if !ok {
		t.Errorf("expected to have deleted entry with key \"%s\", but it was not deleted", "goodbye")
	}

	_, ok = hashmap.Get("goodbye")
	if ok {
		t.Errorf("expected entry with key \"%s\", to have been deleted but it was retrieved", "goodbye")
	}
	spew.Dump(hashmap)
}

func TestDeleteNonexistentKey(t *testing.T) {
	hashmap := NewHashMap[string, string]()
	hashmap.Put("hello", "world")
	ok := hashmap.Delete("goodbye")
	if ok {
		t.Error("expected attempt to delete nonexistent key to return false")
	}
}
