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

// checks if the DescribeProgramsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeProgramsRequest{}

// DescribeProgramsRequest struct for DescribeProgramsRequest
type DescribeProgramsRequest struct {
	// 关键字搜索：项目名
	Keyword *string `json:"Keyword,omitempty"`
	// 请求页数
	PageNumber int32 `json:"PageNumber"`
	// 请求条数
	PageSize int32 `json:"PageSize"`
}

type _DescribeProgramsRequest DescribeProgramsRequest

// NewDescribeProgramsRequest instantiates a new DescribeProgramsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeProgramsRequest(pageNumber int32, pageSize int32) *DescribeProgramsRequest {
	this := DescribeProgramsRequest{}
	this.PageNumber = pageNumber
	this.PageSize = pageSize
	return &this
}

// NewDescribeProgramsRequestWithDefaults instantiates a new DescribeProgramsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeProgramsRequestWithDefaults() *DescribeProgramsRequest {
	this := DescribeProgramsRequest{}
	return &this
}

// GetKeyword returns the Keyword field value if set, zero value otherwise.
func (o *DescribeProgramsRequest) GetKeyword() string {
	if o == nil || utils.IsNil(o.Keyword) {
		var ret string
		return ret
	}
	return *o.Keyword
}

// GetKeywordOk returns a tuple with the Keyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeProgramsRequest) GetKeywordOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Keyword) {
		return nil, false
	}
	return o.Keyword, true
}

// HasKeyword returns a boolean if a field has been set.
func (o *DescribeProgramsRequest) HasKeyword() bool {
	if o != nil && !utils.IsNil(o.Keyword) {
		return true
	}

	return false
}

// SetKeyword gets a reference to the given string and assigns it to the Keyword field.
func (o *DescribeProgramsRequest) SetKeyword(v string) {
	o.Keyword = &v
}

// GetPageNumber returns the PageNumber field value
func (o *DescribeProgramsRequest) GetPageNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageNumber
}

// GetPageNumberOk returns a tuple with the PageNumber field value
// and a boolean to check if the value has been set.
func (o *DescribeProgramsRequest) GetPageNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageNumber, true
}

// SetPageNumber sets field value
func (o *DescribeProgramsRequest) SetPageNumber(v int32) {
	o.PageNumber = v
}

// GetPageSize returns the PageSize field value
func (o *DescribeProgramsRequest) GetPageSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value
// and a boolean to check if the value has been set.
func (o *DescribeProgramsRequest) GetPageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PageSize, true
}

// SetPageSize sets field value
func (o *DescribeProgramsRequest) SetPageSize(v int32) {
	o.PageSize = v
}

func (o DescribeProgramsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeProgramsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Keyword) {
		toSerialize["Keyword"] = o.Keyword
	}
	toSerialize["PageNumber"] = o.PageNumber
	toSerialize["PageSize"] = o.PageSize
	return toSerialize, nil
}

func (o *DescribeProgramsRequest) UnmarshalJSON(data []byte) (err error) {
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

	varDescribeProgramsRequest := _DescribeProgramsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeProgramsRequest)

	if err != nil {
		return err
	}

	*o = DescribeProgramsRequest(varDescribeProgramsRequest)

	return err
}

type NullableDescribeProgramsRequest struct {
	value *DescribeProgramsRequest
	isSet bool
}

func (v NullableDescribeProgramsRequest) Get() *DescribeProgramsRequest {
	return v.value
}

func (v *NullableDescribeProgramsRequest) Set(val *DescribeProgramsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeProgramsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeProgramsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeProgramsRequest(val *DescribeProgramsRequest) *NullableDescribeProgramsRequest {
	return &NullableDescribeProgramsRequest{value: val, isSet: true}
}

func (v NullableDescribeProgramsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeProgramsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


