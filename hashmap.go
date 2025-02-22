package gohashmap

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
)

const DEFAULT_HASHMAP_SIZE int = 16

type Bucket[K comparable, V any] struct {
	key   K
	value V
	next  *Bucket[K, V]
}

type HashMap[K comparable, V any] struct {
	buckets []*Bucket[K, V]
	size    int
}

func NewHashMap[K comparable, V any]() *HashMap[K, V] {
	size := DEFAULT_HASHMAP_SIZE
	buckets := make([]*Bucket[K, V], size)
	return &HashMap[K, V]{buckets, size}
}

func NewHashmapWithSize[K comparable, V any](size int) *HashMap[K, V] {
	buckets := make([]*Bucket[K, V], size)
	return &HashMap[K, V]{buckets, size}
}

func (h *HashMap[K, V]) hash(key K) (uint32, error) {
	hash, err := h.hashKey(key)
	if err != nil {
		return hash, err
	}

	i := hash & uint32((h.size - 1))
	return i, nil
}

func (h *HashMap[K, V]) hashKey(key K) (uint32, error) {
	b, err := anyToBytes(key)
	if err != nil {
		return 0, err
	}

	ha := fnv.New32a()
	ha.Write(b)
	return ha.Sum32(), nil
}

func (h *HashMap[K, V]) Put(k K, v V) error {
	index, err := h.hash(k)
	if err != nil {
		return err
	}

	head := h.buckets[index]

	for bucket := head; bucket != nil; bucket = bucket.next {
		if bucket.key == k {
			bucket.value = v
			return nil
		}
	}

	newBucket := &Bucket[K, V]{
		key:   k,
		value: v,
		next:  head,
	}
	h.buckets[index] = newBucket
	return nil
}

func (h *HashMap[K, V]) Get(k K) (V, bool) {
	var val V

	index, err := h.hash(k)
	if err != nil {
		return val, false
	}

	head := h.buckets[index]
	for bucket := head; bucket != nil; bucket = bucket.next {
		if bucket.key == k {
			return bucket.value, true
		}
	}

	return val, false
}

func (h *HashMap[K, V]) Delete(k K) bool {
	index, err := h.hash(k)
	if err != nil {
		return false
	}

	head := h.buckets[index]

	if head.key == k {
		h.buckets[index] = head.next
		return true
	}

	prev := head

	for bucket := head; bucket != nil; bucket = bucket.next {
		if bucket.key == k {
			prev.next = bucket.next
			return true
		}
		prev = bucket
	}

	return false
}

func anyToBytes(val any) ([]byte, error) {
	switch v := val.(type) {
	case string:
		return []byte(v), nil
	case []byte:
		return v, nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return []byte(fmt.Sprintf("%d", v)), nil
	case float32, float64:
		return []byte(fmt.Sprintf("%f", v)), nil
	case bool:
		return []byte(fmt.Sprintf("%t", v)), nil
	default:
		jsonBytes, err := json.Marshal(v)
		if err != nil {
			return nil, fmt.Errorf("unsupported type: %T, failed to marshal to JSON: %w", v, err)
		}
		return jsonBytes, nil
	}
}
