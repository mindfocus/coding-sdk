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

// checks if the DescribeRelatedCaseListRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeRelatedCaseListRequest{}

// DescribeRelatedCaseListRequest struct for DescribeRelatedCaseListRequest
type DescribeRelatedCaseListRequest struct {
	// 事项ID
	IssueCode int64 `json:"IssueCode"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
}

type _DescribeRelatedCaseListRequest DescribeRelatedCaseListRequest

// NewDescribeRelatedCaseListRequest instantiates a new DescribeRelatedCaseListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeRelatedCaseListRequest(issueCode int64, projectName string) *DescribeRelatedCaseListRequest {
	this := DescribeRelatedCaseListRequest{}
	this.IssueCode = issueCode
	this.ProjectName = projectName
	return &this
}

// NewDescribeRelatedCaseListRequestWithDefaults instantiates a new DescribeRelatedCaseListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeRelatedCaseListRequestWithDefaults() *DescribeRelatedCaseListRequest {
	this := DescribeRelatedCaseListRequest{}
	return &this
}

// GetIssueCode returns the IssueCode field value
func (o *DescribeRelatedCaseListRequest) GetIssueCode() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.IssueCode
}

// GetIssueCodeOk returns a tuple with the IssueCode field value
// and a boolean to check if the value has been set.
func (o *DescribeRelatedCaseListRequest) GetIssueCodeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueCode, true
}

// SetIssueCode sets field value
func (o *DescribeRelatedCaseListRequest) SetIssueCode(v int64) {
	o.IssueCode = v
}

// GetProjectName returns the ProjectName field value
func (o *DescribeRelatedCaseListRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DescribeRelatedCaseListRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DescribeRelatedCaseListRequest) SetProjectName(v string) {
	o.ProjectName = v
}

func (o DescribeRelatedCaseListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeRelatedCaseListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["IssueCode"] = o.IssueCode
	toSerialize["ProjectName"] = o.ProjectName
	return toSerialize, nil
}

func (o *DescribeRelatedCaseListRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDescribeRelatedCaseListRequest := _DescribeRelatedCaseListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeRelatedCaseListRequest)

	if err != nil {
		return err
	}

	*o = DescribeRelatedCaseListRequest(varDescribeRelatedCaseListRequest)

	return err
}

type NullableDescribeRelatedCaseListRequest struct {
	value *DescribeRelatedCaseListRequest
	isSet bool
}

func (v NullableDescribeRelatedCaseListRequest) Get() *DescribeRelatedCaseListRequest {
	return v.value
}

func (v *NullableDescribeRelatedCaseListRequest) Set(val *DescribeRelatedCaseListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeRelatedCaseListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeRelatedCaseListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeRelatedCaseListRequest(val *DescribeRelatedCaseListRequest) *NullableDescribeRelatedCaseListRequest {
	return &NullableDescribeRelatedCaseListRequest{value: val, isSet: true}
}

func (v NullableDescribeRelatedCaseListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeRelatedCaseListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

