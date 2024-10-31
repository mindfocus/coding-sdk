/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeUsersOnResourceAndGrantObject200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeUsersOnResourceAndGrantObject200ResponseResponse{}

// DescribeUsersOnResourceAndGrantObject200ResponseResponse 公共返回体
type DescribeUsersOnResourceAndGrantObject200ResponseResponse struct {
	Data *DescribeUsersOnResourceAndGrantObjectPageData `json:"Data,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeUsersOnResourceAndGrantObject200ResponseResponse instantiates a new DescribeUsersOnResourceAndGrantObject200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeUsersOnResourceAndGrantObject200ResponseResponse() *DescribeUsersOnResourceAndGrantObject200ResponseResponse {
	this := DescribeUsersOnResourceAndGrantObject200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeUsersOnResourceAndGrantObject200ResponseResponseWithDefaults instantiates a new DescribeUsersOnResourceAndGrantObject200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeUsersOnResourceAndGrantObject200ResponseResponseWithDefaults() *DescribeUsersOnResourceAndGrantObject200ResponseResponse {
	this := DescribeUsersOnResourceAndGrantObject200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DescribeUsersOnResourceAndGrantObject200ResponseResponse) GetData() DescribeUsersOnResourceAndGrantObjectPageData {
	if o == nil || utils.IsNil(o.Data) {
		var ret DescribeUsersOnResourceAndGrantObjectPageData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeUsersOnResourceAndGrantObject200ResponseResponse) GetDataOk() (*DescribeUsersOnResourceAndGrantObjectPageData, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DescribeUsersOnResourceAndGrantObject200ResponseResponse) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given DescribeUsersOnResourceAndGrantObjectPageData and assigns it to the Data field.
func (o *DescribeUsersOnResourceAndGrantObject200ResponseResponse) SetData(v DescribeUsersOnResourceAndGrantObjectPageData) {
	o.Data = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeUsersOnResourceAndGrantObject200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeUsersOnResourceAndGrantObject200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeUsersOnResourceAndGrantObject200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeUsersOnResourceAndGrantObject200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeUsersOnResourceAndGrantObject200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeUsersOnResourceAndGrantObject200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeUsersOnResourceAndGrantObject200ResponseResponse struct {
	value *DescribeUsersOnResourceAndGrantObject200ResponseResponse
	isSet bool
}

func (v NullableDescribeUsersOnResourceAndGrantObject200ResponseResponse) Get() *DescribeUsersOnResourceAndGrantObject200ResponseResponse {
	return v.value
}

func (v *NullableDescribeUsersOnResourceAndGrantObject200ResponseResponse) Set(val *DescribeUsersOnResourceAndGrantObject200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeUsersOnResourceAndGrantObject200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeUsersOnResourceAndGrantObject200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeUsersOnResourceAndGrantObject200ResponseResponse(val *DescribeUsersOnResourceAndGrantObject200ResponseResponse) *NullableDescribeUsersOnResourceAndGrantObject200ResponseResponse {
	return &NullableDescribeUsersOnResourceAndGrantObject200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeUsersOnResourceAndGrantObject200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeUsersOnResourceAndGrantObject200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

