/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeDifferentBetween2Commits200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeDifferentBetween2Commits200ResponseResponse{}

// DescribeDifferentBetween2Commits200ResponseResponse 公共返回体
type DescribeDifferentBetween2Commits200ResponseResponse struct {
	DiffFileInfo *QcloudApiGitDepotDiffFileInfo `json:"DiffFileInfo,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeDifferentBetween2Commits200ResponseResponse instantiates a new DescribeDifferentBetween2Commits200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeDifferentBetween2Commits200ResponseResponse() *DescribeDifferentBetween2Commits200ResponseResponse {
	this := DescribeDifferentBetween2Commits200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeDifferentBetween2Commits200ResponseResponseWithDefaults instantiates a new DescribeDifferentBetween2Commits200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeDifferentBetween2Commits200ResponseResponseWithDefaults() *DescribeDifferentBetween2Commits200ResponseResponse {
	this := DescribeDifferentBetween2Commits200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetDiffFileInfo returns the DiffFileInfo field value if set, zero value otherwise.
func (o *DescribeDifferentBetween2Commits200ResponseResponse) GetDiffFileInfo() QcloudApiGitDepotDiffFileInfo {
	if o == nil || utils.IsNil(o.DiffFileInfo) {
		var ret QcloudApiGitDepotDiffFileInfo
		return ret
	}
	return *o.DiffFileInfo
}

// GetDiffFileInfoOk returns a tuple with the DiffFileInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDifferentBetween2Commits200ResponseResponse) GetDiffFileInfoOk() (*QcloudApiGitDepotDiffFileInfo, bool) {
	if o == nil || utils.IsNil(o.DiffFileInfo) {
		return nil, false
	}
	return o.DiffFileInfo, true
}

// HasDiffFileInfo returns a boolean if a field has been set.
func (o *DescribeDifferentBetween2Commits200ResponseResponse) HasDiffFileInfo() bool {
	if o != nil && !utils.IsNil(o.DiffFileInfo) {
		return true
	}

	return false
}

// SetDiffFileInfo gets a reference to the given QcloudApiGitDepotDiffFileInfo and assigns it to the DiffFileInfo field.
func (o *DescribeDifferentBetween2Commits200ResponseResponse) SetDiffFileInfo(v QcloudApiGitDepotDiffFileInfo) {
	o.DiffFileInfo = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeDifferentBetween2Commits200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeDifferentBetween2Commits200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeDifferentBetween2Commits200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeDifferentBetween2Commits200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeDifferentBetween2Commits200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeDifferentBetween2Commits200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.DiffFileInfo) {
		toSerialize["DiffFileInfo"] = o.DiffFileInfo
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeDifferentBetween2Commits200ResponseResponse struct {
	value *DescribeDifferentBetween2Commits200ResponseResponse
	isSet bool
}

func (v NullableDescribeDifferentBetween2Commits200ResponseResponse) Get() *DescribeDifferentBetween2Commits200ResponseResponse {
	return v.value
}

func (v *NullableDescribeDifferentBetween2Commits200ResponseResponse) Set(val *DescribeDifferentBetween2Commits200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeDifferentBetween2Commits200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeDifferentBetween2Commits200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeDifferentBetween2Commits200ResponseResponse(val *DescribeDifferentBetween2Commits200ResponseResponse) *NullableDescribeDifferentBetween2Commits200ResponseResponse {
	return &NullableDescribeDifferentBetween2Commits200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeDifferentBetween2Commits200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeDifferentBetween2Commits200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


