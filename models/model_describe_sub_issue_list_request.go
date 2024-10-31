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

// checks if the DescribeSubIssueListRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeSubIssueListRequest{}

// DescribeSubIssueListRequest struct for DescribeSubIssueListRequest
type DescribeSubIssueListRequest struct {
	// 事项 Code
	IssueCode int64 `json:"IssueCode"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
}

type _DescribeSubIssueListRequest DescribeSubIssueListRequest

// NewDescribeSubIssueListRequest instantiates a new DescribeSubIssueListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeSubIssueListRequest(issueCode int64, projectName string) *DescribeSubIssueListRequest {
	this := DescribeSubIssueListRequest{}
	this.IssueCode = issueCode
	this.ProjectName = projectName
	return &this
}

// NewDescribeSubIssueListRequestWithDefaults instantiates a new DescribeSubIssueListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeSubIssueListRequestWithDefaults() *DescribeSubIssueListRequest {
	this := DescribeSubIssueListRequest{}
	return &this
}

// GetIssueCode returns the IssueCode field value
func (o *DescribeSubIssueListRequest) GetIssueCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IssueCode
}

// GetIssueCodeOk returns a tuple with the IssueCode field value
// and a boolean to check if the value has been set.
func (o *DescribeSubIssueListRequest) GetIssueCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueCode, true
}

// SetIssueCode sets field value
func (o *DescribeSubIssueListRequest) SetIssueCode(v int64) {
	o.IssueCode = v
}

// GetProjectName returns the ProjectName field value
func (o *DescribeSubIssueListRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DescribeSubIssueListRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DescribeSubIssueListRequest) SetProjectName(v string) {
	o.ProjectName = v
}

func (o DescribeSubIssueListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeSubIssueListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["IssueCode"] = o.IssueCode
	toSerialize["ProjectName"] = o.ProjectName
	return toSerialize, nil
}

func (o *DescribeSubIssueListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"IssueCode",
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

	varDescribeSubIssueListRequest := _DescribeSubIssueListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeSubIssueListRequest)

	if err != nil {
		return err
	}

	*o = DescribeSubIssueListRequest(varDescribeSubIssueListRequest)

	return err
}

type NullableDescribeSubIssueListRequest struct {
	value *DescribeSubIssueListRequest
	isSet bool
}

func (v NullableDescribeSubIssueListRequest) Get() *DescribeSubIssueListRequest {
	return v.value
}

func (v *NullableDescribeSubIssueListRequest) Set(val *DescribeSubIssueListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeSubIssueListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeSubIssueListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeSubIssueListRequest(val *DescribeSubIssueListRequest) *NullableDescribeSubIssueListRequest {
	return &NullableDescribeSubIssueListRequest{value: val, isSet: true}
}

func (v NullableDescribeSubIssueListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeSubIssueListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

