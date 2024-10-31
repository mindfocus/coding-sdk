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

// checks if the DescribeGitFileRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitFileRequest{}

// DescribeGitFileRequest struct for DescribeGitFileRequest
type DescribeGitFileRequest struct {
	// 仓库Id
	DepotId int64 `json:"DepotId"`
	// 仓库路径，DepotId与DepotPath二选一即可
	DepotPath *string `json:"DepotPath,omitempty"`
	// 文件路径
	Path *string `json:"Path,omitempty"`
	// 分支名
	Ref string `json:"Ref"`
}

type _DescribeGitFileRequest DescribeGitFileRequest

// NewDescribeGitFileRequest instantiates a new DescribeGitFileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitFileRequest(depotId int64, ref string) *DescribeGitFileRequest {
	this := DescribeGitFileRequest{}
	this.DepotId = depotId
	this.Ref = ref
	return &this
}

// NewDescribeGitFileRequestWithDefaults instantiates a new DescribeGitFileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitFileRequestWithDefaults() *DescribeGitFileRequest {
	this := DescribeGitFileRequest{}
	return &this
}

// GetDepotId returns the DepotId field value
func (o *DescribeGitFileRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *DescribeGitFileRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *DescribeGitFileRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *DescribeGitFileRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitFileRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *DescribeGitFileRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *DescribeGitFileRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *DescribeGitFileRequest) GetPath() string {
	if o == nil || utils.IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitFileRequest) GetPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *DescribeGitFileRequest) HasPath() bool {
	if o != nil && !utils.IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *DescribeGitFileRequest) SetPath(v string) {
	o.Path = &v
}

// GetRef returns the Ref field value
func (o *DescribeGitFileRequest) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *DescribeGitFileRequest) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *DescribeGitFileRequest) SetRef(v string) {
	o.Ref = v
}

func (o DescribeGitFileRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitFileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	if !utils.IsNil(o.Path) {
		toSerialize["Path"] = o.Path
	}
	toSerialize["Ref"] = o.Ref
	return toSerialize, nil
}

func (o *DescribeGitFileRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotId",
		"Ref",
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

	varDescribeGitFileRequest := _DescribeGitFileRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeGitFileRequest)

	if err != nil {
		return err
	}

	*o = DescribeGitFileRequest(varDescribeGitFileRequest)

	return err
}

type NullableDescribeGitFileRequest struct {
	value *DescribeGitFileRequest
	isSet bool
}

func (v NullableDescribeGitFileRequest) Get() *DescribeGitFileRequest {
	return v.value
}

func (v *NullableDescribeGitFileRequest) Set(val *DescribeGitFileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitFileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitFileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitFileRequest(val *DescribeGitFileRequest) *NullableDescribeGitFileRequest {
	return &NullableDescribeGitFileRequest{value: val, isSet: true}
}

func (v NullableDescribeGitFileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitFileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


