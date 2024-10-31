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

// checks if the DescribeDifferentBetweenTwoCommitsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeDifferentBetweenTwoCommitsRequest{}

// DescribeDifferentBetweenTwoCommitsRequest struct for DescribeDifferentBetweenTwoCommitsRequest
type DescribeDifferentBetweenTwoCommitsRequest struct {
	// 仓库id
	DepotId int64 `json:"DepotId"`
	// 仓库路径，DepotId与DepotPath二选一即可
	DepotPath *string `json:"DepotPath,omitempty"`
	// 文件位置
	Path string `json:"Path"`
	// 源请求 Sha 值,分支名称
	Source string `json:"Source"`
	// 目标请求 Sha 值,分支名称
	Target string `json:"Target"`
}

type _DescribeDifferentBetweenTwoCommitsRequest DescribeDifferentBetweenTwoCommitsRequest

// NewDescribeDifferentBetweenTwoCommitsRequest instantiates a new DescribeDifferentBetweenTwoCommitsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDifferentBetweenTwoCommitsRequest(depotId int64, path string, source string, target string) *DescribeDifferentBetweenTwoCommitsRequest {
	this := DescribeDifferentBetweenTwoCommitsRequest{}
	this.DepotId = depotId
	this.Path = path
	this.Source = source
	this.Target = target
	return &this
}

// NewDescribeDifferentBetweenTwoCommitsRequestWithDefaults instantiates a new DescribeDifferentBetweenTwoCommitsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDifferentBetweenTwoCommitsRequestWithDefaults() *DescribeDifferentBetweenTwoCommitsRequest {
	this := DescribeDifferentBetweenTwoCommitsRequest{}
	return &this
}

// GetDepotId returns the DepotId field value
func (o *DescribeDifferentBetweenTwoCommitsRequest) GetDepotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepotId
}

// GetDepotIdOk returns a tuple with the DepotId field value
// and a boolean to check if the value has been set.
func (o *DescribeDifferentBetweenTwoCommitsRequest) GetDepotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotId, true
}

// SetDepotId sets field value
func (o *DescribeDifferentBetweenTwoCommitsRequest) SetDepotId(v int64) {
	o.DepotId = v
}

// GetDepotPath returns the DepotPath field value if set, zero value otherwise.
func (o *DescribeDifferentBetweenTwoCommitsRequest) GetDepotPath() string {
	if o == nil || utils.IsNil(o.DepotPath) {
		var ret string
		return ret
	}
	return *o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDifferentBetweenTwoCommitsRequest) GetDepotPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotPath) {
		return nil, false
	}
	return o.DepotPath, true
}

// HasDepotPath returns a boolean if a field has been set.
func (o *DescribeDifferentBetweenTwoCommitsRequest) HasDepotPath() bool {
	if o != nil && !utils.IsNil(o.DepotPath) {
		return true
	}

	return false
}

// SetDepotPath gets a reference to the given string and assigns it to the DepotPath field.
func (o *DescribeDifferentBetweenTwoCommitsRequest) SetDepotPath(v string) {
	o.DepotPath = &v
}

// GetPath returns the Path field value
func (o *DescribeDifferentBetweenTwoCommitsRequest) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DescribeDifferentBetweenTwoCommitsRequest) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DescribeDifferentBetweenTwoCommitsRequest) SetPath(v string) {
	o.Path = v
}

// GetSource returns the Source field value
func (o *DescribeDifferentBetweenTwoCommitsRequest) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DescribeDifferentBetweenTwoCommitsRequest) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *DescribeDifferentBetweenTwoCommitsRequest) SetSource(v string) {
	o.Source = v
}

// GetTarget returns the Target field value
func (o *DescribeDifferentBetweenTwoCommitsRequest) GetTarget() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *DescribeDifferentBetweenTwoCommitsRequest) GetTargetOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *DescribeDifferentBetweenTwoCommitsRequest) SetTarget(v string) {
	o.Target = v
}

func (o DescribeDifferentBetweenTwoCommitsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDifferentBetweenTwoCommitsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotId"] = o.DepotId
	if !utils.IsNil(o.DepotPath) {
		toSerialize["DepotPath"] = o.DepotPath
	}
	toSerialize["Path"] = o.Path
	toSerialize["Source"] = o.Source
	toSerialize["Target"] = o.Target
	return toSerialize, nil
}

func (o *DescribeDifferentBetweenTwoCommitsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotId",
		"Path",
		"Source",
		"Target",
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

	varDescribeDifferentBetweenTwoCommitsRequest := _DescribeDifferentBetweenTwoCommitsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeDifferentBetweenTwoCommitsRequest)

	if err != nil {
		return err
	}

	*o = DescribeDifferentBetweenTwoCommitsRequest(varDescribeDifferentBetweenTwoCommitsRequest)

	return err
}

type NullableDescribeDifferentBetweenTwoCommitsRequest struct {
	value *DescribeDifferentBetweenTwoCommitsRequest
	isSet bool
}

func (v NullableDescribeDifferentBetweenTwoCommitsRequest) Get() *DescribeDifferentBetweenTwoCommitsRequest {
	return v.value
}

func (v *NullableDescribeDifferentBetweenTwoCommitsRequest) Set(val *DescribeDifferentBetweenTwoCommitsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDifferentBetweenTwoCommitsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDifferentBetweenTwoCommitsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDifferentBetweenTwoCommitsRequest(val *DescribeDifferentBetweenTwoCommitsRequest) *NullableDescribeDifferentBetweenTwoCommitsRequest {
	return &NullableDescribeDifferentBetweenTwoCommitsRequest{value: val, isSet: true}
}

func (v NullableDescribeDifferentBetweenTwoCommitsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDifferentBetweenTwoCommitsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


