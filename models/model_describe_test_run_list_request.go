/*
CODING OPEN API

  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package models

import (
	"encoding/json"
	"bytes"
	"fmt"; "github.com/mindfocus/coding-sdk/utils"
)

// checks if the DescribeTestRunListRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeTestRunListRequest{}

// DescribeTestRunListRequest struct for DescribeTestRunListRequest
type DescribeTestRunListRequest struct {
	// 执行方式: 1-手动执行 2-自动化流水线执行
	ExecuteType *int32 `json:"ExecuteType,omitempty"`
	// 发布版本状态：0-未发布 1-已发布（与参数IterationId、IterationStatus、SectionId互斥）
	GitReleaseState *int32 `json:"GitReleaseState,omitempty"`
	// 是否已经归档
	IsCompleted *bool `json:"IsCompleted,omitempty"`
	// 迭代 ID（与参数IterationStatus、GitReleaseState、SectionId互斥）
	IterationId []int32 `json:"IterationId,omitempty"`
	// 迭代状态: WAIT_PROCESS、PROCESSING、COMPLETED（与参数IterationId、GitReleaseState、SectionId互斥）
	IterationStatus []string `json:"IterationStatus,omitempty"`
	// 计划名称
	Keyword *string `json:"Keyword,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 分组 ID（与参数IterationId、IterationStatus、GitReleaseState互斥）
	SectionId *int64 `json:"SectionId,omitempty"`
	// 状态: 0-未开始 1-进行中 2-已测完
	State *int32 `json:"State,omitempty"`
}

type _DescribeTestRunListRequest DescribeTestRunListRequest

// NewDescribeTestRunListRequest instantiates a new DescribeTestRunListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeTestRunListRequest(projectName string) *DescribeTestRunListRequest {
	this := DescribeTestRunListRequest{}
	this.ProjectName = projectName
	return &this
}

// NewDescribeTestRunListRequestWithDefaults instantiates a new DescribeTestRunListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeTestRunListRequestWithDefaults() *DescribeTestRunListRequest {
	this := DescribeTestRunListRequest{}
	return &this
}

// GetExecuteType returns the ExecuteType field value if set, zero value otherwise.
func (o *DescribeTestRunListRequest) GetExecuteType() int32 {
	if o == nil || utils.IsNil(o.ExecuteType) {
		var ret int32
		return ret
	}
	return *o.ExecuteType
}

// GetExecuteTypeOk returns a tuple with the ExecuteType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestRunListRequest) GetExecuteTypeOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.ExecuteType) {
		return nil, false
	}
	return o.ExecuteType, true
}

// HasExecuteType returns a boolean if a field has been set.
func (o *DescribeTestRunListRequest) HasExecuteType() bool {
	if o != nil && !utils.IsNil(o.ExecuteType) {
		return true
	}

	return false
}

// SetExecuteType gets a reference to the given int32 and assigns it to the ExecuteType field.
func (o *DescribeTestRunListRequest) SetExecuteType(v int32) {
	o.ExecuteType = &v
}

// GetGitReleaseState returns the GitReleaseState field value if set, zero value otherwise.
func (o *DescribeTestRunListRequest) GetGitReleaseState() int32 {
	if o == nil || utils.IsNil(o.GitReleaseState) {
		var ret int32
		return ret
	}
	return *o.GitReleaseState
}

// GetGitReleaseStateOk returns a tuple with the GitReleaseState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestRunListRequest) GetGitReleaseStateOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.GitReleaseState) {
		return nil, false
	}
	return o.GitReleaseState, true
}

// HasGitReleaseState returns a boolean if a field has been set.
func (o *DescribeTestRunListRequest) HasGitReleaseState() bool {
	if o != nil && !utils.IsNil(o.GitReleaseState) {
		return true
	}

	return false
}

// SetGitReleaseState gets a reference to the given int32 and assigns it to the GitReleaseState field.
func (o *DescribeTestRunListRequest) SetGitReleaseState(v int32) {
	o.GitReleaseState = &v
}

// GetIsCompleted returns the IsCompleted field value if set, zero value otherwise.
func (o *DescribeTestRunListRequest) GetIsCompleted() bool {
	if o == nil || utils.IsNil(o.IsCompleted) {
		var ret bool
		return ret
	}
	return *o.IsCompleted
}

// GetIsCompletedOk returns a tuple with the IsCompleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestRunListRequest) GetIsCompletedOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsCompleted) {
		return nil, false
	}
	return o.IsCompleted, true
}

// HasIsCompleted returns a boolean if a field has been set.
func (o *DescribeTestRunListRequest) HasIsCompleted() bool {
	if o != nil && !utils.IsNil(o.IsCompleted) {
		return true
	}

	return false
}

// SetIsCompleted gets a reference to the given bool and assigns it to the IsCompleted field.
func (o *DescribeTestRunListRequest) SetIsCompleted(v bool) {
	o.IsCompleted = &v
}

// GetIterationId returns the IterationId field value if set, zero value otherwise.
func (o *DescribeTestRunListRequest) GetIterationId() []int32 {
	if o == nil || utils.IsNil(o.IterationId) {
		var ret []int32
		return ret
	}
	return o.IterationId
}

// GetIterationIdOk returns a tuple with the IterationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestRunListRequest) GetIterationIdOk() ([]int32, bool) {
	if o == nil || utils.IsNil(o.IterationId) {
		return nil, false
	}
	return o.IterationId, true
}

// HasIterationId returns a boolean if a field has been set.
func (o *DescribeTestRunListRequest) HasIterationId() bool {
	if o != nil && !utils.IsNil(o.IterationId) {
		return true
	}

	return false
}

// SetIterationId gets a reference to the given []int32 and assigns it to the IterationId field.
func (o *DescribeTestRunListRequest) SetIterationId(v []int32) {
	o.IterationId = v
}

// GetIterationStatus returns the IterationStatus field value if set, zero value otherwise.
func (o *DescribeTestRunListRequest) GetIterationStatus() []string {
	if o == nil || utils.IsNil(o.IterationStatus) {
		var ret []string
		return ret
	}
	return o.IterationStatus
}

// GetIterationStatusOk returns a tuple with the IterationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestRunListRequest) GetIterationStatusOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.IterationStatus) {
		return nil, false
	}
	return o.IterationStatus, true
}

// HasIterationStatus returns a boolean if a field has been set.
func (o *DescribeTestRunListRequest) HasIterationStatus() bool {
	if o != nil && !utils.IsNil(o.IterationStatus) {
		return true
	}

	return false
}

// SetIterationStatus gets a reference to the given []string and assigns it to the IterationStatus field.
func (o *DescribeTestRunListRequest) SetIterationStatus(v []string) {
	o.IterationStatus = v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *DescribeTestRunListRequest) GetKeyword() string {
	if o == nil || utils.IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestRunListRequest) GetKeywordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *DescribeTestRunListRequest) HasKeyword() bool {
	if o != nil && !utils.IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *DescribeTestRunListRequest) SetKeyword(v string) {
	o.Keyword = &v
}

// GetProjectName returns the ProjectName field value
func (o *DescribeTestRunListRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DescribeTestRunListRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DescribeTestRunListRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetSectionId returns the SectionId field value if set, zero value otherwise.
func (o *DescribeTestRunListRequest) GetSectionId() int64 {
	if o == nil || utils.IsNil(o.SectionId) {
		var ret int64
		return ret
	}
	return *o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestRunListRequest) GetSectionIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.SectionId) {
		return nil, false
	}
	return o.SectionId, true
}

// HasSectionId returns a boolean if a field has been set.
func (o *DescribeTestRunListRequest) HasSectionId() bool {
	if o != nil && !utils.IsNil(o.SectionId) {
		return true
	}

	return false
}

// SetSectionId gets a reference to the given int64 and assigns it to the SectionId field.
func (o *DescribeTestRunListRequest) SetSectionId(v int64) {
	o.SectionId = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *DescribeTestRunListRequest) GetState() int32 {
	if o == nil || utils.IsNil(o.State) {
		var ret int32
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTestRunListRequest) GetStateOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *DescribeTestRunListRequest) HasState() bool {
	if o != nil && !utils.IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given int32 and assigns it to the State field.
func (o *DescribeTestRunListRequest) SetState(v int32) {
	o.State = &v
}

func (o DescribeTestRunListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeTestRunListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ExecuteType) {
		toSerialize["ExecuteType"] = o.ExecuteType
	}
	if !utils.IsNil(o.GitReleaseState) {
		toSerialize["GitReleaseState"] = o.GitReleaseState
	}
	if !utils.IsNil(o.IsCompleted) {
		toSerialize["IsCompleted"] = o.IsCompleted
	}
	if !utils.IsNil(o.IterationId) {
		toSerialize["IterationId"] = o.IterationId
	}
	if !utils.IsNil(o.IterationStatus) {
		toSerialize["IterationStatus"] = o.IterationStatus
	}
	if !utils.IsNil(o.Keyword) {
		toSerialize["Keyword"] = o.Keyword
	}
	toSerialize["ProjectName"] = o.ProjectName
	if !utils.IsNil(o.SectionId) {
		toSerialize["SectionId"] = o.SectionId
	}
	if !utils.IsNil(o.State) {
		toSerialize["State"] = o.State
	}
	return toSerialize, nil
}

func (o *DescribeTestRunListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ProjectName",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDescribeTestRunListRequest := _DescribeTestRunListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeTestRunListRequest)

	if err != nil {
		return err
	}

	*o = DescribeTestRunListRequest(varDescribeTestRunListRequest)

	return err
}

type NullableDescribeTestRunListRequest struct {
	value *DescribeTestRunListRequest
	isSet bool
}

func (v NullableDescribeTestRunListRequest) Get() *DescribeTestRunListRequest {
	return v.value
}

func (v *NullableDescribeTestRunListRequest) Set(val *DescribeTestRunListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeTestRunListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeTestRunListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeTestRunListRequest(val *DescribeTestRunListRequest) *NullableDescribeTestRunListRequest {
	return &NullableDescribeTestRunListRequest{value: val, isSet: true}
}

func (v NullableDescribeTestRunListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeTestRunListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


