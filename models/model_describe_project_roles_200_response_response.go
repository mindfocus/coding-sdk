/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeProjectRoles200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectRoles200ResponseResponse{}

// DescribeProjectRoles200ResponseResponse 公共返回体
type DescribeProjectRoles200ResponseResponse struct {
	// 用户组
	Roles []Role `json:"Roles,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeProjectRoles200ResponseResponse instantiates a new DescribeProjectRoles200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectRoles200ResponseResponse() *DescribeProjectRoles200ResponseResponse {
	this := DescribeProjectRoles200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeProjectRoles200ResponseResponseWithDefaults instantiates a new DescribeProjectRoles200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectRoles200ResponseResponseWithDefaults() *DescribeProjectRoles200ResponseResponse {
	this := DescribeProjectRoles200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *DescribeProjectRoles200ResponseResponse) GetRoles() []Role {
	if o == nil || utils.IsNil(o.Roles) {
		var ret []Role
		return ret
	}
	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectRoles200ResponseResponse) GetRolesOk() ([]Role, bool) {
	if o == nil || utils.IsNil(o.Roles) {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *DescribeProjectRoles200ResponseResponse) HasRoles() bool {
	if o != nil && !utils.IsNil(o.Roles) {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []Role and assigns it to the Roles field.
func (o *DescribeProjectRoles200ResponseResponse) SetRoles(v []Role) {
	o.Roles = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeProjectRoles200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectRoles200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeProjectRoles200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeProjectRoles200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeProjectRoles200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectRoles200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Roles) {
		toSerialize["Roles"] = o.Roles
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeProjectRoles200ResponseResponse struct {
	value *DescribeProjectRoles200ResponseResponse
	isSet bool
}

func (v NullableDescribeProjectRoles200ResponseResponse) Get() *DescribeProjectRoles200ResponseResponse {
	return v.value
}

func (v *NullableDescribeProjectRoles200ResponseResponse) Set(val *DescribeProjectRoles200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectRoles200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectRoles200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectRoles200ResponseResponse(val *DescribeProjectRoles200ResponseResponse) *NullableDescribeProjectRoles200ResponseResponse {
	return &NullableDescribeProjectRoles200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeProjectRoles200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectRoles200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


