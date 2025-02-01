package kvs

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

type SerializedObject struct {
	data []byte
}

type Bucket struct {
	data map[string]*SerializedObject
}

type Storage struct {
	buckets []Bucket
}

func NewStorage(numBuckets int) *Storage {
	buckets := make([]Bucket, numBuckets)

	for i := range buckets {
		buckets[i] = Bucket{data: make(map[string]*SerializedObject)}
	}
	return &Storage{
		buckets: buckets,
	}
}

func serialize(value string) (*SerializedObject, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(value)

	if err != nil {
		return nil, err
	}

	return &SerializedObject{
		data: buffer.Bytes(),
	}, nil
}

func deserialize(serializedValue *SerializedObject) (string, error) {
	buffer := bytes.NewBuffer(serializedValue.data)
	decoder := gob.NewDecoder(buffer)

	var value string
	err := decoder.Decode(&value)

	if err != nil {
		return "", err
	}

	return value, nil
}

func (s *Storage) getHash(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key))
	return h.Sum64()
}

func (s *Storage) getBucketIndex(key string) int {
	return int(s.getHash(key) % uint64(len(s.buckets)))
}

func (s *Storage) Get(key string) (string, error) {
	val := s.buckets[s.getBucketIndex(key)].data[key]
	if val == nil {
		return "", nil
	}

	return deserialize(val)
}

func (s *Storage) Set(key string, value string) error {
	serializedValue, err := serialize(value)
	if err != nil {
		return err
	}

	s.buckets[s.getBucketIndex(key)].data[key] = serializedValue
	return nil
}

func (s *Storage) Delete(key string) {
	delete(s.buckets[s.getBucketIndex(key)].data, key)
}
