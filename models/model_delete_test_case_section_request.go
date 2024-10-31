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

// checks if the DeleteTestCaseSectionRequest type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeleteTestCaseSectionRequest{}

// DeleteTestCaseSectionRequest struct for DeleteTestCaseSectionRequest
type DeleteTestCaseSectionRequest struct {
	// 项目名称
	ProjectName string `json:"ProjectName"`
	// 分组 ID
	SectionId int32 `json:"SectionId"`
}

type _DeleteTestCaseSectionRequest DeleteTestCaseSectionRequest

// NewDeleteTestCaseSectionRequest instantiates a new DeleteTestCaseSectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteTestCaseSectionRequest(projectName string, sectionId int32) *DeleteTestCaseSectionRequest {
	this := DeleteTestCaseSectionRequest{}
	this.ProjectName = projectName
	this.SectionId = sectionId
	return &this
}

// NewDeleteTestCaseSectionRequestWithDefaults instantiates a new DeleteTestCaseSectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteTestCaseSectionRequestWithDefaults() *DeleteTestCaseSectionRequest {
	this := DeleteTestCaseSectionRequest{}
	return &this
}

// GetProjectName returns the ProjectName field value
func (o *DeleteTestCaseSectionRequest) GetProjectName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value
// and a boolean to check if the value has been set.
func (o *DeleteTestCaseSectionRequest) GetProjectNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectName, true
}

// SetProjectName sets field value
func (o *DeleteTestCaseSectionRequest) SetProjectName(v string) {
	o.ProjectName = v
}

// GetSectionId returns the SectionId field value
func (o *DeleteTestCaseSectionRequest) GetSectionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SectionId
}

// GetSectionIdOk returns a tuple with the SectionId field value
// and a boolean to check if the value has been set.
func (o *DeleteTestCaseSectionRequest) GetSectionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SectionId, true
}

// SetSectionId sets field value
func (o *DeleteTestCaseSectionRequest) SetSectionId(v int32) {
	o.SectionId = v
}

func (o DeleteTestCaseSectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteTestCaseSectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ProjectName"] = o.ProjectName
	toSerialize["SectionId"] = o.SectionId
	return toSerialize, nil
}

func (o *DeleteTestCaseSectionRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ProjectName",
		"SectionId",
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

	varDeleteTestCaseSectionRequest := _DeleteTestCaseSectionRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteTestCaseSectionRequest)

	if err != nil {
		return err
	}

	*o = DeleteTestCaseSectionRequest(varDeleteTestCaseSectionRequest)

	return err
}

type NullableDeleteTestCaseSectionRequest struct {
	value *DeleteTestCaseSectionRequest
	isSet bool
}

func (v NullableDeleteTestCaseSectionRequest) Get() *DeleteTestCaseSectionRequest {
	return v.value
}

func (v *NullableDeleteTestCaseSectionRequest) Set(val *DeleteTestCaseSectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteTestCaseSectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteTestCaseSectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteTestCaseSectionRequest(val *DeleteTestCaseSectionRequest) *NullableDeleteTestCaseSectionRequest {
	return &NullableDeleteTestCaseSectionRequest{value: val, isSet: true}
}

func (v NullableDeleteTestCaseSectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteTestCaseSectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

