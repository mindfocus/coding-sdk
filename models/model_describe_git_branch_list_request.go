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

// checks if the DescribeGitBranchListRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeGitBranchListRequest{}

// DescribeGitBranchListRequest struct for DescribeGitBranchListRequest
type DescribeGitBranchListRequest struct {
	// 仓库路径
	DepotPath string `json:"DepotPath"`
	// 查询的关键词
	KeyWord *string `json:"KeyWord,omitempty"`
	// 分页页码
	PageNumber int64 `json:"PageNumber"`
	// 分页页距,默认为10
	PageSize int64 `json:"PageSize"`
}

type _DescribeGitBranchListRequest DescribeGitBranchListRequest

// NewDescribeGitBranchListRequest instantiates a new DescribeGitBranchListRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeGitBranchListRequest(depotPath string, pageNumber int64, pageSize int64) *DescribeGitBranchListRequest {
	this := DescribeGitBranchListRequest{}
	this.DepotPath = depotPath
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	return &this
}

// NewDescribeGitBranchListRequestWithDefaults instantiates a new DescribeGitBranchListRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeGitBranchListRequestWithDefaults() *DescribeGitBranchListRequest {
	this := DescribeGitBranchListRequest{}
	return &this
}

// GetDepotPath returns the DepotPath field value
func (o *DescribeGitBranchListRequest) GetDepotPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DepotPath
}

// GetDepotPathOk returns a tuple with the DepotPath field value
// and a boolean to check if the value has been set.
func (o *DescribeGitBranchListRequest) GetDepotPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DepotPath, true
}

// SetDepotPath sets field value
func (o *DescribeGitBranchListRequest) SetDepotPath(v string) {
	o.DepotPath = v
}

// GetKeyWord returns the KeyWord field value if set, zero value otherwise.
func (o *DescribeGitBranchListRequest) GetKeyWord() string {
	if o == nil || utils.IsNil(o.KeyWord) {
		var ret string
		return ret
	}
	return *o.KeyWord
}

// GetKeyWordOk returns a tuple with the KeyWord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeGitBranchListRequest) GetKeyWordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.KeyWord) {
		return nil, false
	}
	return o.KeyWord, true
}

// HasKeyWord returns a boolean if a field has been set.
func (o *DescribeGitBranchListRequest) HasKeyWord() bool {
	if o != nil && !utils.IsNil(o.KeyWord) {
		return true
	}

	return false
}

// SetKeyWord gets a reference to the given string and assigns it to the KeyWord field.
func (o *DescribeGitBranchListRequest) SetKeyWord(v string) {
	o.KeyWord = &v
}

// GetPageNumber returns the PageNumber field value
func (o *DescribeGitBranchListRequest) GetPageNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *DescribeGitBranchListRequest) GetPageNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *DescribeGitBranchListRequest) SetPageNumber(v int64) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *DescribeGitBranchListRequest) GetPageSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DescribeGitBranchListRequest) GetPageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DescribeGitBranchListRequest) SetPageSize(v int64) {
	o.PageSize = v
}

func (o DescribeGitBranchListRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeGitBranchListRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["DepotPath"] = o.DepotPath
	if !utils.IsNil(o.KeyWord) {
		toSerialize["KeyWord"] = o.KeyWord
	}
	toSerialize["PageNumber"] = o.PageNumber
	toSerialize["PageSize"] = o.PageSize
	return toSerialize, nil
}

func (o *DescribeGitBranchListRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"DepotPath",
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

	varDescribeGitBranchListRequest := _DescribeGitBranchListRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeGitBranchListRequest)

	if err != nil {
		return err
	}

	*o = DescribeGitBranchListRequest(varDescribeGitBranchListRequest)

	return err
}

type NullableDescribeGitBranchListRequest struct {
	value *DescribeGitBranchListRequest
	isSet bool
}

func (v NullableDescribeGitBranchListRequest) Get() *DescribeGitBranchListRequest {
	return v.value
}

func (v *NullableDescribeGitBranchListRequest) Set(val *DescribeGitBranchListRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeGitBranchListRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeGitBranchListRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeGitBranchListRequest(val *DescribeGitBranchListRequest) *NullableDescribeGitBranchListRequest {
	return &NullableDescribeGitBranchListRequest{value: val, isSet: true}
}

func (v NullableDescribeGitBranchListRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeGitBranchListRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


