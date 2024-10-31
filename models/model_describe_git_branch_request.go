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

// checks if the DescribeGitBranchRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitBranchRequest{}

// DescribeGitBranchRequest struct for DescribeGitBranchRequest
type DescribeGitBranchRequest struct {
	// 分支名称
	BranchName string `json:"BranchName"`
	// 仓库id
	DepotId int64 `json:"DepotId"`
	// 仓库路径，与DepotId选其一就可以
	DepotPath *string `json:"DepotPath,omitempty"`
}

type _DescribeGitBranchRequest DescribeGitBranchRequest

// NewDescribeGitBranchRequest instantiates a new DescribeGitBranchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitBranchRequest(branchName string, depotId int64) *DescribeGitBranchRequest {
	this := DescribeGitBranchRequest{}
	this.BranchName = branchName
	this.DepotId = depotId
	return &this
}

// NewDescribeGitBranchRequestWithDefaults instantiates a new DescribeGitBranchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitBranchRequestWithDefaults() *DescribeGitBranchRequest {
	this := DescribeGitBranchRequest{}
	return &this
}

// GetBranchName returns the BranchName field value
func (o *DescribeGitBranchRequest) GetBranchName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BranchName
}

// GetBranchNameOk returns a tuple with the BranchName field value
// and a boolean to check if the value has been set.
func (o *DescribeGitBranchRequest) GetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BranchName, true
}

// SetBranchName sets field value
func (o *DescribeGitBranchRequest) SetBranchName(v string) {
	o.BranchName = v
}

// GetDepotId returns the DepotId field value
func (o *DescribeGitBranchRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *DescribeGitBranchRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *DescribeGitBranchRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *DescribeGitBranchRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitBranchRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *DescribeGitBranchRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *DescribeGitBranchRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

func (o DescribeGitBranchRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitBranchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["BranchName"] = o.BranchName
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	return toSerialize, nil
}

func (o *DescribeGitBranchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"BranchName",
		"DepotId",
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

	varDescribeGitBranchRequest := _DescribeGitBranchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeGitBranchRequest)

	if err != nil {
		return err
	}

	*o = DescribeGitBranchRequest(varDescribeGitBranchRequest)

	return err
}

type NullableDescribeGitBranchRequest struct {
	value *DescribeGitBranchRequest
	isSet bool
}

func (v NullableDescribeGitBranchRequest) Get() *DescribeGitBranchRequest {
	return v.value
}

func (v *NullableDescribeGitBranchRequest) Set(val *DescribeGitBranchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitBranchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitBranchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitBranchRequest(val *DescribeGitBranchRequest) *NullableDescribeGitBranchRequest {
	return &NullableDescribeGitBranchRequest{value: val, isSet: true}
}

func (v NullableDescribeGitBranchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitBranchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


