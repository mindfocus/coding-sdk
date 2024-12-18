/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the ReorderCdPipelinesResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ReorderCdPipelinesResponseData{}

// ReorderCdPipelinesResponseData ReorderCdPipelinesResponseData 结构
type ReorderCdPipelinesResponseData struct {
	// 任务执行记录 JSON
	TaskExecutionJsonContent *string `json:"TaskExecutionJsonContent,omitempty"`
	// 任务执行状态
	TaskExecutionStatus *string `json:"TaskExecutionStatus,omitempty"`
}

// NewReorderCdPipelinesResponseData instantiates a new ReorderCdPipelinesResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReorderCdPipelinesResponseData() *ReorderCdPipelinesResponseData {
	this := ReorderCdPipelinesResponseData{}
	var taskExecutionJsonContent string = ""
	this.TaskExecutionJsonContent = &taskExecutionJsonContent
	var taskExecutionStatus string = ""
	this.TaskExecutionStatus = &taskExecutionStatus
	return &this
}

// NewReorderCdPipelinesResponseDataWithDefaults instantiates a new ReorderCdPipelinesResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReorderCdPipelinesResponseDataWithDefaults() *ReorderCdPipelinesResponseData {
	this := ReorderCdPipelinesResponseData{}
	var taskExecutionJsonContent string = ""
	this.TaskExecutionJsonContent = &taskExecutionJsonContent
	var taskExecutionStatus string = ""
	this.TaskExecutionStatus = &taskExecutionStatus
	return &this
}

// GetTaskExecutionJsonContent returns the TaskExecutionJsonContent field value if set, zero value otherwise.
func (o *ReorderCdPipelinesResponseData) GetTaskExecutionJsonContent() string {
	if o == nil || utils.IsNil(o.TaskExecutionJsonContent) {
		var ret string
		return ret
	}
	return *o.TaskExecutionJsonContent
}

// GetTaskExecutionJsonContentOk returns a tuple with the TaskExecutionJsonContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReorderCdPipelinesResponseData) GetTaskExecutionJsonContentOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TaskExecutionJsonContent) {
		return nil, false
	}
	return o.TaskExecutionJsonContent, true
}

// HasTaskExecutionJsonContent returns a boolean if a field has been set.
func (o *ReorderCdPipelinesResponseData) HasTaskExecutionJsonContent() bool {
	if o != nil && !utils.IsNil(o.TaskExecutionJsonContent) {
		return true
	}

	return false
}

// SetTaskExecutionJsonContent gets a reference to the given string and assigns it to the TaskExecutionJsonContent field.
func (o *ReorderCdPipelinesResponseData) SetTaskExecutionJsonContent(v string) {
	o.TaskExecutionJsonContent = &v
}

// GetTaskExecutionStatus returns the TaskExecutionStatus field value if set, zero value otherwise.
func (o *ReorderCdPipelinesResponseData) GetTaskExecutionStatus() string {
	if o == nil || utils.IsNil(o.TaskExecutionStatus) {
		var ret string
		return ret
	}
	return *o.TaskExecutionStatus
}

// GetTaskExecutionStatusOk returns a tuple with the TaskExecutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReorderCdPipelinesResponseData) GetTaskExecutionStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TaskExecutionStatus) {
		return nil, false
	}
	return o.TaskExecutionStatus, true
}

// HasTaskExecutionStatus returns a boolean if a field has been set.
func (o *ReorderCdPipelinesResponseData) HasTaskExecutionStatus() bool {
	if o != nil && !utils.IsNil(o.TaskExecutionStatus) {
		return true
	}

	return false
}

// SetTaskExecutionStatus gets a reference to the given string and assigns it to the TaskExecutionStatus field.
func (o *ReorderCdPipelinesResponseData) SetTaskExecutionStatus(v string) {
	o.TaskExecutionStatus = &v
}

func (o ReorderCdPipelinesResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReorderCdPipelinesResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TaskExecutionJsonContent) {
		toSerialize["TaskExecutionJsonContent"] = o.TaskExecutionJsonContent
	}
	if !utils.IsNil(o.TaskExecutionStatus) {
		toSerialize["TaskExecutionStatus"] = o.TaskExecutionStatus
	}
	return toSerialize, nil
}

type NullableReorderCdPipelinesResponseData struct {
	value *ReorderCdPipelinesResponseData
	isSet bool
}

func (v NullableReorderCdPipelinesResponseData) Get() *ReorderCdPipelinesResponseData {
	return v.value
}

func (v *NullableReorderCdPipelinesResponseData) Set(val *ReorderCdPipelinesResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableReorderCdPipelinesResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableReorderCdPipelinesResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReorderCdPipelinesResponseData(val *ReorderCdPipelinesResponseData) *NullableReorderCdPipelinesResponseData {
	return &NullableReorderCdPipelinesResponseData{value: val, isSet: true}
}

func (v NullableReorderCdPipelinesResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReorderCdPipelinesResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


