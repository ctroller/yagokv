package client

// Transformer defines a type that can transform an input of type I to an output of type T.
type Transformer[I any, T any] interface {
	Transform(in I) T
}

// Client defines the interface for a key-value store client.
type Client interface {
	Get(key string) (string, error)
	GetAndTransform(key string, transformer Transformer[string, any]) (any, error)
	Set(key string, value string) error
	Delete(key string) error
}
