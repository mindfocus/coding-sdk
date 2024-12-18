/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateWikiByZip200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateWikiByZip200ResponseResponse{}

// CreateWikiByZip200ResponseResponse 公共返回体
type CreateWikiByZip200ResponseResponse struct {
	// zip上传wiki 的任务id
	JobId *string `json:"JobId,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewCreateWikiByZip200ResponseResponse instantiates a new CreateWikiByZip200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWikiByZip200ResponseResponse() *CreateWikiByZip200ResponseResponse {
	this := CreateWikiByZip200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewCreateWikiByZip200ResponseResponseWithDefaults instantiates a new CreateWikiByZip200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWikiByZip200ResponseResponseWithDefaults() *CreateWikiByZip200ResponseResponse {
	this := CreateWikiByZip200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *CreateWikiByZip200ResponseResponse) GetJobId() string {
	if o == nil || utils.IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWikiByZip200ResponseResponse) GetJobIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *CreateWikiByZip200ResponseResponse) HasJobId() bool {
	if o != nil && !utils.IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *CreateWikiByZip200ResponseResponse) SetJobId(v string) {
	o.JobId = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CreateWikiByZip200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWikiByZip200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CreateWikiByZip200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CreateWikiByZip200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o CreateWikiByZip200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateWikiByZip200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.JobId) {
		toSerialize["JobId"] = o.JobId
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableCreateWikiByZip200ResponseResponse struct {
	value *CreateWikiByZip200ResponseResponse
	isSet bool
}

func (v NullableCreateWikiByZip200ResponseResponse) Get() *CreateWikiByZip200ResponseResponse {
	return v.value
}

func (v *NullableCreateWikiByZip200ResponseResponse) Set(val *CreateWikiByZip200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWikiByZip200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWikiByZip200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWikiByZip200ResponseResponse(val *CreateWikiByZip200ResponseResponse) *NullableCreateWikiByZip200ResponseResponse {
	return &NullableCreateWikiByZip200ResponseResponse{value: val, isSet: true}
}

func (v NullableCreateWikiByZip200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWikiByZip200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


