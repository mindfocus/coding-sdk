/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeArtifactCreditList200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeArtifactCreditList200ResponseResponse{}

// DescribeArtifactCreditList200ResponseResponse 公共返回体
type DescribeArtifactCreditList200ResponseResponse struct {
	// 授信清单列表
	InstanceSet []ArtifactsOpenApiArtifactCreditsData `json:"InstanceSet,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeArtifactCreditList200ResponseResponse instantiates a new DescribeArtifactCreditList200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeArtifactCreditList200ResponseResponse() *DescribeArtifactCreditList200ResponseResponse {
	this := DescribeArtifactCreditList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeArtifactCreditList200ResponseResponseWithDefaults instantiates a new DescribeArtifactCreditList200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeArtifactCreditList200ResponseResponseWithDefaults() *DescribeArtifactCreditList200ResponseResponse {
	this := DescribeArtifactCreditList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetInstanceSet returns the InstanceSet field value if set, zero value otherwise.
func (o *DescribeArtifactCreditList200ResponseResponse) GetInstanceSet() []ArtifactsOpenApiArtifactCreditsData {
	if o == nil || utils.IsNil(o.InstanceSet) {
		var ret []ArtifactsOpenApiArtifactCreditsData
		return ret
	}
	return o.InstanceSet
}

// GetInstanceSetOk returns a tuple with the InstanceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactCreditList200ResponseResponse) GetInstanceSetOk() ([]ArtifactsOpenApiArtifactCreditsData, bool) {
	if o == nil || utils.IsNil(o.InstanceSet) {
		return nil, false
	}
	return o.InstanceSet, true
}

// HasInstanceSet returns a boolean if a field has been set.
func (o *DescribeArtifactCreditList200ResponseResponse) HasInstanceSet() bool {
	if o != nil && !utils.IsNil(o.InstanceSet) {
		return true
	}

	return false
}

// SetInstanceSet gets a reference to the given []ArtifactsOpenApiArtifactCreditsData and assigns it to the InstanceSet field.
func (o *DescribeArtifactCreditList200ResponseResponse) SetInstanceSet(v []ArtifactsOpenApiArtifactCreditsData) {
	o.InstanceSet = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeArtifactCreditList200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactCreditList200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeArtifactCreditList200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeArtifactCreditList200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeArtifactCreditList200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeArtifactCreditList200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.InstanceSet) {
		toSerialize["InstanceSet"] = o.InstanceSet
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeArtifactCreditList200ResponseResponse struct {
	value *DescribeArtifactCreditList200ResponseResponse
	isSet bool
}

func (v NullableDescribeArtifactCreditList200ResponseResponse) Get() *DescribeArtifactCreditList200ResponseResponse {
	return v.value
}

func (v *NullableDescribeArtifactCreditList200ResponseResponse) Set(val *DescribeArtifactCreditList200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeArtifactCreditList200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeArtifactCreditList200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeArtifactCreditList200ResponseResponse(val *DescribeArtifactCreditList200ResponseResponse) *NullableDescribeArtifactCreditList200ResponseResponse {
	return &NullableDescribeArtifactCreditList200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeArtifactCreditList200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeArtifactCreditList200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


