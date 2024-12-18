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

// checks if the DescribeCdDeployTrendByProjectRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DescribeCdDeployTrendByProjectRequest{}

// DescribeCdDeployTrendByProjectRequest struct for DescribeCdDeployTrendByProjectRequest
type DescribeCdDeployTrendByProjectRequest struct {
	// 结束时间（格式：yyyy-MM-dd HH:mm:ss）
	EndAt string `json:"EndAt"`
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 开始时间（格式：yyyy-MM-dd HH:mm:ss）
	StartAt string `json:"StartAt"`
}

type _DescribeCdDeployTrendByProjectRequest DescribeCdDeployTrendByProjectRequest

// NewDescribeCdDeployTrendByProjectRequest instantiates a new DescribeCdDeployTrendByProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDescribeCdDeployTrendByProjectRequest(endAt string, projectName string, startAt string) *DescribeCdDeployTrendByProjectRequest {
	this := DescribeCdDeployTrendByProjectRequest{}
	this.EndAt = endAt
	this.ProjectName = projectName
	this.StartAt = startAt
	return &this
}

// NewDescribeCdDeployTrendByProjectRequestWithDefaults instantiates a new DescribeCdDeployTrendByProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDescribeCdDeployTrendByProjectRequestWithDefaults() *DescribeCdDeployTrendByProjectRequest {
	this := DescribeCdDeployTrendByProjectRequest{}
	return &this
}

// GetEndAt returns the EndAt field value
func (o *DescribeCdDeployTrendByProjectRequest) GetEndAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value
// and a boolean to check if the value has been set.
func (o *DescribeCdDeployTrendByProjectRequest) GetEndAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndAt, true
}

// SetEndAt sets field value
func (o *DescribeCdDeployTrendByProjectRequest) SetEndAt(v string) {
	o.EndAt = v
}

// GetProjectName returns the ProjectName field value
func (o *DescribeCdDeployTrendByProjectRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DescribeCdDeployTrendByProjectRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DescribeCdDeployTrendByProjectRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetStartAt returns the StartAt field value
func (o *DescribeCdDeployTrendByProjectRequest) GetStartAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value
// and a boolean to check if the value has been set.
func (o *DescribeCdDeployTrendByProjectRequest) GetStartAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartAt, true
}

// SetStartAt sets field value
func (o *DescribeCdDeployTrendByProjectRequest) SetStartAt(v string) {
	o.StartAt = v
}

func (o DescribeCdDeployTrendByProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DescribeCdDeployTrendByProjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["EndAt"] = o.EndAt
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["StartAt"] = o.StartAt
	return toSerialize, nil
}

func (o *DescribeCdDeployTrendByProjectRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"EndAt",
		"ProjectName",
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

	varDescribeCdDeployTrendByProjectRequest := _DescribeCdDeployTrendByProjectRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDescribeCdDeployTrendByProjectRequest)

	if err != nil {
		return err
	}

	*o = DescribeCdDeployTrendByProjectRequest(varDescribeCdDeployTrendByProjectRequest)

	return err
}

type NullableDescribeCdDeployTrendByProjectRequest struct {
	value *DescribeCdDeployTrendByProjectRequest
	isSet bool
}

func (v NullableDescribeCdDeployTrendByProjectRequest) Get() *DescribeCdDeployTrendByProjectRequest {
	return v.value
}

func (v *NullableDescribeCdDeployTrendByProjectRequest) Set(val *DescribeCdDeployTrendByProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDescribeCdDeployTrendByProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDescribeCdDeployTrendByProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDescribeCdDeployTrendByProjectRequest(val *DescribeCdDeployTrendByProjectRequest) *NullableDescribeCdDeployTrendByProjectRequest {
	return &NullableDescribeCdDeployTrendByProjectRequest{value: val, isSet: true}
}

func (v NullableDescribeCdDeployTrendByProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDescribeCdDeployTrendByProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


