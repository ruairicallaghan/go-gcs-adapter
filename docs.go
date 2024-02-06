// We want to mock the storage client in main.go - in this example for the call:
//        client.Bucket(bucket).Objects(ctx, nil)
// This removing the need to interact with the real GCS service and allowing us to test the getObjects function in isolation.
//
// 1. Set up a ClientInterface interface with a Bucket method that returns a BucketHandleInterface
// 		- This is the interface that is "injected" into the function that is going to make the call to GCS.
//      - We are going to use this interface to mock the GCS client in our tests.
// 2. Because the GCS call is a chained method call, we need to set up an interface for each part of the call.
//      - The BucketHandleInterface is returned from the Bucket method.
//	  	- The ObjectIteratorInterface is returned from the Objects method.
// 3. Just configuring these Interfaces is not enough - we need to create a type that allows `storage.Client` to satisfy the `ClientInterface` interface.
//      - This is the Adapter pattern in action - we are creating a new set of type that adapts the behaviour of the storage library to our own interfaces - `ClientAdapter`, `BucketHandleAdapter` and `ObjectIteratorAdapter`.
// 4. We can then create a mock implementation of the `ClientInterface` interface and use it in our tests.
//      - We do not need an adapter for the mock implementation because the mock implementation already satisfies the `ClientInterface` interface.
// 5. And we can create a real implementation of the `ClientInterface` interface and use it in our main function.
//      - We do need an adapter for the real implementation because the real implementation (storage.Client) does not satisfy the `ClientInterface` interface.

package main
