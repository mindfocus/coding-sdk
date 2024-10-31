/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeArtifactRepositoryFileList200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeArtifactRepositoryFileList200ResponseResponse{}

// DescribeArtifactRepositoryFileList200ResponseResponse 公共返回体
type DescribeArtifactRepositoryFileList200ResponseResponse struct {
	Data *ArtifactRepositoryFileListData `json:"Data,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeArtifactRepositoryFileList200ResponseResponse instantiates a new DescribeArtifactRepositoryFileList200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeArtifactRepositoryFileList200ResponseResponse() *DescribeArtifactRepositoryFileList200ResponseResponse {
	this := DescribeArtifactRepositoryFileList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeArtifactRepositoryFileList200ResponseResponseWithDefaults instantiates a new DescribeArtifactRepositoryFileList200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeArtifactRepositoryFileList200ResponseResponseWithDefaults() *DescribeArtifactRepositoryFileList200ResponseResponse {
	this := DescribeArtifactRepositoryFileList200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DescribeArtifactRepositoryFileList200ResponseResponse) GetData() ArtifactRepositoryFileListData {
	if o == nil || utils.IsNil(o.Data) {
		var ret ArtifactRepositoryFileListData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactRepositoryFileList200ResponseResponse) GetDataOk() (*ArtifactRepositoryFileListData, bool) {
	if o == nil || utils.IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DescribeArtifactRepositoryFileList200ResponseResponse) HasData() bool {
	if o != nil && !utils.IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ArtifactRepositoryFileListData and assigns it to the Data field.
func (o *DescribeArtifactRepositoryFileList200ResponseResponse) SetData(v ArtifactRepositoryFileListData) {
	o.Data = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeArtifactRepositoryFileList200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeArtifactRepositoryFileList200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeArtifactRepositoryFileList200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeArtifactRepositoryFileList200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeArtifactRepositoryFileList200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeArtifactRepositoryFileList200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Data) {
		toSerialize["Data"] = o.Data
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeArtifactRepositoryFileList200ResponseResponse struct {
	value *DescribeArtifactRepositoryFileList200ResponseResponse
	isSet bool
}

func (v NullableDescribeArtifactRepositoryFileList200ResponseResponse) Get() *DescribeArtifactRepositoryFileList200ResponseResponse {
	return v.value
}

func (v *NullableDescribeArtifactRepositoryFileList200ResponseResponse) Set(val *DescribeArtifactRepositoryFileList200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeArtifactRepositoryFileList200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeArtifactRepositoryFileList200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeArtifactRepositoryFileList200ResponseResponse(val *DescribeArtifactRepositoryFileList200ResponseResponse) *NullableDescribeArtifactRepositoryFileList200ResponseResponse {
	return &NullableDescribeArtifactRepositoryFileList200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeArtifactRepositoryFileList200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeArtifactRepositoryFileList200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

