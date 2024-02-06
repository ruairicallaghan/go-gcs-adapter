package main

import (
	"context"
	"fmt"
	"io"

	"cloud.google.com/go/storage"
)

type ClientInterface interface {
	Bucket(name string) BucketHandleInterface
}

type BucketHandleInterface interface {
	Objects(ctx context.Context, q *storage.Query) ObjectIteratorInterface
}

type ObjectIteratorInterface interface {
	Next() (*storage.ObjectAttrs, error)
}

func main() {
	c, _ := storage.NewClient(context.Background())
	g := NewClientAdapter(c)
	getObjects(g, context.Background(), "bucket", nil)
}

func getObjects(gcs ClientInterface, ctx context.Context, bucket string, w io.Writer) {
	_ = gcs.Bucket(bucket).Objects(ctx, nil)
	// it := s.storageClient.Bucket(bucket).Objects(ctx, nil)
	// for {
	// 	attrs, err := it.Next()
	// 	if err == iterator.Done {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Fprintln(w, attrs.Name)
	// }
}

// Ensure storage.Client can be used where ClientInterface is expected - this is a compile time check to ensure
// that ClientAdapter implements the ClientInterface. Enforcing type-safety at compile time.
var _ ClientInterface = (*ClientAdapter)(nil)

// ClientAdapter is an adapter for the storage.Client type
type ClientAdapter struct {
	client *storage.Client
}

// NewClientAdapter returns a new ClientAdapter. This is the part of the Adapter pattern enabling the creation
// of a ClientAdapter instance that can then adapt the storage.Client behaviour or interface.
func NewClientAdapter(client *storage.Client) *ClientAdapter {
	return &ClientAdapter{client: client}
}

func (c *ClientAdapter) Bucket(name string) BucketHandleInterface {
	return NewBucketHandleAdapter(c.client.Bucket(name))
}

type BucketHandleAdapter struct {
	bucketHandle *storage.BucketHandle
}

func NewBucketHandleAdapter(bucketHandle *storage.BucketHandle) *BucketHandleAdapter {
	return &BucketHandleAdapter{bucketHandle: bucketHandle}
}

func (b *BucketHandleAdapter) Objects(ctx context.Context, q *storage.Query) ObjectIteratorInterface {
	fmt.Println("BucketHandleAdapter.Objects")
	return &ObjectIteratorAdapter{
		objectIterator: b.bucketHandle.Objects(ctx, q),
	}
}

type ObjectIteratorAdapter struct {
	objectIterator *storage.ObjectIterator
}

func (o *ObjectIteratorAdapter) Next() (*storage.ObjectAttrs, error) {
	return o.objectIterator.Next()
}
