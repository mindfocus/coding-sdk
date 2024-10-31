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

// checks if the DescribeProjectMemberPrincipalsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProjectMemberPrincipalsRequest{}

// DescribeProjectMemberPrincipalsRequest struct for DescribeProjectMemberPrincipalsRequest
type DescribeProjectMemberPrincipalsRequest struct {
	// 关键字搜索：成员名/邮箱
	Keyword *string `json:"Keyword,omitempty"`
	// 请求页数
	PageNumber int32 `json:"PageNumber"`
	// 请求条数
	PageSize int32 `json:"PageSize"`
	// 权限组Id
	PolicyId *int64 `json:"PolicyId,omitempty"`
	// 项目Id
	ProjectId int32 `json:"ProjectId"`
}

type _DescribeProjectMemberPrincipalsRequest DescribeProjectMemberPrincipalsRequest

// NewDescribeProjectMemberPrincipalsRequest instantiates a new DescribeProjectMemberPrincipalsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProjectMemberPrincipalsRequest(pageNumber int32, pageSize int32, projectId int32) *DescribeProjectMemberPrincipalsRequest {
	this := DescribeProjectMemberPrincipalsRequest{}
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	this.ProjectId = projectId
	return &this
}

// NewDescribeProjectMemberPrincipalsRequestWithDefaults instantiates a new DescribeProjectMemberPrincipalsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProjectMemberPrincipalsRequestWithDefaults() *DescribeProjectMemberPrincipalsRequest {
	this := DescribeProjectMemberPrincipalsRequest{}
	return &this
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *DescribeProjectMemberPrincipalsRequest) GetKeyword() string {
	if o == nil || utils.IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectMemberPrincipalsRequest) GetKeywordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *DescribeProjectMemberPrincipalsRequest) HasKeyword() bool {
	if o != nil && !utils.IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *DescribeProjectMemberPrincipalsRequest) SetKeyword(v string) {
	o.Keyword = &v
}

// GetPageNumber returns the PageNumber field value
func (o *DescribeProjectMemberPrincipalsRequest) GetPageNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *DescribeProjectMemberPrincipalsRequest) GetPageNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *DescribeProjectMemberPrincipalsRequest) SetPageNumber(v int32) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *DescribeProjectMemberPrincipalsRequest) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DescribeProjectMemberPrincipalsRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DescribeProjectMemberPrincipalsRequest) SetPageSize(v int32) {
	o.PageSize = v
}

// GetPolicyId returns the PolicyId field value if set, zero value otherwise.
func (o *DescribeProjectMemberPrincipalsRequest) GetPolicyId() int64 {
	if o == nil || utils.IsNil(o.PolicyId) {
		var ret int64
		return ret
	}
	return *o.PolicyId
}

// GetPolicyIdOk returns a tuple with the PolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProjectMemberPrincipalsRequest) GetPolicyIdOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PolicyId) {
		return nil, false
	}
	return o.PolicyId, true
}

// HasPolicyId returns a boolean if a field has been set.
func (o *DescribeProjectMemberPrincipalsRequest) HasPolicyId() bool {
	if o != nil && !utils.IsNil(o.PolicyId) {
		return true
	}

	return false
}

// SetPolicyId gets a reference to the given int64 and assigns it to the PolicyId field.
func (o *DescribeProjectMemberPrincipalsRequest) SetPolicyId(v int64) {
	o.PolicyId = &v
}

// GetProjectId returns the ProjectId field value
func (o *DescribeProjectMemberPrincipalsRequest) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeProjectMemberPrincipalsRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeProjectMemberPrincipalsRequest) SetProjectId(v int32) {
	o.ProjectId = v
}

func (o DescribeProjectMemberPrincipalsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProjectMemberPrincipalsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Keyword) {
		toSerialize["Keyword"] = o.Keyword
	}
	toSerialize["PageNumber"] = o.PageNumber
	toSerialize["PageSize"] = o.PageSize
	if !utils.IsNil(o.PolicyId) {
		toSerialize["PolicyId"] = o.PolicyId
	}
	toSerialize["ProjectId"] = o.ProjectId
	return toSerialize, nil
}

func (o *DescribeProjectMemberPrincipalsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"PageNumber",
		"PageSize",
		"ProjectId",
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

	varDescribeProjectMemberPrincipalsRequest := _DescribeProjectMemberPrincipalsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeProjectMemberPrincipalsRequest)

	if err != nil {
		return err
	}

	*o = DescribeProjectMemberPrincipalsRequest(varDescribeProjectMemberPrincipalsRequest)

	return err
}

type NullableDescribeProjectMemberPrincipalsRequest struct {
	value *DescribeProjectMemberPrincipalsRequest
	isSet bool
}

func (v NullableDescribeProjectMemberPrincipalsRequest) Get() *DescribeProjectMemberPrincipalsRequest {
	return v.value
}

func (v *NullableDescribeProjectMemberPrincipalsRequest) Set(val *DescribeProjectMemberPrincipalsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProjectMemberPrincipalsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProjectMemberPrincipalsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProjectMemberPrincipalsRequest(val *DescribeProjectMemberPrincipalsRequest) *NullableDescribeProjectMemberPrincipalsRequest {
	return &NullableDescribeProjectMemberPrincipalsRequest{value: val, isSet: true}
}

func (v NullableDescribeProjectMemberPrincipalsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProjectMemberPrincipalsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


