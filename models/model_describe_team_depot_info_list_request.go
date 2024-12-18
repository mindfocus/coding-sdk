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

// checks if the DescribeTeamDepotInfoListRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeTeamDepotInfoListRequest{}

// DescribeTeamDepotInfoListRequest struct for DescribeTeamDepotInfoListRequest
type DescribeTeamDepotInfoListRequest struct {
	// 仓库名
	DepotName *string `json:"DepotName,omitempty"`
	// 页码
	PageNumber int64 `json:"PageNumber"`
	// 页码大小
	PageSize int64 `json:"PageSize"`
	// 项目名
	ProjectName *string `json:"ProjectName,omitempty"`
}

type _DescribeTeamDepotInfoListRequest DescribeTeamDepotInfoListRequest

// NewDescribeTeamDepotInfoListRequest instantiates a new DescribeTeamDepotInfoListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeTeamDepotInfoListRequest(pageNumber int64, pageSize int64) *DescribeTeamDepotInfoListRequest {
	this := DescribeTeamDepotInfoListRequest{}
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	return &this
}

// NewDescribeTeamDepotInfoListRequestWithDefaults instantiates a new DescribeTeamDepotInfoListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeTeamDepotInfoListRequestWithDefaults() *DescribeTeamDepotInfoListRequest {
	this := DescribeTeamDepotInfoListRequest{}
	return &this
}

// GetDepotName returns the DepotName field value if set, zero value otherwise.
func (o *DescribeTeamDepotInfoListRequest) GetDepotName() string {
	if o == nil || utils.IsNil(o.DepotName) {
		var ret string
		return ret
	}
	return *o.DepotName
}

// GetDepotNameOk returns a tuple with the DepotName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTeamDepotInfoListRequest) GetDepotNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.DepotName) {
		return nil, false
	}
	return o.DepotName, true
}

// HasDepotName returns a boolean if a field has been set.
func (o *DescribeTeamDepotInfoListRequest) HasDepotName() bool {
	if o != nil && !utils.IsNil(o.DepotName) {
		return true
	}

	return false
}

// SetDepotName gets a reference to the given string and assigns it to the DepotName field.
func (o *DescribeTeamDepotInfoListRequest) SetDepotName(v string) {
	o.DepotName = &v
}

// GetPageNumber returns the PageNumber field value
func (o *DescribeTeamDepotInfoListRequest) GetPageNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *DescribeTeamDepotInfoListRequest) GetPageNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *DescribeTeamDepotInfoListRequest) SetPageNumber(v int64) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *DescribeTeamDepotInfoListRequest) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DescribeTeamDepotInfoListRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DescribeTeamDepotInfoListRequest) SetPageSize(v int64) {
	o.PageSize = v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *DescribeTeamDepotInfoListRequest) GetProjectName() string {
	if o == nil || utils.IsNil(o.ProjectName) {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeTeamDepotInfoListRequest) GetProjectNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ProjectName) {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *DescribeTeamDepotInfoListRequest) HasProjectName() bool {
	if o != nil && !utils.IsNil(o.ProjectName) {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *DescribeTeamDepotInfoListRequest) SetProjectName(v string) {
	o.ProjectName = &v
}

func (o DescribeTeamDepotInfoListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeTeamDepotInfoListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.DepotName) {
		toSerialize["DepotName"] = o.DepotName
	}
	toSerialize["PageNumber"] = o.PageNumber
	toSerialize["PageSize"] = o.PageSize
	if !utils.IsNil(o.ProjectName) {
		toSerialize["ProjectName"] = o.ProjectName
	}
	return toSerialize, nil
}

func (o *DescribeTeamDepotInfoListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"PageNumber",
		"PageSize",
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

	varDescribeTeamDepotInfoListRequest := _DescribeTeamDepotInfoListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeTeamDepotInfoListRequest)

	if err != nil {
		return err
	}

	*o = DescribeTeamDepotInfoListRequest(varDescribeTeamDepotInfoListRequest)

	return err
}

type NullableDescribeTeamDepotInfoListRequest struct {
	value *DescribeTeamDepotInfoListRequest
	isSet bool
}

func (v NullableDescribeTeamDepotInfoListRequest) Get() *DescribeTeamDepotInfoListRequest {
	return v.value
}

func (v *NullableDescribeTeamDepotInfoListRequest) Set(val *DescribeTeamDepotInfoListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeTeamDepotInfoListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeTeamDepotInfoListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeTeamDepotInfoListRequest(val *DescribeTeamDepotInfoListRequest) *NullableDescribeTeamDepotInfoListRequest {
	return &NullableDescribeTeamDepotInfoListRequest{value: val, isSet: true}
}

func (v NullableDescribeTeamDepotInfoListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeTeamDepotInfoListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


