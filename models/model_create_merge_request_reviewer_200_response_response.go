/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateMergeRequestReviewer200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateMergeRequestReviewer200ResponseResponse{}

// CreateMergeRequestReviewer200ResponseResponse 公共返回体
type CreateMergeRequestReviewer200ResponseResponse struct {
	// 评审者信息
	Reviewers []CodingUser `json:"Reviewers,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewCreateMergeRequestReviewer200ResponseResponse instantiates a new CreateMergeRequestReviewer200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMergeRequestReviewer200ResponseResponse() *CreateMergeRequestReviewer200ResponseResponse {
	this := CreateMergeRequestReviewer200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewCreateMergeRequestReviewer200ResponseResponseWithDefaults instantiates a new CreateMergeRequestReviewer200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMergeRequestReviewer200ResponseResponseWithDefaults() *CreateMergeRequestReviewer200ResponseResponse {
	this := CreateMergeRequestReviewer200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetReviewers returns the Reviewers field value if set, zero value otherwise.
func (o *CreateMergeRequestReviewer200ResponseResponse) GetReviewers() []CodingUser {
	if o == nil || utils.IsNil(o.Reviewers) {
		var ret []CodingUser
		return ret
	}
	return o.Reviewers
}

// GetReviewersOk returns a tuple with the Reviewers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMergeRequestReviewer200ResponseResponse) GetReviewersOk() ([]CodingUser, bool) {
	if o == nil || utils.IsNil(o.Reviewers) {
		return nil, false
	}
	return o.Reviewers, true
}

// HasReviewers returns a boolean if a field has been set.
func (o *CreateMergeRequestReviewer200ResponseResponse) HasReviewers() bool {
	if o != nil && !utils.IsNil(o.Reviewers) {
		return true
	}

	return false
}

// SetReviewers gets a reference to the given []CodingUser and assigns it to the Reviewers field.
func (o *CreateMergeRequestReviewer200ResponseResponse) SetReviewers(v []CodingUser) {
	o.Reviewers = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CreateMergeRequestReviewer200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMergeRequestReviewer200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CreateMergeRequestReviewer200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CreateMergeRequestReviewer200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o CreateMergeRequestReviewer200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMergeRequestReviewer200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Reviewers) {
		toSerialize["Reviewers"] = o.Reviewers
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableCreateMergeRequestReviewer200ResponseResponse struct {
	value *CreateMergeRequestReviewer200ResponseResponse
	isSet bool
}

func (v NullableCreateMergeRequestReviewer200ResponseResponse) Get() *CreateMergeRequestReviewer200ResponseResponse {
	return v.value
}

func (v *NullableCreateMergeRequestReviewer200ResponseResponse) Set(val *CreateMergeRequestReviewer200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMergeRequestReviewer200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMergeRequestReviewer200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMergeRequestReviewer200ResponseResponse(val *CreateMergeRequestReviewer200ResponseResponse) *NullableCreateMergeRequestReviewer200ResponseResponse {
	return &NullableCreateMergeRequestReviewer200ResponseResponse{value: val, isSet: true}
}

func (v NullableCreateMergeRequestReviewer200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMergeRequestReviewer200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


