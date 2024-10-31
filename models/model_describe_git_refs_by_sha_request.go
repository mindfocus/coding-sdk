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

// checks if the DescribeGitRefsByShaRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitRefsByShaRequest{}

// DescribeGitRefsByShaRequest struct for DescribeGitRefsByShaRequest
type DescribeGitRefsByShaRequest struct {
	// 仓库id
	DepotId int64 `json:"DepotId"`
	// 仓库路径，DepotId与DepotPath二选一即可
	DepotPath *string `json:"DepotPath,omitempty"`
	// 提交id
	Sha string `json:"Sha"`
	// all：列表中既包含分支也包含标签，tag：列表中只包含标签，branch：列表中只包含分支
	Type string `json:"Type"`
}

type _DescribeGitRefsByShaRequest DescribeGitRefsByShaRequest

// NewDescribeGitRefsByShaRequest instantiates a new DescribeGitRefsByShaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitRefsByShaRequest(depotId int64, sha string, type_ string) *DescribeGitRefsByShaRequest {
	this := DescribeGitRefsByShaRequest{}
	this.DepotId = depotId
	this.Sha = sha
	this.Type = type_
	return &this
}

// NewDescribeGitRefsByShaRequestWithDefaults instantiates a new DescribeGitRefsByShaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitRefsByShaRequestWithDefaults() *DescribeGitRefsByShaRequest {
	this := DescribeGitRefsByShaRequest{}
	return &this
}

// GetDepotId returns the DepotId field value
func (o *DescribeGitRefsByShaRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *DescribeGitRefsByShaRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *DescribeGitRefsByShaRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *DescribeGitRefsByShaRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitRefsByShaRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *DescribeGitRefsByShaRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *DescribeGitRefsByShaRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetSha returns the Sha field value
func (o *DescribeGitRefsByShaRequest) GetSha() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sha
}

// GetShaOk returns a tuple with the Sha field value
// and a boolean to check if the value has been set.
func (o *DescribeGitRefsByShaRequest) GetShaOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sha, true
}

// SetSha sets field value
func (o *DescribeGitRefsByShaRequest) SetSha(v string) {
	o.Sha = v
}

// GetType returns the Type field value
func (o *DescribeGitRefsByShaRequest) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DescribeGitRefsByShaRequest) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DescribeGitRefsByShaRequest) SetType(v string) {
	o.Type = v
}

func (o DescribeGitRefsByShaRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitRefsByShaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	toSerialize["Sha"] = o.Sha
	toSerialize["Type"] = o.Type
	return toSerialize, nil
}

func (o *DescribeGitRefsByShaRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotId",
		"Sha",
		"Type",
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

	varDescribeGitRefsByShaRequest := _DescribeGitRefsByShaRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeGitRefsByShaRequest)

	if err != nil {
		return err
	}

	*o = DescribeGitRefsByShaRequest(varDescribeGitRefsByShaRequest)

	return err
}

type NullableDescribeGitRefsByShaRequest struct {
	value *DescribeGitRefsByShaRequest
	isSet bool
}

func (v NullableDescribeGitRefsByShaRequest) Get() *DescribeGitRefsByShaRequest {
	return v.value
}

func (v *NullableDescribeGitRefsByShaRequest) Set(val *DescribeGitRefsByShaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitRefsByShaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitRefsByShaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitRefsByShaRequest(val *DescribeGitRefsByShaRequest) *NullableDescribeGitRefsByShaRequest {
	return &NullableDescribeGitRefsByShaRequest{value: val, isSet: true}
}

func (v NullableDescribeGitRefsByShaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitRefsByShaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


