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

// checks if the DescribeDepartmentRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeDepartmentRequest{}

// DescribeDepartmentRequest struct for DescribeDepartmentRequest
type DescribeDepartmentRequest struct {
	// 部门ID
	DepartmentId int64 `json:"DepartmentId"`
	// 是否获取部门树
	GetTree *bool `json:"GetTree,omitempty"`
}

type _DescribeDepartmentRequest DescribeDepartmentRequest

// NewDescribeDepartmentRequest instantiates a new DescribeDepartmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDepartmentRequest(departmentId int64) *DescribeDepartmentRequest {
	this := DescribeDepartmentRequest{}
	this.DepartmentId = departmentId
	var getTree bool = false
	this.GetTree = &getTree
	return &this
}

// NewDescribeDepartmentRequestWithDefaults instantiates a new DescribeDepartmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDepartmentRequestWithDefaults() *DescribeDepartmentRequest {
	this := DescribeDepartmentRequest{}
	var departmentId int64 = 0
	this.DepartmentId = departmentId
	var getTree bool = false
	this.GetTree = &getTree
	return &this
}

// GetDepartmentId returns the DepartmentId field value
func (o *DescribeDepartmentRequest) GetDepartmentId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.DepartmentId
}

// GetDepartmentIdOk returns a tuple with the DepartmentId field value
// and a boolean to check if the value has been set.
func (o *DescribeDepartmentRequest) GetDepartmentIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepartmentId, true
}

// SetDepartmentId sets field value
func (o *DescribeDepartmentRequest) SetDepartmentId(v int64) {
	o.DepartmentId = v
}

// GetGetTree returns the GetTree field value if set, zero value otherwise.
func (o *DescribeDepartmentRequest) GetGetTree() bool {
	if o == nil || utils.IsNil(o.GetTree) {
		var ret bool
		return ret
	}
	return *o.GetTree
}

// GetGetTreeOk returns a tuple with the GetTree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepartmentRequest) GetGetTreeOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.GetTree) {
		return nil, false
	}
	return o.GetTree, true
}

// HasGetTree returns a boolean if a field has been set.
func (o *DescribeDepartmentRequest) HasGetTree() bool {
	if o != nil && !utils.IsNil(o.GetTree) {
		return true
	}

	return false
}

// SetGetTree gets a reference to the given bool and assigns it to the GetTree field.
func (o *DescribeDepartmentRequest) SetGetTree(v bool) {
	o.GetTree = &v
}

func (o DescribeDepartmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDepartmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepartmentId"] = o.DepartmentId
	if !utils.IsNil(o.GetTree) {
		toSerialize["GetTree"] = o.GetTree
	}
	return toSerialize, nil
}

func (o *DescribeDepartmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepartmentId",
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

	varDescribeDepartmentRequest := _DescribeDepartmentRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeDepartmentRequest)

	if err != nil {
		return err
	}

	*o = DescribeDepartmentRequest(varDescribeDepartmentRequest)

	return err
}

type NullableDescribeDepartmentRequest struct {
	value *DescribeDepartmentRequest
	isSet bool
}

func (v NullableDescribeDepartmentRequest) Get() *DescribeDepartmentRequest {
	return v.value
}

func (v *NullableDescribeDepartmentRequest) Set(val *DescribeDepartmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDepartmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDepartmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDepartmentRequest(val *DescribeDepartmentRequest) *NullableDescribeDepartmentRequest {
	return &NullableDescribeDepartmentRequest{value: val, isSet: true}
}

func (v NullableDescribeDepartmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDepartmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

