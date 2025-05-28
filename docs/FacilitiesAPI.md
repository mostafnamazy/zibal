# \FacilitiesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FacilityAccountToIbanPost**](FacilitiesAPI.md#FacilityAccountToIbanPost) | **Post** /facility/accountToIban | Bank Account To IBAN
[**FacilityCardToAccountPost**](FacilitiesAPI.md#FacilityCardToAccountPost) | **Post** /facility/cardToAccount | Card To Bank Account
[**FacilityCardToIbanPost**](FacilitiesAPI.md#FacilityCardToIbanPost) | **Post** /facility/cardToIban | Card To IBAN
[**FacilityCheckCardWithNationalCodePost**](FacilitiesAPI.md#FacilityCheckCardWithNationalCodePost) | **Post** /facility/checkCardWithNationalCode | Check Card With National Code
[**FacilityNationalIdentityInquiryPost**](FacilitiesAPI.md#FacilityNationalIdentityInquiryPost) | **Post** /facility/nationalIdentityInquiry | National Identity Inquiry
[**FacilityShahkarInquiryPost**](FacilitiesAPI.md#FacilityShahkarInquiryPost) | **Post** /facility/shahkarInquiry | Shahkar Inquiry
[**WalletCheckoutPlusPost**](FacilitiesAPI.md#WalletCheckoutPlusPost) | **Post** /wallet/checkout/plus | Wallet Checkout Plus



## FacilityAccountToIbanPost

> BankAccountToIbanResponse FacilityAccountToIbanPost(ctx).BankAccountToIbanRequest(bankAccountToIbanRequest).Execute()

Bank Account To IBAN

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
	bankAccountToIbanRequest := *openapiclient.NewBankAccountToIbanRequest("BankAccount_example", "BankCode_example") // BankAccountToIbanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacilitiesAPI.FacilityAccountToIbanPost(context.Background()).BankAccountToIbanRequest(bankAccountToIbanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacilitiesAPI.FacilityAccountToIbanPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FacilityAccountToIbanPost`: BankAccountToIbanResponse
	fmt.Fprintf(os.Stdout, "Response from `FacilitiesAPI.FacilityAccountToIbanPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFacilityAccountToIbanPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankAccountToIbanRequest** | [**BankAccountToIbanRequest**](BankAccountToIbanRequest.md) |  | 

### Return type

[**BankAccountToIbanResponse**](BankAccountToIbanResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FacilityCardToAccountPost

> CardToAccountResponse FacilityCardToAccountPost(ctx).CardToAccountRequest(cardToAccountRequest).Execute()

Card To Bank Account

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
	cardToAccountRequest := *openapiclient.NewCardToAccountRequest("CardNumber_example") // CardToAccountRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacilitiesAPI.FacilityCardToAccountPost(context.Background()).CardToAccountRequest(cardToAccountRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacilitiesAPI.FacilityCardToAccountPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FacilityCardToAccountPost`: CardToAccountResponse
	fmt.Fprintf(os.Stdout, "Response from `FacilitiesAPI.FacilityCardToAccountPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFacilityCardToAccountPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardToAccountRequest** | [**CardToAccountRequest**](CardToAccountRequest.md) |  | 

### Return type

[**CardToAccountResponse**](CardToAccountResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FacilityCardToIbanPost

> CardToIbanResponse FacilityCardToIbanPost(ctx).CardToIbanRequest(cardToIbanRequest).Execute()

Card To IBAN

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
	cardToIbanRequest := *openapiclient.NewCardToIbanRequest("CardNumber_example") // CardToIbanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacilitiesAPI.FacilityCardToIbanPost(context.Background()).CardToIbanRequest(cardToIbanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacilitiesAPI.FacilityCardToIbanPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FacilityCardToIbanPost`: CardToIbanResponse
	fmt.Fprintf(os.Stdout, "Response from `FacilitiesAPI.FacilityCardToIbanPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFacilityCardToIbanPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cardToIbanRequest** | [**CardToIbanRequest**](CardToIbanRequest.md) |  | 

### Return type

[**CardToIbanResponse**](CardToIbanResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FacilityCheckCardWithNationalCodePost

> CheckCardWithNationalCodeResponse FacilityCheckCardWithNationalCodePost(ctx).CheckCardWithNationalCodeRequest(checkCardWithNationalCodeRequest).Execute()

Check Card With National Code

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
	checkCardWithNationalCodeRequest := *openapiclient.NewCheckCardWithNationalCodeRequest("NationalCode_example", "BirthDate_example", "CardNumber_example") // CheckCardWithNationalCodeRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacilitiesAPI.FacilityCheckCardWithNationalCodePost(context.Background()).CheckCardWithNationalCodeRequest(checkCardWithNationalCodeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacilitiesAPI.FacilityCheckCardWithNationalCodePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FacilityCheckCardWithNationalCodePost`: CheckCardWithNationalCodeResponse
	fmt.Fprintf(os.Stdout, "Response from `FacilitiesAPI.FacilityCheckCardWithNationalCodePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFacilityCheckCardWithNationalCodePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **checkCardWithNationalCodeRequest** | [**CheckCardWithNationalCodeRequest**](CheckCardWithNationalCodeRequest.md) |  | 

### Return type

[**CheckCardWithNationalCodeResponse**](CheckCardWithNationalCodeResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FacilityNationalIdentityInquiryPost

> NationalIdentityInquiryResponse FacilityNationalIdentityInquiryPost(ctx).NationalIdentityInquiryRequest(nationalIdentityInquiryRequest).Execute()

National Identity Inquiry

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
	nationalIdentityInquiryRequest := *openapiclient.NewNationalIdentityInquiryRequest("NationalCode_example", "BirthDate_example") // NationalIdentityInquiryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacilitiesAPI.FacilityNationalIdentityInquiryPost(context.Background()).NationalIdentityInquiryRequest(nationalIdentityInquiryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacilitiesAPI.FacilityNationalIdentityInquiryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FacilityNationalIdentityInquiryPost`: NationalIdentityInquiryResponse
	fmt.Fprintf(os.Stdout, "Response from `FacilitiesAPI.FacilityNationalIdentityInquiryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFacilityNationalIdentityInquiryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nationalIdentityInquiryRequest** | [**NationalIdentityInquiryRequest**](NationalIdentityInquiryRequest.md) |  | 

### Return type

[**NationalIdentityInquiryResponse**](NationalIdentityInquiryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FacilityShahkarInquiryPost

> ShahkarInquiryResponse FacilityShahkarInquiryPost(ctx).ShahkarInquiryRequest(shahkarInquiryRequest).Execute()

Shahkar Inquiry

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
	shahkarInquiryRequest := *openapiclient.NewShahkarInquiryRequest("NationalCode_example", "Mobile_example") // ShahkarInquiryRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacilitiesAPI.FacilityShahkarInquiryPost(context.Background()).ShahkarInquiryRequest(shahkarInquiryRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacilitiesAPI.FacilityShahkarInquiryPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FacilityShahkarInquiryPost`: ShahkarInquiryResponse
	fmt.Fprintf(os.Stdout, "Response from `FacilitiesAPI.FacilityShahkarInquiryPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFacilityShahkarInquiryPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shahkarInquiryRequest** | [**ShahkarInquiryRequest**](ShahkarInquiryRequest.md) |  | 

### Return type

[**ShahkarInquiryResponse**](ShahkarInquiryResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletCheckoutPlusPost

> WalletCheckoutPlusResponse WalletCheckoutPlusPost(ctx).WalletCheckoutPlusRequest(walletCheckoutPlusRequest).Execute()

Wallet Checkout Plus

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
	walletCheckoutPlusRequest := *openapiclient.NewWalletCheckoutPlusRequest(int32(123), int32(123), "BankAccount_example") // WalletCheckoutPlusRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FacilitiesAPI.WalletCheckoutPlusPost(context.Background()).WalletCheckoutPlusRequest(walletCheckoutPlusRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FacilitiesAPI.WalletCheckoutPlusPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WalletCheckoutPlusPost`: WalletCheckoutPlusResponse
	fmt.Fprintf(os.Stdout, "Response from `FacilitiesAPI.WalletCheckoutPlusPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletCheckoutPlusPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletCheckoutPlusRequest** | [**WalletCheckoutPlusRequest**](WalletCheckoutPlusRequest.md) |  | 

### Return type

[**WalletCheckoutPlusResponse**](WalletCheckoutPlusResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

