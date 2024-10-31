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

// checks if the DescribeReportListRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeReportListRequest{}

// DescribeReportListRequest struct for DescribeReportListRequest
type DescribeReportListRequest struct {
	// 创建时间
	EndAt *string `json:"EndAt,omitempty"`
	// 报告名称关键词
	Keyword *string `json:"Keyword,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 创建时间
	StartAt *string `json:"StartAt,omitempty"`
	// 报告状态：CREATING 创建中，AVAILABLE 可用，UNAVAILABLE 不可用
	Status *string `json:"Status,omitempty"`
}

type _DescribeReportListRequest DescribeReportListRequest

// NewDescribeReportListRequest instantiates a new DescribeReportListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeReportListRequest(projectName string) *DescribeReportListRequest {
	this := DescribeReportListRequest{}
	this.ProjectName = projectName
	return &this
}

// NewDescribeReportListRequestWithDefaults instantiates a new DescribeReportListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeReportListRequestWithDefaults() *DescribeReportListRequest {
	this := DescribeReportListRequest{}
	return &this
}

// GetEndAt returns the EndAt field value if set, zero value otherwise.
func (o *DescribeReportListRequest) GetEndAt() string {
	if o == nil || utils.IsNil(o.EndAt) {
		var ret string
		return ret
	}
	return *o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReportListRequest) GetEndAtOk() (*string, bool) {
	if o == nil || utils.IsNil(o.EndAt) {
		return nil, false
	}
	return o.EndAt, true
}

// HasEndAt returns a boolean if a field has been set.
func (o *DescribeReportListRequest) HasEndAt() bool {
	if o != nil && !utils.IsNil(o.EndAt) {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given string and assigns it to the EndAt field.
func (o *DescribeReportListRequest) SetEndAt(v string) {
	o.EndAt = &v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *DescribeReportListRequest) GetKeyword() string {
	if o == nil || utils.IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReportListRequest) GetKeywordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *DescribeReportListRequest) HasKeyword() bool {
	if o != nil && !utils.IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *DescribeReportListRequest) SetKeyword(v string) {
	o.Keyword = &v
}

// GetProjectName returns the ProjectName field value
func (o *DescribeReportListRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DescribeReportListRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DescribeReportListRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetStartAt returns the StartAt field value if set, zero value otherwise.
func (o *DescribeReportListRequest) GetStartAt() string {
	if o == nil || utils.IsNil(o.StartAt) {
		var ret string
		return ret
	}
	return *o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReportListRequest) GetStartAtOk() (*string, bool) {
	if o == nil || utils.IsNil(o.StartAt) {
		return nil, false
	}
	return o.StartAt, true
}

// HasStartAt returns a boolean if a field has been set.
func (o *DescribeReportListRequest) HasStartAt() bool {
	if o != nil && !utils.IsNil(o.StartAt) {
		return true
	}

	return false
}

// SetStartAt gets a reference to the given string and assigns it to the StartAt field.
func (o *DescribeReportListRequest) SetStartAt(v string) {
	o.StartAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DescribeReportListRequest) GetStatus() string {
	if o == nil || utils.IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReportListRequest) GetStatusOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DescribeReportListRequest) HasStatus() bool {
	if o != nil && !utils.IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DescribeReportListRequest) SetStatus(v string) {
	o.Status = &v
}

func (o DescribeReportListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeReportListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.EndAt) {
		toSerialize["EndAt"] = o.EndAt
	}
	if !utils.IsNil(o.Keyword) {
		toSerialize["Keyword"] = o.Keyword
	}
	toSerialize["ProjectName"] = o.ProjectName
	if !utils.IsNil(o.StartAt) {
		toSerialize["StartAt"] = o.StartAt
	}
	if !utils.IsNil(o.Status) {
		toSerialize["Status"] = o.Status
	}
	return toSerialize, nil
}

func (o *DescribeReportListRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDescribeReportListRequest := _DescribeReportListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeReportListRequest)

	if err != nil {
		return err
	}

	*o = DescribeReportListRequest(varDescribeReportListRequest)

	return err
}

type NullableDescribeReportListRequest struct {
	value *DescribeReportListRequest
	isSet bool
}

func (v NullableDescribeReportListRequest) Get() *DescribeReportListRequest {
	return v.value
}

func (v *NullableDescribeReportListRequest) Set(val *DescribeReportListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeReportListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeReportListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeReportListRequest(val *DescribeReportListRequest) *NullableDescribeReportListRequest {
	return &NullableDescribeReportListRequest{value: val, isSet: true}
}

func (v NullableDescribeReportListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeReportListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


