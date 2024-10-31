/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeCodingCIJobs200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIJobs200ResponseResponse{}

// DescribeCodingCIJobs200ResponseResponse 公共返回体
type DescribeCodingCIJobs200ResponseResponse struct {
	// CI 任务列表
	JobList []CodingCIJob `json:"JobList,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeCodingCIJobs200ResponseResponse instantiates a new DescribeCodingCIJobs200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIJobs200ResponseResponse() *DescribeCodingCIJobs200ResponseResponse {
	this := DescribeCodingCIJobs200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeCodingCIJobs200ResponseResponseWithDefaults instantiates a new DescribeCodingCIJobs200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIJobs200ResponseResponseWithDefaults() *DescribeCodingCIJobs200ResponseResponse {
	this := DescribeCodingCIJobs200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetJobList returns the JobList field value if set, zero value otherwise.
func (o *DescribeCodingCIJobs200ResponseResponse) GetJobList() []CodingCIJob {
	if o == nil || utils.IsNil(o.JobList) {
		var ret []CodingCIJob
		return ret
	}
	return o.JobList
}

// GetJobListOk returns a tuple with the JobList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIJobs200ResponseResponse) GetJobListOk() ([]CodingCIJob, bool) {
	if o == nil || utils.IsNil(o.JobList) {
		return nil, false
	}
	return o.JobList, true
}

// HasJobList returns a boolean if a field has been set.
func (o *DescribeCodingCIJobs200ResponseResponse) HasJobList() bool {
	if o != nil && !utils.IsNil(o.JobList) {
		return true
	}

	return false
}

// SetJobList gets a reference to the given []CodingCIJob and assigns it to the JobList field.
func (o *DescribeCodingCIJobs200ResponseResponse) SetJobList(v []CodingCIJob) {
	o.JobList = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeCodingCIJobs200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIJobs200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeCodingCIJobs200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeCodingCIJobs200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeCodingCIJobs200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIJobs200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.JobList) {
		toSerialize["JobList"] = o.JobList
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeCodingCIJobs200ResponseResponse struct {
	value *DescribeCodingCIJobs200ResponseResponse
	isSet bool
}

func (v NullableDescribeCodingCIJobs200ResponseResponse) Get() *DescribeCodingCIJobs200ResponseResponse {
	return v.value
}

func (v *NullableDescribeCodingCIJobs200ResponseResponse) Set(val *DescribeCodingCIJobs200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIJobs200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIJobs200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIJobs200ResponseResponse(val *DescribeCodingCIJobs200ResponseResponse) *NullableDescribeCodingCIJobs200ResponseResponse {
	return &NullableDescribeCodingCIJobs200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeCodingCIJobs200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIJobs200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

