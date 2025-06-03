# PaymentVerifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaidAt** | Pointer to **string** | ISO 8601 Date | [optional] 
**Amount** | **int32** | مبلغ سفارش (به ریال) | 
**Result** | **int32** | result | 
**Status** | **int32** | verify status | 
**RefNumber** | Pointer to **string** | refrence number | [optional] 
**Description** | Pointer to **string** | transaction description | [optional] 
**CardNumber** | **string** | Masked bank card | 
**OrderId** | Pointer to **string** | order id | [optional] 
**Message** | **string** | message | 

## Methods

### NewPaymentVerifyResponse

`func NewPaymentVerifyResponse(amount int32, result int32, status int32, cardNumber string, message string, ) *PaymentVerifyResponse`

NewPaymentVerifyResponse instantiates a new PaymentVerifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentVerifyResponseWithDefaults

`func NewPaymentVerifyResponseWithDefaults() *PaymentVerifyResponse`

NewPaymentVerifyResponseWithDefaults instantiates a new PaymentVerifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaidAt

`func (o *PaymentVerifyResponse) GetPaidAt() string`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *PaymentVerifyResponse) GetPaidAtOk() (*string, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *PaymentVerifyResponse) SetPaidAt(v string)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *PaymentVerifyResponse) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### GetAmount

`func (o *PaymentVerifyResponse) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentVerifyResponse) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentVerifyResponse) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetResult

`func (o *PaymentVerifyResponse) GetResult() int32`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *PaymentVerifyResponse) GetResultOk() (*int32, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *PaymentVerifyResponse) SetResult(v int32)`

SetResult sets Result field to given value.


### GetStatus

`func (o *PaymentVerifyResponse) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentVerifyResponse) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentVerifyResponse) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetRefNumber

`func (o *PaymentVerifyResponse) GetRefNumber() string`

GetRefNumber returns the RefNumber field if non-nil, zero value otherwise.

### GetRefNumberOk

`func (o *PaymentVerifyResponse) GetRefNumberOk() (*string, bool)`

GetRefNumberOk returns a tuple with the RefNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefNumber

`func (o *PaymentVerifyResponse) SetRefNumber(v string)`

SetRefNumber sets RefNumber field to given value.

### HasRefNumber

`func (o *PaymentVerifyResponse) HasRefNumber() bool`

HasRefNumber returns a boolean if a field has been set.

### GetDescription

`func (o *PaymentVerifyResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentVerifyResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentVerifyResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentVerifyResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCardNumber

`func (o *PaymentVerifyResponse) GetCardNumber() string`

GetCardNumber returns the CardNumber field if non-nil, zero value otherwise.

### GetCardNumberOk

`func (o *PaymentVerifyResponse) GetCardNumberOk() (*string, bool)`

GetCardNumberOk returns a tuple with the CardNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardNumber

`func (o *PaymentVerifyResponse) SetCardNumber(v string)`

SetCardNumber sets CardNumber field to given value.


### GetOrderId

`func (o *PaymentVerifyResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaymentVerifyResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaymentVerifyResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PaymentVerifyResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetMessage

`func (o *PaymentVerifyResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PaymentVerifyResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PaymentVerifyResponse) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


