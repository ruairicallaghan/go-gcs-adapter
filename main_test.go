package main

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/storage"
)

type MockObjectIterator struct{}

func (m *MockObjectIterator) Next() (*storage.ObjectAttrs, error) {
	// Return a mock object attribute or an error as needed
	return nil, nil
}

type MockBucketHandle struct{}

func (m *MockBucketHandle) Objects(_ context.Context, _ *storage.Query) ObjectIteratorInterface {
	fmt.Println("Running Mocked action.")
	return &MockObjectIterator{}
}

type MockClient struct{}

func (m *MockClient) Bucket(name string) BucketHandleInterface {
	return &MockBucketHandle{}
}

func TestGetObjects(t *testing.T) {
	c := &MockClient{}
	getObjects(c, context.Background(), "bucket", nil)
}
