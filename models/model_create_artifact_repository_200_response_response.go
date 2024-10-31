/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateArtifactRepository200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateArtifactRepository200ResponseResponse{}

// CreateArtifactRepository200ResponseResponse 公共返回体
type CreateArtifactRepository200ResponseResponse struct {
	// 创建成功的仓库 ID
	Id *int32 `json:"Id,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewCreateArtifactRepository200ResponseResponse instantiates a new CreateArtifactRepository200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateArtifactRepository200ResponseResponse() *CreateArtifactRepository200ResponseResponse {
	this := CreateArtifactRepository200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewCreateArtifactRepository200ResponseResponseWithDefaults instantiates a new CreateArtifactRepository200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateArtifactRepository200ResponseResponseWithDefaults() *CreateArtifactRepository200ResponseResponse {
	this := CreateArtifactRepository200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateArtifactRepository200ResponseResponse) GetId() int32 {
	if o == nil || utils.IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactRepository200ResponseResponse) GetIdOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateArtifactRepository200ResponseResponse) HasId() bool {
	if o != nil && !utils.IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *CreateArtifactRepository200ResponseResponse) SetId(v int32) {
	o.Id = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CreateArtifactRepository200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactRepository200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CreateArtifactRepository200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CreateArtifactRepository200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o CreateArtifactRepository200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateArtifactRepository200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Id) {
		toSerialize["Id"] = o.Id
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableCreateArtifactRepository200ResponseResponse struct {
	value *CreateArtifactRepository200ResponseResponse
	isSet bool
}

func (v NullableCreateArtifactRepository200ResponseResponse) Get() *CreateArtifactRepository200ResponseResponse {
	return v.value
}

func (v *NullableCreateArtifactRepository200ResponseResponse) Set(val *CreateArtifactRepository200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateArtifactRepository200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateArtifactRepository200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateArtifactRepository200ResponseResponse(val *CreateArtifactRepository200ResponseResponse) *NullableCreateArtifactRepository200ResponseResponse {
	return &NullableCreateArtifactRepository200ResponseResponse{value: val, isSet: true}
}

func (v NullableCreateArtifactRepository200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateArtifactRepository200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


