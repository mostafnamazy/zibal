# \GatewayAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RequestPost**](GatewayAPI.md#RequestPost) | **Post** /request | Initiate Payment
[**StartTrackIdGet**](GatewayAPI.md#StartTrackIdGet) | **Get** /start/{trackId} | Get payment page
[**VerifyPost**](GatewayAPI.md#VerifyPost) | **Post** /verify | Verify Payment



## RequestPost

> PaymentInitResponse RequestPost(ctx).PaymentInitRequest(paymentInitRequest).Execute()

Initiate Payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mostafnamazy/zibal"
)

func main() {
	paymentInitRequest := *openapiclient.NewPaymentInitRequest("Merchant_example", int64(123), "CallbackUrl_example") // PaymentInitRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.RequestPost(context.Background()).PaymentInitRequest(paymentInitRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.RequestPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RequestPost`: PaymentInitResponse
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.RequestPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRequestPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitRequest** | [**PaymentInitRequest**](PaymentInitRequest.md) |  | 

### Return type

[**PaymentInitResponse**](PaymentInitResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartTrackIdGet

> StartTrackIdGet(ctx, trackId).Execute()

Get payment page



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mostafnamazy/zibal"
)

func main() {
	trackId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GatewayAPI.StartTrackIdGet(context.Background(), trackId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.StartTrackIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trackId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartTrackIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VerifyPost

> PaymentVerifyResponse VerifyPost(ctx).PaymentVerifyRequest(paymentVerifyRequest).Execute()

Verify Payment

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mostafnamazy/zibal"
)

func main() {
	paymentVerifyRequest := *openapiclient.NewPaymentVerifyRequest("Merchant_example", int64(123)) // PaymentVerifyRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.VerifyPost(context.Background()).PaymentVerifyRequest(paymentVerifyRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.VerifyPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VerifyPost`: PaymentVerifyResponse
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.VerifyPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVerifyPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentVerifyRequest** | [**PaymentVerifyRequest**](PaymentVerifyRequest.md) |  | 

### Return type

[**PaymentVerifyResponse**](PaymentVerifyResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

