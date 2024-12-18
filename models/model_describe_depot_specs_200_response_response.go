/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeDepotSpecs200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeDepotSpecs200ResponseResponse{}

// DescribeDepotSpecs200ResponseResponse 公共返回体
type DescribeDepotSpecs200ResponseResponse struct {
	// 仓库规范列表
	DepotSpecs []DepotSpec `json:"DepotSpecs,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeDepotSpecs200ResponseResponse instantiates a new DescribeDepotSpecs200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDepotSpecs200ResponseResponse() *DescribeDepotSpecs200ResponseResponse {
	this := DescribeDepotSpecs200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeDepotSpecs200ResponseResponseWithDefaults instantiates a new DescribeDepotSpecs200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDepotSpecs200ResponseResponseWithDefaults() *DescribeDepotSpecs200ResponseResponse {
	this := DescribeDepotSpecs200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetDepotSpecs returns the DepotSpecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DescribeDepotSpecs200ResponseResponse) GetDepotSpecs() []DepotSpec {
	if o == nil {
		var ret []DepotSpec
		return ret
	}
	return o.DepotSpecs
}

// GetDepotSpecsOk returns a tuple with the DepotSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DescribeDepotSpecs200ResponseResponse) GetDepotSpecsOk() ([]DepotSpec, bool) {
	if o == nil || utils.IsNil(o.DepotSpecs) {
		return nil, false
	}
	return o.DepotSpecs, true
}

// HasDepotSpecs returns a boolean if a field has been set.
func (o *DescribeDepotSpecs200ResponseResponse) HasDepotSpecs() bool {
	if o != nil && !utils.IsNil(o.DepotSpecs) {
		return true
	}

	return false
}

// SetDepotSpecs gets a reference to the given []DepotSpec and assigns it to the DepotSpecs field.
func (o *DescribeDepotSpecs200ResponseResponse) SetDepotSpecs(v []DepotSpec) {
	o.DepotSpecs = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeDepotSpecs200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDepotSpecs200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeDepotSpecs200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeDepotSpecs200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeDepotSpecs200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDepotSpecs200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.DepotSpecs != nil {
		toSerialize["DepotSpecs"] = o.DepotSpecs
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeDepotSpecs200ResponseResponse struct {
	value *DescribeDepotSpecs200ResponseResponse
	isSet bool
}

func (v NullableDescribeDepotSpecs200ResponseResponse) Get() *DescribeDepotSpecs200ResponseResponse {
	return v.value
}

func (v *NullableDescribeDepotSpecs200ResponseResponse) Set(val *DescribeDepotSpecs200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDepotSpecs200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDepotSpecs200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDepotSpecs200ResponseResponse(val *DescribeDepotSpecs200ResponseResponse) *NullableDescribeDepotSpecs200ResponseResponse {
	return &NullableDescribeDepotSpecs200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeDepotSpecs200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDepotSpecs200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


