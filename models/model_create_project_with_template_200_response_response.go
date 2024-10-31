/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateProjectWithTemplate200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateProjectWithTemplate200ResponseResponse{}

// CreateProjectWithTemplate200ResponseResponse 公共返回体
type CreateProjectWithTemplate200ResponseResponse struct {
	// 项目Id
	ProjectId int64 `json:"ProjectId,omitempty"`
	// 请求id
	RequestId string `json:"RequestId,omitempty"`
}

// NewCreateProjectWithTemplate200ResponseResponse instantiates a new CreateProjectWithTemplate200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateProjectWithTemplate200ResponseResponse() *CreateProjectWithTemplate200ResponseResponse {
	this := CreateProjectWithTemplate200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = requestId
	return &this
}

// NewCreateProjectWithTemplate200ResponseResponseWithDefaults instantiates a new CreateProjectWithTemplate200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateProjectWithTemplate200ResponseResponseWithDefaults() *CreateProjectWithTemplate200ResponseResponse {
	this := CreateProjectWithTemplate200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = requestId
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CreateProjectWithTemplate200ResponseResponse) GetProjectId() int64 {
	if o == nil || utils.IsNil(o.ProjectId) {
		var ret int64
		return ret
	}
	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectWithTemplate200ResponseResponse) GetProjectIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.ProjectId) {
		return nil, false
	}
	return &o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CreateProjectWithTemplate200ResponseResponse) HasProjectId() bool {
	if o != nil && !utils.IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given int32 and assigns it to the ProjectId field.
func (o *CreateProjectWithTemplate200ResponseResponse) SetProjectId(v int64) {
	o.ProjectId = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *CreateProjectWithTemplate200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateProjectWithTemplate200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return &o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *CreateProjectWithTemplate200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *CreateProjectWithTemplate200ResponseResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o CreateProjectWithTemplate200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateProjectWithTemplate200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ProjectId) {
		toSerialize["ProjectId"] = o.ProjectId
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableCreateProjectWithTemplate200ResponseResponse struct {
	value *CreateProjectWithTemplate200ResponseResponse
	isSet bool
}

func (v NullableCreateProjectWithTemplate200ResponseResponse) Get() *CreateProjectWithTemplate200ResponseResponse {
	return v.value
}

func (v *NullableCreateProjectWithTemplate200ResponseResponse) Set(val *CreateProjectWithTemplate200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateProjectWithTemplate200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateProjectWithTemplate200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateProjectWithTemplate200ResponseResponse(val *CreateProjectWithTemplate200ResponseResponse) *NullableCreateProjectWithTemplate200ResponseResponse {
	return &NullableCreateProjectWithTemplate200ResponseResponse{value: val, isSet: true}
}

func (v NullableCreateProjectWithTemplate200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateProjectWithTemplate200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


