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

// checks if the DescribeReleaseListRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeReleaseListRequest{}

// DescribeReleaseListRequest struct for DescribeReleaseListRequest
type DescribeReleaseListRequest struct {
	// 筛选条件列表，每一个值都是一个筛选条件，条件取值可以参考页面上的对应的HTTP接口 
	Conditions []IssueCondition `json:"Conditions,omitempty"`
	// 分页查询中的页数，page从1开始计数 
	Page *int64 `json:"Page,omitempty"`
	// 分页查询中每页的大小 
	PageSize *int64 `json:"PageSize,omitempty"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 排序KEY 
	SortKey *string `json:"SortKey,omitempty"`
	// 排序VALUE 
	SortValue *string `json:"SortValue,omitempty"`
}

type _DescribeReleaseListRequest DescribeReleaseListRequest

// NewDescribeReleaseListRequest instantiates a new DescribeReleaseListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeReleaseListRequest(projectName string) *DescribeReleaseListRequest {
	this := DescribeReleaseListRequest{}
	this.ProjectName = projectName
	return &this
}

// NewDescribeReleaseListRequestWithDefaults instantiates a new DescribeReleaseListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeReleaseListRequestWithDefaults() *DescribeReleaseListRequest {
	this := DescribeReleaseListRequest{}
	return &this
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *DescribeReleaseListRequest) GetConditions() []IssueCondition {
	if o == nil || utils.IsNil(o.Conditions) {
		var ret []IssueCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseListRequest) GetConditionsOk() ([]IssueCondition, bool) {
	if o == nil || utils.IsNil(o.Conditions) {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *DescribeReleaseListRequest) HasConditions() bool {
	if o != nil && !utils.IsNil(o.Conditions) {
		return true
	}

	return false
}

// SetConditions gets a reference to the given []IssueCondition and assigns it to the Conditions field.
func (o *DescribeReleaseListRequest) SetConditions(v []IssueCondition) {
	o.Conditions = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *DescribeReleaseListRequest) GetPage() int64 {
	if o == nil || utils.IsNil(o.Page) {
		var ret int64
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseListRequest) GetPageOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.Page) {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *DescribeReleaseListRequest) HasPage() bool {
	if o != nil && !utils.IsNil(o.Page) {
		return true
	}

	return false
}

// SetPage gets a reference to the given int64 and assigns it to the Page field.
func (o *DescribeReleaseListRequest) SetPage(v int64) {
	o.Page = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *DescribeReleaseListRequest) GetPageSize() int64 {
	if o == nil || utils.IsNil(o.PageSize) {
		var ret int64
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseListRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil || utils.IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *DescribeReleaseListRequest) HasPageSize() bool {
	if o != nil && !utils.IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int64 and assigns it to the PageSize field.
func (o *DescribeReleaseListRequest) SetPageSize(v int64) {
	o.PageSize = &v
}

// GetProjectName returns the ProjectName field value
func (o *DescribeReleaseListRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DescribeReleaseListRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DescribeReleaseListRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetSortKey returns the SortKey field value if set, zero value otherwise.
func (o *DescribeReleaseListRequest) GetSortKey() string {
	if o == nil || utils.IsNil(o.SortKey) {
		var ret string
		return ret
	}
	return *o.SortKey
}

// GetSortKeyOk returns a tuple with the SortKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseListRequest) GetSortKeyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SortKey) {
		return nil, false
	}
	return o.SortKey, true
}

// HasSortKey returns a boolean if a field has been set.
func (o *DescribeReleaseListRequest) HasSortKey() bool {
	if o != nil && !utils.IsNil(o.SortKey) {
		return true
	}

	return false
}

// SetSortKey gets a reference to the given string and assigns it to the SortKey field.
func (o *DescribeReleaseListRequest) SetSortKey(v string) {
	o.SortKey = &v
}

// GetSortValue returns the SortValue field value if set, zero value otherwise.
func (o *DescribeReleaseListRequest) GetSortValue() string {
	if o == nil || utils.IsNil(o.SortValue) {
		var ret string
		return ret
	}
	return *o.SortValue
}

// GetSortValueOk returns a tuple with the SortValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeReleaseListRequest) GetSortValueOk() (*string, bool) {
	if o == nil || utils.IsNil(o.SortValue) {
		return nil, false
	}
	return o.SortValue, true
}

// HasSortValue returns a boolean if a field has been set.
func (o *DescribeReleaseListRequest) HasSortValue() bool {
	if o != nil && !utils.IsNil(o.SortValue) {
		return true
	}

	return false
}

// SetSortValue gets a reference to the given string and assigns it to the SortValue field.
func (o *DescribeReleaseListRequest) SetSortValue(v string) {
	o.SortValue = &v
}

func (o DescribeReleaseListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeReleaseListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Conditions) {
		toSerialize["Conditions"] = o.Conditions
	}
	if !utils.IsNil(o.Page) {
		toSerialize["Page"] = o.Page
	}
	if !utils.IsNil(o.PageSize) {
		toSerialize["PageSize"] = o.PageSize
	}
	toSerialize["ProjectName"] = o.ProjectName
	if !utils.IsNil(o.SortKey) {
		toSerialize["SortKey"] = o.SortKey
	}
	if !utils.IsNil(o.SortValue) {
		toSerialize["SortValue"] = o.SortValue
	}
	return toSerialize, nil
}

func (o *DescribeReleaseListRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDescribeReleaseListRequest := _DescribeReleaseListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeReleaseListRequest)

	if err != nil {
		return err
	}

	*o = DescribeReleaseListRequest(varDescribeReleaseListRequest)

	return err
}

type NullableDescribeReleaseListRequest struct {
	value *DescribeReleaseListRequest
	isSet bool
}

func (v NullableDescribeReleaseListRequest) Get() *DescribeReleaseListRequest {
	return v.value
}

func (v *NullableDescribeReleaseListRequest) Set(val *DescribeReleaseListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeReleaseListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeReleaseListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeReleaseListRequest(val *DescribeReleaseListRequest) *NullableDescribeReleaseListRequest {
	return &NullableDescribeReleaseListRequest{value: val, isSet: true}
}

func (v NullableDescribeReleaseListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeReleaseListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


