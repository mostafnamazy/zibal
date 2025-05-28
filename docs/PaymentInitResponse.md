# PaymentInitResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackId** | **int64** |  | 
**Result** | **int32** | result | 
**Message** | **string** |  | 

## Methods

### NewPaymentInitResponse

`func NewPaymentInitResponse(trackId int64, result int32, message string, ) *PaymentInitResponse`

NewPaymentInitResponse instantiates a new PaymentInitResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitResponseWithDefaults

`func NewPaymentInitResponseWithDefaults() *PaymentInitResponse`

NewPaymentInitResponseWithDefaults instantiates a new PaymentInitResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackId

`func (o *PaymentInitResponse) GetTrackId() int64`

GetTrackId returns the TrackId field if non-nil, zero value otherwise.

### GetTrackIdOk

`func (o *PaymentInitResponse) GetTrackIdOk() (*int64, bool)`

GetTrackIdOk returns a tuple with the TrackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackId

`func (o *PaymentInitResponse) SetTrackId(v int64)`

SetTrackId sets TrackId field to given value.


### GetResult

`func (o *PaymentInitResponse) GetResult() int32`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PaymentInitResponse) GetResultOk() (*int32, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PaymentInitResponse) SetResult(v int32)`

SetResult sets Result field to given value.


### GetMessage

`func (o *PaymentInitResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentInitResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentInitResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


