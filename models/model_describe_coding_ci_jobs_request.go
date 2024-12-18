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

// checks if the DescribeCodingCIJobsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCodingCIJobsRequest{}

// DescribeCodingCIJobsRequest struct for DescribeCodingCIJobsRequest
type DescribeCodingCIJobsRequest struct {
	// 过滤参数
	Filter []Filter `json:"Filter,omitempty"`
	// 项目 ID
	ProjectId int32 `json:"ProjectId"`
}

type _DescribeCodingCIJobsRequest DescribeCodingCIJobsRequest

// NewDescribeCodingCIJobsRequest instantiates a new DescribeCodingCIJobsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCodingCIJobsRequest(projectId int32) *DescribeCodingCIJobsRequest {
	this := DescribeCodingCIJobsRequest{}
	this.ProjectId = projectId
	return &this
}

// NewDescribeCodingCIJobsRequestWithDefaults instantiates a new DescribeCodingCIJobsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCodingCIJobsRequestWithDefaults() *DescribeCodingCIJobsRequest {
	this := DescribeCodingCIJobsRequest{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DescribeCodingCIJobsRequest) GetFilter() []Filter {
	if o == nil || utils.IsNil(o.Filter) {
		var ret []Filter
		return ret
	}
	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIJobsRequest) GetFilterOk() ([]Filter, bool) {
	if o == nil || utils.IsNil(o.Filter) {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DescribeCodingCIJobsRequest) HasFilter() bool {
	if o != nil && !utils.IsNil(o.Filter) {
		return true
	}

	return false
}

// SetFilter gets a reference to the given []Filter and assigns it to the Filter field.
func (o *DescribeCodingCIJobsRequest) SetFilter(v []Filter) {
	o.Filter = v
}

// GetProjectId returns the ProjectId field value
func (o *DescribeCodingCIJobsRequest) GetProjectId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DescribeCodingCIJobsRequest) GetProjectIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DescribeCodingCIJobsRequest) SetProjectId(v int32) {
	o.ProjectId = v
}

func (o DescribeCodingCIJobsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCodingCIJobsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Filter) {
		toSerialize["Filter"] = o.Filter
	}
	toSerialize["ProjectId"] = o.ProjectId
	return toSerialize, nil
}

func (o *DescribeCodingCIJobsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varDescribeCodingCIJobsRequest := _DescribeCodingCIJobsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeCodingCIJobsRequest)

	if err != nil {
		return err
	}

	*o = DescribeCodingCIJobsRequest(varDescribeCodingCIJobsRequest)

	return err
}

type NullableDescribeCodingCIJobsRequest struct {
	value *DescribeCodingCIJobsRequest
	isSet bool
}

func (v NullableDescribeCodingCIJobsRequest) Get() *DescribeCodingCIJobsRequest {
	return v.value
}

func (v *NullableDescribeCodingCIJobsRequest) Set(val *DescribeCodingCIJobsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCodingCIJobsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCodingCIJobsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCodingCIJobsRequest(val *DescribeCodingCIJobsRequest) *NullableDescribeCodingCIJobsRequest {
	return &NullableDescribeCodingCIJobsRequest{value: val, isSet: true}
}

func (v NullableDescribeCodingCIJobsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCodingCIJobsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


