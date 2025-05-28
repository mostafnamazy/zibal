# WalletCheckoutPlusRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** |  | 
**Id** | **int32** |  | 
**BankAccount** | **string** |  | 
**CheckoutDelay** | Pointer to **int32** |  | [optional] 
**UniqueCode** | Pointer to **string** |  | [optional] 
**Bank** | Pointer to **string** |  | [optional] 
**LedgerId** | Pointer to **string** |  | [optional] 

## Methods

### NewWalletCheckoutPlusRequest

`func NewWalletCheckoutPlusRequest(amount int32, id int32, bankAccount string, ) *WalletCheckoutPlusRequest`

NewWalletCheckoutPlusRequest instantiates a new WalletCheckoutPlusRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletCheckoutPlusRequestWithDefaults

`func NewWalletCheckoutPlusRequestWithDefaults() *WalletCheckoutPlusRequest`

NewWalletCheckoutPlusRequestWithDefaults instantiates a new WalletCheckoutPlusRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *WalletCheckoutPlusRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *WalletCheckoutPlusRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *WalletCheckoutPlusRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetId

`func (o *WalletCheckoutPlusRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WalletCheckoutPlusRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WalletCheckoutPlusRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetBankAccount

`func (o *WalletCheckoutPlusRequest) GetBankAccount() string`

GetBankAccount returns the BankAccount field if non-nil, zero value otherwise.

### GetBankAccountOk

`func (o *WalletCheckoutPlusRequest) GetBankAccountOk() (*string, bool)`

GetBankAccountOk returns a tuple with the BankAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankAccount

`func (o *WalletCheckoutPlusRequest) SetBankAccount(v string)`

SetBankAccount sets BankAccount field to given value.


### GetCheckoutDelay

`func (o *WalletCheckoutPlusRequest) GetCheckoutDelay() int32`

GetCheckoutDelay returns the CheckoutDelay field if non-nil, zero value otherwise.

### GetCheckoutDelayOk

`func (o *WalletCheckoutPlusRequest) GetCheckoutDelayOk() (*int32, bool)`

GetCheckoutDelayOk returns a tuple with the CheckoutDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutDelay

`func (o *WalletCheckoutPlusRequest) SetCheckoutDelay(v int32)`

SetCheckoutDelay sets CheckoutDelay field to given value.

### HasCheckoutDelay

`func (o *WalletCheckoutPlusRequest) HasCheckoutDelay() bool`

HasCheckoutDelay returns a boolean if a field has been set.

### GetUniqueCode

`func (o *WalletCheckoutPlusRequest) GetUniqueCode() string`

GetUniqueCode returns the UniqueCode field if non-nil, zero value otherwise.

### GetUniqueCodeOk

`func (o *WalletCheckoutPlusRequest) GetUniqueCodeOk() (*string, bool)`

GetUniqueCodeOk returns a tuple with the UniqueCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueCode

`func (o *WalletCheckoutPlusRequest) SetUniqueCode(v string)`

SetUniqueCode sets UniqueCode field to given value.

### HasUniqueCode

`func (o *WalletCheckoutPlusRequest) HasUniqueCode() bool`

HasUniqueCode returns a boolean if a field has been set.

### GetBank

`func (o *WalletCheckoutPlusRequest) GetBank() string`

GetBank returns the Bank field if non-nil, zero value otherwise.

### GetBankOk

`func (o *WalletCheckoutPlusRequest) GetBankOk() (*string, bool)`

GetBankOk returns a tuple with the Bank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBank

`func (o *WalletCheckoutPlusRequest) SetBank(v string)`

SetBank sets Bank field to given value.

### HasBank

`func (o *WalletCheckoutPlusRequest) HasBank() bool`

HasBank returns a boolean if a field has been set.

### GetLedgerId

`func (o *WalletCheckoutPlusRequest) GetLedgerId() string`

GetLedgerId returns the LedgerId field if non-nil, zero value otherwise.

### GetLedgerIdOk

`func (o *WalletCheckoutPlusRequest) GetLedgerIdOk() (*string, bool)`

GetLedgerIdOk returns a tuple with the LedgerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedgerId

`func (o *WalletCheckoutPlusRequest) SetLedgerId(v string)`

SetLedgerId sets LedgerId field to given value.

### HasLedgerId

`func (o *WalletCheckoutPlusRequest) HasLedgerId() bool`

HasLedgerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


