package client

import (
	"github.com/ctroller/yagokv/pkg/kvs"
)

type InMemoryClient struct {
	storage *kvs.Storage
}

func NewInMemoryClient(numBuckets int) Client {
	return &InMemoryClient{storage: kvs.NewStorage(numBuckets)}
}

func (c *InMemoryClient) Get(key string) (string, error) {
	return c.storage.Get(key)
}

func (c *InMemoryClient) GetAndTransform(key string, transformer Transformer[string, any]) (any, error) {
	value, err := c.Get(key)
	if err != nil {
		return nil, err
	}

	return transformer.Transform(value), nil
}

func (c *InMemoryClient) Set(key string, value string) error {
	return c.storage.Set(key, value)
}

func (c *InMemoryClient) Delete(key string) error {
	c.storage.Delete(key)
	return nil
}
