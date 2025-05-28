# PaymentInitRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Merchant** | **string** |  | 
**Amount** | **int64** |  | 
**CallbackUrl** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**OrderId** | Pointer to **string** |  | [optional] 
**Mobile** | Pointer to **string** |  | [optional] 
**AllowedCards** | Pointer to **[]string** |  | [optional] 
**LedgerId** | Pointer to **string** |  | [optional] 
**NationalCode** | Pointer to **string** |  | [optional] 
**CheckMobileWithCard** | Pointer to **bool** |  | [optional] 

## Methods

### NewPaymentInitRequest

`func NewPaymentInitRequest(merchant string, amount int64, callbackUrl string, ) *PaymentInitRequest`

NewPaymentInitRequest instantiates a new PaymentInitRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitRequestWithDefaults

`func NewPaymentInitRequestWithDefaults() *PaymentInitRequest`

NewPaymentInitRequestWithDefaults instantiates a new PaymentInitRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchant

`func (o *PaymentInitRequest) GetMerchant() string`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *PaymentInitRequest) GetMerchantOk() (*string, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *PaymentInitRequest) SetMerchant(v string)`

SetMerchant sets Merchant field to given value.


### GetAmount

`func (o *PaymentInitRequest) GetAmount() int64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentInitRequest) GetAmountOk() (*int64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentInitRequest) SetAmount(v int64)`

SetAmount sets Amount field to given value.


### GetCallbackUrl

`func (o *PaymentInitRequest) GetCallbackUrl() string`

GetCallbackUrl returns the CallbackUrl field if non-nil, zero value otherwise.

### GetCallbackUrlOk

`func (o *PaymentInitRequest) GetCallbackUrlOk() (*string, bool)`

GetCallbackUrlOk returns a tuple with the CallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallbackUrl

`func (o *PaymentInitRequest) SetCallbackUrl(v string)`

SetCallbackUrl sets CallbackUrl field to given value.


### GetDescription

`func (o *PaymentInitRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PaymentInitRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PaymentInitRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PaymentInitRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrderId

`func (o *PaymentInitRequest) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PaymentInitRequest) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PaymentInitRequest) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PaymentInitRequest) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetMobile

`func (o *PaymentInitRequest) GetMobile() string`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *PaymentInitRequest) GetMobileOk() (*string, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *PaymentInitRequest) SetMobile(v string)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *PaymentInitRequest) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetAllowedCards

`func (o *PaymentInitRequest) GetAllowedCards() []string`

GetAllowedCards returns the AllowedCards field if non-nil, zero value otherwise.

### GetAllowedCardsOk

`func (o *PaymentInitRequest) GetAllowedCardsOk() (*[]string, bool)`

GetAllowedCardsOk returns a tuple with the AllowedCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedCards

`func (o *PaymentInitRequest) SetAllowedCards(v []string)`

SetAllowedCards sets AllowedCards field to given value.

### HasAllowedCards

`func (o *PaymentInitRequest) HasAllowedCards() bool`

HasAllowedCards returns a boolean if a field has been set.

### GetLedgerId

`func (o *PaymentInitRequest) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *PaymentInitRequest) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *PaymentInitRequest) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.

### HasLedgerId

`func (o *PaymentInitRequest) HasLedgerId() bool`

HasLedgerId returns a boolean if a field has been set.

### GetNationalCode

`func (o *PaymentInitRequest) GetNationalCode() string`

GetNationalCode returns the NationalCode field if non-nil, zero value otherwise.

### GetNationalCodeOk

`func (o *PaymentInitRequest) GetNationalCodeOk() (*string, bool)`

GetNationalCodeOk returns a tuple with the NationalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNationalCode

`func (o *PaymentInitRequest) SetNationalCode(v string)`

SetNationalCode sets NationalCode field to given value.

### HasNationalCode

`func (o *PaymentInitRequest) HasNationalCode() bool`

HasNationalCode returns a boolean if a field has been set.

### GetCheckMobileWithCard

`func (o *PaymentInitRequest) GetCheckMobileWithCard() bool`

GetCheckMobileWithCard returns the CheckMobileWithCard field if non-nil, zero value otherwise.

### GetCheckMobileWithCardOk

`func (o *PaymentInitRequest) GetCheckMobileWithCardOk() (*bool, bool)`

GetCheckMobileWithCardOk returns a tuple with the CheckMobileWithCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMobileWithCard

`func (o *PaymentInitRequest) SetCheckMobileWithCard(v bool)`

SetCheckMobileWithCard sets CheckMobileWithCard field to given value.

### HasCheckMobileWithCard

`func (o *PaymentInitRequest) HasCheckMobileWithCard() bool`

HasCheckMobileWithCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


