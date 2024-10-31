/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeMergeRequestLog200ResponseResponse type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeMergeRequestLog200ResponseResponse{}

// DescribeMergeRequestLog200ResponseResponse 公共返回体
type DescribeMergeRequestLog200ResponseResponse struct {
	// 操作记录列表
	Logs []MergeRequestLog `json:"Logs,omitempty"`
	// 请求id
	RequestId *string `json:"RequestId,omitempty"`
}

// NewDescribeMergeRequestLog200ResponseResponse instantiates a new DescribeMergeRequestLog200ResponseResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeMergeRequestLog200ResponseResponse() *DescribeMergeRequestLog200ResponseResponse {
	this := DescribeMergeRequestLog200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// NewDescribeMergeRequestLog200ResponseResponseWithDefaults instantiates a new DescribeMergeRequestLog200ResponseResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeMergeRequestLog200ResponseResponseWithDefaults() *DescribeMergeRequestLog200ResponseResponse {
	this := DescribeMergeRequestLog200ResponseResponse{}
	var requestId string = "xxxxx"
	this.RequestId = &requestId
	return &this
}

// GetLogs returns the Logs field value if set, zero value otherwise.
func (o *DescribeMergeRequestLog200ResponseResponse) GetLogs() []MergeRequestLog {
	if o == nil || utils.IsNil(o.Logs) {
		var ret []MergeRequestLog
		return ret
	}
	return o.Logs
}

// GetLogsOk returns a tuple with the Logs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeMergeRequestLog200ResponseResponse) GetLogsOk() ([]MergeRequestLog, bool) {
	if o == nil || utils.IsNil(o.Logs) {
		return nil, false
	}
	return o.Logs, true
}

// HasLogs returns a boolean if a field has been set.
func (o *DescribeMergeRequestLog200ResponseResponse) HasLogs() bool {
	if o != nil && !utils.IsNil(o.Logs) {
		return true
	}

	return false
}

// SetLogs gets a reference to the given []MergeRequestLog and assigns it to the Logs field.
func (o *DescribeMergeRequestLog200ResponseResponse) SetLogs(v []MergeRequestLog) {
	o.Logs = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *DescribeMergeRequestLog200ResponseResponse) GetRequestId() string {
	if o == nil || utils.IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeMergeRequestLog200ResponseResponse) GetRequestIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *DescribeMergeRequestLog200ResponseResponse) HasRequestId() bool {
	if o != nil && !utils.IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *DescribeMergeRequestLog200ResponseResponse) SetRequestId(v string) {
	o.RequestId = &v
}

func (o DescribeMergeRequestLog200ResponseResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeMergeRequestLog200ResponseResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Logs) {
		toSerialize["Logs"] = o.Logs
	}
	if !utils.IsNil(o.RequestId) {
		toSerialize["RequestId"] = o.RequestId
	}
	return toSerialize, nil
}

type NullableDescribeMergeRequestLog200ResponseResponse struct {
	value *DescribeMergeRequestLog200ResponseResponse
	isSet bool
}

func (v NullableDescribeMergeRequestLog200ResponseResponse) Get() *DescribeMergeRequestLog200ResponseResponse {
	return v.value
}

func (v *NullableDescribeMergeRequestLog200ResponseResponse) Set(val *DescribeMergeRequestLog200ResponseResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeMergeRequestLog200ResponseResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeMergeRequestLog200ResponseResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeMergeRequestLog200ResponseResponse(val *DescribeMergeRequestLog200ResponseResponse) *NullableDescribeMergeRequestLog200ResponseResponse {
	return &NullableDescribeMergeRequestLog200ResponseResponse{value: val, isSet: true}
}

func (v NullableDescribeMergeRequestLog200ResponseResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeMergeRequestLog200ResponseResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

