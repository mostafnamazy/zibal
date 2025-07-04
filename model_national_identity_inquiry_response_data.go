/*
Zibal API

Combined API for Zibal Gateway and Facilities (Verification, Wallet, etc.)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package zibal

import (
	"encoding/json"
	"fmt"
)

// checks if the NationalIdentityInquiryResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NationalIdentityInquiryResponseData{}

// NationalIdentityInquiryResponseData struct for NationalIdentityInquiryResponseData
type NationalIdentityInquiryResponseData struct {
	Matched bool `json:"matched"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	FatherName string `json:"fatherName"`
	Alive bool `json:"alive"`
	AdditionalProperties map[string]interface{}
}

type _NationalIdentityInquiryResponseData NationalIdentityInquiryResponseData

// NewNationalIdentityInquiryResponseData instantiates a new NationalIdentityInquiryResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNationalIdentityInquiryResponseData(matched bool, firstName string, lastName string, fatherName string, alive bool) *NationalIdentityInquiryResponseData {
	this := NationalIdentityInquiryResponseData{}
	this.Matched = matched
	this.FirstName = firstName
	this.LastName = lastName
	this.FatherName = fatherName
	this.Alive = alive
	return &this
}

// NewNationalIdentityInquiryResponseDataWithDefaults instantiates a new NationalIdentityInquiryResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNationalIdentityInquiryResponseDataWithDefaults() *NationalIdentityInquiryResponseData {
	this := NationalIdentityInquiryResponseData{}
	return &this
}

// GetMatched returns the Matched field value
func (o *NationalIdentityInquiryResponseData) GetMatched() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Matched
}

// GetMatchedOk returns a tuple with the Matched field value
// and a boolean to check if the value has been set.
func (o *NationalIdentityInquiryResponseData) GetMatchedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Matched, true
}

// SetMatched sets field value
func (o *NationalIdentityInquiryResponseData) SetMatched(v bool) {
	o.Matched = v
}

// GetFirstName returns the FirstName field value
func (o *NationalIdentityInquiryResponseData) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *NationalIdentityInquiryResponseData) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *NationalIdentityInquiryResponseData) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *NationalIdentityInquiryResponseData) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *NationalIdentityInquiryResponseData) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *NationalIdentityInquiryResponseData) SetLastName(v string) {
	o.LastName = v
}

// GetFatherName returns the FatherName field value
func (o *NationalIdentityInquiryResponseData) GetFatherName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FatherName
}

// GetFatherNameOk returns a tuple with the FatherName field value
// and a boolean to check if the value has been set.
func (o *NationalIdentityInquiryResponseData) GetFatherNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FatherName, true
}

// SetFatherName sets field value
func (o *NationalIdentityInquiryResponseData) SetFatherName(v string) {
	o.FatherName = v
}

// GetAlive returns the Alive field value
func (o *NationalIdentityInquiryResponseData) GetAlive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Alive
}

// GetAliveOk returns a tuple with the Alive field value
// and a boolean to check if the value has been set.
func (o *NationalIdentityInquiryResponseData) GetAliveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alive, true
}

// SetAlive sets field value
func (o *NationalIdentityInquiryResponseData) SetAlive(v bool) {
	o.Alive = v
}

func (o NationalIdentityInquiryResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NationalIdentityInquiryResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["matched"] = o.Matched
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	toSerialize["fatherName"] = o.FatherName
	toSerialize["alive"] = o.Alive

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NationalIdentityInquiryResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"matched",
		"firstName",
		"lastName",
		"fatherName",
		"alive",
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

	varNationalIdentityInquiryResponseData := _NationalIdentityInquiryResponseData{}

	err = json.Unmarshal(data, &varNationalIdentityInquiryResponseData)

	if err != nil {
		return err
	}

	*o = NationalIdentityInquiryResponseData(varNationalIdentityInquiryResponseData)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "matched")
		delete(additionalProperties, "firstName")
		delete(additionalProperties, "lastName")
		delete(additionalProperties, "fatherName")
		delete(additionalProperties, "alive")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNationalIdentityInquiryResponseData struct {
	value *NationalIdentityInquiryResponseData
	isSet bool
}

func (v NullableNationalIdentityInquiryResponseData) Get() *NationalIdentityInquiryResponseData {
	return v.value
}

func (v *NullableNationalIdentityInquiryResponseData) Set(val *NationalIdentityInquiryResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableNationalIdentityInquiryResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableNationalIdentityInquiryResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNationalIdentityInquiryResponseData(val *NationalIdentityInquiryResponseData) *NullableNationalIdentityInquiryResponseData {
	return &NullableNationalIdentityInquiryResponseData{value: val, isSet: true}
}

func (v NullableNationalIdentityInquiryResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNationalIdentityInquiryResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


