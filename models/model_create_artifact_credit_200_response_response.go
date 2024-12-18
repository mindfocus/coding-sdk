/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateArtifactCredit200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateArtifactCredit200ResponseResponse{}

// CreateArtifactCredit200ResponseResponse 公共返回体
type CreateArtifactCredit200ResponseResponse struct {
	// 授信清单ID
	ArtifactCreditId *string `json:"ArtifactCreditId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewCreateArtifactCredit200ResponseResponse instantiates a new CreateArtifactCredit200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateArtifactCredit200ResponseResponse() *CreateArtifactCredit200ResponseResponse {
	this := CreateArtifactCredit200ResponseResponse{}
	var artifactCreditId string = ""
	this.ArtifactCreditId = &artifactCreditId
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewCreateArtifactCredit200ResponseResponseWithDefaults instantiates a new CreateArtifactCredit200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateArtifactCredit200ResponseResponseWithDefaults() *CreateArtifactCredit200ResponseResponse {
	this := CreateArtifactCredit200ResponseResponse{}
	var artifactCreditId string = ""
	this.ArtifactCreditId = &artifactCreditId
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetArtifactCreditId returns the ArtifactCreditId field value if set, zero value otherwise.
func (o *CreateArtifactCredit200ResponseResponse) GetArtifactCreditId() string {
	if o == nil || utils.IsNil(o.ArtifactCreditId) {
		var ret string
		return ret
	}
	return *o.ArtifactCreditId
}

// GetArtifactCreditIdOk returns a tuple with the ArtifactCreditId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactCredit200ResponseResponse) GetArtifactCreditIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ArtifactCreditId) {
		return nil, false
	}
	return o.ArtifactCreditId, true
}

// HasArtifactCreditId returns a boolean if a field has been set.
func (o *CreateArtifactCredit200ResponseResponse) HasArtifactCreditId() bool {
	if o != nil && !utils.IsNil(o.ArtifactCreditId) {
		return true
	}

	return false
}

// SetArtifactCreditId gets a reference to the given string and assigns it to the ArtifactCreditId field.
func (o *CreateArtifactCredit200ResponseResponse) SetArtifactCreditId(v string) {
	o.ArtifactCreditId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CreateArtifactCredit200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateArtifactCredit200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CreateArtifactCredit200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CreateArtifactCredit200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o CreateArtifactCredit200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateArtifactCredit200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ArtifactCreditId) {
		toSerialize["ArtifactCreditId"] = o.ArtifactCreditId
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableCreateArtifactCredit200ResponseResponse struct {
	value *CreateArtifactCredit200ResponseResponse
	isSet bool
}

func (v NullableCreateArtifactCredit200ResponseResponse) Get() *CreateArtifactCredit200ResponseResponse {
	return v.value
}

func (v *NullableCreateArtifactCredit200ResponseResponse) Set(val *CreateArtifactCredit200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateArtifactCredit200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateArtifactCredit200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateArtifactCredit200ResponseResponse(val *CreateArtifactCredit200ResponseResponse) *NullableCreateArtifactCredit200ResponseResponse {
	return &NullableCreateArtifactCredit200ResponseResponse{value: val, isSet: true}
}

func (v NullableCreateArtifactCredit200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateArtifactCredit200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


