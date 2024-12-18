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

// checks if the DescribeConfigTemplateListRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeConfigTemplateListRequest{}

// DescribeConfigTemplateListRequest struct for DescribeConfigTemplateListRequest
type DescribeConfigTemplateListRequest struct {
	// 配置方案协作类型，包括 SCRUM 和 CLASSIC
	CooperateMode *string `json:"CooperateMode,omitempty"`
	// 关键字
	Keyword *string `json:"Keyword,omitempty"`
	// 页码
	PageNumber int64 `json:"PageNumber"`
	// 分页大小
	PageSize int64 `json:"PageSize"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 配置方案类型，全局配置方案取值 GLOBAL ，项目配置方案取值 PROJECT，不填默认为 GLOBAL
	TemplateType *string `json:"TemplateType,omitempty"`
}

type _DescribeConfigTemplateListRequest DescribeConfigTemplateListRequest

// NewDescribeConfigTemplateListRequest instantiates a new DescribeConfigTemplateListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeConfigTemplateListRequest(pageNumber int64, pageSize int64, projectName string) *DescribeConfigTemplateListRequest {
	this := DescribeConfigTemplateListRequest{}
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	this.ProjectName = projectName
	return &this
}

// NewDescribeConfigTemplateListRequestWithDefaults instantiates a new DescribeConfigTemplateListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeConfigTemplateListRequestWithDefaults() *DescribeConfigTemplateListRequest {
	this := DescribeConfigTemplateListRequest{}
	return &this
}

// GetCooperateMode returns the CooperateMode field value if set, zero value otherwise.
func (o *DescribeConfigTemplateListRequest) GetCooperateMode() string {
	if o == nil || utils.IsNil(o.CooperateMode) {
		var ret string
		return ret
	}
	return *o.CooperateMode
}

// GetCooperateModeOk returns a tuple with the CooperateMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeConfigTemplateListRequest) GetCooperateModeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.CooperateMode) {
		return nil, false
	}
	return o.CooperateMode, true
}

// HasCooperateMode returns a boolean if a field has been set.
func (o *DescribeConfigTemplateListRequest) HasCooperateMode() bool {
	if o != nil && !utils.IsNil(o.CooperateMode) {
		return true
	}

	return false
}

// SetCooperateMode gets a reference to the given string and assigns it to the CooperateMode field.
func (o *DescribeConfigTemplateListRequest) SetCooperateMode(v string) {
	o.CooperateMode = &v
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *DescribeConfigTemplateListRequest) GetKeyword() string {
	if o == nil || utils.IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeConfigTemplateListRequest) GetKeywordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *DescribeConfigTemplateListRequest) HasKeyword() bool {
	if o != nil && !utils.IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *DescribeConfigTemplateListRequest) SetKeyword(v string) {
	o.Keyword = &v
}

// GetPageNumber returns the PageNumber field value
func (o *DescribeConfigTemplateListRequest) GetPageNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *DescribeConfigTemplateListRequest) GetPageNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *DescribeConfigTemplateListRequest) SetPageNumber(v int64) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *DescribeConfigTemplateListRequest) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DescribeConfigTemplateListRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DescribeConfigTemplateListRequest) SetPageSize(v int64) {
	o.PageSize = v
}

// GetProjectName returns the ProjectName field value
func (o *DescribeConfigTemplateListRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DescribeConfigTemplateListRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DescribeConfigTemplateListRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetTemplateType returns the TemplateType field value if set, zero value otherwise.
func (o *DescribeConfigTemplateListRequest) GetTemplateType() string {
	if o == nil || utils.IsNil(o.TemplateType) {
		var ret string
		return ret
	}
	return *o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeConfigTemplateListRequest) GetTemplateTypeOk() (*string, bool) {
	if o == nil || utils.IsNil(o.TemplateType) {
		return nil, false
	}
	return o.TemplateType, true
}

// HasTemplateType returns a boolean if a field has been set.
func (o *DescribeConfigTemplateListRequest) HasTemplateType() bool {
	if o != nil && !utils.IsNil(o.TemplateType) {
		return true
	}

	return false
}

// SetTemplateType gets a reference to the given string and assigns it to the TemplateType field.
func (o *DescribeConfigTemplateListRequest) SetTemplateType(v string) {
	o.TemplateType = &v
}

func (o DescribeConfigTemplateListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeConfigTemplateListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.CooperateMode) {
		toSerialize["CooperateMode"] = o.CooperateMode
	}
	if !utils.IsNil(o.Keyword) {
		toSerialize["Keyword"] = o.Keyword
	}
	toSerialize["PageNumber"] = o.PageNumber
	toSerialize["PageSize"] = o.PageSize
	toSerialize["ProjectName"] = o.ProjectName
	if !utils.IsNil(o.TemplateType) {
		toSerialize["TemplateType"] = o.TemplateType
	}
	return toSerialize, nil
}

func (o *DescribeConfigTemplateListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"PageNumber",
		"PageSize",
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

	varDescribeConfigTemplateListRequest := _DescribeConfigTemplateListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeConfigTemplateListRequest)

	if err != nil {
		return err
	}

	*o = DescribeConfigTemplateListRequest(varDescribeConfigTemplateListRequest)

	return err
}

type NullableDescribeConfigTemplateListRequest struct {
	value *DescribeConfigTemplateListRequest
	isSet bool
}

func (v NullableDescribeConfigTemplateListRequest) Get() *DescribeConfigTemplateListRequest {
	return v.value
}

func (v *NullableDescribeConfigTemplateListRequest) Set(val *DescribeConfigTemplateListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeConfigTemplateListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeConfigTemplateListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeConfigTemplateListRequest(val *DescribeConfigTemplateListRequest) *NullableDescribeConfigTemplateListRequest {
	return &NullableDescribeConfigTemplateListRequest{value: val, isSet: true}
}

func (v NullableDescribeConfigTemplateListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeConfigTemplateListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


