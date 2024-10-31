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

// checks if the DescribeGrantUsersOnResourceRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGrantUsersOnResourceRequest{}

// DescribeGrantUsersOnResourceRequest struct for DescribeGrantUsersOnResourceRequest
type DescribeGrantUsersOnResourceRequest struct {
	UserIdScope []int64 `json:"UserIdScope,omitempty"`
	PolicyIdScope []int64 `json:"PolicyIdScope,omitempty"`
	// 请求条数
	PageSize int64 `json:"PageSize"`
	// 请求页数
	PageNumber int64 `json:"PageNumber"`
	Resource RamGrantResourceInfoRequest `json:"Resource"`
}

type _DescribeGrantUsersOnResourceRequest DescribeGrantUsersOnResourceRequest

// NewDescribeGrantUsersOnResourceRequest instantiates a new DescribeGrantUsersOnResourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGrantUsersOnResourceRequest(pageSize int64, pageNumber int64, resource RamGrantResourceInfoRequest) *DescribeGrantUsersOnResourceRequest {
	this := DescribeGrantUsersOnResourceRequest{}
	this.PageSize = pageSize
	this.PageNumber = pageNumber
	this.Resource = resource
	return &this
}

// NewDescribeGrantUsersOnResourceRequestWithDefaults instantiates a new DescribeGrantUsersOnResourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGrantUsersOnResourceRequestWithDefaults() *DescribeGrantUsersOnResourceRequest {
	this := DescribeGrantUsersOnResourceRequest{}
	var pageSize int64 = 0
	this.PageSize = pageSize
	var pageNumber int64 = 0
	this.PageNumber = pageNumber
	return &this
}

// GetUserIdScope returns the UserIdScope field value if set, zero value otherwise.
func (o *DescribeGrantUsersOnResourceRequest) GetUserIdScope() []int64 {
	if o == nil || utils.IsNil(o.UserIdScope) {
		var ret []int64
		return ret
	}
	return o.UserIdScope
}

// GetUserIdScopeOk returns a tuple with the UserIdScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGrantUsersOnResourceRequest) GetUserIdScopeOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.UserIdScope) {
		return nil, false
	}
	return o.UserIdScope, true
}

// HasUserIdScope returns a boolean if a field has been set.
func (o *DescribeGrantUsersOnResourceRequest) HasUserIdScope() bool {
	if o != nil && !utils.IsNil(o.UserIdScope) {
		return true
	}

	return false
}

// SetUserIdScope gets a reference to the given []int64 and assigns it to the UserIdScope field.
func (o *DescribeGrantUsersOnResourceRequest) SetUserIdScope(v []int64) {
	o.UserIdScope = v
}

// GetPolicyIdScope returns the PolicyIdScope field value if set, zero value otherwise.
func (o *DescribeGrantUsersOnResourceRequest) GetPolicyIdScope() []int64 {
	if o == nil || utils.IsNil(o.PolicyIdScope) {
		var ret []int64
		return ret
	}
	return o.PolicyIdScope
}

// GetPolicyIdScopeOk returns a tuple with the PolicyIdScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGrantUsersOnResourceRequest) GetPolicyIdScopeOk() ([]int64, bool) {
	if o == nil || utils.IsNil(o.PolicyIdScope) {
		return nil, false
	}
	return o.PolicyIdScope, true
}

// HasPolicyIdScope returns a boolean if a field has been set.
func (o *DescribeGrantUsersOnResourceRequest) HasPolicyIdScope() bool {
	if o != nil && !utils.IsNil(o.PolicyIdScope) {
		return true
	}

	return false
}

// SetPolicyIdScope gets a reference to the given []int64 and assigns it to the PolicyIdScope field.
func (o *DescribeGrantUsersOnResourceRequest) SetPolicyIdScope(v []int64) {
	o.PolicyIdScope = v
}

// GetPageSize returns the PageSize field value
func (o *DescribeGrantUsersOnResourceRequest) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DescribeGrantUsersOnResourceRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DescribeGrantUsersOnResourceRequest) SetPageSize(v int64) {
	o.PageSize = v
}

// GetPageNumber returns the PageNumber field value
func (o *DescribeGrantUsersOnResourceRequest) GetPageNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *DescribeGrantUsersOnResourceRequest) GetPageNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *DescribeGrantUsersOnResourceRequest) SetPageNumber(v int64) {
	o.PageNumber = v
}

// GetResource returns the Resource field value
func (o *DescribeGrantUsersOnResourceRequest) GetResource() RamGrantResourceInfoRequest {
	if o == nil {
		var ret RamGrantResourceInfoRequest
		return ret
	}

	return o.Resource
}

// GetResourceOk returns a tuple with the Resource field value
// and a boolean to check if the value has been set.
func (o *DescribeGrantUsersOnResourceRequest) GetResourceOk() (RamGrantResourceInfoRequest, bool) {
	if o == nil {
		return RamGrantResourceInfoRequest{}, false
	}
	return o.Resource, true
}

// SetResource sets field value
func (o *DescribeGrantUsersOnResourceRequest) SetResource(v RamGrantResourceInfoRequest) {
	o.Resource = v
}

func (o DescribeGrantUsersOnResourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGrantUsersOnResourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.UserIdScope) {
		toSerialize["UserIdScope"] = o.UserIdScope
	}
	if !utils.IsNil(o.PolicyIdScope) {
		toSerialize["PolicyIdScope"] = o.PolicyIdScope
	}
	toSerialize["PageSize"] = o.PageSize
	toSerialize["PageNumber"] = o.PageNumber
	toSerialize["Resource"] = o.Resource
	return toSerialize, nil
}

func (o *DescribeGrantUsersOnResourceRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"PageSize",
		"PageNumber",
		"Resource",
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

	varDescribeGrantUsersOnResourceRequest := _DescribeGrantUsersOnResourceRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeGrantUsersOnResourceRequest)

	if err != nil {
		return err
	}

	*o = DescribeGrantUsersOnResourceRequest(varDescribeGrantUsersOnResourceRequest)

	return err
}

type NullableDescribeGrantUsersOnResourceRequest struct {
	value *DescribeGrantUsersOnResourceRequest
	isSet bool
}

func (v NullableDescribeGrantUsersOnResourceRequest) Get() *DescribeGrantUsersOnResourceRequest {
	return v.value
}

func (v *NullableDescribeGrantUsersOnResourceRequest) Set(val *DescribeGrantUsersOnResourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGrantUsersOnResourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGrantUsersOnResourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGrantUsersOnResourceRequest(val *DescribeGrantUsersOnResourceRequest) *NullableDescribeGrantUsersOnResourceRequest {
	return &NullableDescribeGrantUsersOnResourceRequest{value: val, isSet: true}
}

func (v NullableDescribeGrantUsersOnResourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGrantUsersOnResourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


