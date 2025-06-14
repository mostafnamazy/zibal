/*
Zibal API

Combined API for Zibal Gateway and Facilities (Verification, Wallet, etc.)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zibal

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WalletCheckoutPlusRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WalletCheckoutPlusRequest{}

// WalletCheckoutPlusRequest struct for WalletCheckoutPlusRequest
type WalletCheckoutPlusRequest struct {
	Amount int32 `json:"amount"`
	Id int32 `json:"id"`
	BankAccount string `json:"bankAccount"`
	CheckoutDelay *int32 `json:"checkoutDelay,omitempty"`
	UniqueCode *string `json:"uniqueCode,omitempty"`
	Bank *string `json:"bank,omitempty"`
	LedgerId *string `json:"ledgerId,omitempty"`
}

type _WalletCheckoutPlusRequest WalletCheckoutPlusRequest

// NewWalletCheckoutPlusRequest instantiates a new WalletCheckoutPlusRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWalletCheckoutPlusRequest(amount int32, id int32, bankAccount string) *WalletCheckoutPlusRequest {
	this := WalletCheckoutPlusRequest{}
	this.Amount = amount
	this.Id = id
	this.BankAccount = bankAccount
	return &this
}

// NewWalletCheckoutPlusRequestWithDefaults instantiates a new WalletCheckoutPlusRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWalletCheckoutPlusRequestWithDefaults() *WalletCheckoutPlusRequest {
	this := WalletCheckoutPlusRequest{}
	return &this
}

// GetAmount returns the Amount field value
func (o *WalletCheckoutPlusRequest) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *WalletCheckoutPlusRequest) GetAmountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *WalletCheckoutPlusRequest) SetAmount(v int32) {
	o.Amount = v
}

// GetId returns the Id field value
func (o *WalletCheckoutPlusRequest) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WalletCheckoutPlusRequest) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WalletCheckoutPlusRequest) SetId(v int32) {
	o.Id = v
}

// GetBankAccount returns the BankAccount field value
func (o *WalletCheckoutPlusRequest) GetBankAccount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BankAccount
}

// GetBankAccountOk returns a tuple with the BankAccount field value
// and a boolean to check if the value has been set.
func (o *WalletCheckoutPlusRequest) GetBankAccountOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BankAccount, true
}

// SetBankAccount sets field value
func (o *WalletCheckoutPlusRequest) SetBankAccount(v string) {
	o.BankAccount = v
}

// GetCheckoutDelay returns the CheckoutDelay field value if set, zero value otherwise.
func (o *WalletCheckoutPlusRequest) GetCheckoutDelay() int32 {
	if o == nil || IsNil(o.CheckoutDelay) {
		var ret int32
		return ret
	}
	return *o.CheckoutDelay
}

// GetCheckoutDelayOk returns a tuple with the CheckoutDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCheckoutPlusRequest) GetCheckoutDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.CheckoutDelay) {
		return nil, false
	}
	return o.CheckoutDelay, true
}

// HasCheckoutDelay returns a boolean if a field has been set.
func (o *WalletCheckoutPlusRequest) HasCheckoutDelay() bool {
	if o != nil && !IsNil(o.CheckoutDelay) {
		return true
	}

	return false
}

// SetCheckoutDelay gets a reference to the given int32 and assigns it to the CheckoutDelay field.
func (o *WalletCheckoutPlusRequest) SetCheckoutDelay(v int32) {
	o.CheckoutDelay = &v
}

// GetUniqueCode returns the UniqueCode field value if set, zero value otherwise.
func (o *WalletCheckoutPlusRequest) GetUniqueCode() string {
	if o == nil || IsNil(o.UniqueCode) {
		var ret string
		return ret
	}
	return *o.UniqueCode
}

// GetUniqueCodeOk returns a tuple with the UniqueCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCheckoutPlusRequest) GetUniqueCodeOk() (*string, bool) {
	if o == nil || IsNil(o.UniqueCode) {
		return nil, false
	}
	return o.UniqueCode, true
}

// HasUniqueCode returns a boolean if a field has been set.
func (o *WalletCheckoutPlusRequest) HasUniqueCode() bool {
	if o != nil && !IsNil(o.UniqueCode) {
		return true
	}

	return false
}

// SetUniqueCode gets a reference to the given string and assigns it to the UniqueCode field.
func (o *WalletCheckoutPlusRequest) SetUniqueCode(v string) {
	o.UniqueCode = &v
}

// GetBank returns the Bank field value if set, zero value otherwise.
func (o *WalletCheckoutPlusRequest) GetBank() string {
	if o == nil || IsNil(o.Bank) {
		var ret string
		return ret
	}
	return *o.Bank
}

// GetBankOk returns a tuple with the Bank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCheckoutPlusRequest) GetBankOk() (*string, bool) {
	if o == nil || IsNil(o.Bank) {
		return nil, false
	}
	return o.Bank, true
}

// HasBank returns a boolean if a field has been set.
func (o *WalletCheckoutPlusRequest) HasBank() bool {
	if o != nil && !IsNil(o.Bank) {
		return true
	}

	return false
}

// SetBank gets a reference to the given string and assigns it to the Bank field.
func (o *WalletCheckoutPlusRequest) SetBank(v string) {
	o.Bank = &v
}

// GetLedgerId returns the LedgerId field value if set, zero value otherwise.
func (o *WalletCheckoutPlusRequest) GetLedgerId() string {
	if o == nil || IsNil(o.LedgerId) {
		var ret string
		return ret
	}
	return *o.LedgerId
}

// GetLedgerIdOk returns a tuple with the LedgerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WalletCheckoutPlusRequest) GetLedgerIdOk() (*string, bool) {
	if o == nil || IsNil(o.LedgerId) {
		return nil, false
	}
	return o.LedgerId, true
}

// HasLedgerId returns a boolean if a field has been set.
func (o *WalletCheckoutPlusRequest) HasLedgerId() bool {
	if o != nil && !IsNil(o.LedgerId) {
		return true
	}

	return false
}

// SetLedgerId gets a reference to the given string and assigns it to the LedgerId field.
func (o *WalletCheckoutPlusRequest) SetLedgerId(v string) {
	o.LedgerId = &v
}

func (o WalletCheckoutPlusRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WalletCheckoutPlusRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["amount"] = o.Amount
	toSerialize["id"] = o.Id
	toSerialize["bankAccount"] = o.BankAccount
	if !IsNil(o.CheckoutDelay) {
		toSerialize["checkoutDelay"] = o.CheckoutDelay
	}
	if !IsNil(o.UniqueCode) {
		toSerialize["uniqueCode"] = o.UniqueCode
	}
	if !IsNil(o.Bank) {
		toSerialize["bank"] = o.Bank
	}
	if !IsNil(o.LedgerId) {
		toSerialize["ledgerId"] = o.LedgerId
	}
	return toSerialize, nil
}

func (o *WalletCheckoutPlusRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"amount",
		"id",
		"bankAccount",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varWalletCheckoutPlusRequest := _WalletCheckoutPlusRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWalletCheckoutPlusRequest)

	if err != nil {
		return err
	}

	*o = WalletCheckoutPlusRequest(varWalletCheckoutPlusRequest)

	return err
}

type NullableWalletCheckoutPlusRequest struct {
	value *WalletCheckoutPlusRequest
	isSet bool
}

func (v NullableWalletCheckoutPlusRequest) Get() *WalletCheckoutPlusRequest {
	return v.value
}

func (v *NullableWalletCheckoutPlusRequest) Set(val *WalletCheckoutPlusRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWalletCheckoutPlusRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWalletCheckoutPlusRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWalletCheckoutPlusRequest(val *WalletCheckoutPlusRequest) *NullableWalletCheckoutPlusRequest {
	return &NullableWalletCheckoutPlusRequest{value: val, isSet: true}
}

func (v NullableWalletCheckoutPlusRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWalletCheckoutPlusRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


