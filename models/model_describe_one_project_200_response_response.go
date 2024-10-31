/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeOneProject200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeOneProject200ResponseResponse{}

// DescribeOneProject200ResponseResponse 公共返回体
type DescribeOneProject200ResponseResponse struct {
	Project *Project `json:"Project,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeOneProject200ResponseResponse instantiates a new DescribeOneProject200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeOneProject200ResponseResponse() *DescribeOneProject200ResponseResponse {
	this := DescribeOneProject200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeOneProject200ResponseResponseWithDefaults instantiates a new DescribeOneProject200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeOneProject200ResponseResponseWithDefaults() *DescribeOneProject200ResponseResponse {
	this := DescribeOneProject200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetProject returns the Project field value if set, zero value otherwise.
func (o *DescribeOneProject200ResponseResponse) GetProject() Project {
	if o == nil || utils.IsNil(o.Project) {
		var ret Project
		return ret
	}
	return *o.Project
}

// GetProjectOk returns a tuple with the Project field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeOneProject200ResponseResponse) GetProjectOk() (*Project, bool) {
	if o == nil || utils.IsNil(o.Project) {
		return nil, false
	}
	return o.Project, true
}

// HasProject returns a boolean if a field has been set.
func (o *DescribeOneProject200ResponseResponse) HasProject() bool {
	if o != nil && !utils.IsNil(o.Project) {
		return true
	}

	return false
}

// SetProject gets a reference to the given Project and assigns it to the Project field.
func (o *DescribeOneProject200ResponseResponse) SetProject(v Project) {
	o.Project = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeOneProject200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeOneProject200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeOneProject200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeOneProject200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeOneProject200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeOneProject200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Project) {
		toSerialize["Project"] = o.Project
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeOneProject200ResponseResponse struct {
	value *DescribeOneProject200ResponseResponse
	isSet bool
}

func (v NullableDescribeOneProject200ResponseResponse) Get() *DescribeOneProject200ResponseResponse {
	return v.value
}

func (v *NullableDescribeOneProject200ResponseResponse) Set(val *DescribeOneProject200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeOneProject200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeOneProject200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeOneProject200ResponseResponse(val *DescribeOneProject200ResponseResponse) *NullableDescribeOneProject200ResponseResponse {
	return &NullableDescribeOneProject200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeOneProject200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeOneProject200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

