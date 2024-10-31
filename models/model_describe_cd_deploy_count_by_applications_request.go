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

// checks if the DescribeCdDeployCountByApplicationsRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdDeployCountByApplicationsRequest{}

// DescribeCdDeployCountByApplicationsRequest struct for DescribeCdDeployCountByApplicationsRequest
type DescribeCdDeployCountByApplicationsRequest struct {
	// 应用名列表
	Application []string `json:"Application"`
	// 结束时间（格式：yyyy-MM-dd HH:mm:ss）
	EndAt string `json:"EndAt"`
	// 开始时间（格式：yyyy-MM-dd HH:mm:ss）
	StartAt string `json:"StartAt"`
}

type _DescribeCdDeployCountByApplicationsRequest DescribeCdDeployCountByApplicationsRequest

// NewDescribeCdDeployCountByApplicationsRequest instantiates a new DescribeCdDeployCountByApplicationsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdDeployCountByApplicationsRequest(application []string, endAt string, startAt string) *DescribeCdDeployCountByApplicationsRequest {
	this := DescribeCdDeployCountByApplicationsRequest{}
	this.Application = application
	this.EndAt = endAt
	this.StartAt = startAt
	return &this
}

// NewDescribeCdDeployCountByApplicationsRequestWithDefaults instantiates a new DescribeCdDeployCountByApplicationsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdDeployCountByApplicationsRequestWithDefaults() *DescribeCdDeployCountByApplicationsRequest {
	this := DescribeCdDeployCountByApplicationsRequest{}
	return &this
}

// GetApplication returns the Application field value
func (o *DescribeCdDeployCountByApplicationsRequest) GetApplication() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *DescribeCdDeployCountByApplicationsRequest) GetApplicationOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Application, true
}

// SetApplication sets field value
func (o *DescribeCdDeployCountByApplicationsRequest) SetApplication(v []string) {
	o.Application = v
}

// GetEndAt returns the EndAt field value
func (o *DescribeCdDeployCountByApplicationsRequest) GetEndAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value
// and a boolean to check if the value has been set.
func (o *DescribeCdDeployCountByApplicationsRequest) GetEndAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndAt, true
}

// SetEndAt sets field value
func (o *DescribeCdDeployCountByApplicationsRequest) SetEndAt(v string) {
	o.EndAt = v
}

// GetStartAt returns the StartAt field value
func (o *DescribeCdDeployCountByApplicationsRequest) GetStartAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
func (o *DescribeCdDeployCountByApplicationsRequest) GetStartAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartAt, true
}

// SetStartAt sets field value
func (o *DescribeCdDeployCountByApplicationsRequest) SetStartAt(v string) {
	o.StartAt = v
}

func (o DescribeCdDeployCountByApplicationsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdDeployCountByApplicationsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Application"] = o.Application
	toSerialize["EndAt"] = o.EndAt
	toSerialize["StartAt"] = o.StartAt
	return toSerialize, nil
}

func (o *DescribeCdDeployCountByApplicationsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"Application",
		"EndAt",
		"StartAt",
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

	varDescribeCdDeployCountByApplicationsRequest := _DescribeCdDeployCountByApplicationsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeCdDeployCountByApplicationsRequest)

	if err != nil {
		return err
	}

	*o = DescribeCdDeployCountByApplicationsRequest(varDescribeCdDeployCountByApplicationsRequest)

	return err
}

type NullableDescribeCdDeployCountByApplicationsRequest struct {
	value *DescribeCdDeployCountByApplicationsRequest
	isSet bool
}

func (v NullableDescribeCdDeployCountByApplicationsRequest) Get() *DescribeCdDeployCountByApplicationsRequest {
	return v.value
}

func (v *NullableDescribeCdDeployCountByApplicationsRequest) Set(val *DescribeCdDeployCountByApplicationsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdDeployCountByApplicationsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdDeployCountByApplicationsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdDeployCountByApplicationsRequest(val *DescribeCdDeployCountByApplicationsRequest) *NullableDescribeCdDeployCountByApplicationsRequest {
	return &NullableDescribeCdDeployCountByApplicationsRequest{value: val, isSet: true}
}

func (v NullableDescribeCdDeployCountByApplicationsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdDeployCountByApplicationsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


