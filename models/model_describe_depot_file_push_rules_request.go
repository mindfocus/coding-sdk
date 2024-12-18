/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"bytes"
	"fmt"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeDepotFilePushRulesRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeDepotFilePushRulesRequest{}

// DescribeDepotFilePushRulesRequest struct for DescribeDepotFilePushRulesRequest
type DescribeDepotFilePushRulesRequest struct {
	// 仓库路径
	DepotPath string `json:"DepotPath"`
}

type _DescribeDepotFilePushRulesRequest DescribeDepotFilePushRulesRequest

// NewDescribeDepotFilePushRulesRequest instantiates a new DescribeDepotFilePushRulesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDepotFilePushRulesRequest(depotPath string) *DescribeDepotFilePushRulesRequest {
	this := DescribeDepotFilePushRulesRequest{}
	this.DepotPath = depotPath
	return &this
}

// NewDescribeDepotFilePushRulesRequestWithDefaults instantiates a new DescribeDepotFilePushRulesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDepotFilePushRulesRequestWithDefaults() *DescribeDepotFilePushRulesRequest {
	this := DescribeDepotFilePushRulesRequest{}
	return &this
}

// GetDepotPath returns the DepotPath field value
func (o *DescribeDepotFilePushRulesRequest) GetDepotPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value
// and a boolean to check if the value has been set.
func (o *DescribeDepotFilePushRulesRequest) GetDepotPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotPath, true
}

// SetDepotPath sets field value
func (o *DescribeDepotFilePushRulesRequest) SetDepotPath(v string) {
	o.DepotPath = v
}

func (o DescribeDepotFilePushRulesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDepotFilePushRulesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotPath"] = o.DepotPath
	return toSerialize, nil
}

func (o *DescribeDepotFilePushRulesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotPath",
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

	varDescribeDepotFilePushRulesRequest := _DescribeDepotFilePushRulesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeDepotFilePushRulesRequest)

	if err != nil {
		return err
	}

	*o = DescribeDepotFilePushRulesRequest(varDescribeDepotFilePushRulesRequest)

	return err
}

type NullableDescribeDepotFilePushRulesRequest struct {
	value *DescribeDepotFilePushRulesRequest
	isSet bool
}

func (v NullableDescribeDepotFilePushRulesRequest) Get() *DescribeDepotFilePushRulesRequest {
	return v.value
}

func (v *NullableDescribeDepotFilePushRulesRequest) Set(val *DescribeDepotFilePushRulesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDepotFilePushRulesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDepotFilePushRulesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDepotFilePushRulesRequest(val *DescribeDepotFilePushRulesRequest) *NullableDescribeDepotFilePushRulesRequest {
	return &NullableDescribeDepotFilePushRulesRequest{value: val, isSet: true}
}

func (v NullableDescribeDepotFilePushRulesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDepotFilePushRulesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


