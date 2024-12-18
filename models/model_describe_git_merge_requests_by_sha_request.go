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

// checks if the DescribeGitMergeRequestsByShaRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitMergeRequestsByShaRequest{}

// DescribeGitMergeRequestsByShaRequest struct for DescribeGitMergeRequestsByShaRequest
type DescribeGitMergeRequestsByShaRequest struct {
	// 仓库id
	DepotId int64 `json:"DepotId"`
	// 与DepotId选择其一即可
	DepotPath *string `json:"DepotPath,omitempty"`
	// 提交 Id
	Sha string `json:"Sha"`
}

type _DescribeGitMergeRequestsByShaRequest DescribeGitMergeRequestsByShaRequest

// NewDescribeGitMergeRequestsByShaRequest instantiates a new DescribeGitMergeRequestsByShaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitMergeRequestsByShaRequest(depotId int64, sha string) *DescribeGitMergeRequestsByShaRequest {
	this := DescribeGitMergeRequestsByShaRequest{}
	this.DepotId = depotId
	this.Sha = sha
	return &this
}

// NewDescribeGitMergeRequestsByShaRequestWithDefaults instantiates a new DescribeGitMergeRequestsByShaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitMergeRequestsByShaRequestWithDefaults() *DescribeGitMergeRequestsByShaRequest {
	this := DescribeGitMergeRequestsByShaRequest{}
	return &this
}

// GetDepotId returns the DepotId field value
func (o *DescribeGitMergeRequestsByShaRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *DescribeGitMergeRequestsByShaRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *DescribeGitMergeRequestsByShaRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *DescribeGitMergeRequestsByShaRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitMergeRequestsByShaRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *DescribeGitMergeRequestsByShaRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *DescribeGitMergeRequestsByShaRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetSha returns the Sha field value
func (o *DescribeGitMergeRequestsByShaRequest) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *DescribeGitMergeRequestsByShaRequest) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *DescribeGitMergeRequestsByShaRequest) SetSha(v string) {
	o.Sha = v
}

func (o DescribeGitMergeRequestsByShaRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitMergeRequestsByShaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	toSerialize["Sha"] = o.Sha
	return toSerialize, nil
}

func (o *DescribeGitMergeRequestsByShaRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotId",
		"Sha",
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

	varDescribeGitMergeRequestsByShaRequest := _DescribeGitMergeRequestsByShaRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeGitMergeRequestsByShaRequest)

	if err != nil {
		return err
	}

	*o = DescribeGitMergeRequestsByShaRequest(varDescribeGitMergeRequestsByShaRequest)

	return err
}

type NullableDescribeGitMergeRequestsByShaRequest struct {
	value *DescribeGitMergeRequestsByShaRequest
	isSet bool
}

func (v NullableDescribeGitMergeRequestsByShaRequest) Get() *DescribeGitMergeRequestsByShaRequest {
	return v.value
}

func (v *NullableDescribeGitMergeRequestsByShaRequest) Set(val *DescribeGitMergeRequestsByShaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitMergeRequestsByShaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitMergeRequestsByShaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitMergeRequestsByShaRequest(val *DescribeGitMergeRequestsByShaRequest) *NullableDescribeGitMergeRequestsByShaRequest {
	return &NullableDescribeGitMergeRequestsByShaRequest{value: val, isSet: true}
}

func (v NullableDescribeGitMergeRequestsByShaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitMergeRequestsByShaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


