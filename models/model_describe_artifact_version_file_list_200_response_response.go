/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeArtifactVersionFileList200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeArtifactVersionFileList200ResponseResponse{}

// DescribeArtifactVersionFileList200ResponseResponse 公共返回体
type DescribeArtifactVersionFileList200ResponseResponse struct {
	// 文件列表
	InstanceSet []ArtifactVersionFileBean `json:"InstanceSet,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeArtifactVersionFileList200ResponseResponse instantiates a new DescribeArtifactVersionFileList200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeArtifactVersionFileList200ResponseResponse() *DescribeArtifactVersionFileList200ResponseResponse {
	this := DescribeArtifactVersionFileList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeArtifactVersionFileList200ResponseResponseWithDefaults instantiates a new DescribeArtifactVersionFileList200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeArtifactVersionFileList200ResponseResponseWithDefaults() *DescribeArtifactVersionFileList200ResponseResponse {
	this := DescribeArtifactVersionFileList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetInstanceSet returns the InstanceSet field value if set, zero value otherwise.
func (o *DescribeArtifactVersionFileList200ResponseResponse) GetInstanceSet() []ArtifactVersionFileBean {
	if o == nil || utils.IsNil(o.InstanceSet) {
		var ret []ArtifactVersionFileBean
		return ret
	}
	return o.InstanceSet
}

// GetInstanceSetOk returns a tuple with the InstanceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactVersionFileList200ResponseResponse) GetInstanceSetOk() ([]ArtifactVersionFileBean, bool) {
	if o == nil || utils.IsNil(o.InstanceSet) {
		return nil, false
	}
	return o.InstanceSet, true
}

// HasInstanceSet returns a boolean if a field has been set.
func (o *DescribeArtifactVersionFileList200ResponseResponse) HasInstanceSet() bool {
	if o != nil && !utils.IsNil(o.InstanceSet) {
		return true
	}

	return false
}

// SetInstanceSet gets a reference to the given []ArtifactVersionFileBean and assigns it to the InstanceSet field.
func (o *DescribeArtifactVersionFileList200ResponseResponse) SetInstanceSet(v []ArtifactVersionFileBean) {
	o.InstanceSet = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeArtifactVersionFileList200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactVersionFileList200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeArtifactVersionFileList200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeArtifactVersionFileList200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeArtifactVersionFileList200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeArtifactVersionFileList200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.InstanceSet) {
		toSerialize["InstanceSet"] = o.InstanceSet
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeArtifactVersionFileList200ResponseResponse struct {
	value *DescribeArtifactVersionFileList200ResponseResponse
	isSet bool
}

func (v NullableDescribeArtifactVersionFileList200ResponseResponse) Get() *DescribeArtifactVersionFileList200ResponseResponse {
	return v.value
}

func (v *NullableDescribeArtifactVersionFileList200ResponseResponse) Set(val *DescribeArtifactVersionFileList200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeArtifactVersionFileList200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeArtifactVersionFileList200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeArtifactVersionFileList200ResponseResponse(val *DescribeArtifactVersionFileList200ResponseResponse) *NullableDescribeArtifactVersionFileList200ResponseResponse {
	return &NullableDescribeArtifactVersionFileList200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeArtifactVersionFileList200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeArtifactVersionFileList200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


