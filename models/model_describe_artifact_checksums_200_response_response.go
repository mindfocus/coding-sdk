/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeArtifactChecksums200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeArtifactChecksums200ResponseResponse{}

// DescribeArtifactChecksums200ResponseResponse 公共返回体
type DescribeArtifactChecksums200ResponseResponse struct {
	// 制品Checksum列表
	InstanceSet []ArtifactChecksum `json:"InstanceSet,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeArtifactChecksums200ResponseResponse instantiates a new DescribeArtifactChecksums200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeArtifactChecksums200ResponseResponse() *DescribeArtifactChecksums200ResponseResponse {
	this := DescribeArtifactChecksums200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeArtifactChecksums200ResponseResponseWithDefaults instantiates a new DescribeArtifactChecksums200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeArtifactChecksums200ResponseResponseWithDefaults() *DescribeArtifactChecksums200ResponseResponse {
	this := DescribeArtifactChecksums200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetInstanceSet returns the InstanceSet field value if set, zero value otherwise.
func (o *DescribeArtifactChecksums200ResponseResponse) GetInstanceSet() []ArtifactChecksum {
	if o == nil || utils.IsNil(o.InstanceSet) {
		var ret []ArtifactChecksum
		return ret
	}
	return o.InstanceSet
}

// GetInstanceSetOk returns a tuple with the InstanceSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactChecksums200ResponseResponse) GetInstanceSetOk() ([]ArtifactChecksum, bool) {
	if o == nil || utils.IsNil(o.InstanceSet) {
		return nil, false
	}
	return o.InstanceSet, true
}

// HasInstanceSet returns a boolean if a field has been set.
func (o *DescribeArtifactChecksums200ResponseResponse) HasInstanceSet() bool {
	if o != nil && !utils.IsNil(o.InstanceSet) {
		return true
	}

	return false
}

// SetInstanceSet gets a reference to the given []ArtifactChecksum and assigns it to the InstanceSet field.
func (o *DescribeArtifactChecksums200ResponseResponse) SetInstanceSet(v []ArtifactChecksum) {
	o.InstanceSet = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeArtifactChecksums200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactChecksums200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeArtifactChecksums200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeArtifactChecksums200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeArtifactChecksums200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeArtifactChecksums200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.InstanceSet) {
		toSerialize["InstanceSet"] = o.InstanceSet
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeArtifactChecksums200ResponseResponse struct {
	value *DescribeArtifactChecksums200ResponseResponse
	isSet bool
}

func (v NullableDescribeArtifactChecksums200ResponseResponse) Get() *DescribeArtifactChecksums200ResponseResponse {
	return v.value
}

func (v *NullableDescribeArtifactChecksums200ResponseResponse) Set(val *DescribeArtifactChecksums200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeArtifactChecksums200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeArtifactChecksums200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeArtifactChecksums200ResponseResponse(val *DescribeArtifactChecksums200ResponseResponse) *NullableDescribeArtifactChecksums200ResponseResponse {
	return &NullableDescribeArtifactChecksums200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeArtifactChecksums200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeArtifactChecksums200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


