# CardToAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Result** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**CardToAccountResponseData**](CardToAccountResponseData.md) |  | [optional] 

## Methods

### NewCardToAccountResponse

`func NewCardToAccountResponse() *CardToAccountResponse`

NewCardToAccountResponse instantiates a new CardToAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardToAccountResponseWithDefaults

`func NewCardToAccountResponseWithDefaults() *CardToAccountResponse`

NewCardToAccountResponseWithDefaults instantiates a new CardToAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResult

`func (o *CardToAccountResponse) GetResult() int32`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CardToAccountResponse) GetResultOk() (*int32, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CardToAccountResponse) SetResult(v int32)`

SetResult sets Result field to given value.

### HasResult

`func (o *CardToAccountResponse) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetMessage

`func (o *CardToAccountResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CardToAccountResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CardToAccountResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CardToAccountResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetData

`func (o *CardToAccountResponse) GetData() CardToAccountResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CardToAccountResponse) GetDataOk() (*CardToAccountResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CardToAccountResponse) SetData(v CardToAccountResponseData)`

SetData sets Data field to given value.

### HasData

`func (o *CardToAccountResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


