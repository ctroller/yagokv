// File: internal/kvs/storage_test.go
package kvs

import (
	"strconv"
	"sync"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	s := NewStorage(16)
	key := "testKey"
	value := "testValue"

	// Set value and retrieve it
	err := s.Set(key, value)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	got, err := s.Get(key)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	if got != value {
		t.Errorf("expected value %q, got %q", value, got)
	}
}

func TestGetNonExistingKey(t *testing.T) {
	s := NewStorage(16)
	key := "nonExistent"

	// Get a key that was never set. Expect empty string.
	got, err := s.Get(key)
	if err != nil {
		t.Fatalf("Get returned error for non-existing key: %v", err)
	}

	if got != "" {
		t.Errorf("expected empty string for non-existing key, got %q", got)
	}
}

func TestDeleteKey(t *testing.T) {
	s := NewStorage(16)
	key := "deleteMe"
	value := "toBeDeleted"

	// Set, then delete, then get.
	err := s.Set(key, value)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	s.Delete(key)

	got, err := s.Get(key)
	if err != nil {
		t.Fatalf("Get failed after delete: %v", err)
	}

	if got != "" {
		t.Errorf("expected empty string after deleting key, got %q", got)
	}
}

func TestConcurrentAccess(t *testing.T) {
	s := NewStorage(32)
	numGoroutines := 50
	numIterations := 100
	var wg sync.WaitGroup

	// Concurrent sets
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				key := "key_" + strconv.Itoa(id) + "_" + strconv.Itoa(j)
				value := "value_" + strconv.Itoa(id) + "_" + strconv.Itoa(j)
				if err := s.Set(key, value); err != nil {
					t.Errorf("concurrent Set failed: %v", err)
				}
			}
		}(i)
	}
	wg.Wait()

	// Concurrent gets
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numIterations; j++ {
				key := "key_" + strconv.Itoa(id) + "_" + strconv.Itoa(j)
				expected := "value_" + strconv.Itoa(id) + "_" + strconv.Itoa(j)
				got, err := s.Get(key)
				if err != nil {
					t.Errorf("concurrent Get failed: %v", err)
				}
				if got != expected {
					t.Errorf("for key %q, expected %q but got %q", key, expected, got)
				}
			}
		}(i)
	}
	wg.Wait()
}
