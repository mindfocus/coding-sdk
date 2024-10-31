/*
CODING OPEN API

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the CreateCdPipelineResponseData type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CreateCdPipelineResponseData{}

// CreateCdPipelineResponseData CreateCdPipelineResponseData 结构
type CreateCdPipelineResponseData struct {
	// 任务执行记录 ID
	TaskExecutionId *string `json:"TaskExecutionId,omitempty"`
	// 任务执行记录引用
	TaskExecutionRef *string `json:"TaskExecutionRef,omitempty"`
}

// NewCreateCdPipelineResponseData instantiates a new CreateCdPipelineResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateCdPipelineResponseData() *CreateCdPipelineResponseData {
	this := CreateCdPipelineResponseData{}
	var taskExecutionId string = ""
	this.TaskExecutionId = &taskExecutionId
	var taskExecutionRef string = ""
	this.TaskExecutionRef = &taskExecutionRef
	return &this
}

// NewCreateCdPipelineResponseDataWithDefaults instantiates a new CreateCdPipelineResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateCdPipelineResponseDataWithDefaults() *CreateCdPipelineResponseData {
	this := CreateCdPipelineResponseData{}
	var taskExecutionId string = ""
	this.TaskExecutionId = &taskExecutionId
	var taskExecutionRef string = ""
	this.TaskExecutionRef = &taskExecutionRef
	return &this
}

// GetTaskExecutionId returns the TaskExecutionId field value if set, zero value otherwise.
func (o *CreateCdPipelineResponseData) GetTaskExecutionId() string {
	if o == nil || utils.IsNil(o.TaskExecutionId) {
		var ret string
		return ret
	}
	return *o.TaskExecutionId
}

// GetTaskExecutionIdOk returns a tuple with the TaskExecutionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCdPipelineResponseData) GetTaskExecutionIdOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TaskExecutionId) {
		return nil, false
	}
	return o.TaskExecutionId, true
}

// HasTaskExecutionId returns a boolean if a field has been set.
func (o *CreateCdPipelineResponseData) HasTaskExecutionId() bool {
	if o != nil && !utils.IsNil(o.TaskExecutionId) {
		return true
	}

	return false
}

// SetTaskExecutionId gets a reference to the given string and assigns it to the TaskExecutionId field.
func (o *CreateCdPipelineResponseData) SetTaskExecutionId(v string) {
	o.TaskExecutionId = &v
}

// GetTaskExecutionRef returns the TaskExecutionRef field value if set, zero value otherwise.
func (o *CreateCdPipelineResponseData) GetTaskExecutionRef() string {
	if o == nil || utils.IsNil(o.TaskExecutionRef) {
		var ret string
		return ret
	}
	return *o.TaskExecutionRef
}

// GetTaskExecutionRefOk returns a tuple with the TaskExecutionRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateCdPipelineResponseData) GetTaskExecutionRefOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TaskExecutionRef) {
		return nil, false
	}
	return o.TaskExecutionRef, true
}

// HasTaskExecutionRef returns a boolean if a field has been set.
func (o *CreateCdPipelineResponseData) HasTaskExecutionRef() bool {
	if o != nil && !utils.IsNil(o.TaskExecutionRef) {
		return true
	}

	return false
}

// SetTaskExecutionRef gets a reference to the given string and assigns it to the TaskExecutionRef field.
func (o *CreateCdPipelineResponseData) SetTaskExecutionRef(v string) {
	o.TaskExecutionRef = &v
}

func (o CreateCdPipelineResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateCdPipelineResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.TaskExecutionId) {
		toSerialize["TaskExecutionId"] = o.TaskExecutionId
	}
	if !utils.IsNil(o.TaskExecutionRef) {
		toSerialize["TaskExecutionRef"] = o.TaskExecutionRef
	}
	return toSerialize, nil
}

type NullableCreateCdPipelineResponseData struct {
	value *CreateCdPipelineResponseData
	isSet bool
}

func (v NullableCreateCdPipelineResponseData) Get() *CreateCdPipelineResponseData {
	return v.value
}

func (v *NullableCreateCdPipelineResponseData) Set(val *CreateCdPipelineResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateCdPipelineResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateCdPipelineResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateCdPipelineResponseData(val *CreateCdPipelineResponseData) *NullableCreateCdPipelineResponseData {
	return &NullableCreateCdPipelineResponseData{value: val, isSet: true}
}

func (v NullableCreateCdPipelineResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateCdPipelineResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

