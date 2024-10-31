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

// checks if the DescribeProjectAnnouncementsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectAnnouncementsRequest{}

// DescribeProjectAnnouncementsRequest struct for DescribeProjectAnnouncementsRequest
type DescribeProjectAnnouncementsRequest struct {
	// 每页数量
	PageSize int64 `json:"PageSize"`
	// 页数
	PageNumber int64 `json:"PageNumber"`
	// 项目名
	ProjectName string `json:"ProjectName"`
}

type _DescribeProjectAnnouncementsRequest DescribeProjectAnnouncementsRequest

// NewDescribeProjectAnnouncementsRequest instantiates a new DescribeProjectAnnouncementsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectAnnouncementsRequest(pageSize int64, pageNumber int64, projectName string) *DescribeProjectAnnouncementsRequest {
	this := DescribeProjectAnnouncementsRequest{}
	this.PageSize = pageSize
	this.PageNumber = pageNumber
	this.ProjectName = projectName
	return &this
}

// NewDescribeProjectAnnouncementsRequestWithDefaults instantiates a new DescribeProjectAnnouncementsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectAnnouncementsRequestWithDefaults() *DescribeProjectAnnouncementsRequest {
	this := DescribeProjectAnnouncementsRequest{}
	var pageSize int64 = 0
	this.PageSize = pageSize
	var pageNumber int64 = 0
	this.PageNumber = pageNumber
	var projectName string = ""
	this.ProjectName = projectName
	return &this
}

// GetPageSize returns the PageSize field value
func (o *DescribeProjectAnnouncementsRequest) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DescribeProjectAnnouncementsRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DescribeProjectAnnouncementsRequest) SetPageSize(v int64) {
	o.PageSize = v
}

// GetPageNumber returns the PageNumber field value
func (o *DescribeProjectAnnouncementsRequest) GetPageNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *DescribeProjectAnnouncementsRequest) GetPageNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *DescribeProjectAnnouncementsRequest) SetPageNumber(v int64) {
	o.PageNumber = v
}

// GetProjectName returns the ProjectName field value
func (o *DescribeProjectAnnouncementsRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DescribeProjectAnnouncementsRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DescribeProjectAnnouncementsRequest) SetProjectName(v string) {
	o.ProjectName = v
}

func (o DescribeProjectAnnouncementsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectAnnouncementsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["PageSize"] = o.PageSize
	toSerialize["PageNumber"] = o.PageNumber
	toSerialize["ProjectName"] = o.ProjectName
	return toSerialize, nil
}

func (o *DescribeProjectAnnouncementsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"PageSize",
		"PageNumber",
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

	varDescribeProjectAnnouncementsRequest := _DescribeProjectAnnouncementsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeProjectAnnouncementsRequest)

	if err != nil {
		return err
	}

	*o = DescribeProjectAnnouncementsRequest(varDescribeProjectAnnouncementsRequest)

	return err
}

type NullableDescribeProjectAnnouncementsRequest struct {
	value *DescribeProjectAnnouncementsRequest
	isSet bool
}

func (v NullableDescribeProjectAnnouncementsRequest) Get() *DescribeProjectAnnouncementsRequest {
	return v.value
}

func (v *NullableDescribeProjectAnnouncementsRequest) Set(val *DescribeProjectAnnouncementsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectAnnouncementsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectAnnouncementsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectAnnouncementsRequest(val *DescribeProjectAnnouncementsRequest) *NullableDescribeProjectAnnouncementsRequest {
	return &NullableDescribeProjectAnnouncementsRequest{value: val, isSet: true}
}

func (v NullableDescribeProjectAnnouncementsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectAnnouncementsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

