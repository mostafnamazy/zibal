# PaymentVerifyRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merchant** | **string** |  | 
**TrackId** | **int64** |  | 

## Methods

### NewPaymentVerifyRequest

`func NewPaymentVerifyRequest(merchant string, trackId int64, ) *PaymentVerifyRequest`

NewPaymentVerifyRequest instantiates a new PaymentVerifyRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentVerifyRequestWithDefaults

`func NewPaymentVerifyRequestWithDefaults() *PaymentVerifyRequest`

NewPaymentVerifyRequestWithDefaults instantiates a new PaymentVerifyRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchant

`func (o *PaymentVerifyRequest) GetMerchant() string`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *PaymentVerifyRequest) GetMerchantOk() (*string, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *PaymentVerifyRequest) SetMerchant(v string)`

SetMerchant sets Merchant field to given value.


### GetTrackId

`func (o *PaymentVerifyRequest) GetTrackId() int64`

GetTrackId returns the TrackId field if non-nil, zero value otherwise.

### GetTrackIdOk

`func (o *PaymentVerifyRequest) GetTrackIdOk() (*int64, bool)`

GetTrackIdOk returns a tuple with the TrackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackId

`func (o *PaymentVerifyRequest) SetTrackId(v int64)`

SetTrackId sets TrackId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


